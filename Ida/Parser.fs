module Ida.Parser

open System
open System.Text.RegularExpressions
open System.Web
open Ida.Contest
open TextCopy

type ParsedVotes =
    { Votes: Vote list
      Warnings: string list }

type Parser =
    { Contest: Contest }
    member this.Parse(voter: Country option, lines: string list) =
        let votes =
            lines
            |> List.map (this.unescapeHTML)
            |> List.map (this.getVotes)
            |> List.choose id
            |> List.sortByDescending (fun vote -> vote.Points)

        let warnings = this.getWarnings (voter, votes)
        { Votes = votes; Warnings = warnings }

    member this.CopyVotes(voter: Country option, votes: ParsedVotes) =
        let getVoteString (entry: Entry) =
            let getEntryString =
                let vote =
                    votes.Votes
                    |> List.tryFind (fun vote -> vote.Entry.Country = entry.Country)

                match vote with
                | Some (vote) -> vote.Points.ToString()
                | None -> ""

            match voter with
            | Some (voter) -> if voter = entry.Country then "X" else getEntryString
            | None -> getEntryString

        let votesString =
            this.Contest.Entries
            |> List.map getVoteString
            |> String.concat "\n"

        ClipboardService.SetText(votesString)

    member private this.unescapeHTML(text: string) = HttpUtility.HtmlDecode text

    member private this.getVotes(line: string) =
        let entry =
            [ this.Contest.FindEntryByCountryForum(line)
              this.Contest.FindEntryByCountryName(line)
              this.Contest.FindEntryByArtist(line)
              this.Contest.FindEntryBySong(line) ]
            |> List.choose id
            |> List.tryHead

        let points = this.getPoints (line)

        match (entry, points) with
        | (Some (entry), Some (points)) -> Some({ Entry = entry; Points = points })
        | _ -> None

    member private this.getPoints(line: string) =
        match Int32.TryParse(Regex.Match(line, "\\d+").Value) with
        | true, int -> Some(int)
        | _ -> None

    member private this.getWarnings(voter: Country option, votes: Vote list) =
        let checkForSelfVoting =
            let voteCountries =
                votes |> List.map (fun vote -> vote.Entry.Country)

            match voter with
            | Some (voter) ->
                if List.contains voter voteCountries then
                    let voterName =
                        voter.PrimaryName |> Option.defaultValue "Voter"

                    Some($"{voterName} voted for self")
                else
                    None
            | None -> None

        let checkForDuplicateVotes =
            if votes.Length <> Seq.length (Seq.distinct (votes))
            then Some("At least one entry received points more than once")
            else None

        let checkForPointsTotal =
            let pointsTotal =
                votes |> List.sumBy (fun vote -> vote.Points)

            if pointsTotal <> 58
            then Some($"Total number of points was not 58: {pointsTotal}")
            else None

        [ checkForSelfVoting
          checkForDuplicateVotes
          checkForPointsTotal ]
        |> List.choose id
