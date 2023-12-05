package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

type Stop struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Z           float64 `json:"z"`
	Theta       float64 `json:"theta"`
	Visible     int     `json:"visible"`
	Lines       []int   `json:"lines"`
	Type        int     `json:"type"`
	Subtype     int     `json:"subtype"`
	Extend      string  `json:"extend"`
	Floor       int     `json:"floor,omitempty"`
}

type Line struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	IsCircle  int         `json:"is_circle"`
	Waypoints [][]float64 `json:"waypoints"`
}

type Map struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Image   struct {
		Url          string  `json:"url"`
		LatitudeMax  float64 `json:"latitude_max"`
		LatitudeMin  float64 `json:"latitude_min"`
		LongitudeMax float64 `json:"longitude_max"`
		LongitudeMin float64 `json:"longitude_min"`
	} `json:"image"`
	Lines []Line `json:"lines"`
	Stops []Stop `json:"stops"`
}

func TestReadJson(t *testing.T) {
	jsonFilePath := "/home/swk/WorkSpace/MapSpace/179_v7/map.json"
	mapJson := ReadJson[[]Map](jsonFilePath)
	fmt.Printf("%v\n", Array2String(FetchPoint(15, mapJson[0].Stops)))
}

func ReadJson[T any](path string) T {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload T
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func FetchPoint(num int, stops []Stop) [][]float64 {
	var res [][]float64
	for i := 0; i < num; i++ {
		res = append(res, []float64{stops[i].X, stops[i].Y, stops[i].Theta})
	}

	return res
}

func Array2String(arr [][]float64) string {
	var res []string
	for _, item := range arr {
		strList := func() []string {
			var r []string
			for _, num := range item {
				r = append(r, fmt.Sprint(num))
			}
			return r
		}()
		res = append(res, strings.Join(strList, ","))
	}

	for i, item := range res {
		if i != len(res)-1 {
			res[i] = fmt.Sprintf("\"car%v_init_position\": \"[%v]\",", i, item)
		} else {
			res[i] = fmt.Sprintf("\"car%v_init_position\": \"[%v]\"", i, item)
		}

	}

	return strings.Join(res, "\n")
}
