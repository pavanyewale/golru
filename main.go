package main

import (
	"fmt"

	"github.com/pavanyewale/LRU-in-golang/lru"
)

func main() {
	lru, _ := lru.NewLRU(5)
	fmt.Println("initiated lru with size 5")
	fmt.Println("Enter choice 1)set value 2)get value")
	var choice, key, value int
	for {
		fmt.Print("enter choice =>")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("Enter key value integer")
			fmt.Scanf("%d%d", &key, &value)
			lru.SetItem(key, value)
		case 2:
			fmt.Println("Enter key")
			fmt.Scanf("%d", &key)
			fmt.Println(lru.GetItem(key))
		}
	}
}
