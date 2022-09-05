package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sanidad/alortiz/katy/telegram"

	"text/template"

	"github.com/gin-gonic/gin"
)

// DebugearAlerta Para ver que esta enviando como parte de la alerta
func DebugearAlerta() func(*gin.Context) {
	return func(c *gin.Context) {
		datos, err := c.GetRawData()

		if err != nil {
			c.IndentedJSON(500, gin.H{"error": fmt.Sprintf("%s", err)})
			return
		}

		fmt.Println(string(datos))
	}
}

func ejecutarTemplate(plantilla *template.Template, contenido map[string]any) (string, string, error) {
	var destino bytes.Buffer
	indice := contenido["_check_name"]
	nombrePlantilla := fmt.Sprintf("%s.tpl", indice)

	plantilla = plantilla.Lookup(nombrePlantilla)
	if plantilla == nil {
		plantilla = plantilla.Lookup(nombrePlantilla)
	}

	err := plantilla.Execute(&destino, contenido)
	if err != nil {
		return "", "", err
	}

	return destino.String(), fmt.Sprintf("%v", indice), nil
}

// RecibirAlerta Pues nada, que maneja la alerta que InfluxDB envía
func RecibirAlerta(plantilla *template.Template, token string, chat_id string) func(*gin.Context) {
	return func(c *gin.Context) {
		datos, err := c.GetRawData()

		if err != nil {
			c.IndentedJSON(500, gin.H{"error": fmt.Sprintf("%s", err)})
			return
		}

		contenido := make(map[string]any)
		json.Unmarshal(datos, &contenido)

		contenidoAlerta, indice, err := ejecutarTemplate(plantilla, contenido)
		if err != nil {
			c.IndentedJSON(500, gin.H{"error plantilla:> ": fmt.Sprintf("%v", err)})
			return
		}

		codigo, err := telegram.EnviarPeticion(token, chat_id, contenidoAlerta)
		if err != nil {
			fmt.Println(codigo)
			c.IndentedJSON(codigo, gin.H{"error envio:> ": fmt.Sprintf("%v", err)})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"tipo": indice, "codigo": codigo})

	}
}
