// Texture Loading and Rendering
// Adapted from http://lazyfoo.net/tutorials/SDL/07_texture_loading_and_rendering/index.php

package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	err      error
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture

	quit  bool
	event sdl.Event
)

func initSDL() error {
	err = sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		return err
	}

	window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}

	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

	imgFlags := img.INIT_PNG
	imgInitResult := img.Init(imgFlags)
	if (imgInitResult & imgFlags) != imgFlags {
		return img.GetError()
	}

	return nil
}

func loadTexture(path string) (*sdl.Texture, error) {
	var newTexture *sdl.Texture

	loadedSurface, err := img.Load(path)
	if err != nil {
		return nil, err
	}

	newTexture, err = renderer.CreateTextureFromSurface(loadedSurface)
	if err != nil {
		return nil, err
	}

	loadedSurface.Free()

	return newTexture, nil
}

func loadMedia() error {
	texture, err = loadTexture("texture.png")
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	texture.Destroy()
	renderer.Destroy()
	window.Destroy()
	img.Quit()
	sdl.Quit()
}

func main() {
	err = initSDL()
	if err != nil {
		log.Fatal("Error initializing SDL:", err)
	}

	err = loadMedia()
	if err != nil {
		log.Fatal("Error loading Media:", err)
	}

	quit = false
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}

		renderer.Clear()
		renderer.Copy(texture, nil, nil)
		renderer.Present()
	}

	Close()
}
