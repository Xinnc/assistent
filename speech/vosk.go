package speech

import (
	"assistent/config"
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/alphacep/vosk-api/go"
)

type Result struct {
	Text string `json:"text"`
}

func InitRecognizer() (*vosk.VoskModel, *vosk.VoskRecognizer, error) {
	model, err := vosk.NewModel(config.VoskRuModel)
	if err != nil {
		return nil, nil, err
	}
	recognizer, err := vosk.NewRecognizerGrm(
		model,
		16000,
		`[
		"открой пхп шторм",
		"открыть пхп шторм",
		"запустить пхп шторм",
		"запуск пхп шторм",

		"открой браузер",
		"открыть браузер",
		"запустить браузер",
		"запусти браузер",

		"запусти музыку",
		"запустить музыку",
		"открой музыку",
		"открыть музыку",

		"включи музыку",
		"включить музыку",
		"играй музыку",
		"играй",
		"пауза",

		"следующий трек",
		"дальше",

		"предыдущий трек",
		"назад",

		"сделай громче",
		"громче",

		"сделай тише",
		"тише",

		"выключи звук",
		"без звука",
		"мут",


		"стоп",
		"закрыть программу",
		"закрой программу",

		"что ты умеешь"
	]`,
	)
	if err != nil {
		return nil, nil, err
	}

	return model, recognizer, nil
}

func Recognize(
	recognizer *vosk.VoskRecognizer,
	buffer []int16,
) (string, error) {

	bytesBuffer := new(bytes.Buffer)

	err := binary.Write(
		bytesBuffer,
		binary.LittleEndian,
		buffer,
	)

	if err != nil {
		return "", err
	}

	byteBuf := bytesBuffer.Bytes()

	if recognizer.AcceptWaveform(byteBuf) == 1 {

		var result Result

		err := json.Unmarshal(
			[]byte(recognizer.Result()),
			&result,
		)

		if err != nil {
			return "", err
		}

		return result.Text, nil
	}

	return "", nil
}
