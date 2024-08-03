# go-media-devices-state

Forked from: github.com/antonfisher/go-media-devices-state

Go module to get camera/microphone state -- checks if camera/microphone is ON.

Module uses `cgo` to call native specific API on darwin (MacOS) platform and uses `registry` to get the microphone and webcam windows registry.


## Installation

```shell
go get github.com/dqle/go-media-devices-state
```

## Usage

```go
package main

import (
	"fmt"

	mediaDevices "github.com/dqle/go-media-devices-state"
)

func main() {
	isCameraOn, err := mediaDevices.IsCameraOn()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Is camera on:", isCameraOn)
	}

	isMicrophoneOn, err := mediaDevices.IsMicrophoneOn()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Is microphone on:", isMicrophoneOn)
	}
}
```

## Implemented APIs

```go
// IsCameraOn returns true is any camera in the system is ON
func IsCameraOn() (bool, error)

// IsMicrophoneOn returns true is any microphone in the system is ON
func IsMicrophoneOn() (bool, error)

// Debug calls all available device functions and prints the results with logging enabled
func Debug()
```

## Demo

Demo prints out all system video devices and their states (ON/OFF):

```shell
git clone https://github.com/dqle/go-media-devices-state-darwin.git
cd go-media-devices-state-darwin
go run -a cmd/demo.go
```

Output example:

```
Debug go-media-devices-state module...

2021-01-12 23:04:53.674 demo[70272:606077] C.IsCameraOn()
2021-01-12 23:04:53.698 demo[70272:606077] C.IsCameraOn(): found devices: 1
2021-01-12 23:04:53.698 demo[70272:606077] C.IsCameraOn(): # | is used | description
2021-01-12 23:04:53.725 demo[70272:606077] C.IsCameraOn(): 0 | NO      | 0x8020000005ac1234 (name: 'FaceTime HD Camera (Built-in)', model: 'UVC Camera VendorID_1234 ProductID_12345', is exclusively used: 0)
2021-01-12 23:04:53.725 demo[70272:606077] C.IsCameraOn(): failed devices: 0
2021-01-12 23:04:53.725 demo[70272:606077] C.IsCameraOn(): ignored devices (always on): 0
2021-01-12 23:04:53.725 demo[70272:606077] C.IsCameraOn(): is any camera on: NO

Camera state: OFF

2021-01-12 23:04:53.725 demo[70272:606077] C.IsMicrophoneOn()
2021-01-12 23:04:53.725 demo[70272:606077] C.IsMicrophoneOn(): found devices: 2
2021-01-12 23:04:53.725 demo[70272:606077] C.IsMicrophoneOn(): # | is used | description
2021-01-12 23:04:53.725 demo[70272:606077] C.IsMicrophoneOn(): 0 | NO      | BuiltInMicrophoneDevice (name: 'MacBook Pro Microphone', model: 'Digital Mic', is exclusively used: 0)
2021-01-12 23:04:53.726 demo[70272:606077] C.IsMicrophoneOn(): failed audio devices: 0
2021-01-12 23:04:53.726 demo[70272:606077] C.IsMicrophoneOn(): ignored devices (speakers): 1
2021-01-12 23:04:53.726 demo[70272:606077] C.IsMicrophoneOn(): is any microphone on: NO

Microphone state: OFF
```

## Troubleshooting

List all registered cameras:
```shell
// macOS
system_profiler SPCameraDataType
```

List all registered microphones and speakers:
```shell
// macOS
system_profiler SPAudioDataType
```

## License

MIT License.
