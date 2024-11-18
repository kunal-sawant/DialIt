package main

import (
	"fmt"
	"lib/terminalview"
	"strconv"
	"strings"
	"time"

	"github.com/gek64/displayController"
	"github.com/itchyny/volume-go"
	"github.com/jacobsa/go-serial/serial"
)

func main() {
	config := terminalview.GetSerialConfig()

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

	compositeMonitors, err := displayController.GetCompositeMonitors()
	if err != nil {
		fmt.Println(err)
	}
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

			if brightnessGlobal != brightnessVal && absDiffInt(brightnessGlobal, brightnessVal) > 5 {
				for _, compositeMonitor := range compositeMonitors {
					// Set the brightness of the current display to current value
					err = displayController.SetVCPFeature(compositeMonitor.PhysicalInfo.Handle, displayController.Brightness, brightnessVal)
					fmt.Printf("Setting brightness to: %d\n", brightnessVal)
					if err != nil {
						fmt.Println(err)
					} else {
						brightnessGlobal = brightnessVal
					}
				}
			}

			if volumeGloabal != volumeVal && absDiffInt(volumeGloabal, volumeVal) > 5 {
				err = volume.SetVolume(volumeVal)
				fmt.Printf("Setting volume to: %d\n", volumeVal)
				if err != nil {
					fmt.Printf("set volume failed: %+v\n", err)
					continue
				} else {
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
