package main

import (
	"assistent/audio"
	"assistent/commands"
	"assistent/speech"
	"assistent/tts"
	"log"
	"os"
	"time"

	_ "embed"

	"github.com/getlantern/systray"
)

//go:embed assets/icon.ico
var icon []byte

var statusItem *systray.MenuItem

func onReady() {
	statusItem = systray.AddMenuItem("Запущен", "")
	statusItem.Disable()
	systray.AddSeparator()

	systray.SetIcon(icon)
	systray.SetTitle("Assistant")
	systray.SetTooltip("Voice Assistant")

	mQuit := systray.AddMenuItem("Выход", "Закрыть ассистента")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		os.Exit(0)
	}()
}

func onExit() {
}

func StartTray() {
	systray.Run(onReady, onExit)
}

func SetStatus(status string) {
	if statusItem != nil {
		statusItem.SetTitle(status)
	}
}

func StartAssistant() {
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

	logFile, _ := os.OpenFile(
		"assistant.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	log.SetOutput(logFile)
	log.Println("Ассистент запущен")
	tts.Speak("Ассистент запущен!")

	for {
		err := stream.Read()
		if err != nil {
			log.Println(err)
			return
		}

		if !listening {

			clapCount, clap, detected =
				audio.DetectDoubleClap(
					buf,
					clapCount,
					clap,
					recognizer,
				)

			if detected {
				log.Println("Слушаю...")
				SetStatus("Слушаю")
				tts.Speak("Слушаю")
				listening = true
			}

			continue
		} else {
			text, err := speech.Recognize(
				recognizer,
				buf,
			)

			if err != nil {
				log.Println(err)
				continue
			}

			if text != "" {

				log.Println(text)

				quit, stop :=
					commands.HandleCommand(text)

				if quit {
					systray.Quit()
					os.Exit(0)
					return
				}

				if stop {
					listening = false
					SetStatus("Ожидаю")
					tts.Speak("Ожидаю")
				}
			}
		}

	}

}
