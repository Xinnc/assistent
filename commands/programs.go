package commands

import (
	"assistent/config"
	"assistent/tts"
	"log"
	"os/exec"
)

func OpenPhpStorm() {
	log.Println("Запускаю PhpStorm...")
	tts.Speak("Запускаю PhpStorm...")
	exec.Command("cmd", "/c", "start", "", config.PhpPath).Start()
}
func OpenBrowser() {
	log.Println("Открываю бразуер...")
	tts.Speak("Открываю бразуер...")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://google.com").Start()
}
func OpenGitHub() {
	log.Println("Открываю GitHub...")
	tts.Speak("Открываю GitHub...")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/").Start()
}
func OpenYouTube() {
	log.Println("Открываю YouTube...")
	tts.Speak("Открываю YouTube...")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.youtube.com/").Start()
}
func OpenTelegram() {
	log.Println("Открываю Telegram...")
	tts.Speak("Открываю Telegram...")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://web.telegram.org/k/").Start()
}
func OpenChatgpt() {
	log.Println("Открываю Chatgpt...")
	tts.Speak("Открываю Chatgpt...")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://chatgpt.com/").Start()
}
func OpenYandexMusic() {
	log.Println("Открываю Yandex.Music...")
	tts.Speak("Открываю Yandex.Music...")
	exec.Command(
		"explorer.exe",
		"shell:AppsFolder\\A025C540.Yandex.Music_vfvw9svesycw6!App",
	).Start()
}
