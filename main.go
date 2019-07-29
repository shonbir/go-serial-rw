package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/shonbir/go-serial-rw/event"
	"github.com/shonbir/go-serial-rw/reader"

	"github.com/tarm/serial"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	ctrlCSignal := make(chan os.Signal, 1)
	signal.Notify(ctrlCSignal, os.Interrupt)
	readTimeOut, _ := time.ParseDuration("2000ms")
	config := &serial.Config{
		Name:        "COM1",
		Baud:        19200,
		StopBits:    serial.Stop1,
		ReadTimeout: readTimeOut,
		Parity:      serial.ParityEven,
		Size:        serial.DefaultSize,
	}

	fmt.Println("Press CTRL+C to end program")

	go func() {
		<-ctrlCSignal
		defer wg.Done()
	}()

	go func() {
		readChan := make(chan []byte)
		readEvent := event.NewEvents()
		readEvent.AddListener(readChan)
		reader.ReadAndPublishData(readEvent, config)
		for {
			select {
			case msg := <-readChan:
				fmt.Println(string(msg))

			}
		}
	}()

	wg.Wait()
}
