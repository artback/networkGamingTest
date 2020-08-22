package scoring

func ScoreGuess(actual int, guess [2]int) int {
	var score int
	if guess[0] > guess[1] {
		guess0 := guess[0]
		guess[0] = guess[1]
		guess[1] = guess0
	}
	if actual == guess[0] || actual == guess[1] {
		score = count(guess, actual) * 5
	} else if actual >= guess[0] && actual <= guess[1] {
		score += 5 - (guess[1] - guess[0])
	} else {
		score--
	}
	return score
}

func count(arr [2]int, value int) int {
	var count = 0
	for _, v := range arr {
		if v == value {
			count++
		}
	}
	return count
}
