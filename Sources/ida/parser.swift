import Foundation
import HTMLString
import Rainbow

extension String {
    var digits: String {
        components(separatedBy: CharacterSet.decimalDigits.inverted).joined()
    }
}

extension Array where Element: Hashable {
    var isUnique: Bool {
        var seen = Set<Int>()
        return allSatisfy {
            seen.insert($0.hashValue).inserted
        }
    }
}

class Parser {
    let contest: Contest

    init(contest: Contest) {
        self.contest = contest
    }

    func parse(voter: Country?, lines: [String]) {
        let votes = lines
                .map({ $0.removingHTMLEntities })
                .map({ getVotes(line: $0) })
                .compactMap({ $0 })

        printVotes(votes: votes)
        self.contest.copyVotesToClipboard(voter: voter, votes: votes)

        if voter != nil && votes.first(where: { $0.entry.country == voter }) != nil {
            print("\(voter!.primaryName()) voted for self.".red)
        }

        let voteRecipients = votes.map({ $0.entry })
        if !voteRecipients.isUnique {
            print("At least one entry received points more than once.".red)
        }

        let pointsTotal = votes.map({ $0.points }).reduce(0, +)
        if pointsTotal != 58 {
            print("Total number of points was not 58: \(pointsTotal).".red)
        }
    }

    func getVotes(line: String) -> Vote? {
        var entry = self.contest.findEntryByCountryForum(line: line)
        if entry == nil {
            entry = self.contest.findEntryByCountryName(countryName: line)
        }
        if entry == nil {
            entry = self.contest.findEntryByArtist(line: line)
        }
        if entry == nil {
            entry = self.contest.findEntryBySong(line: line)
        }
        if entry == nil {
            return nil
        }

        if let points = getPoints(line: line) {
            return Vote(entry: entry!, points: points)
        } else {
            return nil
        }
    }

    func getPoints(line: String) -> Int? {
        Int(line.digits)
    }

    func printVotes(votes: [Vote]) {
        print("Found the following votes:")
        for vote in votes.sorted(by: { $0.points > $1.points }) {
            print("\(String(vote.points).padding(toLength: 2, withPad: " ", startingAt: 0)) | \(vote.entry.country.primaryName()): \(vote.entry.artist) – \(vote.entry.song)")
        }
    }
}