from dataclasses import dataclass
from typing import List
from html import unescape
from string import digits
from collections import Counter
from colorama import init
from termcolor import colored

import re
import json
import argparse
import pyperclip

init()
digits = frozenset(digits)

@dataclass(frozen=True)
class Country:
    forum: str
    iso: str
    names: List[str]

    def primary_name(self):
        return self.names[0]

    def contains_name(self, line):
        return any(name.lower() in line.lower() for name in self.names)


@dataclass(frozen=True)
class Entry:
    country: Country
    artist: str
    song: str

    def flag(self):
        return "World/{}.png".format(self.country.iso)


@dataclass(frozen=True)
class Contest:
    countries: List[Country]
    entries: List[Entry]
    voters: List[Country]

    def find_voter_by_country_name(self, voter_name):
        return next((voter for voter in self.voters if voter.contains_name(voter_name)), None)

    def find_entry_by_artist(self, line):
        return next((entry for entry in self.entries if entry.artist.lower() in line.lower()), None)

    def find_entry_by_song(self, line):
        return next((entry for entry in self.entries if entry.song.lower() in line.lower()), None)

    def find_entry_by_country_name(self, line):
        return next((entry for entry in self.entries if entry.country.contains_name(line)), None)

    def find_entry_by_country_forum(self, line):
        return next((entry for entry in self.entries if entry.country.forum.lower() in line.lower()), None)

    def copy_votes_to_clipboard(self, voter, votes):
        votes_lst = [""] * len(self.voters)

        if voter:
            countries = [entry.country for entry in self.entries]
            voter_idx = countries.index(voter)
            votes_lst[voter_idx] = "X"

        for vote in votes:
            entry_idx = self.entries.index(vote.entry)
            votes_lst[entry_idx] = str(vote.points)

        votes_str = "\n".join(votes_lst)
        pyperclip.copy(votes_str)


@dataclass
class Vote:
    entry: Entry
    points: int

    @staticmethod
    def _print_votes(votes):
        print("Found the following votes:")
        for vote in sorted(votes, key=lambda x: [-x.points, x.entry.country.primary_name()]):
            print("{:2} | {}: {} - {}".format(vote.points,
                                              vote.entry.country.primary_name(), vote.entry.artist, vote.entry.song))


def load_contest(file_name):
    def load_json():
        with open(file_name, "r") as f:
            return json.load(f)

    def parse_country(country):
        return Country(
            forum=country["forum"],
            iso=country["iso"],
            names=country["names"]
        )

    def parse_entry(entry):
        return Entry(
            country=parse_country(entry["country"]),
            artist=entry["artist"],
            song=entry["song"]
        )

    contest = load_json()
    countries = [parse_country(country) for country in contest["countries"]]
    voters = [parse_country(country) for country in contest["voters"]]
    entries = [parse_entry(entry) for entry in contest["entries"]]
    return Contest(countries=countries, voters=voters, entries=entries)


class Parser:
    def __init__(self, contest):
        self.contest = contest

    def parse(self, voter, lines):
        votes = [self._get_votes(unescape(line)) for line in lines]
        votes = [vote for vote in votes if vote]

        if any(voter == vote.entry.country for vote in votes):
            print(colored("{} voted for themselves".format(
                voter.primary_name()), "red"))

        vote_recipients = [vote.entry.country.primary_name() for vote in votes]
        duplicate_recipients = [country for country, count in Counter(
            vote_recipients).items() if count > 1]
        for country in duplicate_recipients:
            print(colored("{} received points more than once".format(country), "red"))

        points_total = sum([vote.points for vote in votes])
        if points_total != 58:
            print(colored("Total number of points was not 58: {}".format(points_total), "red"))

        print()
        Vote._print_votes(votes)
        self.contest.copy_votes_to_clipboard(voter, votes)

    def _get_votes(self, line):
        entry = self.contest.find_entry_by_country_forum(line)
        if not entry:
            entry = self.contest.find_entry_by_country_name(line)
        if not entry:
            entry = self.contest.find_entry_by_artist(line)
        if not entry:
            entry = self.contest.find_entry_by_song(line)
        if not entry:
            return None

        points = self._get_points(line)
        if points:
            return Vote(entry=entry, points=points)
        else:
            return None

    def _get_points(self, line):
        try:
            return int(''.join(c for c in line if c in digits))
        except ValueError:
            return None

def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("contest", help="JSON file containing contest details")
    return parser.parse_args()


def main():
    args = parse_args()
    contest = load_contest(args.contest)
    parser = Parser(contest)

    while True:
        country = input("Country Name:\n> ")
        voter = contest.find_voter_by_country_name(country)

        lines = []
        while True:
            try:
                lines.append(input())
            except EOFError:
                break

        print()
        parser.parse(voter, lines)
        print()


if __name__ == "__main__":
    main()
