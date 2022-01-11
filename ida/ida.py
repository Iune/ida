from os import get_terminal_size
from pathlib import Path
from tabulate import tabulate
from ida.contest import load_contest
from ida.parser import ParsedVotes, Parser
import typer

app = typer.Typer()


def print_votes(votes: ParsedVotes) -> None:
    # Print table of parsed votes
    headers = ["Points", "Country", "Artist", "Song"]
    rows = [
        [
            vote.points,
            vote.entry.country.primary_name(),
            vote.entry.artist,
            vote.entry.song,
        ]
        for vote in votes.votes
    ]
    print(tabulate(rows, headers=headers, tablefmt="fancy_grid"))

    # Print warnings
    for warning in votes.warnings:
        typer.secho(warning, fg=typer.colors.RED, bold=True)


def voter_loop(parser: Parser, first_run: bool = False) -> None:
    if not first_run:
        terminal_size = get_terminal_size()
        print("─" * terminal_size.columns)

    voter = input("Voter Name:\n> ")
    voter = parser.contest.find_voter(voter)

    print()
    print("Votes:")

    lines = []
    while True:
        try:
            lines.append(input())
        except EOFError:
            break

    votes = parser.parse(lines, voter)
    parser.copy_votes(votes, voter)

    print()
    print_votes(votes)


@app.command()
def main(contest_file: Path):
    contest = load_contest(contest_file)
    parser = Parser(contest)

    first_run = True
    while True:
        voter_loop(parser, first_run)
        first_run = False


def run() -> None:
    app()
