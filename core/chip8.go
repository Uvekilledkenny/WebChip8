package core

import "time"

// System :
type System struct {
	Communication Communication
	ROM           []byte
	cpu           CPU
	gfx           gfx
	clock         *time.Ticker
	delay         *time.Ticker
	input         [16]bool
}

// New :
func New() System {
	return System{
		Communication: newChan(),
		cpu:           newCPU(),
		gfx:           newGFX(),
		clock:         time.NewTicker(time.Second / 500),
		delay:         time.NewTicker(time.Second / 60),
		input:         [16]bool{},
	}
}

// LoadROM :
func (s *System) LoadROM(r []byte) {
	s.cpu.copyROM(r)
	s.ROM = r
}

// Stop :
func (s *System) Stop() {
	<-s.Communication.shutdown
}

// PressKey :
func (s *System) PressKey(i int) {
	s.Communication.keypress <- i
}

// ChangeClock :
func (s *System) ChangeClock(c int) {
	s.clock = time.NewTicker(time.Second / time.Duration(c))
}

// Run :
func (s *System) Run() {
	if s.ROM == nil {
		s.Communication.Screen <- s.gfx.Screen
		return
	}

	for {
		select {
		case <-s.Communication.shutdown:
			return
		case i := <-s.Communication.keypress:
			s.input[i] = true
		case <-s.delay.C:
			s.timer()
			s.drawing()
			s.sound()
		case <-s.clock.C:
			s.cpu.fetchOpcode()

			switch s.cpu.Opcode.ID {
			case 0x8:
				instructions8[s.cpu.Opcode.N](s)
			case 0xF:
				instructionsF[s.cpu.Opcode.NN](s)
			default:
				instructions[s.cpu.Opcode.ID](s)
			}

			s.Communication.CPU <- s.cpu
		}
	}
}

func (s *System) timer() {
	if s.cpu.DelayTimer > 0 {
		s.cpu.DelayTimer--
	}
}

func (s *System) drawing() {
	if s.gfx.DrawFlag {
		s.Communication.Screen <- s.gfx.Screen
		s.gfx.DrawFlag = false
	}
}

func (s *System) sound() {
	if s.cpu.SoundTimer > 0 {
		if s.cpu.SoundTimer == 1 {
			s.Communication.Sound <- true
		}
		s.cpu.SoundTimer--
	}
}
