# Golang UI - Fyne
UI learning process for golang



## Cross Compiling for Windows in Linux
Using this link https://developer.fyne.io/started/cross-compiling
- install required binaries by using below commands
   - apt-get install gcc-multilib
   - apt install gcc-mingw-w64


- example command
  `env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build app.go`
