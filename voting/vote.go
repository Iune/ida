package voting

import "github.com/iune/ida/contest"

type Vote struct {
	Points int
	Entry  contest.Entry
}
