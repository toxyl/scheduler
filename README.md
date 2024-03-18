# Scheduler
This is a simple lib to schedule the execution of functions.

## Usage Example
This will run a schedule every 5 seconds with 1 second offset, ie. at 00:00:01, 00:00:06, 00:00:11, 00:00:16, and so on. The cycle function prints the current time, waits a random amount of seconds (max 5) and then rolls dice to decide whether to stop the schedule.  
  
After 30 seconds the schedule will be stopped either way and 10 seconds later the program exits.

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