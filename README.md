# katy
Recibe alertas de InfluxDB y las envia a Telegram

## Construcción
La construcción del binario debe hacerse en un sistema independiente al servidor en producción. Un servidor en producción no debería tener herramientas para compilación de paquetes.
El truco esta en usar una contenedor igual al sistema destino (**Debian Bullseye**) y con la misma versión de Golang (**1.18**), de allí el nombre de la imagen usada **golang:1.18-bullseye**
```bash
podman run  -it  --rm -v "$PWD":/go/src/myapp -w /go/src/myapp golang:1.18-bullseye go build .
```

## Instalación
El binario se envía al servidor destino. SCP bastaría
scp gaby root@servidor:/usr/local/sbin


## Configuración
```bash
export GIN_MODE=release
export TELEGRAM_BOT_TOKEN="bot123:ABDCEFGHIJQLMNOPQRT"
export TELEGRAM_CHAT_ID="-576489013"
export KATY_PROXY_IP="127.0.0.1"
export KATY_SOCKET="127.0.0.1:8080"
```

## Prueba inicial
```bash
localhost:8080/alertas -H 'Content-Type: application/json' -d @contenido.json
```
El contenido de json podría ser el siguiente
```json
{
    "_check_name": "Temperatura alta",
    "_level": "Crit",
    "host": "xibalba",
    "parametro": "/var",
    "path": "/var",
    "temp2": 38.2,
    "used_percent": "38%"
}
```
