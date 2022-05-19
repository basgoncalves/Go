package main

import "fmt"

func main() {
	// 3 different options
	// var color map[string]string
	//colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
	}
	// delete an element in the map
	delete(colors, "red")

	// add an element to the map
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
