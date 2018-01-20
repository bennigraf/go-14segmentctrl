package main

import "fmt"
import "github.com/joeblubaugh/adafruit-led/HT16K33"
import "time"
import "net"
import "os"

var device HT16K33.Device

type MainState struct {
	introMode bool
}

var mainState MainState

func main() {
	mainState.introMode = true

	fmt.Println("Running!")
	// device = new(HT16K33.Device)
	device.Init(0x70, 1)

	showIntro()

	fmt.Println("GOing on...")

	// tcp server code taken from https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
	listener, _ := net.Listen("tcp", ":3024")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func showIntro() {
	writeToDisplay("Ok, Go!")
	time.Sleep(time.Second * 2)

	writeToDisplay("")
	time.Sleep(time.Millisecond * 200)

	for mainState.introMode {
		writeToDisplay("Great")
		time.Sleep(time.Second)

		writeToDisplay("")
		time.Sleep(time.Millisecond * 200)

		writeToDisplay("Atlantic")
		time.Sleep(time.Second)

		writeToDisplay("")
		time.Sleep(time.Millisecond * 200)

		// remove me once you figured out how to do this in a goroutine
		mainState.introMode = false
	}

	writeToDisplay(".")
}

// tcp server code taken from https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// write bytes as string to display
	writeToDisplay(string(buf[0 : reqLen-1]))
	// Send a response back to person contacting us.
	// conn.Write([]byte("Message received.\n"))
	// Close the connection when you're done with it.
	conn.Close()
}

func writeToDisplay(displayString string) {
	if len(displayString) > 8 {
		fmt.Println("Can only print 8 chars to display!")
	}

	for digit := 0; digit < 8; digit++ {
		if digit < len(displayString) {
			charAtDigit, charDoesExist := chars[displayString[digit]]
			if charDoesExist {
				device.SetBufferRow(byte(digit), charAtDigit)
			}
		} else {
			device.SetBufferRow(byte(digit), chars[' '])
		}
	}
}

func stepper() {
	displayString := "Hello, World! ... This is CRAZY! The quick brown fox jumps over an angry lizzard or something. "
	for stepper := 0; stepper >= 0; stepper++ {
		for digit := 0; digit < 8; digit++ {
			charAtDigit := chars[displayString[(digit+stepper)%len(displayString)]]
			device.SetBufferRow(byte(digit), charAtDigit)
		}

		time.Sleep(time.Millisecond * 150)
	}
}
