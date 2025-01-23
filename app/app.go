package app

import (
	"fmt"
	"sync"
)

type ConcurrentStruct struct {
	A User
	B User
}

type ASample struct {
	obj ConcurrentStruct
}

func (s *ASample) Process() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go s.populateA(&s.obj, &wg)

	wg.Add(1)
	go s.populateB(&s.obj, &wg)

	wg.Wait()

	fmt.Println(s.obj)
}

func (s *ASample) populateA(obj *ConcurrentStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("POPULATE A ")
	obj.A = User{
		Name: "A Name",
		Age:  uint64(19),
	}
}

func (s *ASample) populateB(obj *ConcurrentStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("POPULATE B ")
	obj.B = User{
		Name: "B Name",
		Age:  uint64(20),
	}
}
