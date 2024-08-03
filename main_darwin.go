package mediadevicesstate

import (
	"github.com/dqle/go-media-devices-state/pkg/camera"
	"github.com/dqle/go-media-devices-state/pkg/debug"
	"github.com/dqle/go-media-devices-state/pkg/microphone"
)

// IsCameraOn returns true if any camera in the darwin system is ON
func IsCameraOn() (bool, error) {
	return camera.IsCameraOnDarwin()
}

// IsMicrophoneOn returns true if any camera in the darwin system is ON
func IsMicrophoneOn() (bool, error) {
	return microphone.IsMicrophoneOnDarwin()
}

// Debug calls all available device functions and prints the results
func Debug() {
	debug.DebugDarwin()
}
