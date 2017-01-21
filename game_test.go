package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	updateCalled bool
	renderCalled bool
	running      chan bool
}

func (g *game) Render() {
	g.renderCalled = true
}

func (g *game) Update() {
	g.updateCalled = true
}

func (g *game) Running() bool {
	return <-g.running
}

func Test(t *testing.T) {
	Convey("Given a new game", t, func() {
		g := &game{}
		g.running = make(chan bool, 10)

		Convey("game is running", func() {
			g.running <- true
			mainLoop(g)
			Convey("update is called", func() {
				So(g.updateCalled, ShouldBeTrue)
			})

			Convey("render is called", func() {
				So(g.renderCalled, ShouldBeTrue)
			})
		})

		Convey("game is not running", func() {
			g.running <- false
			mainLoop(g)
			Convey("update is not called", func() {
				So(g.updateCalled, ShouldBeFalse)
			})

			Convey("render is not called", func() {
				So(g.renderCalled, ShouldBeFalse)
			})

		})
	})

}
