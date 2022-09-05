# katy
Recibe alertas de InfluxDB y las envia a Telegram

## Construcción
La construcción del binario debe hacerse en un sistema independiente al servidor en producción. Un servidor en producción no debería tener herramientas para compilación de paquetes.

El truco esta en usar una contenedor igual al sistema destino (**Debian Bullseye**) y con la misma versión de Golang (**1.18**), de allí el nombre de la imagen usada **golang:1.18-bullseye**
```bash
podman run  -it  --rm -v "$PWD":/go/src/myapp -w /go/src/myapp golang:1.18-bullseye go build .
```

Comando que, como ya es sabido, es totalmente compatible con docker
```bash
docker run  -it  --rm -v "$PWD":/go/src/myapp -w /go/src/myapp golang:1.18-bullseye go build .
```

## Instalación
El binario se envía al servidor destino. SCP bastaría
```bash
scp gaby root@servidor:/usr/local/sbin
```

## Configuración
```bash
# Creamos el directorio para la aplicación
mkdir /var/lib/katy/
scp -r plantillas/ root@servidor:/var/lib/katy/

# Creamos el fichero de configuración
cat <<MAFI >/etc/default/katy
GIN_MODE=release
TELEGRAM_BOT_TOKEN="bot123:ABDCEFGHIJQLMNOPQRT"
TELEGRAM_CHAT_ID="-576489013"
KATY_PROXY_IP="127.0.0.1"
KATY_SOCKET="127.0.0.1:8080"
KATY_PLANTILLAS="/var/lib/katy/plantillas"
MAFI

# Configuramos el servicio
cat <<MAFI> /lib/systemd/system/katy.service 
[Unit]
Description=Gaby, que toma los mensajes y los reenvia a todo mundo
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/sbin/katy
EnvironmentFile=/etc/default/katy

[Install]
WantedBy=multi-user.target
MAFI
```

# Activamos e iniciamos el servicio por primera vez
```bash
systemctl enable --now katy.service
```

## Prueba inicial
```bash
curl localhost:8080/alertas -H 'Content-Type: application/json' -d @contenido.json
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
