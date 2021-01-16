module Ida.Console

open System
open Ida.Parser
open Spectre.Console

let printVotes (votes: ParsedVotes) =
    let mutable ptsCol = new TableColumn("Points")
    ptsCol <- ptsCol.Centered()

    let mutable table = new Table()
    table <- table.AddColumn(ptsCol)
    table <- table.AddColumn("Country")
    table <- table.AddColumn("Artist")
    table <- table.AddColumn("Song")

    votes.Votes
    |> List.iter (fun vote ->
        let countryName =
            vote.Entry.Country.PrimaryName
            |> Option.defaultValue ""

        table <- table.AddRow(vote.Points.ToString(), countryName, vote.Entry.Artist, vote.Entry.Song))

    AnsiConsole.Render(table)
    AnsiConsole.WriteLine()

    votes.Warnings
    |> List.iter (fun warning -> AnsiConsole.MarkupLine($"[red]Warning: {warning}[/]"))


// See https://stackoverflow.com/a/27796307
let readLines () =
    let read _ = Console.ReadLine()

    let isValid =
        function
        | null -> false
        | _ -> true

    Seq.initInfinite read
    |> Seq.takeWhile isValid
    |> Seq.toList

let rec voterLoop (parser: Parser, firstVoter: bool) =
    if not firstVoter then
        let rule = new Rule()
        AnsiConsole.Render(rule)

    let voter =
        AnsiConsole.Ask<string>("Voter Name")
        |> parser.Contest.FindVoter

    AnsiConsole.WriteLine()
    AnsiConsole.MarkupLine("Votes:")

    let lines = readLines ()
    let votes = parser.Parse(voter, lines)

    AnsiConsole.WriteLine()
    printVotes (votes)
    parser.CopyVotes(voter, votes)
    
    voterLoop (parser, false)
