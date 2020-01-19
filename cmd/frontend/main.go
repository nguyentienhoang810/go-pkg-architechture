package main

import(
	"fmt"
)

func main() {
	fmt.Println("start app ...")
	s := initServices()
	s.HTTP.ListenAndServe()
}