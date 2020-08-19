package scoreboard

import (
	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/asc"
)

type Scoreboard struct {
	sortedMap *sortedmap.SortedMap
}

func NewScoreBoard(size int) *Scoreboard {
	m := sortedmap.New(size, asc.Int)
	return &Scoreboard{m}
}

//Adds or create a new score for player
func (s *Scoreboard) AddScore(score int, player string) {
	s.addScore(score, player)
}
func (s *Scoreboard) addScore(score int, name string) {
	currI, has := s.sortedMap.Get(name)
	if has {
		curr, ok := currI.(int)
		if ok {
			s.sortedMap.Replace(name, score+curr)
		}
	} else {
		s.sortedMap.Replace(name, score)
	}
}
func (s *Scoreboard) Get(player string) (int, bool) {
	sI, _ := s.sortedMap.Get(player)
	score, ok := sI.(int)
	return score, ok
}
func (s *Scoreboard) Join(s2 *Scoreboard) {
	scoreMap := s2.sortedMap.Map()
	for na, sc := range scoreMap {
		score, okScore := sc.(int)
		name, okName := na.(string)
		if okScore && okName {
			s.addScore(score, name)
		}
	}
}
func (s *Scoreboard) Map() map[string]int {
	convertedMap := make(map[string]int)
	for key, value := range s.sortedMap.Map() {
		k, kOk := key.(string)
		v, vOk := value.(int)
		if kOk && vOk {
			convertedMap[k] = v
		}
	}
	return convertedMap
}
