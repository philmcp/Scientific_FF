package models

type PlayerList []Player

// Pick a random fantasy football lineup for a sepcific position (e.g. Defender)
func (players PlayerList) RandomPosition(ignore map[string]bool, minValue float64, num int) PlayerList {
	out := PlayerList{}

	for {

		rand := players[random(0, len(players))]
		if minValue > rand.Selected {
			continue
		}

		fullName := rand.getFullName()
		if _, exist := ignore[fullName]; !exist {

			out = append(out, rand)
			ignore[fullName] = true
		}
		if len(out) >= num {
			break
		}
	}

	return out
}