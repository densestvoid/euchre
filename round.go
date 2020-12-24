package main

import (
	"strings"
)

type Round struct {
	numPlayers int
	games      []*Game
}

func (r *Round) possibleGames(games []*Game) []*Game {
	players := r.players()

	var possibleGames []*Game
	for _, game := range games {
		if !in(players, game.players()) {
			possibleGames = append(possibleGames, game)
		}
	}
	return possibleGames
}

func (r *Round) possibleTeams(teams []*Team) []*Team {
	players := r.players()

	var possibleTeams []*Team
	for _, team := range teams {
		if !in(players, team.players()) {
			possibleTeams = append(possibleTeams, team)
		}
	}
	return possibleTeams
}

func (r *Round) Copy(other *Round) {
	r.numPlayers = other.numPlayers
	r.games = CopyGames(other.games)
}

func NewRounds(rounds []*Round) []*Round {
	var newRounds []*Round
	for _, r := range rounds {
		newR := &Round{numPlayers: numPlayers}
		for _, g := range r.games {
			newR.games = append(newR.games, &Game{g.teamA, g.teamB})
		}
		newRounds = append(newRounds, newR)
	}
	return newRounds
}

func (r *Round) full() bool {
	return r != nil && int(r.numPlayers/4) == len(r.games) && r.games[len(r.games)-1].full()
}

func (r *Round) empty() bool {
	return r == nil || len(r.games) == 0
}

func (r *Round) valid() bool {
	for i, Game := range r.games {
		for _, other := range r.games[i+1:] {
			if in(Game.players(), other.players()) {
				return false
			}
		}
	}

	return true
}

func (r *Round) players() []int {
	if r == nil {
		return nil
	}

	var players []int
	for _, Game := range r.games {
		players = append(players, Game.players()...)
	}
	return players
}

func (r *Round) hasPlayed(hasPlayed [][]int) [][]int {
	for _, Game := range r.games {
		hasPlayed = Game.hasPlayed(hasPlayed)
	}
	return hasPlayed
}

func (r *Round) reflectiveHasPlayed() bool {
	hasPlayed := r.hasPlayed(emptyHasPlayed(r.numPlayers))
	for x := 0; x < int(r.numPlayers); x++ {
		for y := 0; y < int(r.numPlayers); y++ {
			if hasPlayed[x][y] != hasPlayed[int(r.numPlayers)-1-y][int(r.numPlayers)-1-x] {
				return false
			}
		}
	}
	return true
}

func (r *Round) groups() []Group {
	var groups []Group
	for _, g := range r.games {
		groups = append(groups, g.groups()...)
	}
	return groups
}

func (r *Round) addTeam(t *Team) {
	gameIdx := len(r.games) - 1
	if gameIdx < 0 || r.games[gameIdx].full() {
		r.games = append(r.games, &Game{teamA: t})
		return
	}

	r.games[gameIdx].teamB = t
}

func (r *Round) popTeam() {
	gameIdx := len(r.games) - 1
	if gameIdx < 0 {
		return
	}
	g := r.games[gameIdx]

	g.popTeam()
	if g.empty() {
		r.games = r.games[:gameIdx]
	}
}

func (r *Round) String() string {
	var gameStrings []string
	for _, Game := range r.games {
		gameStrings = append(gameStrings, Game.String())
	}
	return strings.Join(gameStrings, "|")
}
