package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maxNums = 26 * 26 * 1000
)

var (
	index int
	names []string
)

// Define the Robot type here.
type Robot struct {
	name string
}

func init() {
	index = 0
	names = []string{}

	for _, l1 := range letters {
		for _, l2 := range letters {
			for i := 0; i < 1000; i++ {
				name := fmt.Sprintf("%c%c%03d", l1, l2, i)
				names = append(names, name)
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if index < maxNums {
		r.name = names[index]
		index++
		return r.name, nil
	}

	return "", errors.New("Run out of robot names")
}

func (r *Robot) Reset() {
	r.name = ""
}
