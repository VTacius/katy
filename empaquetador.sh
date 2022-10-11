#!/bin/bash

# Creamos el directorio de trabajo si no existe
[ -d output ] || mkdir output

DEBIAN_VERSION=(11.4 11.5)
KATY_VERSION=$(git describe --abbrev=0 --tags)

for dv in ${DEBIAN_VERSION[*]}; do 
    [ -f katy ] && rm katy
    podman run --rm -it -v $PWD:/usr/local/src/ alortiz/constructor-golang:$dv-1.18 go build
    tar czvf output/katy-${KATY_VERSION}-$dv.tgz utils/instalador.sh utils/contenido.json katy plantillas --transform 's,^utils/,,' --transform 's,^,katy/,' 
done
