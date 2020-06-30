require "colorize"
require "html"
require "./contest"

module Parser
  class Vote
    getter entry : Contest::Entry
    getter points : Int32

    def initialize(@entry, @points)
    end

    def self.print_votes(votes)
      puts "Found the following votes:"
      votes
        .sort_by{|vote| {-vote.points, vote.entry.country.primary_name}}
        .each{|vote| printf "%2d | %s: %s - %s\n", vote.points, vote.entry.country.primary_name, vote.entry.artist, vote.entry.song}
    
      # TODO
    end
  end

  class Parser
    getter contest : Contest::Contest

    def initialize(@contest)
    end

    def parse(voter, lines : Array(String))
      votes = lines.map { |line| get_votes(HTML.unescape(line)) }.compact
      
      if votes.find { |vote| voter.primary_name == vote.entry.country.primary_name }
        puts "#{voter.primary_name} voted for themselves".colorize(:red)
      end

      vote_recipients = votes.map { |vote| vote.entry.country.primary_name }
      duplicate_vote_recipients = vote_recipients.select { |country| vote_recipients.count(country) > 1 }.uniq
      duplicate_vote_recipients.each { |country| puts "#{country} received points more than once".colorize(:red) }

      points_total = votes.map{|vote| vote.points}.sum
      if points_total != 58
        puts "Total number of points was not 58: #{points_total}".colorize(:red)
      end      

      puts ""
      Vote.print_votes(votes)
      @contest.copy_votes_to_clipboard(voter, votes)
    end

    private def get_votes(line)
      entry = @contest.find_entry_by_country_forum(line)
      if !entry
        entry = @contest.find_entry_by_country_name(line)
      end
      if !entry
        entry = @contest.find_entry_by_artist(line)
      end
      if !entry
        entry = @contest.find_entry_by_song(line)
      end
      if !entry
        return nil
      end

      points = get_points(line)
      if !points
        return nil
      end

      Vote.new(entry, points)
    end

    private def get_points(line)
      digits = line.chars.select { |char| char.ascii_number? }
      digits.join.to_i { nil }
    end
  end
end
