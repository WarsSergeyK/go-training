package main

import (
	"collection/domain"
	"fmt"
)

func main() {

	collection := domain.Collection{}
	for i := 0; i < 10; i++ {
		collection.Add(i)
	}

	fmt.Println(collection.Get(-999))
	fmt.Println(collection.Get(999))

	collection.Print()
	collection.Remove(3)
	collection.Print()
	collection.Remove(999)
	collection.Print()

	fmt.Println(collection.Get(7))

	fmt.Println(collection.First())
	fmt.Println(collection.Last())

	fmt.Println(collection.Length())
}
