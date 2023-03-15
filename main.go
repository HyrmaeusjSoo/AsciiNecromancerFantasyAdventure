package main

import "necromancer/game"

func main() {
	g := game.NewGame()
	g.Graph()
	defer func() {
		maybePanic := recover()
		g.Screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}()

	game.Input(g)
}
