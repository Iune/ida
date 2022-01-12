from dataclasses import dataclass
from string import digits
from typing import List, Optional
from ida.contest import Contest, Country, Entry
import html
import pyperclip

_DIGITS = frozenset(digits)


@dataclass
class Vote:
    entry: Entry
    points: int


@dataclass
class ParsedVotes:
    votes: List[Vote]
    warnings: List[str]


@dataclass
class Parser:
    contest: Contest

    def parse(self, lines: List[str], voter: Optional[Country]) -> ParsedVotes:
        lines = [html.unescape(line) for line in lines]

        votes = [self._get_votes(line) for line in lines]
        votes = [vote for vote in votes if vote is not None]
        votes = sorted(votes, key=lambda x: x.points, reverse=True)
        warnings = self._get_warnings(votes, voter)

        return ParsedVotes(votes, warnings)

    def _get_votes(self, line: str) -> Optional[Vote]:
        search_functions = (
            self.contest.find_entry_by_country_forum,
            self.contest.find_entry_by_country_name,
            self.contest.find_entry_by_artist,
            self.contest.find_entry_by_song,
        )

        entry = None
        for search in search_functions:
            if entry := search(line):
                break

        if not entry:
            return None

        if points := self._get_points(line):
            return Vote(entry, points)

        return None

    def _get_points(self, line: str) -> Optional[int]:
        try:
            return int("".join(c for c in line if c in _DIGITS))
        except ValueError:
            return None

    def _get_warnings(self, votes: List[Vote], voter: Optional[Country]) -> List[str]:
        def check_for_self_voting() -> Optional[str]:
            vote_countries = [vote.entry.country for vote in votes]
            for vote in votes:
                if voter in vote_countries:
                    voter_name = voter.primary_name if voter else "Voter"
                    return f"{voter_name} voted for self"
            return None

        def check_for_duplicate_votes() -> Optional[str]:
            vote_entries = {vote.entry for vote in votes}
            if len(votes) != len(vote_entries):
                return "At least one entry received poitns more than once"
            return None

        def check_for_points_total() -> Optional[str]:
            points_total = sum([vote.points for vote in votes])
            if points_total != 58:
                return f"Total number of points was not 58: {points_total}"
            return None

        warnings = [
            check_for_self_voting(),
            check_for_duplicate_votes(),
            check_for_points_total(),
        ]
        warnings = [warning for warning in warnings if warning is not None]
        return warnings

    def copy_votes(self, votes: ParsedVotes, voter: Optional[Country]) -> None:
        def get_vote_string(entry: Entry) -> str:
            def get_entry_string() -> str:
                vote = next(
                    (vote for vote in votes.votes if vote.entry == entry),
                    None,
                )

                return str(vote.points) if vote else ""

            if voter:
                return "X" if voter == entry.country else get_entry_string()
            return get_entry_string()

        votes_str = "\n".join(
            [get_vote_string(entry) for entry in self.contest.entries]
        )
        pyperclip.copy(votes_str)
