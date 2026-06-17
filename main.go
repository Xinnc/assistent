package main

import (
	"assistent/audio"
	"assistent/commands"
	"assistent/speech"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ассистент запущен")
	fmt.Println("Дважды хлопните, чтобы начать")

	stream, buf, err := audio.NewStream()
	if err != nil {
		return
	}

	defer audio.Close(stream)

	model, recognizer, err := speech.InitRecognizer()
	if err != nil {
		return
	}
	defer model.Free()
	defer recognizer.Free()

	clapCount := 0
	clap := time.Now()
	listening := false
	detected := false

	for {
		err := stream.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !listening {

			clapCount, clap, detected =
				audio.DetectDoubleClap(
					buf,
					clapCount,
					clap,
				)

			if detected {
				fmt.Println("Listening...")
				listening = true
			}

			continue
		} else {
			text, err := speech.Recognize(
				recognizer,
				buf,
			)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if text != "" {

				fmt.Println(text)

				quit, stop :=
					commands.HandleCommand(text)

				if quit {
					return
				}

				if stop {
					listening = false
				}
			}
		}

	}
}
