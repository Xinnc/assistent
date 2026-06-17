package tts

import (
	"log"
	"os/exec"
)

func Speak(text string) error {

	cmd := exec.Command(
		"tts/piper/piper.exe",
		"--model",
		"models/piper-ruslan-medium/ru_RU-ruslan-medium.onnx",
		"--output_file",
		"speech.wav",
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
	}

	_, err = stdin.Write([]byte(text))
	if err != nil {
		log.Println(err)
	}

	err = stdin.Close()
	if err != nil {
		log.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
	}

	return exec.Command(
		"powershell",
		"-c",
		"(New-Object Media.SoundPlayer 'speech.wav').PlaySync()",
	).Run()
}
