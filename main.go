package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/codingame/inodebeat/beater"
)

func main() {
	err := beat.Run("inodebeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
