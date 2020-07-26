import Cocoa

struct Country: Equatable, Hashable {
    let forum: String
    let iso: String
    let names: [String]

    init(forum: String, iso: String, names: [String]) {
        self.forum = forum
        self.iso = iso
        self.names = names
    }

    func primaryName() -> String {
        self.names.first ?? ""
    }

    func containsName(line: String) -> Bool {
        self.names.contains { name in
            return line.lowercased().contains(name.lowercased())
        }
    }
}

struct Entry: Hashable {
    let country: Country
    let artist: String
    let song: String

    init(country: Country, artist: String, song: String) {
        self.country = country
        self.artist = artist
        self.song = song
    }

    func flag() -> String {
        "World/\(self.country.iso).png"
    }
}

struct Vote: Hashable {
    let entry: Entry
    let points: Int
}

struct Contest {
    let countries: [Country]
    let entries: [Entry]
    let voters: [Country]

    init(countries: [Country], entries: [Entry], voters: [Country]) {
        self.countries = countries
        self.entries = entries
        self.voters = voters
    }

    func findVoterByCountryName(countryName: String) -> Country? {
        self.voters.first { voter in
            voter.containsName(line: countryName)
        }
    }

    func findEntryByCountryName(countryName: String) -> Entry? {
        self.entries.first { entry in
            entry.country.containsName(line: countryName)
        }
    }

    func findEntryByArtist(line: String) -> Entry? {
        self.entries.first { entry in
            line.lowercased().contains(entry.artist.lowercased())
        }
    }

    func findEntryBySong(line: String) -> Entry? {
        self.entries.first { entry in
            line.lowercased().contains(entry.song.lowercased())
        }
    }

    func findEntryByCountryForum(line: String) -> Entry? {
        self.entries.first { entry in
            line.lowercased().contains(entry.country.forum.lowercased())
        }
    }

    func copyVotesToClipboard(voter: Country?, votes: [Vote]) {
        var votesList = Array(repeating: "", count: self.entries.count)
        if voter != nil {
            if let index = self.entries.map({$0.country}).firstIndex(of: voter) {
                votesList[index] = "X"
            }
        }

        for vote in votes {
            if let index = self.entries.firstIndex(of: vote.entry) {
                votesList[index] = String(vote.points)
            }
        }

        let toCopy = votesList.joined(separator: "\n")
        let pasteBoard = NSPasteboard.general
        pasteBoard.clearContents()
        pasteBoard.setString(toCopy, forType: .string)
    }
}