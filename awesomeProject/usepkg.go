package main

import (
	bbbpkg "awesomeProject/aaa/bbbdir"
	"awesomeProject/aaa/cccdir"
	"awesomeProject/usepkg"
	"fmt"
	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	usepkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	bbbpkg.PrintCustom()
	cccdir.PrintCustom()

}
