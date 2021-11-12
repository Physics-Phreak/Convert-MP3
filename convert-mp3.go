package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func listDir() []fs.FileInfo {
	files, err := ioutil.ReadDir(".")

	var songList []fs.FileInfo

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if !file.IsDir() {
			switch filepath.Ext(file.Name()) {
			case ".mp3":
				songList = append(songList, file)

			case ".wav":
				songList = append(songList, file)

			case ".flac":
				songList = append(songList, file)

			case ".alac":
				songList = append(songList, file)

			case ".m4a":
				songList = append(songList, file)

			default:
				fmt.Printf("%v is not a music file \n", file.Name())
			}
		}

	}

	if len(songList) == 0 {

		fmt.Println("No Music Files in this folder")

	}

	return songList
}

func convertToMp3(file fs.FileInfo) {

	os.Mkdir("MP3", 0755)

	name := file.Name()

	output := strings.TrimSuffix(name, filepath.Ext(name))
	output += ".mp3"
	output = "./MP3/" + output

	command := exec.Command("ffmpeg", "-i", name, output)
	command.Run()

}

func main() {
	songList := listDir()

	for _, file := range songList {

		convertToMp3(file)

		fmt.Println("Converted File: ", file.Name())
	}

}
