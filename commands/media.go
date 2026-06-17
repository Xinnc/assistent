package commands

import (
	"assistent/config"

	"golang.org/x/sys/windows"
)

var (
	user32         = windows.NewLazySystemDLL("user32.dll")
	keybdEventProc = user32.NewProc("keybd_event")
)

func pressKey(key byte) {
	keybdEventProc.Call(
		uintptr(key),
		0,
		0,
		0,
	)

	keybdEventProc.Call(
		uintptr(key),
		0,
		config.KEYEVENTF_KEYUP,
		0,
	)
}

func PlayPause() {
	pressKey(config.MEDIA_PLAY_PAUSE)
}

func NextTrack() {
	pressKey(config.MEDIA_NEXT_TRACK)
}

func PrevTrack() {
	pressKey(config.MEDIA_PREV_TRACK)
}
func VolumeUp() {
	for i := 1; i < 5; i++ {
		pressKey(config.MEDIA_VOLUME_UP)
	}
}
func VolumeDown() {
	for i := 1; i < 5; i++ {
		pressKey(config.MEDIA_VOLUME_DOWN)
	}
}
func VolumeMute() {
	pressKey(config.MEDIA_VOLUME_MUTE)
}
