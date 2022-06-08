package main

import (
	"context"
	"fmt"
	pb "separategrpc/createproto"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	data := BlogItem{
		Id:       in.Id,
		Authorid: in.Authorid,
		Title:    in.Title,
		Content:  in.Content,
	}
	if err := Session.Query("insert into allblog(id,authorid,title,content)Values(?,?,?,?);", data.Id, data.Authorid, data.Title, data.Content).Exec(); err != nil {
		fmt.Println("error while create a new blog")
		fmt.Println(err)
	}
	return &pb.BlogId{
		Id: in.Id,
	}, nil
}
