open Argu
open Ida.Contest
open Ida.Parser
open Ida.Console

type Arguments =
    | [<Mandatory>] Contest of path: string

    interface IArgParserTemplate with
        member s.Usage =
            match s with
            | Contest _ -> "JSON file containing contest details"

let parseArguments (argv: string []) =
    let errorHandler = ProcessExiter()

    let argumentParser =
        ArgumentParser.Create<Arguments>(programName = "ida", errorHandler = errorHandler)

    let results = argumentParser.Parse(argv)
    results.GetResult Contest

[<EntryPoint>]
let main argv =
    let contestFile = parseArguments argv
    let contest = loadContest (contestFile)
    let parser = { Contest = contest }
    voterLoop (parser, true)
    0
