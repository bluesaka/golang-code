package main

import (
	"log"
	"go-code/study/options"
)

func main()  {
	s := options.NewStuffClientWithOptions(options.Connection{})
	log.Printf("%+v\n\n", s)

	s1 := options.NewStuffClientWithOptions(options.Connection{}, options.WithTimeout(77), options.WithRetries(66))
	log.Printf("%+v\n\n", s1)
}
