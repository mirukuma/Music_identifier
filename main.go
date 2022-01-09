package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
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

	var music Music
	json.Unmarshal(reqBody, &music)

	Musics = append(Musics, music)

	json.NewEncoder(w).Encode(music)
}

func deleteMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  key := vars["id"]

  for index, music := range Musics {
  	if music.Id == key {
  		Musics = append(Musics[:index], Musics[index+1:]...)
  	}
  }
}

func updateMusic(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var new_music Music
	json.Unmarshal(reqBody, &new_music)

	vars := mux.Vars(r)
  key := vars["id"]

  for index, music := range Musics {
  	if music.Id == key {
  		Musics[index] = new_music
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
  myRouter.HandleFunc("/music/{id}", updateMusic).Methods("PUT")
  myRouter.HandleFunc("/music/{id}", returnSingleMusic)


  db, err := sql.Open("mysql", "root:TestPassword1!@/goods")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close() // 関数がリターンする直前に呼び出される

  rows, err := db.Query("SELECT * FROM goods.musics;") // 
  if err != nil {
    panic(err.Error())
  }

  columns, err := rows.Columns() // カラム名を取得
  if err != nil {
    panic(err.Error())
  }

  values := make([]sql.RawBytes, len(columns))

  scanArgs := make([]interface{}, len(values))
  for i := range values {
    scanArgs[i] = &values[i]
  }

  for rows.Next() {
    err = rows.Scan(scanArgs...)
    if err != nil {
      panic(err.Error())
    }

    var value string
    for i, col := range values {
      // Here we can check if the value is nil (NULL value)
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      fmt.Println(columns[i], ": ", value)
    }
    fmt.Println("-----------------------------------")
  }

  log.Fatal(http.ListenAndServe(":10000", myRouter))
}


/*

CREATE TABLE musics(ID int AUTO_INCREMENT NOT NULL PRIMARY KEY, TITLE varchar(100),   DESCRIPTION varchar(100),  CONTENT varchar(100));
INSERT INTO goods.musics(TITLE, DESCRIPTION, CONTENT) value(‘Koizakura’, ‘favorite hentai game song in Apple Music’, ‘form of melody is undefined.’);

*/


