
# signet-challenge-backend

## easy setup
download the latest release
```bash
unzip signet-challenge-backend-{your-arch-here}.zip
cd signet-challenge-backend
./signet-challenge-backend serve
```

## building from source
install go version 1.9 or higher

then download and build this program
```bash
go get github.com/ahouts/signet-challenge-backend
cd $GOPATH/src/github.com/ahouts/signet-challenge-backend
go build
./signet-challenge-backend serve
```

if you want to change the port or use a different schedule json file, look at the help
```bash
./signet-challenge-backend --help
./signet-challenge-backend serve --help
```
