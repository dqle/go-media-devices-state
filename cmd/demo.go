package main

import (
	"log"

	mediaDevices "github.com/dqle/go-media-devices-state-darwin"
)

func main() {
	for {
		isCameraOn, err := mediaDevices.IsCameraOnDarwin(false)
		if err != nil {
			log.Println("Error")
		} else {
			if isCameraOn == true {
				log.Println("Camera is on")
			} else {
				log.Println("Camera is off")
			}
		}
	}
}
