package config

import (
	"time"
)

const (
	AssistantName = "Джарвис"

	MEDIA_NEXT_TRACK  = 0xB0
	MEDIA_PREV_TRACK  = 0xB1
	MEDIA_PLAY_PAUSE  = 0xB3
	MEDIA_VOLUME_MUTE = 0xAD
	MEDIA_VOLUME_DOWN = 0xAE
	MEDIA_VOLUME_UP   = 0xAF
	KEYEVENTF_KEYUP   = 0x0002

	PhpPath           = "C:\\Users\\Xinnc\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\JetBrains\\PhpStorm 2025.2.3.lnk\""
	VoskEnModel       = `D:\GoProjects\Assistent\models\vosk-model-small-en-us-0.15`
	VoskRuModel       = `D:\GoProjects\Assistent\models\vosk-model-small-ru-0.22`
	SampleRate        = 16000
	BufferSize        = 4000
	ClapThreshold     = 28000
	ClapCooldown      = time.Millisecond * 150
	DoubleClapTimeout = time.Second * 1
)
