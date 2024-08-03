package mediadevicesstate

// IsCameraOn returns true is any camera in the system is ON
func IsCameraOnWindows(logging bool) (bool, error) {
	return false, nil
}

// IsMicrophoneOn returns true is any camera in the system is ON
func IsMicrophoneOnWindows(logging bool) (bool, error) {
	return false, nil
}

// Debug calls all available device functions and prints the results
func Debug() {
}
