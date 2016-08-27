package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/uvekilledkenny/WebChip8/core"
)

type msg struct {
	Type int
}

type msgCPU struct {
	msg // 0
	core.CPU
}

type msgScreen struct {
	msg    // 1
	Screen string
}

type msgSound struct {
	msg   // 2
	Sound string
}

type msgKeys struct {
	msg // 3
	keyState
}

type keyState struct {
	Num   int
	State bool
}

type msgRom struct {
	msg // 4
	Rom string
}

type msgClock struct {
	msg   // 6
	Clock int
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

func encode(i image.Image) string {
	b := bytes.Buffer{}
	err := png.Encode(&b, i)
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func readLoop(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Paused")
			conn.Close()
			c.Stop()
			break
		}

		tmp := msg{}

		err = json.Unmarshal(p, &tmp)
		if err != nil {
			conn.Close()
			c.Stop()
			break
		}

		switch tmp.Type {
		case 3:
			k := msgKeys{}
			_ = json.Unmarshal(p, &k)
			c.PressKey(k.Num)

			log.Printf("Pressed Key %v", k.Num)
		case 5:
			r := msgRom{}
			_ = json.Unmarshal(p, &r)

			rom, err := base64.StdEncoding.DecodeString(r.Rom)
			if err != nil {
				log.Panicln(err)
				conn.Close()
				c.Stop()
				break
			}

			c = core.New()
			c.LoadROM(rom)

			log.Println("Rom Loaded")
		case 4:
			rom := c.ROM
			c = core.New()
			c.LoadROM(rom)

			log.Println("Reset")
		case 6:
			r := msgClock{}
			_ = json.Unmarshal(p, &r)

			c.ChangeClock(r.Clock)

			log.Printf("Changed Clock to %v Hz", r.Clock)
		}
	}
}

func writeLoop(conn *websocket.Conn) {
	for {
		select {
		case cpu := <-c.Communication.CPU:
			v := msgCPU{
				msg: msg{
					Type: 0,
				},
				CPU: cpu,
			}
			if err := conn.WriteJSON(v); err != nil {
				conn.Close()
				c.Stop()
				return
			}
		case gfx := <-c.Communication.Screen:
			v := msgScreen{
				msg: msg{
					Type: 1,
				},
				Screen: encode(gfx),
			}
			if err := conn.WriteJSON(v); err != nil {
				conn.Close()
				c.Stop()
				return
			}
		case <-c.Communication.Sound:
			v := msgSound{
				msg: msg{
					Type: 2,
				},
				Sound: "beep",
			}
			if err := conn.WriteJSON(v); err != nil {
				conn.Close()
				c.Stop()
				return
			}
		}
	}
}

func chipHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Play")
	go readLoop(conn)
	go writeLoop(conn)
	c.Run()
}
