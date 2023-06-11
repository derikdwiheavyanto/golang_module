package latihan

import "fmt"

func Map() {
	var mapData = make(map[string]string)
	mapData["title"] = "Test Map"
	mapData["desc"] = "Holla"
	mapData["voice"] = "Moo"

	fmt.Println(mapData["title"])
	fmt.Println(mapData["desc"])
	fmt.Println(mapData["voice"])
}
