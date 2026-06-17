package audio

import (
	"assistent/config"
	"assistent/speech"
	"time"

	vosk "github.com/alphacep/vosk-api/go"
)

func DetectDoubleClap(
	buffer []int16,
	clapCount int,
	lastClap time.Time,
	recognizer *vosk.VoskRecognizer,
) (int, time.Time, bool) {

	var maxVol int16

	for _, sample := range buffer {
		value := int32(sample)

		if value < 0 {
			value = -value
		}

		if sample > maxVol {
			maxVol = sample
		}
	}

	text, _ := speech.Recognize(
		recognizer,
		buffer,
	)
	if text == config.AssistantName || text == "подонок" {
		return 0, lastClap, true
	}

	if maxVol > config.ClapThreshold &&
		time.Since(lastClap) > config.ClapCooldown {

		clapCount++
		lastClap = time.Now()

		if clapCount == 2 {
			return 0, lastClap, true
		}
	}

	if clapCount == 1 &&
		time.Since(lastClap) > config.DoubleClapTimeout {

		clapCount = 0
	}

	return clapCount, lastClap, false
}
