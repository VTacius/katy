#!/bin/bash

TOKEN=${1:-bot123:ABDCEFGHIJQLMNOPQRT}
CHAT=${2:-"-576489013"}

# Creamos el directorio para la aplicación
[ -d /var/lib/katy/ ] || mkdir /var/lib/katy/
cp -r plantillas/ /var/lib/katy/

# Movemos a Katy a su nueva ubicación
cp katy /usr/local/sbin

# Creamos el fichero de configuración
cat <<MAFI >/etc/default/katy
GIN_MODE=release
TELEGRAM_BOT_TOKEN="${TOKEN}"
TELEGRAM_CHAT_ID="${CHAT}"
KATY_PROXY_IP="127.0.0.1"
KATY_SOCKET="127.0.0.1:8080"
KATY_PLANTILLAS="/var/lib/katy/plantillas"
MAFI

# Configuramos el servicio
cat <<MAFI> /lib/systemd/system/katy.service 
[Unit]
Description=Gaby, reenvío de mensajes a Telegram 
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/sbin/katy
EnvironmentFile=/etc/default/katy

[Install]
WantedBy=multi-user.target
MAFI

# Activamos e iniciamos el servicio por primera vez
systemctl enable --now katy.service
