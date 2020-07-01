require "json"

module Contest
  class Country
    getter forum : String
    getter iso : String
    getter names : Array(String)

    def initialize(@forum, @iso, @names)
    end

    def primary_name
      @names[0]
    end

    def contains_name?(line)
      @names
        .map { |name| name.downcase }
        .includes?(line.downcase)
    end
  end

  class Entry
    getter country : Country
    getter artist : String
    getter song : String

    def initialize(@country, @artist, @song)
    end

    def flag
      "World/#{@country.iso}.png"
    end
  end

  class Contest
    getter countries : Array(Country)
    getter entries : Array(Entry)
    getter voters : Array(Country)

    def initialize(@countries, @entries, @voters)
    end

    def find_voter_by_country_name(voter_name)
      @voters.find { |voter| voter.contains_name?(voter_name) }
    end

    def find_entry_by_artist(line)
      @entries.find { |entry| line.downcase.includes?(entry.artist.downcase) }
    end

    def find_entry_by_song(line)
      @entries.find { |entry| line.downcase.includes?(entry.song.downcase) }
    end

    def find_entry_by_country_name(line)
      @entries.find { |entry| entry.country.contains_name?(line) }
    end

    def find_entry_by_country_forum(line)
      @entries.find { |entry| line.downcase.includes?(entry.country.forum.downcase) }
    end

    def copy_votes_to_clipboard(voter, votes)
      votes_arr = Array.new(@voters.size, "")

      if voter
        voter_idx = @entries
          .map { |entry| entry.country }
          .index { |country| voter.primary_name == country.primary_name }
        if voter_idx
          votes_arr[voter_idx] = "X"
        end
      end

      votes.each { |vote|
        entry_idx = @entries
          .map { |entry| entry.country }
          .index { |country| vote.entry.country.primary_name == country.primary_name }
        if entry_idx
          votes_arr[entry_idx] = vote.points.to_s
        end
      }

      votes_str = votes_arr.join("\n")
      puts votes_str
      system "\"#{votes_str}\" | pbcopy"
    end

    private def self.parse_country(country)
      names = country["names"].as_a.map { |name| name.as_s }
      Country.new(country["forum"].as_s, country["iso"].as_s, names)
    end

    private def self.parse_entry(entry)
      country = parse_country(entry["country"].as_h)
      Entry.new(country, entry["artist"].as_s, entry["song"].as_s)
    end

    def self.from_file(file_name)
      json = JSON.parse(File.read(file_name))
      countries = json["countries"].as_a.map { |country| parse_country(country) }
      entries = json["entries"].as_a.map { |entry| parse_entry(entry) }
      voters = json["voters"].as_a.map { |voter| parse_country(voter) }
      Contest.new(countries, entries, voters)
    end
  end
end
