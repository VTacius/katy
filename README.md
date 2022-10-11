# katy
Recibe alertas de InfluxDB y las envia a Telegram

## Escogiendo la versión
Los paquetes se encuentran en [la página de lanzamientos](https://github.com/VTacius/katy/releases/latest). Por ahora, solo hay paquetes para el último lanzamiento de Debian, y de ellos, debe escoger la versión para su sistema según el esquema de versionado:
``` 
katy-{version-aplicativo}-{version-debian}.tgz 
```

Así por ejemplo, para la versión 
```
katy-v0.9.5-11.4.tgz 
```

Significa que es la versión `v0.9.5` de la aplicación para la versión `11.4` de Debian

## Instalación
Una vez el paquete esté en el servidor, de descomprime:
```bash
tar xzvf katy-v0.9.5-11.5.tgz
```

Se entra al directorio resultante
```bash
cd katy/
```

Y se ejecuta el instalador
```bash
bash instalador.sh bot5352586733:AAGdHYn348vV6YWm-iogrVjB1Noa0uHjkgA -576489013
```
Donde el primero parametro es el token de telegram, y el segundo es el ID del chat al cuál se van a enviar los mensajes.

Alternativamente, puede pasar ejecutarse el instalador sin los parámetros, para luego configurarse a mano en `/etc/default/katy`

## Prueba inicial
Se suministra un json de prueba `contenido.json`, que puede usarse para una prueba inicial
```bash
curl localhost:8080/alertas -H 'Content-Type: application/json' -d @contenido.json
```

## El sistema de plantillas
En `/var/lib/katy/plantilas/` se crea una plantila con el nombre de la alerta en específico. Dicho nombre se específica en el atributo `_check_name`, que el asigna InfluxDB cuando se configura la alerta.

Así por ejemplo, si se tiene algo como lo siguiente:

```json
{
    "_check_name": "Temperatura alta",
    "_level": "Crit",
    "host": "xibalba",
    "parametro": "/var",
    ...
```

Debe crearse una plantilla con nombre `/var/lib/katy/plantillas/Temperatura alta.tpl` que use los atributos disponibles.

Para ver que atributos se encuentran disponibles en la alerta, es posible usar el endpoint `/alertas/debug`, que no envía alerta por telegram, sino que muestra todo su contenido en consola