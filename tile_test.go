package farmer_test

import (
	"fmt"
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
	tile := farmer.Tile{X: 1, Y: 1, Z: 0, Solar: 1200, Wind: 9, Temp: 130, Moisture: 10}
	genTile := tile.GenerateNextTile(3) //generate chunk to Northeast
	err := true
	for _, y := range genTile.Resources {
		if strings.Contains(y.Name, "Sand") {
			err = false
		}
		fmt.Printf("Checking %v\n", y.Name)
	}
	if err {
		t.Error("No sand in desert.")
	}
}
