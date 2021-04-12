package main

import (
	"fmt"
	"reflect"
)

type ErrWrath struct {
}

func (e ErrWrath) Error() string {
	return "Err generated for an error wrath"
}
func main() {
	var err error
	var man int
	man, err = angelFunc()
	fmt.Printf("Err type: %s\n", reflect.TypeOf(err))
	fmt.Printf("Err value: %s\n", reflect.ValueOf(err))
	fmt.Printf("Err type: %s\n", reflect.TypeOf(nil))
	fmt.Printf("Err value: %s\n", reflect.ValueOf(nil))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got %d\n", man)
}

func angelFunc() (int, error) {
	return 9001, nil
}
