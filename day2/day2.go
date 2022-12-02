package day2

import (
	"fmt"
	"strings"

	"github.com/alenius/aoctools"
)

const ROCK_OPPONENT string = "A"
const PAPER_OPPONENT string = "B"
const SCISSOR_OPPONENT string = "C"

const ROCK_ME string = "X"
const PAPER_ME string = "Y"
const SCISSOR_ME string = "Z"

const NEEDS_LOSS string = "X"
const NEEDS_DRAW string = "Y"
const NEEDS_WIN string = "Z"

const ROCK_PTS int = 1
const PAPER_PTS int = 2
const SCISSOR_PTS int = 3

type play struct {
	me           string
	predictedEnd string
	opponent     string
	res          int
	pts          int
}

func parsePlay(s string) play {
	sp := strings.Split(s, " ")
	return play{
		me:       sp[1],
		opponent: sp[0],
	}
}

func parsePlay2(s string) play {
	sp := strings.Split(s, " ")
	return play{
		predictedEnd: sp[1],
		opponent:     sp[0],
	}
}

func Day2() {
	lines := aoctools.ReadNewlineSeparatedFile("./day2/input.txt")

	plays := []*play{}
	for _, line := range lines {
		p := parsePlay(line)
		plays = append(plays, &p)
	}

	for _, play := range plays {
		result := checkMyWin(play.me, play.opponent)
		play.res = result
		score := calcScore(play.me, result)
		play.pts = score
	}

	totScore := 0
	for _, p := range plays {
		totScore += p.pts
	}
	fmt.Printf("%+v\n", totScore)

	plays2 := []*play{}
	for _, line := range lines {
		p := parsePlay2(line)
		plays2 = append(plays2, &p)
	}

	tot2 := 0
	for _, play := range plays2 {
		need := checkNeeded(play.predictedEnd, play.opponent)

		res := 0
		switch play.predictedEnd {
		case NEEDS_LOSS:
			res = LOSS
		case NEEDS_DRAW:
			res = EQUAL
		case NEEDS_WIN:
			res = WIN
		default:
			panic("what")
		}
		score := calcScore(need, res)
		tot2 += score
	}
	fmt.Printf("%+v\n", tot2)
}

func calcScore(sel string, res int) int {
	winpts := 6
	drawpts := 3
	losspts := 0
	switch res {
	case WIN:
		switch sel {
		case ROCK_ME:
			return ROCK_PTS + winpts
		case PAPER_ME:
			return PAPER_PTS + winpts
		case SCISSOR_ME:
			return SCISSOR_PTS + winpts
		default:
			panic("not valid me value")
		}
	case EQUAL:
		switch sel {
		case ROCK_ME:
			return ROCK_PTS + drawpts
		case PAPER_ME:
			return PAPER_PTS + drawpts
		case SCISSOR_ME:
			return SCISSOR_PTS + drawpts
		default:
			panic("not valid me value")
		}
	case LOSS:
		switch sel {
		case ROCK_ME:
			return ROCK_PTS + losspts
		case PAPER_ME:
			return PAPER_PTS + losspts
		case SCISSOR_ME:
			return SCISSOR_PTS + losspts
		default:
			panic("not valid me value")
		}
	default:
		panic("not valid res")
	}
}

const WIN = 1
const EQUAL = 0
const LOSS = -1

func checkMyWin(x, y string) int {
	switch x {
	case ROCK_ME:
		switch y {
		case ROCK_OPPONENT:
			return EQUAL
		case PAPER_OPPONENT:
			return LOSS
		case SCISSOR_OPPONENT:
			return WIN
		default:
			panic("not correct opponent string")
		}
	case PAPER_ME:
		switch y {
		case ROCK_OPPONENT:
			return WIN
		case PAPER_OPPONENT:
			return EQUAL
		case SCISSOR_OPPONENT:
			return LOSS
		default:
			panic("not correct opponent string")
		}
	case SCISSOR_ME:
		switch y {
		case ROCK_OPPONENT:
			return LOSS
		case PAPER_OPPONENT:
			return WIN
		case SCISSOR_OPPONENT:
			return EQUAL
		default:
			panic("not correct opponent string")
		}
	default:
		panic("not correct me string")
	}
}

func checkNeeded(x, y string) string {
	switch x {
	case NEEDS_LOSS:
		switch y {
		case ROCK_OPPONENT:
			return SCISSOR_ME
		case PAPER_OPPONENT:
			return ROCK_ME
		case SCISSOR_OPPONENT:
			return PAPER_ME
		default:
			panic("not correct opponent string")
		}
	case NEEDS_DRAW:
		switch y {
		case ROCK_OPPONENT:
			return ROCK_ME
		case PAPER_OPPONENT:
			return PAPER_ME
		case SCISSOR_OPPONENT:
			return SCISSOR_ME
		default:
			panic("not correct opponent string")
		}
	case NEEDS_WIN:
		switch y {
		case ROCK_OPPONENT:
			return PAPER_ME
		case PAPER_OPPONENT:
			return SCISSOR_ME
		case SCISSOR_OPPONENT:
			return ROCK_ME
		default:
			panic("not correct opponent string")
		}
	default:
		panic(fmt.Sprintf("not correct me string: %v", x))
	}
}
