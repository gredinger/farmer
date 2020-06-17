package farmer

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Tile is a basic unit in the game. It can be modified.
type Tile struct {
	gorm.Model
	Wind       int8
	Solar      uint16
	X          uint64
	Y          uint64
	Z          int16
	Temp       int16
	Owner      uint
	SpriteIcon uint
	Moisture   float64
	Resources  []Resource
}

//Resource is an item within a tile.
type Resource struct {
	Name     string
	Quantity float64
}

//generateResources calculates resources of the tile
func (t *Tile) generateResources() {

	// Sand for above 1000 solar and lower than 1000 moisture

	// Silt for 500-1500 solar + 1000 moisture

	// Ice from solar under 500 at least 750 moisture
}

//GenerateNextTile creates a tile in the appointed direction
// Valid entries for direction are 1 through 9 excluding 5
func (t Tile) GenerateNextTile(direction int) Tile {
	genTile := t
	if direction < 3 {
		genTile.Y++
	}
	if direction > 5 {
		genTile.Y--
	}
	if direction%3 == 0 {
		genTile.X++
	}
	if direction%3 == 1 {
		genTile.X--
	}
	genTile.generateResources()
	return genTile
}

func (t Tile) String() string {
	s := fmt.Sprintf("X: %v Y: %v Z: %v SpriteIcon: %v Owner: %v", t.X,
		t.Y, t.Z, t.SpriteIcon, t.Owner)
	fmt.Println(s)
	return s
}
