package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar += sendsToAdd
	return sendsSoFar
}
