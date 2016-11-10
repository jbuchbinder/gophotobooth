package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

/*
func main() {
	// Testing
	err := InitCamera()
	if err != nil {
		panic(err)
	}
	err = TakePhoto()
	if err != nil {
		panic(err)
	}
	folders := camera.RListFolders("/")
	fmt.Printf("folders = %#v\n", folders)
	for _, folder := range folders {
		files, _ := camera.ListFiles(folder)
		fmt.Printf("files = %#v\n", files)
		for _, fileName := range files {
			cameraFileReader := camera.FileReader(folder, fileName)
			fmt.Printf("%s/%s\n", folder, fileName)
			fileWriter, err := os.Create("/tmp/" + fileName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			io.Copy(fileWriter, cameraFileReader)

			// Important, since there is memory used in the transfer that needs to be freed up
			cameraFileReader.Close()
		}
	}
}
*/

var (
	Extension   = flag.String("extension", "CR2", "File extension for output files")
	Debug       = flag.Bool("debug", false, "Debug mode")
	Port        = flag.Int("port", 8888, "HTTP port")
	StoragePath = flag.String("storage-path", "./storage", "Path to media storage")
)

func main() {
	flag.Parse()

	if !*Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Printf("Initializing web services")
	m := gin.New()
	m.Use(gin.Logger())
	m.Use(gin.Recovery())
	// Enable gzip compression
	m.Use(gzip.Gzip(gzip.DefaultCompression))

	initApi(m)

	log.Print("[static] Initializing with local unbundled resources")
	m.Use(static.Serve("/", static.LocalFile("ui", false)))
	m.StaticFile("/", "ui"+string(os.PathSeparator)+"index.html")

	log.Printf("Initializing on :%d", *Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *Port), m); err != nil {
		log.Fatal(err)
	}
}
