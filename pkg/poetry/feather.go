package poetry

import (
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
)

type Feather struct {
	Content  string `json:"content"`
	Origin   string `json:"origin"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func HttpFeather() (*Feather, error) {
	var feather Feather
	client := &http.Client{}
	url := "https://v1.jinrishici.com/all.json"
	req, err := http.NewRequest("GET", url, nil) //GET大写
	if err != nil {
		log.Fatal(err)
	}
	rep, err := client.Do(req) //发起请求
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//使用mapstructure解析到结构体中
	if err := mapstructure.Decode(data, &feather); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &feather, nil
}
