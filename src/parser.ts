import { decode } from "html-entities";
import clipboardy from "clipboardy";
import { Contest, Country, Vote } from "./contest";
import { notEmpty } from "./utilities";

export interface ParsedVotes {
    votes: Vote[];
    warnings: string[];
}

export class Parser {
    contest: Contest

    constructor(contest: Contest) {
        this.contest = contest;
    }

    parse(lines: string[], voter?: Country): ParsedVotes {
        const votes: Vote[] = lines
            .map(line => decode(line))
            .map(line => this.getVoteFromLine(line))
            .filter(notEmpty);
        const warnings = this.getWarnings(votes, voter);
        return { votes, warnings };
    }

    private getVoteFromLine(line: string): Vote | undefined {
        // Search line for an entry
        const entry = [
            this.contest.findEntryByCountryForum(line),
            this.contest.findEntryByCountryName(line),
            this.contest.findEntryByArtist(line),
            this.contest.findEntryBySong(line),
        ].find(match => match !== undefined);
        // Search line for points
        const points = this.getPointsFromLine(line);
        // Check if we found an entry and points
        if (entry === undefined || points === undefined) {
            return undefined;
        }
        return new Vote(entry, points);
    }

    private getPointsFromLine(line: string): number | undefined {
        const search = line.match(/\d/g);
        if (search !== null) {
            const parsed = parseInt(search.join(''), 10);
            if (isNaN(parsed)) {
                return undefined;
            }
            return parsed;
        }
        return undefined;
    }

    private getWarnings(votes: Vote[], voter?: Country): string[] {
        function checkForSelfVoting(): string | undefined {
            if (voter === undefined) {
                return undefined;
            }
            const votedCountries = votes.map(vote => vote.entry.country);
            if (votedCountries.includes(voter)) {
                return `${voter.primaryName() ?? 'Voter'} voted for self`;
            }
            return undefined;
        }

        function checkForDuplicates(): string | undefined {
            if (new Set(votes).size !== votes.length) {
                return `At least one entry received points more than once`;
            }
            return undefined;
        }

        function checkPointTotals(): string | undefined {
            const total = votes.map(vote => vote.points).reduce((a, b) => a + b, 0);
            if (total !== 58) {
                return `Total number of points was not 58: ${total}`;
            }
            return undefined;
        }

        return [
            checkForSelfVoting(),
            checkForDuplicates(),
            checkPointTotals()
        ].filter(notEmpty);
    }

    copyVotesToClipboard(votes: Vote[], voter?: Country) {
        const votesArray: string[] = new Array(this.contest.entries.length).fill('');
        // Set voter cell to 'X' if the voter has an entry
        if (voter !== undefined) {
            const countries = this.contest.entries.map(entry => entry.country);
            const index = countries.indexOf(voter);
            votesArray[index] = 'X';
        }
        // Set the points awarded by the voter
        votes.forEach(vote => {
            const index = this.contest.entries.indexOf(vote.entry);
            votesArray[index] = vote.points.toString();
        })
        // Copy to clipboard
        const toCopy = votesArray.join('\n');
        clipboardy.writeSync(toCopy);
    }
}