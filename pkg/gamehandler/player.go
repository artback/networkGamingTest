package gamehandler

type Player struct {
	Name    string
	Channel chan int
	Guess   [2]int
}
