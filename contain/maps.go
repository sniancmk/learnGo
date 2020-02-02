package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "tips",
		"course":  "golang",
		"quality": "notbad",
		"aim":     "changework",
	}

	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 ==nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
