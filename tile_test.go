package farmer_test

import (
	"strings"
	"testing"

	"github.com/gredinger/farmer"
)

func TestTileString(t *testing.T) {
	tile := farmer.Tile{X: 0, Y: 0, Z: 0}
	if !strings.Contains(tile.String(), "X: 0 Y: 0 Z: 0") {
		t.Error("no match")
	}
}

func TestTileGenerationDirection(t *testing.T) {
	tile := farmer.Tile{X: 1, Y: 1, Z: 0}
	if !strings.Contains(tile.GenerateNextTile(1).String(), "X: 0 Y: 2 Z: 0") {
		t.Error("Northwest tile failure")
	}

}

func TestTileResourceGenerator(t *testing.T) {
	tile := farmer.Tile{X: 1, Y: 1, Z: 0, Solar: 1200, Wind: 90, Temp: 130}
	tile.GenerateNextTile(3)
	err := true
	for _, y := range tile.Resources {
		if strings.Contains(y.Name, "sand") {
			err = false
		}
	}
	if err {
		t.Error("No sand in desert.")
	}
}
