#!/bin/bash

mkdir build dist
cp -r ./schedule_data.json ./build

# linux builds
env GOOS=linux GOARCH=amd64 go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-linux-amd64.zip ./*
cd ..

env GOOS=linux GOARCH=386 go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-linux-i386.zip ./*
cd ..

env GOOS=linux GOARCH=arm go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-linux-arm.zip ./*
cd ..

# mac builds
env GOOS=darwin GOARCH=amd64 go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-macOS-amd64.zip ./*
cd ..

env GOOS=darwin GOARCH=386 go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-macOS-i386.zip ./*
cd ..

# windows builds
env GOOS=windows GOARCH=amd64 go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-windows-amd64.zip ./*
cd ..

env GOOS=windows GOARCH=386 go build
cp ./signet-challenge-backend ./build
cd ./build
zip ../dist/signet-challenge-backend-windows-i386.zip ./*
cd ..
