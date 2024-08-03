package mediadevicesstate

import (
	"github.com/dqle/go-media-devices-state/pkg/microphone"
)

// IsCameraOn returns true is any camera in the system is ON
func IsCameraOn(logging bool) (bool, error) {
	return false, nil
}

// IsMicrophoneOn returns true is any camera in the system is ON
func IsMicrophoneOn(logging bool) (bool, error) {
	return microphone.IsMicrophoneOnWindows(logging)
}

// Debug calls all available device functions and prints the results
func Debug() {
}
