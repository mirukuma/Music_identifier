package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Music struct {
	  Id string `json:"id"` // 曲のid
    Title string `json:"Title"` // 曲のタイトル
    Desc string `json:"desc"` // 曲の説明や，YouTubeなりApple MusicなりのURLなど
    Content []int `json:"content"` // メロディーライン（形式は要検討）
}

var Musics []Music

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello, world!\n")
}

func returnAllMusics(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Musics)
}

func returnSingleMusic(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
    key := vars["id"]

    for _, music := range Musics {
    	if music.Id == key {
    		fmt.Fprint(w, music)
    	}
    }
}

func createNewMusic(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprint(w, reqBody)

	var music Music
	json.Unmarshal(reqBody, &music)

	Musics = append(Musics, music)

	json.NewEncoder(w).Encode(music)
}

func deleteMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  key := vars["id"]
  fmt.Println("sex")

  for index, music := range Musics {
  	if music.Id == key {
  		Musics = append(Musics[:index], Musics[index+1:]...)
  	}
  }
}


func main() {
	Musics = []Music{
      Music{Id: "0", Title: "いつもの曲", Desc: "https://music.apple.com/jp/album/tema-obu-isutansutori/1437763647?i=1437763752&l=en", Content: []int{57, 62, 64, 67, 64, 62, 59, 60, 59, 55}},
      Music{Id: "1", Title: "夏影", Desc: "https://music.apple.com/jp/album/natukage/818439884?i=818439888&l=en", Content: []int{60, 55, 60, 62, 64, 67, 62, 62, 64, 60, 62, 64}},
  }

  myRouter := mux.NewRouter()

  myRouter.HandleFunc("/", hello)
  myRouter.HandleFunc("/all", returnAllMusics)

  myRouter.HandleFunc("/music", createNewMusic).Methods("POST")

  myRouter.HandleFunc("/music/{id}", deleteMusic).Methods("DELETE")
  myRouter.HandleFunc("/music/{id}", returnSingleMusic)
  
  log.Fatal(http.ListenAndServe(":10000", myRouter))
}



