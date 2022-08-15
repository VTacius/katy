package telegram

import (
    "fmt"
    "bytes"
    "errors"
    "net/http"
    "encoding/json"
)

// enviarPeticion trata directamente con la API de telegram
// TODO: Tenemos que manejar mejor los mensajes, sobre todo los de error, que la API pudiera enviar
func EnviarPeticion(token string, chat_id string, contenido string) (int, error){
    valores := map[string]interface{}{
        "chat_id": chat_id, 
        "text": contenido,
        "disable_notification": true,
        "parse_mode": "markdown",
    }
    
    mensaje, err := json.Marshal(valores)
    
    if err != nil {
        return 500, err  
    }
    
    // TODO: Sacar la construcción del endpoint a otra parte
    endpoint := fmt.Sprintf("https://api.telegram.org/%s/sendMessage", token)
    respuesta, err := http.Post(endpoint, "application/json", bytes.NewBuffer(mensaje))
    
    if err != nil {
        return 500, errors.New(fmt.Sprintf("%s", err))
    }
    
    if respuesta.StatusCode != 200 {
        return respuesta.StatusCode, errors.New(fmt.Sprintf("%s", respuesta.Body))
    }

    return respuesta.StatusCode, nil
}

