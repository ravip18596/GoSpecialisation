/*
Implement the dining philosopher’s problem with the following constraints/modifications.

    There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
    Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
    The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
    In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
    The host allows no more than 2 philosophers to eat concurrently.
    Each philosopher is numbered, 1 through 5.
    When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
    When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS,rightCS *Chopstick
	philosNo int
}

func (p Philosopher) Eat(){
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Println("starting to eat ",p.philosNo)
	fmt.Println("finishing eating ",p.philosNo)
	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

/*
In Dining philosopher problem, if each professor takes left chopstick first and then right chopstick then there might be a deadlock state
where each philosopher has left chopstick and every one is waiting for right chopstick.

To solve this, Dijkstra, proposed a solution where philosopher always have to pick lowest number chopstick first in every case. This will
prevent deadlock but philosopher 4 may starve.
 */

func min(a,b int) int{
	if a<b { return a }
	return b
}

func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}

func startHost(hostChan chan *Philosopher, quit chan struct{},finished chan struct{},wg *sync.WaitGroup) {
	fmt.Println("host starting")
	for {
		select {
		case p := <-hostChan:
			// running it as go routine, otherwise it will block
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				var wg2 sync.WaitGroup
				for i:=0;i<3;i++ {
					wg2.Add(1)
					go func(wg2 *sync.WaitGroup){
						defer wg2.Done()
						p.Eat()
					}(&wg2)
				}
				wg2.Wait()
			}(wg)
		case <-quit:
			fmt.Println("host shutting down")
			finished <- struct{}{}
			break
		}
	}
}

func main() {
	// initializing chopstick
	CSticks := make([]*Chopstick,5)
	for i:=0;i<5;i++{
		CSticks[i] = new(Chopstick)
	}
	// init philosopher
	philos := make([]*Philosopher,5)
	for i:=0;i<5;i++{
		philos[i] = &Philosopher{
			leftCS:  CSticks[min(i,(i+1)%5)],
			rightCS: CSticks[max(i,(i+1)%5)],
			philosNo: i+1,
		}
	}
	//making host
	host := make(chan *Philosopher,2)
	quit := make(chan struct{})
	finished := make(chan struct{})
	var wg sync.WaitGroup
	go startHost(host,quit,finished,&wg)
	for i:=0;i<5;i++{
		wg.Add(1)
		host <- philos[i]
	}
	wg.Wait()
	//close the host goroutine
	quit <- struct{}{}
	<-finished
}

