import { Contest, Country, Entry } from "./contest";
import { createWorkbook, writeWorkbook } from "./spreadsheet";

async function main() {
    const countries = Country.fromFile("resources/countries.json")
    const entries = Entry.fromFile("resources/entries.txt", countries);
    // const contest = Contest.fromFile("resources/sample.json")
    // console.log(contest);

    const workbook = createWorkbook(entries);
    await writeWorkbook(workbook, "resources/output.xlsx");
}

(async () => {
    try {
        await main();
    } catch (err) {
        console.error(err);
    }
})();

