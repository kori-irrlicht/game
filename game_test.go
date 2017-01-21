package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	updateCalled int
	renderCalled int
	running      chan bool
}

func (g *game) Render() {
	g.renderCalled++
}

func (g *game) Update() {
	g.updateCalled++
}

func (g *game) Running() bool {
	return <-g.running
}

func Test(t *testing.T) {
	Convey("Given a new game", t, func() {
		g := &game{}
		g.running = make(chan bool, 10)

		Convey("game is running twice", func() {
			g.running <- true
			g.running <- true
			g.running <- false
			mainLoop(g)
			Convey("update is called twice", func() {
				So(g.updateCalled, ShouldEqual, 2)
			})

			Convey("render is called twice", func() {
				So(g.renderCalled, ShouldEqual, 2)
			})
		})

		Convey("game is running", func() {
			g.running <- true
			g.running <- false
			mainLoop(g)
			Convey("update is called", func() {
				So(g.updateCalled, ShouldEqual, 1)
			})

			Convey("render is called", func() {
				So(g.renderCalled, ShouldEqual, 1)
			})
		})

		Convey("game is not running", func() {
			g.running <- false
			mainLoop(g)
			Convey("update is not called", func() {
				So(g.updateCalled, ShouldEqual, 0)
			})

			Convey("render is not called", func() {
				So(g.renderCalled, ShouldEqual, 0)
			})

		})
	})

}
