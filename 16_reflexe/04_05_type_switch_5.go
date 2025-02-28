// Reflexe v programovacím jazyku Go
//
// - praktická ukázka rozvětvení a konverze na základě
//   požadovaného datového typu v aplikaci využívající
//   multimediální knihovnu SDL2

package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func eventLoop() {
	var event sdl.Event
	done := false

	for !done {
		event = sdl.PollEvent()
		// události mohou být různého typu,
		// takže se můžeme rozhodnout, o kterou událost
		// se jedná, následujícím rozvětvením
		switch t := event.(type) {
		case *sdl.QuitEvent:
			done = true
		case *sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			switch t.State {
			case sdl.PRESSED:
				switch keyCode {
				case sdl.K_ESCAPE:
					done = true
				case sdl.K_q:
					done = true
				}
			}
		}
		state.moveNPC()
		state.redraw()
		sdl.Delay(10)
	}
	log.Println("Quitting")
}
