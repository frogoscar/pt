package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	. "github.com/fogleman/pt/pt"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run volume.go DIRECTORY")
		return
	}
	dirname := args[0]
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	var images []image.Image
	for _, info := range infos {
		filename := path.Join(dirname, info.Name())
		im, err := LoadPNG(filename)
		if err != nil {
			continue
			// panic(err)
		}
		images = append(images, im)
	}

	scene := Scene{}
	// scene.SetColor(Color{1, 1, 1})

	colors := []Color{
		HexColor(0xFFF8E3),

		// HexColor(0x004358),
		// HexColor(0x1F8A70),
		// HexColor(0xBEDB39),
		// HexColor(0xFFE11A),
		// HexColor(0xFD7400),

		// HexColor(0xFFE11A),
		// HexColor(0xBEDB39),
		// HexColor(0x1F8A70),
		// HexColor(0x004358),

		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
		// Color{1, 1, 1},
	}
	const (
		start = 0.8
		size  = 0.25
		step  = 0.02
	)
	var windows []VolumeWindow
	for i := 0; i < len(colors); i++ {
		lo := start + step*float64(i)
		hi := lo + size
		material := GlossyMaterial(colors[i], 1.1, Radians(20))
		w := VolumeWindow{lo, hi, material}
		windows = append(windows, w)
	}
	box := Box{Vector{-1, -0.5, -1.3}, Vector{1, 0.65, 1.3}}
	volume := NewVolume(box, images, 3.4/0.9765625, windows)
	scene.Add(volume)

	wall := GlossyMaterial(Color{1, 1, 1}, 1.1, Radians(20))
	scene.Add(NewCube(V(-10, 0.65, -10), V(10, 10, 10), wall))

	light := LightMaterial(Color{1, 1, 1}, 20, NoAttenuation)
	scene.Add(NewSphere(V(1, -5, -1), 1, light))

	fmt.Println(volume.W, volume.H, volume.D)

	camera := LookAt(V(0, -5, 0), V(0, 0, 0), V(0, 0, -1), 35)
	sampler := DefaultSampler{4, 4}
	IterativeRender("out%03d.png", 1000, &scene, &camera, &sampler, 1600, 1600, -1)
}
