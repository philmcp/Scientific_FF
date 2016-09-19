package main

// Takes a league, and throws out any players that aren't playing
func (dk *League) playing(l League) {

	newFinalPlayers := []Player{}
	for pos, players := range dk.FinalPlayers {
		dk.FinalPlayers[pos] = []Player{}

		for _, player := range players {
			player, err := l.getPlayer(player)

			if err != nil {

				newFinalPlayers = append(newFinalPlayers, player)
			}
		}

		dk.FinalPlayers[pos] = newFinalPlayers
	}

}
