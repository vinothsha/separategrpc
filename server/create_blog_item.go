package main

import (
	pb "separategrpc/createproto"
)

type BlogItem struct {
	Id       string `json:"id"`
	Authorid string `json:"authorid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func JsonToText(p *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       p.Id,
		Authorid: p.Authorid,
		Title:    p.Title,
		Content:  p.Content,
	}
}
