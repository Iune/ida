from dataclasses import dataclass
from pathlib import Path
from typing import List, Optional
import json
import re


@dataclass(eq=True, frozen=True)
class Country:
    forum: str
    names: List[str]

    def primary_name(self) -> Optional[str]:
        return self.names[0] if self.names else None

    def contained_in(self, text: str) -> bool:
        return any(name.lower() in text.lower() for name in self.names)

    def contained_in_regex(self, text: str) -> bool:
        search_regex = r"\b({0})\b"
        return any(
            re.search(re.compile(search_regex.format(name), flags=re.IGNORECASE), text)
            for name in self.names
        )

    def __hash__(self):
        return hash(self.primary_name())


@dataclass(eq=True, frozen=True)
class Entry:
    country: Country
    artist: str
    song: str


@dataclass
class Contest:
    entries: List[Entry]
    countries: List[Country]
    voters: List[Country]

    def find_voter(self, text: str) -> Optional[Country]:
        return next((voter for voter in self.voters if voter.contained_in(text)), None)

    def find_entry_by_country_name(self, text: str) -> Optional[Entry]:
        # We need to search country names in whole words, or we can have clashes such as UK <-> (Uk)raine
        return next(
            (entry for entry in self.entries if entry.country.contained_in_regex(text)),
            None,
        )

    def find_entry_by_country_forum(self, text: str) -> Optional[Entry]:
        return next(
            (
                entry
                for entry in self.entries
                if entry.country.forum.lower() in text.lower()
            ),
            None,
        )

    def find_entry_by_artist(self, text: str) -> Optional[Entry]:
        return next(
            (entry for entry in self.entries if entry.artist.lower() in text.lower()),
            None,
        )

    def find_entry_by_song(self, text: str) -> Optional[Entry]:
        return next(
            (entry for entry in self.entries if entry.song.lower() in text.lower()),
            None,
        )


def load_contest(contest_file: Path) -> Contest:
    def parse_country(record: dict) -> Country:
        return Country(record["forum"], record["names"])

    def parse_entry(record: dict) -> Entry:
        return Entry(parse_country(record["country"]), record["artist"], record["song"])

    with open(contest_file, "r") as f:
        data = json.load(f)

    entries = [parse_entry(record) for record in data["entries"]]
    countries = [parse_country(record) for record in data["countries"]]
    voters = [parse_country(record) for record in data["voters"]]
    contest = Contest(entries=entries, countries=countries, voters=voters)
    return contest
