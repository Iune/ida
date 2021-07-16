import { loadJSON, loadTextLines } from "./utilities";

export class Country {
    forum: string;
    names: string[];
    flag?: string;

    constructor(forum: string, names: string[], flag?: string) {
        this.forum = forum;
        this.names = names;
        this.flag = flag;
    }

    static fromFile(path: string): Country[] {
        interface CountriesFile {
            countries: Country[]
        }

        const data = loadJSON(path) as CountriesFile;
        return data.countries.map(c => new Country(c.forum, c.names, c.flag));
    }

    primaryName(): string | undefined {
        if (this.names.length > 0) {
            return this.names[0];
        }
        return undefined;
    }

    containedIn(text: string): boolean {
        return this.names.some(name => text.toLowerCase().includes(name.toLowerCase()));
    }
}


export class Entry {
    country: Country;
    artist: string;
    song: string;

    constructor(country: Country, artist: string, song: string) {
        this.country = country;
        this.artist = artist;
        this.song = song;
    }

    static fromFile(path: string, countries: Country[]): Entry[] {
        const countriesMap: Record<string, Country> = {};
        countries.forEach(country => countriesMap[country.forum] = country);
        return loadTextLines(path)
            .filter(line => line) // empty strings '' evaluate to false and is filtered out
            .map(line => {
                const forum = line.match(/:([A-z]*):/)![0].replace(/:/g, '');
                const country: Country = countriesMap[forum];
                const [artist, song] = line.match(/\[B\](.*)\[\/B\]/)![0]
                    .replace('[B]', '').replace('[/B]', '')
                    .split(' - ')
                return new Entry(country, artist, song);
            });
    }

    flag(): string {
        if (this.country.flag === undefined) { return ''; }
        else { return `World/${this.country.flag}.png`; }
    }
}

export class Vote {
    constructor(
        public entry: Entry,
        public points: number
    ) { }
}

export class Contest {
    entries: Entry[];
    countries: Country[];
    voters: Country[];

    constructor(entries: Entry[], countries: Country[], voters: Country[]) {
        this.entries = entries;
        this.countries = countries;
        this.voters = voters;
    }

    static fromFile(path: string): Contest {
        const data = loadJSON(path) as Contest;
        const entries = data.entries.map(e => new Entry(new Country(e.country.forum, e.country.names, e.country.flag), e.artist, e.song));
        const countries = data.countries.map(c => new Country(c.forum, c.names, c.flag));
        const voters = data.voters.map(c => new Country(c.forum, c.names, c.flag));
        return new Contest(entries, countries, voters);
    }

    findVoterByCountryName(voterName: string): Country | undefined {
        return this.voters.find(voter => voter.containedIn(voterName));
    }

    findEntryByCountryName(text: string): Entry | undefined {
        return this.entries.find(entry => entry.country.containedIn(text));
    }

    findEntryByCountryForum(text: string): Entry | undefined {
        return this.entries.find(entry => text.toLowerCase().includes(entry.country.forum.toLowerCase()));
    }

    findEntryByArtist(text: string): Entry | undefined {
        return this.entries.find(entry => text.toLowerCase().includes(entry.artist.toLowerCase()));
    }

    findEntryBySong(text: string): Entry | undefined {
        return this.entries.find(entry => text.toLowerCase().includes(entry.song.toLowerCase()));
    }
}