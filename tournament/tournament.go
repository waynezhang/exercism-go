package tournament

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
)

type team struct {
	name string
	win  int
	draw int
	loss int
}

type board map[string]*team

func (team *team) point() int {
	return team.win*3 + team.draw*1
}

func (team *team) matches() int {
	return team.win + team.draw + team.loss
}

func processBoard(board board, output io.Writer) {
	var head = fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team")
	output.Write([]byte(head))

	var teams = make([]*team, 0, len(board))
	for _, v := range board {
		teams = append(teams, v)
	}

	sort.Slice(teams, func(i, j int) bool {
		t1, t2 := teams[i], teams[j]
		return t1.point() > t2.point() || (t1.point() == t2.point() && t1.name < t2.name)
	})

	for _, team := range teams {
		var line = fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", team.name, team.matches(), team.win, team.draw, team.loss, team.point())
		output.Write([]byte(line))
	}
}

func parseMatch(match []byte) ([]team, error) {
	var cols = bytes.Split(match, []byte(";"))
	if len(cols) != 3 {
		return nil, errors.New("")
	}

	var team1, team2 = team{name: string(cols[0])}, team{name: string(cols[1])}
	switch string(cols[2]) {
	case "win":
		team1.win++
		team2.loss++
	case "draw":
		team1.draw++
		team2.draw++
	case "loss":
		team1.loss++
		team2.win++
	default:
		return []team{}, errors.New("")
	}
	return []team{team1, team2}, nil
}

// Tally document here
func Tally(input io.Reader, output io.Writer) error {
	var board = make(board)

	var reader = bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if bytes.HasPrefix(line, []byte("#")) || len(bytes.Trim(line, " ")) == 0 {
			continue
		}

		teams, err := parseMatch(line)
		if err != nil {
			return err
		}
		for _, t := range teams {
			if _, ok := board[t.name]; !ok {
				board[t.name] = &(team{name: t.name})
			}
			board[t.name].win += t.win
			board[t.name].draw += t.draw
			board[t.name].loss += t.loss
		}
	}

	processBoard(board, output)

	return nil
}
