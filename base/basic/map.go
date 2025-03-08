package basic

import "fmt"

func TestMap() {
	m := map[string]string{ //无序
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbaf",
	}
	m2 := make(map[string]int) // m2 == empyt map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	causeName, ok := m["cause"]
	fmt.Println(causeName, ok)

	fmt.Println("Deleting values")
	delete(m, "name")
	name, ok := m["name"]
	fmt.Println(name, ok)
}
