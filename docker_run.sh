#!/bin/bash

which docker

if [ $? -ne 0 ]; then
    printf "Docker is not installed. You can't run this script without Docker\n"
    printf "If you have GO installed on your machine you can try running:\n make run TEXT=\"Something awesome\"\n"
    exit 1
fi

docker build -t printer . && docker run -it \
 -v $(pwd)/assets:/assets \
 --rm printer \
 -fontPath /assets/FiraSans-Light.ttf \
 -bgImg /assets/00-instagram-background.png \
 -output "/assets/cool_img.png" \
 -text "$1"
