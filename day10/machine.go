package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	LightTarget   uint
	JoltageTarget []uint
	Buttons       []uint
}

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

	return &Machine{lightTarget, joltageTarget, buttons}
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

func (m *Machine) allIndicators() map[uint][]uint {
	LightPaths := make(map[uint][]uint)
	LightPaths[0] = make([]uint, 0)

	for i := range uint(1 << len(m.Buttons)) {
		var pos uint = 0
		for p := 0; p < len(m.Buttons); p++ {
			if i&(1<<p) > 0 {
				pos = pos ^ m.Buttons[p]
			}
		}
		if LightPaths[pos] == nil {
			LightPaths[pos] = make([]uint, 0)
		}
		LightPaths[pos] = append(LightPaths[pos], i)
	}

	return LightPaths
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

func (m *Machine) Joltage() uint {
	all_i := m.allIndicators()
	var target func([]uint) uint

	step := func(u []uint, mov uint) ([]uint, error) {
		uc := make([]uint, len(u))
		copy(uc, u)

		for p := 0; p < len(m.Buttons); p++ {
			if mov&(1<<p) > 0 {
				for i := range u {
					if (1<<i)&m.Buttons[p] > 0 {
						if uc[i] == 0 {
							return nil, fmt.Errorf("Overflow")
						}
						uc[i] -= 1
					}
				}
			}
		}

		for i := range uc {
			uc[i] /= 2
		}
		return uc, nil
	}

	target = func(u []uint) uint {
		var pos uint = 0
		all_zeros := true
		for i, val := range u {
			if val%2 == 1 {
				pos |= 1 << i
			}
			if val != 0 {
				all_zeros = false
			}
		}

		if all_zeros {
			return 0
		}

		if combos, ok := all_i[pos]; ok {
			var m uint = math.MaxUint - 1000
			for _, c := range combos {
				s, err := step(u, c)
				if err != nil {
					continue
				}

				var btn_presses uint = 0
				for c > 0 {
					c = c & (c - 1)
					btn_presses++
				}

				m = min(m, target(s)*2+btn_presses)
			}

			var mu uint = 0
			for _, uv := range u {
				mu = max(uv, mu)	
			}
			return m
		}

		return math.MaxUint - 1000
	}

	val := target(m.JoltageTarget)
	return val
}
