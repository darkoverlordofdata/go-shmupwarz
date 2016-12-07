package main

// main
func main() {

	game := NewShmupWarz(800, 600, "ShmupWarz")
	defer game.Destroy()
	game.Run(game)
}
