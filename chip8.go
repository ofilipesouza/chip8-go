package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type CPU struct {
}

type Chip8 struct {
	Screen tcell.Screen
	Memory [4096]uint
	CPU    *CPU
}

func NewChip8() *Chip8 {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("Azedou com a tela")
	}
	s.Init()
	c := &Chip8{Screen: s}
	return c
}
