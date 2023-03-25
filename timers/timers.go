package main

import "time"

const firstTimerSeconds = 2

func main() {
	timer1 := time.NewTimer(firstTimerSeconds * time.Second)

	<-timer1.C
	println("(First timer) Emit after", firstTimerSeconds, " seconds")

	timer2 := time.NewTimer(firstTimerSeconds * time.Second)

	go func() {
		<-timer2.C
		println("(Second timer) Emit after", firstTimerSeconds, " seconds")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		println("(Second timer) Stopped")
	}

	time.Sleep( 3 * time.Second)
}
