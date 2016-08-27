package core

import (
	"math/rand"
	"time"
)

// instruction :
type instruction map[uint8]func(*System)

// instructions
var instructions = instruction{
	0x0: func(s *System) {
		switch s.cpu.Opcode.NN {
		case 0xE0:
			s.gfx.clearScreen()
			s.cpu.PC += 2
		case 0xEE:
			s.cpu.SP--
			s.cpu.PC = s.cpu.Stack[s.cpu.SP]
			s.cpu.PC += 2
		}
	},
	0x1: func(s *System) {
		s.cpu.PC = s.cpu.Opcode.NNN
	},
	0x2: func(s *System) {
		s.cpu.Stack[s.cpu.SP] = s.cpu.PC
		s.cpu.SP++
		s.cpu.PC = s.cpu.Opcode.NNN
	},
	0x3: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.X] == s.cpu.Opcode.NN {
			s.cpu.PC += 4
		} else {
			s.cpu.PC += 2
		}
	},
	0x4: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.X] != s.cpu.Opcode.NN {
			s.cpu.PC += 4
		} else {
			s.cpu.PC += 2
		}
	},
	0x5: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.X] == s.cpu.V[s.cpu.Opcode.Y] {
			s.cpu.PC += 4
		} else {
			s.cpu.PC += 2
		}
	},
	0x6: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] = s.cpu.Opcode.NN
		s.cpu.PC += 2
	},
	0x7: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] += s.cpu.Opcode.NN
		s.cpu.PC += 2
	},
	0x9: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.X] != s.cpu.V[s.cpu.Opcode.Y] {
			s.cpu.PC += 4
		} else {
			s.cpu.PC += 2
		}
	},
	0xA: func(s *System) {
		s.cpu.Index = s.cpu.Opcode.NNN
		s.cpu.PC += 2
	},
	0xB: func(s *System) {
		s.cpu.PC = uint16(s.cpu.V[0]) + s.cpu.Opcode.NNN
	},
	0xC: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] = uint8(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(255)) & s.cpu.Opcode.NN
		s.cpu.PC += 2
	},
	0xD: func(s *System) {
		s.cpu.V[0xF] = s.gfx.drawSprite(s.cpu.Memory[s.cpu.Index:s.cpu.Index+uint16(s.cpu.Opcode.N)],
			s.cpu.V[s.cpu.Opcode.X], s.cpu.V[s.cpu.Opcode.Y], s.cpu.Opcode.N)
		s.cpu.PC += 2
	},
	0xE: func(s *System) {
		switch s.cpu.Opcode.NN {
		case 0x9E:
			if s.input[s.cpu.V[s.cpu.Opcode.X]] {
				s.input[s.cpu.V[s.cpu.Opcode.X]] = false
				s.cpu.PC += 4
			} else {
				s.cpu.PC += 2
			}
		case 0xA1:
			if !s.input[s.cpu.V[s.cpu.Opcode.X]] {
				s.cpu.PC += 4
			} else {
				s.input[s.cpu.V[s.cpu.Opcode.X]] = false
				s.cpu.PC += 2
			}
		}
	},
}

