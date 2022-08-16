package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "sanidad/alortiz/katy/telegram"
    "sanidad/alortiz/katy/alertas"
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
func RecibirAlerta(token string, chat_id string) func(*gin.Context) {
    return func (c *gin.Context) {
        datos, err := c.GetRawData()

        if err != nil {
            c.IndentedJSON(500, gin.H{"error": fmt.Sprintf("%s", err)})
            return 
        }

        contenido := make(map[string]any)
        json.Unmarshal(datos, &contenido)
       
        indice, _ := contenido["_check_name"]
        formateador, encontrado := alertas.Formateadores[fmt.Sprintf("%s", indice)]
        if !encontrado {
            formateador = alertas.FormatearAlertaBase
        }
        
        contenidoAlerta := formateador(contenido)
        
        codigo, err := telegram.EnviarPeticion(token, chat_id, contenidoAlerta)

        if err != nil {
            c.IndentedJSON(codigo, gin.H{"error": fmt.Sprintf("%s", err)})
            return
        }  
        c.IndentedJSON(http.StatusCreated, gin.H{"tipo": indice})

    }
}

