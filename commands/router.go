package commands

import (
	"fmt"
)

var commandMap = map[string]func(){
	"открыть пхп шторм": OpenPhpStorm,

	"открыть браузер": OpenBrowser,

	"запустить музыку": OpenYandexMusic,

	"включи музыку":   PlayPause,
	"следующий трек":  NextTrack,
	"предыдущий трек": PrevTrack,

	"сделай громче": VolumeUp,
	"сделай тише":   VolumeDown,
	"выключи звук":  VolumeMute,
}

func HandleCommand(text string) (bool, bool) {

	text = NormalizeCommand(text)

	if command, exists := commandMap[text]; exists {
		command()
		return false, false
	}

	switch text {
	case "что ты умеешь":
		fmt.Println("-----------------------------------------------------")
		fmt.Println("открыть пхп шторм")
		fmt.Println("открыть браузер")
		fmt.Println("открыть музыку")
		fmt.Println("включить музыку")
		fmt.Println("переключать треки")
		fmt.Println("стоп")
		fmt.Println("закрыть программу")

	case "закрыть программу":
		fmt.Println("Закрытие программы...")
		return true, false

	case "стоп":
		fmt.Println("Остнавливаю прослушивание...")
		return false, true
	}
	return false, false
}

func NormalizeCommand(text string) string {

	switch text {

	case "открой браузер",
		"открыть браузер",
		"запустить браузер",
		"запусти браузер":
		return "открыть браузер"

	case "открой пхп шторм",
		"открыть пхп шторм",
		"запустить пхп шторм",
		"запуск пхп шторм":
		return "открыть пхп шторм"

	case "закрыть программу",
		"закрой программу":
		return "закрыть программу"

	case "запусти музыку",
		"запустить музыку",
		"открой музыку",
		"открыть музыку":
		return "запустить музыку"

	case "включи музыку",
		"включить музыку",
		"играй музыку",
		"играй",
		"пауза":
		return "включи музыку"

	case "следующий трек",
		"дальше":
		return "следующий трек"

	case "предыдущий трек",
		"назад":
		return "предыдущий трек"

	case "сделай громче",
		"громче":
		return "сделай громче"

	case "сделай тише",
		"тише":
		return "сделай тише"

	case "выключи звук",
		"без звука",
		"мут":
		return "выключи звук"

	}
	return text
}
