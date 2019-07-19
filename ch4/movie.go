package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//json序列化与反序列化

func main() {
 	movies:=[]Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

 	//序列化
 	//json.MarshalIndent带缩进序列化
 	//json.Marshal不带缩进序列化
 	data,err:=json.MarshalIndent(movies,"","	")
 	if err!=nil{
 		log.Fatalf("JSON marshaling failed: %s\n",err)
	}
 	fmt.Printf("%s\n",data)

 	//再将data反序列化为另外一个对象
 	var titles []struct{Title string}
 	if err:=json.Unmarshal(data,&titles);err!=nil{
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
 	fmt.Println(titles)
}

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"` //omitempty忽略零值
	Actors []string
}
