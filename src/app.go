package main

import "runtime"

// main
func main() {
	runtime.LockOSThread()

	game := NewShmupWarz(800, 600, "ShmupWarz")
	defer game.Destroy()
	game.Run(game)
}
