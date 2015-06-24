package main

import (
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/gl"
)

//START OMIT
import "golang.org/x/mobile/app"

func main() {
	app.Run(app.Callbacks{
		Start: start,
		Stop:  stop,
		Draw:  draw,
		Touch: touch,
	})
}

//END OMIT

func draw() {
	gl.ClearColor(1, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func start() {
}

func stop() {
}

func touch(event.Touch) {
}
