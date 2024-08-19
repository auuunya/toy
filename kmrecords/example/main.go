package main

import (
	"fmt"
	"kmrecords/keyboard"
	"kmrecords/mouse"
	"time"
)

func main() {
	mouseChan := make(chan mouse.MouseEvent, 100)
	mouse.Use(mouseChan)
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				fmt.Println("Received timeout signal")
			case m := <-mouseChan:
				fmt.Printf("m: %#v\n", m)
				continue
			}
		}
	}()
	keyChan := make(chan keyboard.KeyBoardEvent, 100)
	keyboard.Use(keyChan)
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				fmt.Println("Received timeout signal")
			case m := <-keyChan:
				fmt.Printf("k: %#v\n", m)
				continue
			}
		}
	}()
	select {}
}
