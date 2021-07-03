import { Command } from 'commander';
import { spreadsheet } from "./commands";

const program = new Command();

program
    .version('0.4.0')
    .command('spreadsheet <countriesFile> <entriesFile> <outputFilePrefix>')
    .description('Generate Excel spreadsheet and contest JSON file based on input country and entry details')
    .action(async (countriesFile, entriesFile, outputFilePrefix) => {
        spreadsheet(countriesFile, entriesFile, outputFilePrefix)
    });

program.parse();

