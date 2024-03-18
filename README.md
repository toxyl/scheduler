# Scheduler
This is a simple lib to schedule the execution of functions.

## Usage Example
```golang
package main

import "github.com/toxyl/scheduler"

func main() {
	stop := scheduler.Run(
		time.Second*5,
		time.Second*1,
		func() (stop bool) {
			fmt.Println(time.Now().String())
			sleep := time.Duration(rand.Int31n(5)) * time.Second
			time.Sleep(sleep)
			return rand.Int31n(9) == 3 // if random dice roll is 3, then we stop the schedule
		}, func() {
			fmt.Println("schedule stopped by cycle function")
			os.Exit(0)
		},
	)
	time.Sleep(30 * time.Second)
	fmt.Println("time exceeded, stopping...")
	stop()
	fmt.Println("waiting a while to be sure...")
	time.Sleep(10 * time.Second)
	fmt.Println("goodbye!")
}
```
Run multiple times to see the random dice rolls take effect.