package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    *time.Time `json:"create_at"` //hiện time null json nếu không có
	UpdateAt    *time.Time `json:"update_at"`
}

//`id` int NOT NULL AUTO_INCREMENT,
//`title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
//`description` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci,
//`image` json DEFAULT NULL,
//`craete_at` datetime DEFAULT CURRENT_TIMESTAMP,
//`update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`status` enum('Doing','Done','Deleted') DEFAULT NULL,

func main() {
	fmt.Println("Hello World")

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Hello World",
		Description: "This is a test",
		CreateAt:    &now, // lấy con trỏ
		UpdateAt:    nil,  // nil là null
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}
