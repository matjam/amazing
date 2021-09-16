# amazing

A game of life simulator written in Go, using ebiten.

![Animation](https://user-images.githubusercontent.com/578676/133537968-7bd38b7e-8140-4148-a67c-a7462522727f.gif)

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

You can click around anywhere and add more live cells to the simulation. Right click will remove cells.

## Motivation

This was really just a thing to do to test Ebiten and go. I wanted to see how well it could handle just dumbly drawing all of the pixels every frame. It does pretty well!
