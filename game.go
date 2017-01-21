package main

/*
Game contains the calls to the game logic
*/
type Game interface {
	Update()
	Render()
}

func main() {

}

func mainLoop(g Game) {
	g.Update()
	g.Render()
}
