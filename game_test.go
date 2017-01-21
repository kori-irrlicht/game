package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	updateCalled bool
	renderCalled bool
}

func (g *game) Render() {
	g.renderCalled = true
}

func (g *game) Update() {
	g.updateCalled = true
}

func Test(t *testing.T) {
	Convey("Given a new game", t, func() {
		g := &game{}
		mainLoop(g)

		Convey("update is called", func() {
			So(g.updateCalled, ShouldBeTrue)
		})

		Convey("render is called", func() {
			So(g.renderCalled, ShouldBeTrue)
		})
	})

}
