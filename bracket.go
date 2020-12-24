package main

import (
	"strings"
)

type Bracket struct {
	numPlayers int
	rounds     []*Round
}

func NewBracket(b *Bracket) *Bracket {
	return &Bracket{numPlayers: b.numPlayers, rounds: NewRounds(b.rounds)}
}

func (bracket *Bracket) possibleTeams(teams []*Team) []*Team {
	roundIdx := len(bracket.rounds) - 1
	if roundIdx < 0 || bracket.rounds[roundIdx].full() {
		return teams
	}
	return bracket.rounds[roundIdx].possibleTeams(teams)
}

func (b *Bracket) Copy(other *Bracket) {
	b.numPlayers = other.numPlayers
	b.rounds = NewRounds(other.rounds)
}

func (b *Bracket) full() bool {
	return b != nil && int(b.numPlayers-1) == len(b.rounds) && b.rounds[len(b.rounds)-1].full()
}

func (b *Bracket) valid() bool {
	if b == nil {
		return false
	}

	var testReflectiveness bool = true
	for _, Round := range b.rounds {
		if !Round.valid() {
			return false
		}

		if testReflectiveness && !Round.full() {
			testReflectiveness = false
		}
	}

	groups := b.groups()
	for i, group := range groups {
		for _, other := range groups[i+1:] {
			if group.Equal(other) {
				return false
			}
		}
	}

	hasPlayed := b.hasPlayed()

	if testReflectiveness && !b.reflectiveHasPlayed(hasPlayed) {
		return false
	}

	for _, row := range hasPlayed {
		for _, column := range row {
			if column > 2 {
				return false
			}
		}
	}

	return true
}

func (b *Bracket) hasPlayed() [][]int {
	var hasPlayed = emptyHasPlayed(b.numPlayers)
	for _, Round := range b.rounds {
		hasPlayed = Round.hasPlayed(hasPlayed)
	}
	return hasPlayed
}

func (b *Bracket) reflectiveHasPlayed(hasPlayed [][]int) bool {
	for x := 0; x < int(b.numPlayers); x++ {
		for y := 0; y < int(b.numPlayers); y++ {
			if hasPlayed[x][y] != hasPlayed[int(b.numPlayers)-1-y][int(b.numPlayers)-1-x] {
				return false
			}
		}
	}
	return true
}

func (b *Bracket) groups() []Group {
	var groups []Group
	for _, r := range b.rounds {
		groups = append(groups, r.groups()...)
	}
	return groups
}

func (b *Bracket) addTeam(t *Team) bool {
	roundIdx := len(b.rounds) - 1
	if roundIdx < 0 || b.rounds[roundIdx].full() {
		b.rounds = append(b.rounds, &Round{
			numPlayers: b.numPlayers,
			games: []*Game{
				{
					teamA: &Team{
						playerA: 1,
						playerB: roundIdx + 3,
					},
					teamB: t,
				},
			},
		})
		return true
	}
	r := b.rounds[roundIdx]

	if in(r.players(), t.players()) {
		return false
	}
	r.addTeam(t)

	return b.valid()
}

func (b *Bracket) String() string {
	var roundStrings []string
	for _, Round := range b.rounds {
		roundStrings = append(roundStrings, Round.String())
	}
	return "----------\n" + strings.Join(roundStrings, "\n") + "\n----------"
}
