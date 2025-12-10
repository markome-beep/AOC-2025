package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	LightTarget uint
	// JoltageTargetKey string
	JoltageTarget []uint
	Buttons       []uint
	// LightPath        map[uint]uint
	// JoltagePath      map[string]JoltageNode
}

// type JoltageNode struct {
// 	dist uint
// 	val []uint
// }

func NewMachine(line string) *Machine {
	parts := strings.Split(line, " ")
	var lightTarget uint = 0
	buttons := make([]uint, 0)
	joltageTarget := make([]uint, 0)
	for _, p := range parts {
		switch p[0] {
		case '[':
			for i, char := range p[1 : len(p)-1] {
				if char == '#' {
					lightTarget |= 1 << i
				}
			}
		case '(':
			var btn uint = 0
			for val := range strings.SplitSeq(p[1:len(p)-1], ",") {
				num, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("RIP")
				}
				btn += 1 << num
			}
			buttons = append(buttons, btn)
		case '{':
			for val := range strings.SplitSeq(p[1:len(p)-1], ",") {
				num, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("RIP")
				}
				joltageTarget = append(joltageTarget, uint(num))
			}
		}
	}

	// lightPath := make(map[uint]uint)
	// lightPath[0] = 0

	// joltagePath := make(map[string]JoltageNode)
	// start := make([]uint, len(joltageTarget))
	// key := fmt.Sprintf("%v", start)
	// joltagePath[key] = JoltageNode{0, start}

	return &Machine{lightTarget, joltageTarget, buttons}
	// return &Machine{lightTarget, fmt.Sprintf("%v", joltageTarget), joltageTarget, buttons, lightPath, joltagePath}
}

func (m *Machine) Indicator() uint {

	LightPath := make(map[uint]uint)
	LightPath[0] = 0

	for i := range len(m.Buttons) {
		for pos, dist := range LightPath {
			if dist != uint(i) {
				continue
			}
			for _, btn := range m.Buttons {
				move := pos ^ btn
				if move == m.LightTarget {
					return dist + 1
				}
				_, ok := LightPath[move]
				if !ok {
					LightPath[move] = dist + 1
				}
			}
		}
	}
	return 0
}

func (m *Machine) Joltage_BFS() uint {

	start := make([]uint, len(m.JoltageTarget))
	key := fmt.Sprintf("%v", start)

	JP_1 := make(map[string][]uint)
	JP_1[key] = start
	JP_2 := make(map[string][]uint)

	JoltagePaths := [2]map[string][]uint{JP_1, JP_2}

	JoltageTargetKey := fmt.Sprintf("%v", m.JoltageTarget)
	var mx uint = 0
	for _, v := range m.JoltageTarget {
		mx += v
	}

	for i := range mx {
		JoltagePathPrev := JoltagePaths[i%2]
		JoltagePathNext := JoltagePaths[(i+1)%2]

		for k := range JoltagePathNext {
			delete(JoltagePathNext, k)
		}

		for _, pos := range JoltagePathPrev {
		Btn:
			for _, btn := range m.Buttons {
				move := make([]uint, len(m.JoltageTarget))

				for i := range move {
					move[i] = (btn>>i)&1 + pos[i]
					if move[i] > m.JoltageTarget[i] {
						continue Btn
					}
				}
				moveKey := fmt.Sprintf("%v", move)

				if moveKey == JoltageTargetKey {
					return i + 1
				}
				JoltagePathNext[moveKey] = move
			}
		}
	}

	return 0
}

func (m *Machine) Joltage_MAT() uint {
	return 0
}

func debugTravel(t map[uint]uint) {
	var str strings.Builder
	str.WriteString("map[")
	for pos, dist := range t {
		str.WriteString(littleEndian(uint64(pos)))
		str.WriteRune(':')
		str.WriteString(fmt.Sprint(dist))
		str.WriteRune(' ')
	}
	str.WriteRune(']')
	fmt.Println(str.String())
}

func littleEndian(n uint64) string {
	var sb strings.Builder

	for i := 0; n>>i > 0; i++ {
		if (n>>i)&1 == 1 {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

func (m *Machine) String() string {
	// var str strings.Builder
	// str.WriteRune('{')
	// str.WriteString(littleEndian(uint64(m.LightTarget)))
	// str.WriteString(" [")
	// for i, btn := range m.Buttons {
	// 	if i != 0 {
	// 		str.WriteRune(' ')
	// 	}
	// 	str.WriteString(littleEndian(uint64(btn)))
	// }
	// str.WriteString("]}")
	// return str.String()
	return fmt.Sprintf("%v", m.JoltageTarget)
}
