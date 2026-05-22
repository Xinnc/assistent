package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	vosk "github.com/alphacep/vosk-api/go"
	port "github.com/gordonklaus/portaudio"
)

type Result struct {
	Text string
}

const VOSK_EN_MODEL = `D:\GoProjects\Assistent\models\vosk-model-small-en-us-0.15`

func main() {
	port.Initialize()
	defer port.Terminate()

	model, err := vosk.NewModel(VOSK_EN_MODEL)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]int16, 4000)

	recognizer, err := vosk.NewRecognizer(model, 16000)
	if err != nil {
		fmt.Println(err)
		return
	}

	stream, err := port.OpenDefaultStream(1, 0, 16000, 64, &buf)

	if err != nil {
		fmt.Println(err)
	}

	err = stream.Start()
	if err != nil {
		return
	}

	cooldown := time.Millisecond * 150
	doubleClap := time.Second * 1
	clapCount := 0
	clap := time.Now()
	listening := false

	fmt.Println("Program started!")
	for {
		err := stream.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !listening {
			var maxVol int16 = 0
			for _, sample := range buf {
				if sample > maxVol {
					maxVol = sample
				}
			}
			if maxVol > 30000 && time.Since(clap) > cooldown {
				clap = time.Now()
				clapCount++
				if clapCount == 2 {
					clapCount = 0
					fmt.Println("Listening...")
					listening = true
				}
			} else if clapCount == 1 && time.Since(clap) > doubleClap {
				clapCount = 0
			}
		} else {
			bytesBuffer := new(bytes.Buffer)
			binary.Write(bytesBuffer, binary.LittleEndian, buf)
			byteBuf := bytesBuffer.Bytes()
			if recognizer.AcceptWaveform(byteBuf) == 1 {
				text := Result{}

				result := recognizer.Result()
				err := json.Unmarshal([]byte(result), &text)
				if err != nil {
					return
				}
				if text.Text != "" {
					fmt.Println(text.Text)
				}
				switch text.Text {
				case "close the program":
					fmt.Println("Closing program...!")
					return
				case "close":
					fmt.Println("I can't hear you anymore(")
					listening = false
				case "what can you do":
					fmt.Println("I repeat you, i can also:")
					fmt.Println("-----------------------------------------------------")
					fmt.Println("  wash your belly ----> I'll wash your belly")
					fmt.Println("  close the program --> I'll close this window")
					fmt.Println("  close --------------> You'll have to clap again")
				case "wash my belly":
					fmt.Println("YEAAAAA")
				}
			}
		}
	}

	//devices, err := port.Devices()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, d := range devices {
	//	Println(d)
	//}
	//
	//Println("-------------")
	//
	//devices1, err1 := port.DefaultInputDevice()
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	//Println(devices1)

	//buf1 := make([]float32, 512)
	//
	//stream1, err1 := port.OpenDefaultStream(1, 0, 44100, 64, &buf)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = stream.Start()
	//if err != nil {
	//	return
	//}

	//cooldown := time.Millisecond * 150
	//doubleClap := time.Second * 1
	//clapCount := 0
	//clap := time.Now()
	//for {
	//	err := stream.Read()
	//	if err != nil {
	//		log.Fatal(err)
	//	} else {
	//		fmt.Println("You can clap!")
	//	}
	//
	//	maxVol := 0.0
	//	for _, sample := range buf {
	//		sound := math.Abs(float64(sample))
	//		if sound > maxVol {
	//			maxVol = sound
	//		}
	//	}
	//	if maxVol > 0.9 && time.Since(clap) > cooldown {
	//		clap = time.Now()
	//		clapCount++
	//		fmt.Println(clapCount)
	//		if clapCount == 2 {
	//			clapCount = 0
	//			fmt.Println("Start Program")
	//		}
	//	} else if clapCount == 1 && time.Since(clap) > doubleClap {
	//		clapCount = 0
	//	}
	//}
}
