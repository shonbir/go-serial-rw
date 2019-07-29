package reader

import (
	"fmt"

	"github.com/shonbir/go-serial-rw/event"
	"github.com/tarm/serial"
)

// ReadAndPublishData connects to the serail port and start publishing data
func ReadAndPublishData(readEvent *event.SerialReadEvent, serialConfig *serial.Config) {
	go func(serialConfig *serial.Config) {
		serialConn, serialConnFail := serial.OpenPort(serialConfig)
		if serialConnFail != nil {
			fmt.Println("connection to serial port failed ", serialConfig.Name)
		}
		for {
			buf := make([]byte, 128)
			numRead, err := serialConn.Read(buf)
			if err != nil {
				continue
			}
			readEvent.PublishEventRecieved(buf[:numRead])
		}
	}(serialConfig)
}
