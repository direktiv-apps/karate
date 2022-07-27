#!/bin/sh

docker build -t karate . && docker run -p 9191:8080 karate