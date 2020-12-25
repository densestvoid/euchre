package main

import (
	"fmt"
	"time"
)

func in(playersA []int, playersB []int) bool {
	for _, playerA := range playersA {
		for _, playerB := range playersB {
			if playerA == playerB {
				return true
			}
		}
	}
	return false
}

func NewTeams(teams []*Team) []*Team {
	var newTeams []*Team
	for _, t := range teams {
		newTeams = append(newTeams, t)
	}
	return newTeams
}

func emptyHasPlayed(numPlayers int) [][]int {
	var hasPlayed = make([][]int, numPlayers)
	for i := 0; i < int(numPlayers); i++ {
		hasPlayed[i] = make([]int, numPlayers)
	}
	return hasPlayed
}

func fillBracket(teams []*Team, b *Bracket) *Bracket {
	var possibleTeams []*Team
	if b != nil {
		possibleTeams = b.possibleTeams(teams)
	} else {
		possibleTeams = teams
	}

	for i, team := range possibleTeams {
		newB := NewBracket(b)
		if !newB.addTeam(team) {
			continue
		}

		var newTeams []*Team
		for j := 0; i < len(teams); i++ {
			if teams[j].playerA == team.playerA && teams[j].playerB == team.playerB {
				newTeams := NewTeams(teams)
				newTeams = append(newTeams[:j], newTeams[j+1:]...)
			}
		}

		if len(newTeams) == 0 {
			return newB
		}
		if newB := fillBracket(newTeams, newB); newB.valid() {
			return newB
		}
	}
	return nil
}

const numPlayers = 4

func main() {
	teamGroupings := [][]*Team{
		{
			{1,2},
			{3,4},
			{5,6},
			{7,8},
			{9,10},
			{11,12},
			{13,14},
			{15,16},
		},
		{
			{1,3},
			{2,4},
			{5,7},
			{6,8},
			{9,11},
			{10,12},
			{13,15},
			{14,16},
		},
		{
			{1,4},
			{2,3},
			{5,8},
			{6,7},
			{9,12},
			{10,11},
			{13,16},
			{14,15},
		},
		{
			{1,5},
			{2,6},
			{3,7},
			{4,8},
			{9,13},
			{10,14},
			{11,15},
			{12,16},
		},
		{
			{1,6},
			{2,5},
			{3,8},
			{4,7},
			{9,14},
			{10,13},
			{11,16},
			{12,15},
		},
		{
			{1,7},
			{2,8},
			{3,5},
			{4,6},
			{9,15},
			{10,16},
			{11,13},
			{12,14},
		},
		{
			{1,8},
			{2,7},
			{3,6},
			{4,5},
			{9,16},
			{10,15},
			{11,14},
			{12,13},
		},
		{
			{1,9},
			{2,10},
			{3,11},
			{4,12},
			{5,13},
			{6,14},
			{7,15},
			{8,16},
		},
		{
			{1,10},
			{2,9},
			{3,12},
			{4,11},
			{5,14},
			{6,13},
			{7,16},
			{8,15},
		},
		{
			{1,11},
			{2,12},
			{3,9},
			{4,10},
			{5,15},
			{6,16},
			{7,13},
			{8,14},
		},
		{
			{1,12},
			{2,11},
			{3,10},
			{4,9},
			{5,16},
			{6,15},
			{7,14},
			{8,13},
		},
		{
			{1,13},
			{2,14},
			{3,15},
			{4,16},
			{5,9},
			{6,10},
			{7,11},
			{8,12},
		},
		{
			{1,14},
			{2,13},
			{3,16},
			{4,15},
			{5,10},
			{6,9},
			{7,12},
			{8,11},
		},
		{
			{1,15},
			{2,16},
			{3,13},
			{4,14},
			{5,11},
			{6,12},
			{7,9},
			{8,10},
		},
		{
			{1,16},
			{2,15},
			{3,14},
			{4,13},
			{5,12},
			{6,11},
			{7,10},
			{8,9},
		},
	}

	var roundCombosPerTeamGroupings [][]*Round
	for _, teamGrouping := range teamGroupings {
		roundCombosPerTeamGroupings = append(roundCombosPerTeamGroupings, teamsToRoundsCombos(teamGrouping, &Round{numPlayers:16}))
	}
	
	fmt.Println(bracketFromRounds(&Bracket{numPlayers:16}, roundCombosPerTeamGroupings))
}

func NewRoundCombos(combos [][]*Round) [][]*Round{
	return append([][]*Round{}, combos...)
}

func bracketFromRounds(bracket *Bracket, roundCombos [][]*Round) *Bracket{
	if len(roundCombos) == 0 {
		return bracket
	}

	rounds := roundCombos[0]
	newRoundCombos := NewRoundCombos(roundCombos[1:])

	start := time.Now()
	for _, round := range rounds {
		newBracket := NewBracket(bracket)
		newBracket.rounds = append(newBracket.rounds, round)
		if newBracket.valid() {
			b := bracketFromRounds(newBracket, newRoundCombos)
			if b != nil {
				return b
			}
		}
	}
	fmt.Println(len(bracket.rounds), time.Now().Sub(start))
	return nil
}

func roundsFromTeamPerms(perms [][]*Team) []*Round {
	var rounds []*Round
	for _, perm := range perms {
		var r = &Round{numPlayers: int(len(perm)) * 2}
		for _, t := range perm {
			r.addTeam(t)
		}
		rounds = append(rounds, r)
	}
	return rounds
}

func createPerms(baseTeams, pickTeams []*Team, count int) [][]*Team {
	if len(pickTeams) == 0 || count == 0 {
		return [][]*Team{baseTeams}
	}

	var perms [][]*Team
	for i, t := range pickTeams {
		newTs := append(baseTeams, t)
		newPicks := append([]*Team{}, pickTeams...)
		newPicks = append(newPicks[:i], newPicks[i+1:]...)
		perms = append(perms, createPerms(newTs, newPicks, count-1)...)
	}
	return perms
}

func createIntPerms(base, picks []int, count int) [][]int {
	if len(picks) == 0 || count == 0 {
		return [][]int{base}
	}

	var perms [][]int
	for i, v := range picks {
		newBase := append(base, v)
		newPicks := append([]int{}, picks...)
		newPicks = append(newPicks[:i], newPicks[i+1:]...)
		perms = append(perms, createIntPerms(newBase, newPicks, count-1)...)
	}
	return perms
}

func newTeams(teams []*Team) []*Team {
	var newTeams []*Team
	for _, t := range teams {
		newTeams = append(newTeams, t)
	}
	return newTeams
}
