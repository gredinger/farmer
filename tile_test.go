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

type TileTests struct {
	Tile             farmer.Tile
	ExpectedResource string
}

func TestTileResourceGenerator(t *testing.T) {
	tt := make([]TileTests, 0)
	tt = append(tt, TileTests{
		farmer.Tile{X: 1, Y: 1, Z: 0, Solar: 1200, Wind: 9, Temp: 130, Moisture: 10},
		"Sand",
	})
	tt = append(tt, TileTests{
		farmer.Tile{Y: 1, X: 1, Z: 10, Solar: 450, Wind: 1, Temp: 25, Moisture: 900},
		"Ice",
	})

	for _, x := range tt {
		genTile := x.Tile.GenerateNextTile(3) //generate chunk to Northeast
		err := true
		for _, y := range genTile.Resources {
			if strings.Contains(y.Name, x.ExpectedResource) {
				err = false
			}
		}
		if err {
			t.Error(fmt.Sprintf("Generated tile lacks %s", x.ExpectedResource))
		}
	}
}
