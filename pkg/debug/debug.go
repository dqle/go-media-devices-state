package debug

import (
	"fmt"

	"github.com/dqle/go-media-devices-state-darwin/pkg/camera"
	"github.com/dqle/go-media-devices-state-darwin/pkg/microphone"
)

func formatBool(value bool) string {
	if value {
		return "ON"
	}
	return "OFF"
}

// Debug calls all available device functions and prints the results
func Debug() {
	fmt.Println("Debug go-media-devices-state module...")
	fmt.Println()

	isCameraOn, err := camera.IsCameraOn()
	fmt.Println()
	if err != nil {
		fmt.Println("Is camera on: ERROR:", err)
	} else {
		fmt.Println("Camera state:", formatBool(isCameraOn))
	}

	fmt.Println()

	isMicrophoneOn, err := microphone.IsMicrophoneOn()
	fmt.Println()
	if err != nil {
		fmt.Println("isMicrophoneOn(): ERROR:", err)
	} else {
		fmt.Println("Microphone state:", formatBool(isMicrophoneOn))
	}
}
