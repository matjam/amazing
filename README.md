# amazing

A game of life simulator written in Go, using ebiten.

## Building

1. Install msys2 using chocolatey.
2. Open a msys2 shell. Install the development packages

```bash
pacman -S --needed base-devel mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain
```

3. Install the go compiler using msys2 or chocolatey, it doesn't matter. Make sure it's in your path and make sure gcc
   from msys2 is in your path too.

`go run cmd/amazing/main.go`

## Keyboard

* **F1** Randomize
* **F2** Clear the screen
* **SPACE** pause
* **ESC** quit

## Mouse

You can click around anywhere and add more live cells to the simulation.
