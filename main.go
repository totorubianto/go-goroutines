package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main(){

	showCurrentTime("Start")

	repeater(20, "repeater", 2, showMessageWithDelay)
	pauseExecution(3)

	showCurrentTime("End")
}

/*
	Show message in console after waiting a certain amount of time
 */
func showMessageWithDelay(msg string, seconds float64) {
	pauseExecution(seconds)
	fmt.Printf("%s - %s\n", time.Now(), msg)
}

/*
	Show current time in console
 */
func showCurrentTime(initmsg string)  {
	fmt.Printf("%s - %s\n", time.Now(), initmsg)
}

/*
	Pause execution for given time
 */
func pauseExecution(seconds float64){
	time.Sleep(time.Duration(seconds * math.Pow10(9)))
}

/*
	Repeat calling on GoRoutines
 */
func repeater(repeats int, msg string, millis float64, myFunc func(msg string, millis float64)){
	for i:=0 ; i<repeats ; i++ {
		go myFunc(msg + "_" + strconv.Itoa(i), millis)
	}
}