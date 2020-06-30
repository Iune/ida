require "option_parser"

require "./ida/contest"
require "./ida/parser"

module Ida
  VERSION = "0.4.0"

  def self.parse_args
    contest_file_path = ""
    OptionParser.parse do |parser|
      parser.banner = "Usage: ida [arguments]"
      parser.on("-c CONTEST", "--contest=CONTEST", "(Required) JSON file containing contest details") { |input| contest_file_path = input }
      parser.on("-h", "--help", "Show this help") {
        puts parser
        exit 0
      }
    end

    if contest_file_path.size == 0
      puts "Contest file was not specified"
      exit 1
    else
      return contest_file_path
    end
  end

  def self.run
    contest_file_path = parse_args()
    contest = Contest::Contest.from_file(contest_file_path)
    parser = Parser::Parser.new(contest)

    while true
      printf "Country Name:\n> "
      country = read_line
      voter = contest.find_voter_by_country_name(country)

      lines = [] of String
      while true
        begin
          lines.push(read_line)
        rescue IO::EOFError
          break
        end
      end

      puts ""
      parser.parse(contest.countries[0], lines)
      puts ""
    end
  end
end

Ida.run
