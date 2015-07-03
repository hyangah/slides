// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on golang.org/x/mobile/example/audio

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
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/audio"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
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

	eng   = glsprite.Engine()
	scene *sprite.Node

	player *audio.Player

	activate    = false
	acceptTouch = false
	touchLoc    geom.Point

	started = false
)

func touch(t event.Touch, c event.Config) {
	if t.Change != event.ChangeOn {
		return
	}

	touchLoc = t.Loc
	acceptTouch = true
	if !activate {
		activate = true
	}
}

func start() {
	rc, err := asset.Open("hello.wav")
	if err != nil {
		log.Fatal(err)
	}
	player, err = audio.NewPlayer(rc, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

}

func stop() {
	player.Close()
}

func draw(c event.Config) {
	if !started {
		touchLoc = geom.Point{c.Width / 2, c.Height / 2}
		started = true
	}
	if scene == nil {
		loadScene(c)
	}

	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(startClock) * 60 / time.Second)
	eng.Render(scene, now, c)
	debug.DrawFPS(c)
}

func newNode() *sprite.Node {
	n := &sprite.Node{}
	eng.Register(n)
	scene.AppendChild(n)
	return n
}

func loadScene(c event.Config) {
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
			dx = (float32(touchLoc.X) - x) / 60
			dy = (float32(touchLoc.Y) - y) / 60
			acceptTouch = false
			hello()
			x += dx
			y += dy
		} else if activate {
			if x < 0 {
				dx = 1
			}
			if y < 0 {
				dy = 1
			}
			if x+width > float32(c.Width) {
				dx = -1
			}
			if y+height > float32(c.Height) {
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
	player.Seek(0)
	player.Play()
}

func loadGopher() sprite.SubTex {
	a, err := asset.Open("gopher.jpeg")
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
