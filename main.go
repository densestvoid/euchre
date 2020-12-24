package main

import "fmt"

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
	// fmt.Println("Players:", numPlayers)
	// teams := playerToTeamCombos(numPlayers)
	// fmt.Println("Teams:", len(teams))
	// games := teamsToGamesCombos(teams)
	// fmt.Println("Games:", len(games))
	// rounds := gamesToRoundsCombos(numPlayers, games, nil)
	// fmt.Println("Rounds:", len(rounds))
	// brackets := roundsToBracketsCombos(numPlayers, rounds, nil)
	// fmt.Println(brackets)

	teams := playerToTeamCombos(numPlayers)
	fmt.Println(fillBracket(teams, &Bracket{}))
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
