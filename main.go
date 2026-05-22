package main

import (
	"fmt"
	"log"
	"math"
	"time"

	port "github.com/gordonklaus/portaudio"
)

func main() {
	port.Initialize()
	defer port.Terminate()

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

	buf := make([]float32, 512)

	stream, err := port.OpenDefaultStream(1, 0, 44100, 64, &buf)
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
	for {
		err := stream.Read()
		if err != nil {
			log.Fatal(err)
		}

		maxVol := 0.0
		for _, sample := range buf {
			sound := math.Abs(float64(sample))
			if sound > maxVol {
				maxVol = sound
			}
		}
		if maxVol > 0.9 && time.Since(clap) > cooldown {
			clap = time.Now()
			clapCount++
			fmt.Println(clapCount)
			if clapCount == 2 {
				clapCount = 0
				fmt.Println("Start Program")
			}
		} else if clapCount == 1 && time.Since(clap) > doubleClap {
			clapCount = 0
		}
	}
}
