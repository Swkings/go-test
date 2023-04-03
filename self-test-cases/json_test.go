package test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type Post struct { //带结构标签，反引号来包围字符串
	Id      int       `json:"id"`
	Content string    `json:"content"`
	Author  Author    `json:"author"`
	Comment []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func readFromJsonFile(path string) {
	jsonFile, err := os.Open("path")
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading json file")
		return
	}
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}

func jsonFileToStruct(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF { //当所有的Json文件的数据被解码完毕以后，Decode方法会返回一个EOF，这个时候退出循环
			break
		}
		if err != nil {
			fmt.Println("解码失败")
			return
		}
		fmt.Printf("json: %#v", post)
	}
}

func writeStructDataToJsonFile(path string) {
	post := Post{
		Id:      1,
		Content: "努力做一个技术达人",
		Author: Author{
			Id:   2,
			Name: "商鞅",
		},
		Comment: []Comment{
			Comment{
				Id:      3,
				Content: "this is my hard day",
				Author:  "小不点",
			},
			Comment{
				Id:      4,
				Content: "good good study",
				Author:  "day day up",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	err = ioutil.WriteFile(path, output, 0644)
	if err != nil {
		fmt.Println("写文件失败")
		return
	}
}

func TestJson(t *testing.T) {
	writeStructDataToJsonFile("out.json")
	jsonFileToStruct("out.json")
}
