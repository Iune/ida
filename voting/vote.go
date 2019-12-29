package voting

import "github.com/Iune/ida/contest"

type Vote struct {
	Points int
	Entry  contest.Entry
}
