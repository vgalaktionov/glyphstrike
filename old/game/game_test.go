//go:build !js
// +build !js

package game_test

import (
	"testing"

	"github.com/vgalaktionov/roguelike-go/game"
)

func BenchmarkRandomGameWithRoomsAndCorridors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		game := game.NewGame(true)
		game.RunN(500)
	}
}
