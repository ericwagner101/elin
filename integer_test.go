package main

import "fmt"

func main() {
	m := makeMap()
	fmt.Println("v1: ", getValue(m, "k2"))
	printValue(m, "k2")
}

func makeMap() map[string]int {
	m := make(map[string]int)
	m["k1"] = 8
	m["k2"] = 13
	m["k3"] = 22
	return m
}

// get value from map

func getValue(m map[string]int, key string) int {
	return m[key]
}

func printValue(m map[string]int, key string) {
	fmt.Println("v2: ", m[key])
}
