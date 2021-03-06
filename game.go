package main

/*
Game contains the calls to the game logic
*/
type Game interface {
	Update()
	Render()
	Running() bool
}

func main() {

}

func mainLoop(g Game) {
	for g.Running() {
		g.Update()
		g.Render()
	}
}