var instructions8 = instruction{
	0x0: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] = s.cpu.V[s.cpu.Opcode.Y]
		s.cpu.PC += 2
	},
	0x1: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] |= s.cpu.V[s.cpu.Opcode.Y]
		s.cpu.PC += 2
	},
	0x2: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] &= s.cpu.V[s.cpu.Opcode.Y]
		s.cpu.PC += 2
	},
	0x3: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] ^= s.cpu.V[s.cpu.Opcode.Y]
		s.cpu.PC += 2
	},
	0x4: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.Y] > (0xFF - s.cpu.V[s.cpu.Opcode.X]) {
			s.cpu.V[0xF] = 1
		} else {
			s.cpu.V[0xF] = 0
		}
		s.cpu.V[s.cpu.Opcode.X] += s.cpu.V[s.cpu.Opcode.Y]
		s.cpu.PC += 2

	},
	0x5: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.Y] > s.cpu.V[s.cpu.Opcode.X] {
			s.cpu.V[0xF] = 0
		} else {
			s.cpu.V[0xF] = 1
		}
		s.cpu.V[s.cpu.Opcode.X] -= s.cpu.V[s.cpu.Opcode.Y]
		s.cpu.PC += 2
	},
	0x6: func(s *System) {
		s.cpu.V[0xF] = s.cpu.V[s.cpu.Opcode.X] & 1
		s.cpu.V[s.cpu.Opcode.X] /= 2
		s.cpu.PC += 2
	},
	0x7: func(s *System) {
		if s.cpu.V[s.cpu.Opcode.X] > s.cpu.V[s.cpu.Opcode.Y] {
			s.cpu.V[0xF] = 0
		} else {
			s.cpu.V[0xF] = 1
		}
		s.cpu.V[s.cpu.Opcode.X] = s.cpu.V[s.cpu.Opcode.Y] - s.cpu.V[s.cpu.Opcode.X]
		s.cpu.PC += 2
	},
	0xE: func(s *System) {
		s.cpu.V[0xF] = s.cpu.V[s.cpu.Opcode.X] >> 7
		s.cpu.V[s.cpu.Opcode.X] *= 2
		s.cpu.PC += 2
	},
}

var instructionsF = instruction{
	0x07: func(s *System) {
		s.cpu.V[s.cpu.Opcode.X] = s.cpu.DelayTimer
		s.cpu.PC += 2
	},
	0x0A: func(s *System) {
		var (
			i        uint8
			keyPress = false
		)

		for i = 0; i < 16; i++ {
			if s.input[i] {
				s.cpu.V[s.cpu.Opcode.X] = i
				s.input[i] = false
				keyPress = true
			}
		}

		if !keyPress {
			return
		}

		s.cpu.PC += 2
	},
	0x15: func(s *System) {
		s.cpu.DelayTimer = s.cpu.V[s.cpu.Opcode.X]
		s.cpu.PC += 2
	},
	0x18: func(s *System) {
		s.cpu.SoundTimer = s.cpu.V[s.cpu.Opcode.X]
		s.cpu.PC += 2
	},
	0x1E: func(s *System) {
		if (s.cpu.Index + uint16(s.cpu.V[s.cpu.Opcode.X])) > 0xFFF {
			s.cpu.V[0xF] = 1
		} else {
			s.cpu.V[0xF] = 0
		}
		s.cpu.Index += uint16(s.cpu.V[s.cpu.Opcode.X])
		s.cpu.PC += 2
	},
	0x29: func(s *System) {
		s.cpu.Index = uint16(s.cpu.V[s.cpu.Opcode.X]) * 0x5
		s.cpu.PC += 2
	},
	0x33: func(s *System) {
		s.cpu.Memory[s.cpu.Index] = s.cpu.V[s.cpu.Opcode.X] / 100
		s.cpu.Memory[s.cpu.Index+1] = (s.cpu.V[s.cpu.Opcode.X] / 10) % 10
		s.cpu.Memory[s.cpu.Index+2] = (s.cpu.V[s.cpu.Opcode.X] % 100) % 10
		s.cpu.PC += 2
	},
	0x55: func(s *System) {
		var i uint16
		for i = 0; i <= uint16(s.cpu.Opcode.X); i++ {
			s.cpu.Memory[s.cpu.Index+i] = s.cpu.V[i]
		}
		s.cpu.Index += uint16(s.cpu.Opcode.X) + 1
		s.cpu.PC += 2
	},
	0x65: func(s *System) {
		var i uint16
		for i = 0; i <= uint16(s.cpu.Opcode.X); i++ {
			s.cpu.V[i] = s.cpu.Memory[s.cpu.Index+i]
		}
		s.cpu.Index += uint16(s.cpu.Opcode.X) + 1
		s.cpu.PC += 2
	},
}
