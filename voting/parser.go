package voting

import (
	"regexp"
	"strconv"

	"github.com/iune/ida/contest"
)

func Find(c contest.Contest, lines []string) []Vote {
	var votes []Vote
	for _, line := range lines {
		vote, found := findEntryOnLine(c, line)
		if found {
			votes = append(votes, vote)
		}
	}
	return votes
}

func findEntryOnLine(contest contest.Contest, line string) (vote Vote, found bool) {
	// Check if flagicon present
	entry, found := contest.FindCountryByForum(line)
	if !found {
		// Check if country name is present
		entry, found = contest.FindCountryByName(line)
		if !found {
			// Check artist is present
			entry, found = contest.FindArtist(line)
			if !found {
				// Check if song is present
				entry, found = contest.FindSong(line)
			}
		}
	}

	if !found {
		return Vote{}, false
	}
	// Now we check to see if there were any points on this line
	points, foundPoints := findPointsOnLine(line)
	if !foundPoints {
		return Vote{}, false
	}
	return Vote{Entry: entry, Points: points}, found
}

func findPointsOnLine(line string) (points int, found bool) {
	re, _ := regexp.Compile("[^0-9]")
	numericLine := re.ReplaceAllString(line, "")
	points, _ = strconv.Atoi(numericLine)
	if (points > 0 && points < 10) || (points == 10 || points == 12) {
		return points, true
	}
	return -1, false
}
