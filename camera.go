package main

/*
import (
	"errors"
	gphoto "github.com/micahwedemeyer/gphoto2go"
)

var (
	camera *gphoto.Camera
)

func InitCamera() error {
	camera = new(gphoto.Camera)
	err := camera.Init()
	if err < 0 {
		return errors.New(gphoto.CameraResultToString(err))
	}
	return nil
}

func TakePhoto() error {
	err := camera.TriggerCapture()
	if err < 0 {
		return errors.New(gphoto.CameraResultToString(err))
	}
	return nil
}
*/

import (
	"fmt"
	"log"
	"os"
)

func CapturePhoto(basepath, path, slug string) error {
	os.Mkdir(basepath+string(os.PathSeparator)+path, 0755)
	out, err := RunWithTimeout([]string{
		"/usr/bin/gphoto2",
		"--capture-image-and-download",
		fmt.Sprintf("--filename='%s/%s/%s.%s'", basepath, path, slug, *Extension),
	}, 10)
	log.Printf("CapturePhoto(): %s", out)
	return err
}
