#bin/bash

docker build -t super-heroes . --no-cache
docker run -p 3222:3222 --rm -it --network=host super-heroes:latest 
