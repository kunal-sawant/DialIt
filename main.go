package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/itchyny/volume-go"
	"github.com/jacobsa/go-serial/serial"
	"github.com/kunal-sawant/DialIt/systemcontrols"
	"github.com/kunal-sawant/DialIt/ui"
)

func main() {

	config := ui.GetSerialConfig()

	controller, err := systemcontrols.New()

	if err != nil {
		panic(err)
	}

	options := serial.OpenOptions{
		PortName:        config.ComPort,
		BaudRate:        config.BaudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		fmt.Printf("serial.Open: %v\n", err)
	}
	defer port.Close()

	brightnessGlobal := 0
	volumeGloabal, _ := volume.GetVolume()

	buf := make([]byte, 10)
	for {
		time.Sleep(10 * time.Millisecond)
		n, err := port.Read(buf)
		if err != nil {
			// fmt.Printf("port.Read: %v\n", err)
			panic(err)
		}
		serialData := string(buf[:n])
		data := strings.Split(serialData, "\r\n")[0]
		data = strings.TrimSpace(data)
		if len(data) > 0 {
			vals := getValues(data)
			fmt.Printf("serialVal: %v\n", vals)
			brightnessVal := vals[0]
			volumeVal := vals[1]

			if brightnessGlobal != brightnessVal {
				err := controller.SetBrightness(brightnessVal)
				if err != nil {
					log.Printf("Set brightness failed: %v\n", err)
				} else {
					log.Printf("Setting brightness to: %d\n", brightnessVal)
					brightnessGlobal = brightnessVal
				}
			}

			if volumeGloabal != volumeVal {
				err = controller.SetVolume(volumeVal)
				if err != nil {
					log.Printf("Set volume failed: %v\n", err)
				} else {
					log.Printf("Setting volume to: %d\n", volumeVal)
					volumeGloabal = volumeVal
				}
			}
		}

	}
}

func getValues(serialVal string) []int {
	valuesArr := strings.Split(serialVal, ",")
	var intSlice []int = make([]int, 2)
	var err error
	for i, v := range valuesArr {
		intSlice[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Error in parsing serial value: %s.\nError: %v", serialVal, err)
		}
	}
	return intSlice
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
