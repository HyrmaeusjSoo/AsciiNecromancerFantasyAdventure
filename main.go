package main

import (
	"necromancer/game"
)

func main() {
	g := game.NewGame()
	defer func() {
		g.Screen.Fini()
		if maybe := recover(); maybe != nil {
			panic(maybe)
		}
	}()

	g.Start()
}
