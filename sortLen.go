//Given a list of names, you need to organize each name within a slice based on its length.

package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	max := len(names[0])
	for i := 1; i < len(names); i++ {
		if len(names[i]) > max {
			max = len(names[i])
		}
	} 
	
	list := make([][]string, max)
	
	for _, v := range names {
		h := len(v)
		list[h-1] = append(list[h-1], v)
	}
	
	fmt.Println(list)
	
}
