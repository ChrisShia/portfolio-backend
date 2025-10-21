package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
)

func (app *App) codingSkills(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = ""
	payload.Data = codingSkills

	app.writeJSON(w, http.StatusOK, payload)
}

var codingSkills = []struct {
	Title string `json:"title"`
	Image string `json:"image"`
}{
	{"Golang", imageString64("project/images/skills/mygopher.png")},
	{"Java", imageString64("project/images/skills/java.png")},
	{"PostgreSQL", imageString64("project/images/skills/postgresql.png")},
	{"Redis", imageString64("project/images/skills/redis.png")},
	{"Docker", imageString64("project/images/skills/docker.png")},
	{"Nats", imageString64("project/images/skills/nats.png")},
	{"React", imageString64("project/images/skills/react.png")},
}

func imageString64(path string) string {
	img := openImage(path)
	imgBuf := new(bytes.Buffer)

	//TODO: abstract the encoding method

	_ = png.Encode(imgBuf, img)
	return base64.StdEncoding.EncodeToString(imgBuf.Bytes())
}

func openImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
