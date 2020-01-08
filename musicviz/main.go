package main

import (
	"fmt"
	"io"
	"log"
	"os"

	gomp3 "github.com/hajimehoshi/go-mp3"
)

func main() {
	//For finishing REST : https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
	// server := &server{}
	// http.Handle("/", server)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	token := authenticate()
	getSongInformation(&token, "Toro y moi", "Freelance")
	getAudioAnalysis(&token, "spotify:track:2HsKkeVWys5Ts20z3e5lT0")
}

/**
Gets the bytes from the supplied file using gomp3
*/
func getBytes(path string) []byte {
	file, er := os.Open(path)
	if er != nil {
		log.Fatal("Could not open file")

	}
	// streamer, format, err := mp3.Decode(file)
	decodedFile, _ := gomp3.NewDecoder(file)
	//Audio will be backwards in decodedFile struct
	lengthOfFile := decodedFile.Length()
	data := make([]byte, lengthOfFile)
	for {
		n, err := decodedFile.Read(data)
		if n != 0 {
			fmt.Println(n)
		}
		if err != io.EOF {
			fmt.Println(err)
			break
		}
	}
	for _, theByte := range data {
		if theByte != 0 {
			fmt.Println("A BYTE HAS BEEN FOUND")
		}
	}

	return data
}
