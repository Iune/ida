import Foundation
import Rainbow
import ArgumentParser

enum IdaError: Error {
    case runtimeError(String)
}

struct Ida: ParsableCommand {
    @Argument(help: "JSON file containing contest details")
    var contestPath: String

    mutating func run() throws {
        let contest = try loadContest(atPath: contestPath)
        let parser = Parser(contest: contest)

        print("Country Name:\n> ", terminator: "")
        let voterName = readLine()
        let voter = contest.findVoterByCountryName(countryName: voterName!)
        if voter == nil {
            print("No voter with the name '\(voterName!)' found.".yellow)
        }

        print()
        var lines: [String] = []
        while let readString = readLine() {
            guard readString.lowercased() != "done" else {
                break
            }
            lines.append(readString)
        }

        print()
        parser.parse(voter: voter, lines: lines)
    }
}

Ida.main()