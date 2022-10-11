package main

import (
	"fmt"
	"os"
	"sanidad/alortiz/katy/handlers"
	"text/template"

	"github.com/gin-gonic/gin"
)

var plantilla *template.Template

func init() {
	directorioPlantillas := os.Getenv("KATY_PLANTILLAS")
	if directorioPlantillas == "" {
		directorioPlantillas = "plantillas/"
	}
	contenido, err := template.ParseGlob(fmt.Sprintf("%s/*.tpl", directorioPlantillas))
	if err != nil {
		fmt.Println("Error con las plantillas: %w", err)
		os.Exit(1)
	}
	plantilla = template.Must(contenido, err)
}

func main() {
	TELEGRAM_BOT_TOKEN := os.Getenv("TELEGRAM_BOT_TOKEN")
	TELEGRAM_CHAT_ID := os.Getenv("TELEGRAM_CHAT_ID")
	KATY_PROXY_IP := os.Getenv("KATY_PROXY_IP")
	KATY_SOCKET := os.Getenv("KATY_SOCKET")

	newEnviarAlerta := handlers.RecibirAlerta(plantilla, TELEGRAM_BOT_TOKEN, TELEGRAM_CHAT_ID)
	newDebugearAlerta := handlers.DebugearAlerta()

	router := gin.Default()
	router.SetTrustedProxies([]string{KATY_PROXY_IP})

	router.POST("/alertas", newEnviarAlerta)
	router.POST("/alertas/debug", newDebugearAlerta)

	router.Run(KATY_SOCKET)
}
