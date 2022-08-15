package main
import (
    "os"
    "sanidad/alortiz/katy/alertas"
    "sanidad/alortiz/katy/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    TELEGRAM_BOT_TOKEN := os.Getenv("TELEGRAM_BOT_TOKEN")
    TELEGRAM_CHAT_ID := os.Getenv("TELEGRAM_CHAT_ID")
    KATY_PROXY_IP := os.Getenv("KATY_PROXY_IP")
    KATY_SOCKET := os.Getenv("KATY_SOCKET")
    
    var newAlertaDisco alertas.AlertaDisco
    newEnviarAlertaDisco := handlers.RecibirAlerta(TELEGRAM_BOT_TOKEN, TELEGRAM_CHAT_ID, newAlertaDisco)
    var newAlertaBase alertas.AlertaBase
    newEnviarAlertaBase := handlers.RecibirAlerta(TELEGRAM_BOT_TOKEN, TELEGRAM_CHAT_ID, newAlertaBase)
    newDebugearAlerta := handlers.DebugearAlerta() 
    
    router := gin.Default()
    router.SetTrustedProxies([]string{KATY_PROXY_IP})
    
    router.POST("/alertas", newEnviarAlertaBase)
    router.POST("/alertas/disco", newEnviarAlertaDisco)
    router.POST("/alertas/debug", newDebugearAlerta)

    router.Run(KATY_SOCKET)
}
