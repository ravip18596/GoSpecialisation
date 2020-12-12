/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
*/
package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	x = x + 1 //race condition occurs because of concurrent access of two go-routines (7,8) at same time.
	//x = x+1 is not atomic instructions at machine code level. It comprises of LOAD X; INCR X; SAVE X.
	// interleaving of these 3 instructions from two different go-routines at run-time is non-deterministic.
	//final result vary acc. to interleaving, so it is not deterministic. It can output 2 or 3
}

func main() {
	//global var initialised as one
	x = 1
	var wg sync.WaitGroup
	wg.Add(1)
	go increment(&wg)
	wg.Add(1)
	go increment(&wg)
	wg.Wait()
	fmt.Println(x)
}

/*
Above program has race condition
WARNING: DATA RACE
Read at 0x000001274d08 by goroutine 8:
  main.increment()
      /Users/raviprakash/go/src/GoSpecialisation/ConcurrencyInGo/Week2/race.go:15 +0x6e

Previous write at 0x000001274d08 by goroutine 7:
  main.increment()
      /Users/raviprakash/go/src/GoSpecialisation/ConcurrencyInGo/Week2/race.go:15 +0x8a

Goroutine 8 (running) created at:
  main.main()
      /Users/raviprakash/go/src/GoSpecialisation/ConcurrencyInGo/Week2/race.go:26 +0xf6

Goroutine 7 (finished) created at:
  main.main()
      /Users/raviprakash/go/src/GoSpecialisation/ConcurrencyInGo/Week2/race.go:24 +0xbd
==================
3
Found 1 data race(s)
exit status 66
*/
