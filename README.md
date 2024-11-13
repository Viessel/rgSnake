# Snake game

## Building from linux
```bash
$ go build .
```

## Cross compiling for windows 
```bash
$ CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" .
```

## Deps for Fedora

```bash
# dnf install mesa-libGL-devel libXi-devel libXcursor-devel libXrandr-devel libXinerama-devel wayland-devel libxkbcommon-devel
# dnf install mingw32-gcc
```

