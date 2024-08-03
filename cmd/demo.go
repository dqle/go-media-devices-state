package main

import (
	"log"

	mediaDevices "github.com/dqle/go-media-devices-state"
)

func main() {
	for {
		isCameraOn, err := mediaDevices.IsCameraOn()
		isMicrophoneOn, err := mediaDevices.IsMicrophoneOn()
		if err != nil {
			log.Println(err)
		} else {
			if isCameraOn {
				log.Println("Camera is on")
			} else {
				log.Println("Camera is off")
			}
			if isMicrophoneOn {
				log.Println("Microphone is on")
			} else {
				log.Println("Microphone is off")
			}
		}
	}
}
