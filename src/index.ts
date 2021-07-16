#!/usr/bin/env node

import { Command } from 'commander';
import { parse, spreadsheet } from "./commands";

const program = new Command();

program
    .version('0.4.0')
    .command('spreadsheet <countriesFile> <entriesFile> <outputFilePrefix>')
    .description('Generate Excel spreadsheet and contest JSON file based on country and entry details')
    .action(async (countriesFile, entriesFile, outputFilePrefix) => {
        spreadsheet(countriesFile, entriesFile, outputFilePrefix)
    });

program
    .command('parse <contestFile>')
    .description('Parse votes for a given contest')
    .action(async (contestFile) => {
        parse(contestFile);
    });

program.parse();

