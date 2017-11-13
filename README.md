
# signet-challenge-backend

## easy setup
download the latest release
```bash
mkdir signet-challenge-backend && cd signet-challenge-backend
unzip /path/to/signet-challenge-backend-{your-arch-here}.zip
./signet-challenge-backend serve
```

## getting help

```bash
./signet-challenge-backend --help
./signet-challenge-backend serve --help
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
