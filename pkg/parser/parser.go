package parser

import "github.com/Iune/ida/pkg/contest"

type Parser struct {
	Contest contest.Contest
}

type Vote struct {
	Entry  contest.Entry
	Points int
}
