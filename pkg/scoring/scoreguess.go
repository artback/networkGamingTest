package scoring

func ScoreGuess(actual int, guess [2]int) int {
	var score int
	if guess[0] > guess[1] {
		guess0 := guess[0]
		guess[0] = guess[1]
		guess[1] = guess0
	}
	if actual == guess[0] {
		score += 5
	} else if actual == guess[1] {
		score += 5
	} else if actual >= guess[0] && actual <= guess[1] {
		score += 5 - (guess[1] - guess[0])
	} else {
		score--
	}
	return score
}
