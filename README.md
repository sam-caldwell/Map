Map
===

## Description

A set of simple tools to work with maps.  This includes data sanitization features.

## Usage

```text
import "github.com/sam-caldwell/map"

func main() {
	// Create two maps using New function
	map1 := New[string, int]()
	(*map1)["one"] = 1
	(*map1)["two"] = 2

	map2 := New[string, int]()
	(*map2)["one"] = 1
	(*map2)["two"] = 3

    // Copy the map
	map3 := New[string, int]()
	map2.Copy(map3)

	// Compare the maps
	fmt.Println("map1 == map2:", map1.Equal(map2)) // true
	fmt.Println("map1 == map3:", map1.Equal(map3)) // false
	
	//Purge the map
	map1.Purge()
	map2.Purge()
	map3.Purge()

	// Create a nil map
	var map4 *Map[string, int]
}

```