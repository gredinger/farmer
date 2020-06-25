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
	Wind       uint8      // Wind class
	Solar      uint16     // Average watts / sq meter
	X          uint64     // X Coordinate - (Left/Right)
	Y          uint64     // Y Coordinate - (Up / Down)
	Z          float64    // Z Coordinate - Depth
	Temp       int16      // Average temp
	Owner      uint       // TODO: Implement land ownership
	SpriteIcon uint       // TODO: Implement rendering
	Moisture   float64    // Grams of water / sq meter (??)
	Resources  []Resource // Resources embedded in the chunk
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
		t.AddResource("Sand", float64(rand.Intn(int(t.Wind)*100)))
	}

	if t.Moisture > 750 && t.Temp <= 32 {
		t.AddResource("Ice", float64(rand.Intn(int(t.Moisture)*30)))
	}

	// Silt for 500-1500 solar + 1000 moisture
	if t.Moisture >= 1000 && t.Solar >= 500 && t.Solar <= 1500 {
		t.AddResource("Silt", float64(rand.Intn(1500)))
	}

	// Ice from solar under 500 at least 750 moisture
	if t.Moisture >= 1000 && t.Solar <= 500 {
		t.AddResource("Ice", float64(rand.Intn(int(t.Moisture))+100))
	}
}

//AddResource appends resources to a tile, combining with existing resource.
func (t *Tile) AddResource(name string, quantity float64) {
	for _, x := range t.Resources {
		if x.Name == name {
			x.Quantity += quantity
			return
		}
	}
	t.Resources = append(t.Resources, Resource{
		Name:     name,
		Quantity: quantity,
	})
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
