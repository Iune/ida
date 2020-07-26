import Foundation
import SwiftyJSON

func loadContest(atPath: String) throws -> Contest {
    func parseCountry(country: JSON) -> Country {
        Country(forum: country["forum"].stringValue, iso: country["iso"].stringValue, names: country["names"].arrayValue.map({ $0.stringValue }))
    }

    func parseEntry(entry: JSON) -> Entry {
        Entry(country: parseCountry(country: entry["country"]), artist: entry["artist"].stringValue, song: entry["song"].stringValue)
    }

    let fm = FileManager.default
    guard let data = fm.contents(atPath: atPath) else {
        throw IdaError.runtimeError("Could not read contents of file at \(atPath)")
    }

    let json = try JSON(data: data)
    let countries = json["countries"].arrayValue.map({ country in parseCountry(country: country) })
    let entries = json["entries"].arrayValue.map({ entries in parseEntry(entry: entries) })
    let voters = json["voters"].arrayValue.map({ country in parseCountry(country: country) })

    return Contest(countries: countries, entries: entries, voters: voters)
}