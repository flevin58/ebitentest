package resources

import (
	"bytes"
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/*/*.png
var resources embed.FS

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetImage(path string) *ebiten.Image {
	image_bytes, err := resources.ReadFile(path)
	failOnError(err)
	image, _, err := image.Decode(bytes.NewReader(image_bytes))
	failOnError(err)
	return ebiten.NewImageFromImage(image)
}

func AsFS() embed.FS {
	return resources
}

func AsBytes(path string) []byte {
	data, err := resources.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func AsString(path string) string {
	data, err := resources.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
