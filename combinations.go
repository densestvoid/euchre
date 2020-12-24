package main

func playerToTeamCombos(numPlayers int) []*Team {
	var teams []*Team
	for i := 0; i < numPlayers; i++ {
		for j := i + 1; j < numPlayers; j++ {
			teams = append(teams, &Team{i, j})
		}
	}
	return teams
}

func teamsToGamesCombos(teams []*Team) []*Game {
	var games []*Game
	for i, teamA := range teams {
		for _, teamB := range teams[i+1:] {
			game := &Game{teamA, teamB}
			if game.valid() {
				games = append(games, &Game{teamA, teamB})
			}
		}
	}
	return games
}

func gamesToRoundsCombos(numPlayers int, games []*Game, baseRound *Round) []*Round {
	if baseRound == nil {
		baseRound = &Round{numPlayers: numPlayers}
	} else if baseRound.full() {
		if baseRound.reflectiveHasPlayed() {
			return []*Round{baseRound}
		}
	}

	var rounds []*Round
	for _, game := range games {
		newRound := &Round{}
		newRound.Copy(baseRound)
		newRound.games = append(newRound.games, game)

		if !newRound.valid() {
			continue
		}

		newGames := newRound.possibleGames(games)

		rounds = append(rounds, gamesToRoundsCombos(numPlayers, newGames, newRound)...)
	}
	return rounds
}

func roundsToBracketsCombos(numPlayers int, rounds []*Round, baseBracket *Bracket) []*Bracket {
	if baseBracket == nil {
		baseBracket = &Bracket{numPlayers: numPlayers}
	} else if baseBracket.full() {
		return []*Bracket{baseBracket}
	}

	var brackets []*Bracket
	for i, round := range rounds {
		newBracket := &Bracket{}
		newBracket.Copy(baseBracket)
		newBracket.rounds = append(newBracket.rounds, round)

		if !newBracket.valid() {
			continue
		}

		newRounds := NewRounds(rounds)
		newRounds = append(newRounds[:i], newRounds[i+1:]...)

		brackets = append(brackets, roundsToBracketsCombos(numPlayers, newRounds, newBracket)...)
	}
	return brackets
}
