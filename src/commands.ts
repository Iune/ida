import { Contest, Country, Entry } from "./contest";
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