package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bits-and-blooms/bitset"
	"github.com/markome-beep/AOC-2025/shared"
)

func (p *point) String() string {
	return fmt.Sprintf("{x: %d y: %d from: %v to: %v}", p.x, p.y, p.from, p.to)
}

func (ps path) String() string {
	var maxX, maxY uint = 0, 0

	for _, p := range ps {
		maxX = max(p.x+1, maxX)
		maxY = max(p.y+1, maxY)
	}

	visualGrid := make([][]rune, maxY)
	for i := range maxY {
		visualGrid[i] = make([]rune, maxX)
		for j := range maxX {
			visualGrid[i][j] = '░'
		}
	}
	for i, p := range ps {
		char := '!'
		switch p.from {
		case [2]int{-1, 0}:
			switch p.to {
			case [2]int{-1, 0}:
			case [2]int{1, 0}:
				char = '═'
			case [2]int{0, 1}:
				char = '╗'
			case [2]int{0, -1}:
				char = '╝'
			}
		case [2]int{1, 0}:
			switch p.to {
			case [2]int{-1, 0}:
				char = '═'
			case [2]int{1, 0}:
			case [2]int{0, 1}:
				char = '╔'
			case [2]int{0, -1}:
				char = '╚'
			}
		case [2]int{0, 1}:
			switch p.to {
			case [2]int{-1, 0}:
				char = '╗'
			case [2]int{1, 0}:
				char = '╔'
			case [2]int{0, 1}:
				char = '║'
			case [2]int{0, -1}:
			}
		case [2]int{0, -1}:
			switch p.to {
			case [2]int{-1, 0}:
				char = '╝'
			case [2]int{1, 0}:
				char = '╚'
			case [2]int{0, 1}:
			case [2]int{0, -1}:
				char = '║'
			}
		}
		visualGrid[p.y][p.x] = char

		x, y := int(p.x)+p.to[0], int(p.y)+p.to[1]
		for x != int(ps[(i+1)%len(ps)].x) || y != int(ps[(i+1)%len(ps)].y) {
			var char rune
			switch p.to {
			case [2]int{-1, 0}:
				char = '─'
			case [2]int{1, 0}:
				char = '─'
			case [2]int{0, 1}:
				char = '│'
			case [2]int{0, -1}:
				char = '│'
			}
			visualGrid[y][x] = char

			x += p.to[0]
			y += p.to[1]
		}
	}

	str := ""
	for _, row := range visualGrid {
		str += "\n"
		str += string(row)
	}

	return str
}

func (g *path) scale() {
	var minX, minY uint = ^uint(0) >> 1, ^uint(0) >> 1

	for _, p := range *g {
		minX = min(p.x, minX)
		minY = min(p.y, minY)
	}

	for _, p := range *g {
		p.x -= minX - 1
		p.y -= minY - 1
	}
}

func parsePoints(file string) path {
	points := make([]*point, 0)
	prevPoint := &point{from: [2]int{0, 0}, to: [2]int{0, 0}, x: 0, y: 0}
	for line := range shared.ReadLines(file, "\n") {
		vals := strings.Split(line, ",")
		xi, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("RIP")
		}
		x := uint(xi)

		yi, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("RIP")
		}
		y := uint(yi)

		var fromDir [2]int
		if x == prevPoint.x {
			if y > prevPoint.y {
				fromDir = [2]int{0, -1}
				prevPoint.to = [2]int{0, 1}
			} else {
				fromDir = [2]int{0, 1}
				prevPoint.to = [2]int{0, -1}
			}
		} else {
			if x > prevPoint.x {
				fromDir = [2]int{-1, 0}
				prevPoint.to = [2]int{1, 0}
			} else {
				fromDir = [2]int{1, 0}
				prevPoint.to = [2]int{-1, 0}
			}
		}

		prevPoint = &point{from: fromDir, to: [2]int{0, 0}, x: x, y: y}
		points = append(points, prevPoint)
	}

	var fromDir [2]int
	x, y := points[0].x, points[0].y
	if x == prevPoint.x {
		if y > prevPoint.y {
			fromDir = [2]int{0, -1}
			prevPoint.to = [2]int{0, 1}
		} else {
			fromDir = [2]int{0, 1}
			prevPoint.to = [2]int{0, -1}
		}
	} else {
		if x > prevPoint.x {
			fromDir = [2]int{-1, 0}
			prevPoint.to = [2]int{1, 0}
		} else {
			fromDir = [2]int{1, 0}
			prevPoint.to = [2]int{-1, 0}
		}
	}
	points[0].from = fromDir

	return points
}

func newGrid(x uint, y uint) *grid {
	return &grid{bitset.MustNew((x + 1) * (y + 1)), x + 1, y + 1}
}

func (g *grid) flipEdge(ps *path) {
	for i, p := range *ps {
		g.data = g.data.Flip(p.y*g.width + p.x)
		x, y := int(p.x)+p.to[0], int(p.y)+p.to[1]
		for x != int((*ps)[(i+1)%len(*ps)].x) || y != int((*ps)[(i+1)%len(*ps)].y) {
			g.data.Flip(uint(y)*g.width + uint(x))

			x += p.to[0]
			y += p.to[1]
		}
	}
}

func (g *grid) fill(ps *path) {
	// g.data.SetAll()
	g.flipEdge(ps)
}

func (g *grid) contains(p1 *point, p2 *point) bool {
	top := max(p1.y, p2.y) - 1
	bot := min(p1.y, p2.y) + 1
	left := min(p1.x, p2.x) + 1
	right := max(p1.x, p2.x) - 1
	for x := left; x <= right; x++ {
		if g.data.Test(top*g.width+x) {
			return false
		}
		if g.data.Test(bot*g.width+x) { 
			return false
		}
	}


	for y := bot + 1; y <= top - 1; y++ { 
		if g.data.Test(y*g.width+left) {
			return false
		}
		if g.data.Test(y*g.width+right) { 
			return false
		}
	}

	return true
}
