package main

import "fmt"

type Game struct {
	teamA, teamB *Team
}

func CopyGames(games []*Game) []*Game {
	return append([]*Game{}, games...)
}

func (g *Game) full() bool {
	return g != nil && g.teamA != nil && g.teamB != nil
}

func (g *Game) empty() bool {
	return g == nil || g.teamA == nil && g.teamB == nil
}

func (g *Game) valid() bool {
	return g == nil || g.teamA.valid() && g.teamB.valid() && !in(g.teamA.players(), g.teamB.players())
}

func (g *Game) players() []int {
	if g == nil {
		return nil
	}
	return append(g.teamA.players(), g.teamB.players()...)
}

func (g *Game) possibleTeams(teams []*Team) []*Team {
	players := g.players()

	var possibleTeams []*Team
	for _, team := range teams {
		if !in(players, team.players()) {
			possibleTeams = append(possibleTeams, team)
		}
	}
	return possibleTeams
}

func (g *Game) hasPlayed(hasPlayed [][]int) [][]int {
	if !g.full() {
		return hasPlayed
	}

	hasPlayed[g.teamA.playerA-1][g.teamB.playerA-1] += 1
	hasPlayed[g.teamA.playerA-1][g.teamB.playerB-1] += 1
	hasPlayed[g.teamA.playerB-1][g.teamB.playerA-1] += 1
	hasPlayed[g.teamA.playerB-1][g.teamB.playerB-1] += 1
	hasPlayed[g.teamB.playerA-1][g.teamA.playerA-1] += 1
	hasPlayed[g.teamB.playerA-1][g.teamA.playerB-1] += 1
	hasPlayed[g.teamB.playerB-1][g.teamA.playerA-1] += 1
	hasPlayed[g.teamB.playerB-1][g.teamA.playerB-1] += 1

	return hasPlayed
}

func (g *Game) groups() []Group {
	players := g.players()
	if len(players) < 3 {
		return nil
	}

	var groups []Group
	for x := 0; x < len(players); x++ {
		for y := x + 1; y < len(players); y++ {
			for z := y + 1; z < len(players); z++ {
				groups = append(groups, Group{players[x], players[y], players[z]})
			}
		}
	}
	return groups
}

func (g *Game) popTeam() {
	if g.teamB == nil {
		g.teamA = nil
	} else {
		g.teamB = nil
	}
}

func (g *Game) String() string {
	return fmt.Sprintf("%sv%s", g.teamA, g.teamB)
}
