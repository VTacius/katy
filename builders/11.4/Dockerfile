FROM debian:11.4-slim

RUN apt update \
        && apt -y install wget git \
        && wget https://golang.org/dl/go1.18.6.linux-amd64.tar.gz \
        && tar -xzvf go1.18.6.linux-amd64.tar.gz -C /usr/local/

ENV PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

WORKDIR /usr/local/src/ 
