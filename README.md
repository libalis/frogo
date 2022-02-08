# FroGo
Fyne + RobotGo
## License
frogo.go: Copyright (C) 2022 Robert Kagan

fyne.png: Copyright (C) 2018 Fyne.io developers
## Base
- [Fyne](https://fyne.io/)
- [RobotGo](https://github.com/go-vgo/robotgo)
- [hotkey](https://github.com/golang-design/hotkey)
## Windows
### Requirements
 - [Go](https://go.dev/)
 - [MinGW-w64](https://sourceforge.net/projects/mingw-w64/)
### Compilation
    git clone https://github.com/libalis/frogo
    cd frogo
    go mod init frogo
    go mod tidy
    go install github.com/tc-hib/go-winres@latest
    go-winres simply --icon fyne.png --manifest gui
    go build -ldflags -H=windowsgui -o FRoGo.exe .
