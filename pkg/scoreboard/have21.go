package scoreboard

func (s *Scoreboard) Have21() *Result {
	for _, result := range *s {
		if result.Score == 21 {
			return result
		}
	}
	return nil
}
