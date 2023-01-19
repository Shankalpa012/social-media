package main

import (
	"github.com/joho/godotenv"
	"github.com/shankalpa12/cmd"
)

func main() {
	_ = godotenv.Load()

	cmd.Execute()
}
