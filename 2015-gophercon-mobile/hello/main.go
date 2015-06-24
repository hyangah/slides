// +build omit

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"log"
	"time"

	_ "image/jpeg"
)

//START OMIT
import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/exp/audio"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)
//END OMIT

func main() {
	app.Run(app.Callbacks{
		Start: start,
		Stop:  stop,
		Draw:  draw,
		Touch: touch,
	})
}

const (
	width  = 72
	height = 60
)

var (
	startClock = time.Now()
	lastClock  = clock.Time(-1)

	eng   = glsprite.Engine()
	scene *sprite.Node

	player *audio.Player

	activate    = false
	acceptTouch = false
	touchLoc    geom.Point
)

func touch(t event.Touch) {
	if t.Type != event.TouchStart {
		return
	}

	touchLoc = t.Loc
	acceptTouch = true
	if !activate {
		activate = true
	}
}

func start() {
	rc, err := app.Open("hello.wav")
	if err != nil {
		log.Fatal(err)
	}
	player, err = audio.NewPlayer(rc, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	touchLoc = geom.Point{geom.Width / 2, geom.Height / 2}
}

func stop() {
	player.Close()
}

func draw() {
	if scene == nil {
		loadScene()
	}

	now := clock.Time(time.Since(startClock) * 60 / time.Second)
	/*if now == lastClock {
		return
	}*/
	lastClock = now

	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	eng.Render(scene, now)
	debug.DrawFPS()
}

func newNode() *sprite.Node {
	n := &sprite.Node{}
	eng.Register(n)
	scene.AppendChild(n)
	return n
}

func loadScene() {
	gopher := loadGopher()
	scene = &sprite.Node{}
	eng.Register(scene)
	eng.SetTransform(scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})

	var x, y float32
	dx, dy := float32(1), float32(1)

	n := newNode()
	n.Arranger = arrangerFunc(func(eng sprite.Engine, n *sprite.Node, t clock.Time) {
		eng.SetSubTex(n, gopher)

		if acceptTouch {
			x = float32(touchLoc.X)
			y = float32(touchLoc.Y)
			acceptTouch = false
			hello()
		} else if activate {
			if x < 0 {
				dx = 1
			}
			if y < 0 {
				dy = 1
			}
			if x+width > float32(geom.Width) {
				dx = -1
			}
			if y+height > float32(geom.Height) {
				dy = -1
			}

			x += dx
			y += dy
		}

		eng.SetTransform(n, f32.Affine{
			{width, 0, x},
			{0, height, y},
		})
	})
}

func hello() {
	// TODO(jbd): the sound explodes at seek, volume down and seek?
	player.Seek(0)
	player.Play()
}

func loadGopher() sprite.SubTex {
	a, err := app.Open("gopher.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}
	t, err := eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}
	return sprite.SubTex{t, image.Rect(0, 0, 360, 300)}
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }
