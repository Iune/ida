import colors from "colors/safe";
import { getBorderCharacters, table, TableUserConfig } from "table";
import { ParsedVotes } from "./parser";
import { Contest, Country, Entry } from "./contest";
import { Parser } from "./parser";
import { createWorkbook, writeWorkbook } from "./spreadsheet";
import { writeJSON } from "./utilities";

export async function spreadsheet(countriesPath: string, entriesPath: string, outputPathPrefix: string) {
    // Load countries and entries
    const countries = Country.fromFile(countriesPath)
    const entries = Entry.fromFile(entriesPath, countries);
    // Create and write Excel file
    const workbook = createWorkbook(entries);
    await writeWorkbook(workbook, `${outputPathPrefix}.xlsx`);
    // Create and write JSON file
    const voters = entries.map(e => e.country);
    const contest = new Contest(entries, countries, voters);
    writeJSON(contest, `${outputPathPrefix}.json`);
}

function printVotes(votes: ParsedVotes) {
    // Print votes table
    const headers: string[] = ['Points', 'Country', 'Artist', 'Song'];
    const data: string[][] = votes.votes.map(vote => [vote.points.toString(), vote.entry.country.primaryName() ?? '', vote.entry.artist, vote.entry.song]);
    const config = {
        border: getBorderCharacters('norc'),
        columns: [
            { alignment: 'center' },
            { alignment: 'left' },
            { alignment: 'left' },
            { alignment: 'left' }
        ],
    } as TableUserConfig;
    console.log(table(data, config));
    console.log();
    // Print warnings
    votes.warnings.forEach(warning => console.log(colors.red(`Warning: ${warning}`)));
}

function voterLoop(parser: Parser, firstVoter: boolean) {
    if (!firstVoter) {
        const line = '-'.repeat(process.stdout.columns)
        console.log(line);
    }

    voterLoop(parser, false);
}

export async function parse(contestFilePath: string) {
    // Load contest
    const contest = Contest.fromFile(contestFilePath);
    const parser = new Parser(contest);
    const lines = [
        '12 Estonia',
        '10 India'
    ]
    const votes = parser.parse(lines);
    printVotes(votes);
}