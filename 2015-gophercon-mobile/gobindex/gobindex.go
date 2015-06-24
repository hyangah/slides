package gobindex

import "fmt"

//START OMIT
type Gopher struct {
	name string
}

func (g Gopher) Name() string {
	return g.name
}

func NewGopher(name string) *Gopher {
	return &Gopher{name: "awesome " + name}
}

//END OMIT

func Hello(g *Gopher) {
	fmt.Printf("Hello, %s!\n", g.Name())
}
