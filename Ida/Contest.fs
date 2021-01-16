module Ida.Contest

open System.IO
open FSharp.Json

type Country =
    { Forum: string
      Names: string list }
    member this.PrimaryName = List.tryHead this.Names

    member this.Contains(text: string) =
        this.Names
        |> List.map (fun name -> name.ToLower())
        |> List.exists (text.ToLower().Contains)

type Entry =
    { Country: Country
      Artist: string
      Song: string }

type Vote = { Entry: Entry; Points: int }

type Contest =
    { Entries: Entry list
      Countries: Country list
      Voters: Country list }

    member this.FindVoter(country: string) =
        this.Voters
        |> List.tryFind (fun voter -> voter.Contains(country))

    member this.FindEntryByCountryName(text: string) =
        this.Entries
        |> List.tryFind (fun entry -> entry.Country.Contains(text))

    member this.FindEntryByCountryForum(text: string) =
        this.Entries
        |> List.tryFind (fun entry ->
            text
                .ToLower()
                .Contains(entry.Country.Forum.ToLower()))

    member this.FindEntryByArtist(text: string) =
        this.Entries
        |> List.tryFind (fun entry -> text.ToLower().Contains(entry.Artist.ToLower()))

    member this.FindEntryBySong(text: string) =
        this.Entries
        |> List.tryFind (fun entry -> text.ToLower().Contains(entry.Song.ToLower()))

let loadContest (path: string) =
    let contents = File.ReadAllText(path)

    let config =
        JsonConfig.create (jsonFieldNaming = Json.lowerCamelCase)

    Json.deserializeEx<Contest> config contents
