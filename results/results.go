package results

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Iune/ida/contest"
	"github.com/Iune/ida/voting"
)

func Output(contest contest.Contest, votes []voting.Vote, voterName string) string {
	votesArray := initializeVotesArray(contest)

	// Set points for entries
	for _, vote := range votes {
		entryIdx, found := contest.GetEntryIndex(vote.Entry)
		if found {
			votesArray[entryIdx] = fmt.Sprintf("%d", vote.Points)
		}
	}

	// Add placeholder 'X' to voter entry, if present
	voterEntry, found := contest.FindCountryByName(voterName)
	if found {
		// Check and warn if voter voted for self
		voterVotedForSelf := checkIfVoterVotedForSelf(votes, voterEntry)
		if voterVotedForSelf {
			log.WithFields(log.Fields{
				"voter": voterEntry.Country.GetPrimaryName(),
			}).Warning("Voter voted for self")
		}

		voterIdx, foundVoter := contest.GetEntryIndex(voterEntry)
		if foundVoter {
			votesArray[voterIdx] = "X"
		}
	}

	return strings.Join(votesArray, "\n")
}

func initializeVotesArray(contest contest.Contest) []string {
	var votesArray []string
	for i := 0; i < len(contest.Entries); i++ {
		votesArray = append(votesArray, "")
	}
	return votesArray
}

func checkIfVoterVotedForSelf(votes []voting.Vote, voter contest.Entry) bool {
	for _, vote := range votes {
		if vote.Entry.Equals(voter) {
			return true
		}
	}
	return false
}

func PrintVotes(votes []voting.Vote) {
	sort.Slice(votes, func(i, j int) bool {
		return votes[i].Points > votes[j].Points
	})

	fmt.Println("Found the following votes:")
	for _, vote := range votes {
		fmt.Printf("%02d | %s\n", vote.Points, vote.Entry.Country.GetPrimaryName())
	}
}
