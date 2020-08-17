package scoreboard

import (
	"github.com/artback/networkGamingTest/pkg/gamehandler"
	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/asc"
)

type Scoreboard struct {
	sortedMap *sortedmap.SortedMap
}

func NewScoreBoard(size int) *Scoreboard {
	sortedmap := sortedmap.New(size, asc.Int)
	return &Scoreboard{sortedmap}
}

//Adds or create a new score for player
func (s *Scoreboard) AddScore(score int, player gamehandler.Player) {
	s.addScore(score, player.Name)
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
func (s *Scoreboard) Get(player gamehandler.Player) (int, bool) {
	sI, _ := s.sortedMap.Get(player.Name)
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
