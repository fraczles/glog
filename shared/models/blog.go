package models

import (
	"time"
)

type Post struct {
	ID      int    `jsonapi:"primary,post" db:"post"`
	Title   string `jsonapi:"attr,title"`
	Body    string `jsonapi:"attr,body"`
	Created int64  "jsonapi:attr,created"
}

func NewPost(title, body string) Post {
	return Post{
		Title:   title,
		Body:    body,
		Created: time.Now().UnixNano(),
	}
}
