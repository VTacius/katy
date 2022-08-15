package handlers

import (
    "fmt"
    "net/http"
    "sanidad/alortiz/katy/telegram"
    "github.com/gin-gonic/gin"
)

// DebugearAlerta Para ver que esta enviando como parte de la alerta
func DebugearAlerta() func(*gin.Context){
    return func (c *gin.Context) {
        datos, err := c.GetRawData()

        if err != nil {
            c.IndentedJSON(500, gin.H{"error": fmt.Sprintf("%s", err)})
            return 
        }
        
        fmt.Println(string(datos))
    }
}

// RecibirAlerta Pues nada, que maneja la alerta que InfluxDB envía
func RecibirAlerta[alertaTipo any](token string, chat_id string, newAlerta alertaTipo) func(*gin.Context) {
    return func (c *gin.Context) {
        if err := c.BindJSON(&newAlerta); err != nil {
            c.IndentedJSON(500, gin.H{"error": fmt.Sprintf("%s", err)})
            return 
        }
        
        contenidoAlerta := fmt.Sprintf("%s", newAlerta)
        codigo, err := telegram.EnviarPeticion(token, chat_id, contenidoAlerta)

        if err != nil {
            c.IndentedJSON(codigo, gin.H{"error": fmt.Sprintf("%s", err)})
            return
        }  
        
        c.IndentedJSON(http.StatusCreated, newAlerta)
    }
}

