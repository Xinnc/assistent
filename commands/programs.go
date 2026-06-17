package commands

import (
	"assistent/config"
	"fmt"
	"os/exec"
)

func OpenPhpStorm() {
	fmt.Println("Запускаю PhpStorm")
	exec.Command("cmd", "/c", "start", "", config.PhpPath).Start()
}
func OpenBrowser() {
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "/").Start()
}
func OpenYandexMusic() {
	exec.Command(
		"explorer.exe",
		"shell:AppsFolder\\A025C540.Yandex.Music_vfvw9svesycw6!App",
	).Start()
}
