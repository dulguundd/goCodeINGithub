package main

import (
	"fmt"
	"log"
	"time"
)

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}

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
