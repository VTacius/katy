package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func crearPeticion(chat_id string, contenido string) (*bytes.Buffer, error) {
	valores := map[string]interface{}{
		"chat_id":              chat_id,
		"text":                 contenido,
		"disable_notification": true,
		"parse_mode":           "markdown",
	}

	mensaje, err := json.Marshal(valores)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(mensaje), nil
}

// enviarPeticion trata directamente con la API de telegram
// TODO: Tenemos que manejar mejor los mensajes, sobre todo los de error, que la API pudiera enviar
func EnviarPeticion(token string, chat_id string, contenido string) (int, error) {

	mensaje, err := crearPeticion(chat_id, contenido)
	if err != nil {
		return 500, fmt.Errorf("error creando la petición")
	}
	// TODO: Sacar la construcción del endpoint a otra parte
	endpoint := fmt.Sprintf("https://api.telegram.org/%s/sendMessage", token)
	respuesta, err := http.Post(endpoint, "application/json", mensaje)

	if err != nil {
		return 500, fmt.Errorf("%s", err)
	}

	if respuesta.StatusCode != 200 {
		cuerpo, _ := ioutil.ReadAll(respuesta.Body)
		contenido := string(cuerpo)
		return respuesta.StatusCode, fmt.Errorf("%s", contenido)
	}

	return respuesta.StatusCode, nil
}
