package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Music struct {
	  Id int `json:"id"` // 曲のid
    Title string `json:"Title"` // 曲のタイトル
    Desc string `json:"desc"` // 曲の説明や，YouTubeなりApple MusicなりのURLなど
    Content []int `json:"content"` // メロディーライン（形式は要検討）
}

var Musics []Music

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello, world!\n")
}

func returnAllMusics(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Musics)
}


func main() {
	Musics = []Music{
      Music{Id: 0, Title: "いつもの曲", Desc: "https://music.apple.com/jp/album/tema-obu-isutansutori/1437763647?i=1437763752&l=en", Content: []int{57, 62, 64, 67, 64, 62, 59, 60, 59, 55}},
      Music{Id: 1, Title: "夏影", Desc: "https://music.apple.com/jp/album/natukage/818439884?i=818439888&l=en", Content: []int{60, 55, 60, 62, 64, 67, 62, 62, 64, 60, 62, 64}},
  }

  http.HandleFunc("/", hello)
  http.HandleFunc("/all", returnAllMusics)
  log.Fatal(http.ListenAndServe(":10000", nil))
}
