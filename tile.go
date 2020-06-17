package farmer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

//Tile is a basic unit in the game. It can be modified. One square meter of space.
type Tile struct {
	gorm.Model
	Wind       uint8  // Wind speed (class)
	Solar      uint16 // Watts
	X          uint64 // X
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

	rand.Seed(time.Now().UnixNano())

	// Sand for above 1000 solar and lower than 1000 moisture
	if t.Moisture <= 1000 && t.Solar >= 1000 {

		t.Resources = append(t.Resources, Resource{
			Name:     "Sand",
			Quantity: float64(rand.Intn(int(t.Wind) * 100)),
		})
	}

	// Silt for 500-1500 solar + 1000 moisture

	// Ice from solar under 500 at least 750 moisture
	if t.Moisture >= 1000 && t.Solar <= 500 {
		t.Resources = append(t.Resources, Resource{
			Name:     "Ice",
			Quantity: float64(rand.Intn(int(t.Moisture)) + 100),
		})
	}
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
