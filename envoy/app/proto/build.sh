#!/bin/bash

docker pull namely/protoc-all
cd helloworld
docker run -v `pwd`:/defs namely/protoc-all -f helloworld.proto -o . -l go
docker run -v `pwd`:/defs namely/protoc-all -f helloworld.proto -o . -l descriptor_set
