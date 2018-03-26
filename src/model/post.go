package model

import "time"

//文章
type Post struct {
	Content string `yaml:"content"`
	Title string `yaml:"title"`
	Path string `yaml:"path"`
	PubTime string `yaml:"pub_time"`
	Tags string `yaml:"tags"`
	PostTime string `yaml:"post_time"`
}

type Meta struct {
	Title string `yaml:"title"`
	Path string `yaml:"path"`
	PubTime time.Time `yaml:"pub_time"`
	Tags []string `yaml:"tags"`
	PostTime time.Time `yaml:"post_time"`
}
