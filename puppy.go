package puppygo

import (
	"fmt"

	doggo "github.com/Andy025824/DogGo"
)

func Bark() string {
	return "Woof!====PuppyGo"
}

func Barks() string {
	s3 := doggo.Bark1()
	s4 := doggo.Bark2()
	fmt.Println(s3)
	fmt.Println(s4)
	return "Woof! Woof! Woof!----PuppyGo"
}
