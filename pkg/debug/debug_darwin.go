package debug

import (
	"fmt"

	"github.com/dqle/go-media-devices-state/pkg/camera"
	"github.com/dqle/go-media-devices-state/pkg/microphone"
)

func formatBool(value bool) string {
	if value {
		return "ON"
	}
	return "OFF"
}

// Debug calls all available device functions and prints the results
func DebugDarwin() {
	fmt.Println("Debug go-media-devices-state module...")
	fmt.Println()

	isCameraOn, err := camera.IsCameraOnDarwin()
	fmt.Println()
	if err != nil {
		fmt.Println("Is camera on: ERROR:", err)
	} else {
		fmt.Println("Camera state:", formatBool(isCameraOn))
	}

	fmt.Println()

	isMicrophoneOn, err := microphone.IsMicrophoneOnDarwin()
	fmt.Println()
	if err != nil {
		fmt.Println("isMicrophoneOn(): ERROR:", err)
	} else {
		fmt.Println("Microphone state:", formatBool(isMicrophoneOn))
	}
}
