package main

import (
	"encoding/json"
	"fmt"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"time"
)

var jsonjoniter = jsoniter.ConfigCompatibleWithStandardLibrary

type Data struct {
	Person struct {
		Id   string `json:"id"`
		Name struct {
			FullName   string `json:"fullName"`
			GivenName  string `json:"givenName"`
			FamilyName string `json:"familyName"`
		} `json:"name"`
		Email    string `json:"email"`
		Gender   string `json:"gender"`
		Location string `json:"location"`
		Geo      struct {
			City    string  `json:"city"`
			State   string  `json:"state"`
			Country string  `json:"country"`
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
		} `json:"geo"`
		Bio        string `json:"bio"`
		Site       string `json:"site"`
		Avatar     string `json:"avatar"`
		Employment struct {
			Name   string `json:"name"`
			Title  string `json:"title"`
			Domain string `json:"domain"`
		} `json:"employment"`
		Facebook struct {
			Handle string `json:"handle"`
		} `json:"facebook"`
		Github struct {
			Handle    string `json:"handle"`
			Id        int    `json:"id"`
			Avatar    string `json:"avatar"`
			Company   string `json:"company"`
			Blog      string `json:"blog"`
			Followers int    `json:"followers"`
			Following int    `json:"following"`
		} `json:"github"`
		Twitter struct {
			Handle    string      `json:"handle"`
			Id        int         `json:"id"`
			Bio       interface{} `json:"bio"`
			Followers int         `json:"followers"`
			Following int         `json:"following"`
			Statuses  int         `json:"statuses"`
			Favorites int         `json:"favorites"`
			Location  string      `json:"location"`
			Site      string      `json:"site"`
			Avatar    interface{} `json:"avatar"`
		} `json:"twitter"`
		Linkedin struct {
			Handle string `json:"handle"`
		} `json:"linkedin"`
		Googleplus struct {
			Handle interface{} `json:"handle"`
		} `json:"googleplus"`
		Angellist struct {
			Handle    string `json:"handle"`
			Id        int    `json:"id"`
			Bio       string `json:"bio"`
			Blog      string `json:"blog"`
			Site      string `json:"site"`
			Followers int    `json:"followers"`
			Avatar    string `json:"avatar"`
		} `json:"angellist"`
		Klout struct {
			Handle interface{} `json:"handle"`
			Score  interface{} `json:"score"`
		} `json:"klout"`
		Foursquare struct {
			Handle interface{} `json:"handle"`
		} `json:"foursquare"`
		Aboutme struct {
			Handle string      `json:"handle"`
			Bio    interface{} `json:"bio"`
			Avatar interface{} `json:"avatar"`
		} `json:"aboutme"`
		Gravatar struct {
			Handle  string        `json:"handle"`
			Urls    []interface{} `json:"urls"`
			Avatar  string        `json:"avatar"`
			Avatars []struct {
				Url  string `json:"url"`
				Type string `json:"type"`
			} `json:"avatars"`
		} `json:"gravatar"`
		Fuzzy bool `json:"fuzzy"`
	} `json:"person"`
	Company interface{} `json:"company"`
}

var TestData = []byte(`
    {
  "person": {
    "id": "d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "name": {
      "fullName": "Leonid Bugaev",
      "givenName": "Leonid",
      "familyName": "Bugaev"
    },
    "email": "leonsbox@gmail.com",
    "gender": "male",
    "location": "Saint Petersburg, Saint Petersburg, RU",
    "geo": {
      "city": "Saint Petersburg",
      "state": "Saint Petersburg",
      "country": "Russia",
      "lat": 59.9342802,
      "lng": 30.3350986
    },
    "bio": "Senior engineer at Granify.com",
    "site": "http://flickfaver.com",
    "avatar": "https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "employment": {
      "name": "www.latera.ru",
      "title": "Software Engineer",
      "domain": "gmail.com"
    },
    "facebook": {
      "handle": "leonid.bugaev"
    },
    "github": {
      "handle": "buger",
      "id": 14009,
      "avatar": "https://avatars.githubusercontent.com/u/14009?v=3",
      "company": "Granify",
      "blog": "http://leonsbox.com",
      "followers": 95,
      "following": 10
    },
    "twitter": {
      "handle": "flickfaver",
      "id": 77004410,
      "bio": null,
      "followers": 2,
      "following": 1,
      "statuses": 5,
      "favorites": 0,
      "location": "",
      "site": "http://flickfaver.com",
      "avatar": null
    },
    "linkedin": {
      "handle": "in/leonidbugaev"
    },
    "googleplus": {
      "handle": null
    },
    "angellist": {
      "handle": "leonid-bugaev",
      "id": 61541,
      "bio": "Senior engineer at Granify.com",
      "blog": "http://buger.github.com",
      "site": "http://buger.github.com",
      "followers": 41,
      "avatar": "https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"
    },
    "klout": {
      "handle": null,
      "score": null
    },
    "foursquare": {
      "handle": null
    },
    "aboutme": {
      "handle": "leonid.bugaev",
      "bio": null,
      "avatar": null
    },
    "gravatar": {
      "handle": "buger",
      "urls": [
      ],
      "avatar": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
      "avatars": [
        {
          "url": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
          "type": "thumbnail"
        }
      ]
    },
    "fuzzy": false
  },
  "company": null
}`)

func Funcjsoniter(in []byte) *Data {
	//start := time.Now()
	var data Data
	for i := 0; i < 1; i++ {
		if err := jsonjoniter.Unmarshal(in, &data); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
	//ServiceLatencyLogger(start)
	//fmt.Println(data)
	return &data
}

func Funcjson(in []byte) *Data {
	//start := time.Now()
	var data Data
	for i := 0; i < 1; i++ {
		if err := json.Unmarshal(in, &data); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
	//ServiceLatencyLogger(start)
	//fmt.Println(data)
	return &data
}

func Funcgojson(in []byte) *Data {
	//start := time.Now()
	var data Data
	for i := 0; i < 1; i++ {
		if err := gojson.Unmarshal(in, &data); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
	//ServiceLatencyLogger(start)
	//fmt.Println(data)
	return &data
}

func ServiceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}
