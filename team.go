package main

import "fmt"

type Team struct {
	playerA, playerB int
}

func (t *Team) full() bool {
	return t != nil && t.playerA != 0 && t.playerB != 0
}

func (t *Team) valid() bool {
	return t == nil || t.playerA != t.playerB
}

func (t *Team) players() []int {
	if t == nil {
		return nil
	}

	return []int{t.playerA, t.playerB}
}

func (t *Team) String() string {
	return fmt.Sprintf("(%d,%d)", t.playerA, t.playerB)
}
