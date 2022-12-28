package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	headerFormat = "%-30s | %2s | %2s | %2s | %2s | %2s\n"
	bodyFormat   = "%-30s | %2d | %2d | %2d | %2d | %2d\n"
)

type team struct {
	name  string
	won   int
	drawn int
	lost  int
}

func (t *team) matchesPlayed() int {
	return t.won + t.drawn + t.lost
}

func (t *team) points() int {
	return t.won*3 + t.drawn*1
}

func Tally(reader io.Reader, writer io.Writer) error {
	lines := read(reader)

	teams, err := buildTeams(lines)
	if err != nil {
		return err
	}

	sortTeams(teams)
	write(teams, writer)

	return nil
}

func read(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func buildTeams(data []string) ([]*team, error) {
	teamsMap := make(map[string]*team)
	for _, line := range data {
		trimedLine := strings.TrimSpace(line)
		if len(trimedLine) == 0 ||
			strings.HasPrefix(trimedLine, "#") {
			continue
		}

		if strings.Count(line, ";") != 2 {
			return nil, fmt.Errorf("invalid format, %s", line)
		}

		record := strings.Split(line, ";")
		if record[2] != "win" && record[2] != "loss" && record[2] != "draw" {
			return nil, fmt.Errorf("invalid value, %s", record[2])
		}

		team1, ok := teamsMap[record[0]]
		if !ok {
			team1 = &team{name: record[0]}
			teamsMap[record[0]] = team1
		}

		team2, ok := teamsMap[record[1]]
		if !ok {
			team2 = &team{name: record[1]}
			teamsMap[record[1]] = team2
		}

		switch record[2] {
		case "win":
			team1.won++
			team2.lost++
		case "draw":
			team1.drawn++
			team2.drawn++
		case "loss":
			team1.lost++
			team2.won++
		}
	}

	teamsSlice := []*team{}
	for _, v := range teamsMap {
		teamsSlice = append(teamsSlice, v)
	}
	return teamsSlice, nil
}

func sortTeams(teams []*team) {
	// implement the less function
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points() == teams[j].points() {
			return teams[i].name < teams[j].name
		}
		return teams[i].points() > teams[j].points()
	})
}

func write(teams []*team, writer io.Writer) {
	w := bufio.NewWriter(writer)

	fmt.Fprintf(w, headerFormat, "Team", "MP", "W", "D", "L", "P")
	for _, t := range teams {
		fmt.Fprintf(w, bodyFormat,
			t.name,
			t.matchesPlayed(),
			t.won,
			t.drawn,
			t.lost,
			t.points())
	}
	w.Flush()
}
