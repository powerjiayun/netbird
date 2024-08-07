//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package main

import (
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

const defaultFontPath = "/Library/Fonts/Arial Unicode.ttf"

func (s *serviceClient) setDefaultFonts() {
	// TODO: add other bsd paths
	if runtime.GOOS != "darwin" {
		return
	}

	if _, err := os.Stat(defaultFontPath); err != nil {
		log.Errorf("Failed to find default font file: %v", err)
		return
	}

	os.Setenv("FYNE_FONT", defaultFontPath)
}
