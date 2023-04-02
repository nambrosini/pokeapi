package models

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"strconv"
	"time"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b Book) Create() {
	DB.Create(&b)
	b.createRedis()
}

func (b Book) createRedis() {
	resB, _ := json.Marshal(b)
	RDB.Set(strconv.Itoa(int(b.ID)), resB, time.Hour)
}

func GetBook(id string) (Book, error) {
	var book Book

	cmd := RDB.Get(id)
	cmdb, err := cmd.Bytes()

	if err == nil {
		b := bytes.NewReader(cmdb)

		if err := gob.NewDecoder(b).Decode(&book); err != nil {
			return Book{}, err
		}
	} else {
		if err := DB.Where("id = ?", id).First(&book).Error; err != nil {
			return Book{}, nil
		}
		book.createRedis()
	}

	return book, nil
}

func (b Book) UpdateBook(input UpdateBookInput) {
	updatedBook := Book{Title: input.Title, Author: input.Author}
	DB.Model(&b).Updates(updatedBook)
	b.createRedis()
}

func (b Book) DeleteBook() {
	DB.Delete(&b)
	RDB.Del(strconv.Itoa(int(b.ID)))
}
