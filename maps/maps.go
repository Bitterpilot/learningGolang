package main

import (
	"fmt"
)

func main()  {
	grades := make(map[string]float32)

	grades["timmy"] = 42
	grades["sam"] = 70
	grades["jess"] = 96

	fmt.Println(grades)
	fmt.Println(grades["timmy"])

	// timmy isn't doing well
	delete(grades,"timmy")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}