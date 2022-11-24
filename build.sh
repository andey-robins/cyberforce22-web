#!/bin/bash
docker build -t energy-web -f Dockerfile .
docker run -it -p 80:80 -d energy-web