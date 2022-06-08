package main

import (
	"context"
	"fmt"
	pb "separategrpc/readproto"
)

func jsonToBlog(data *BlogItem) *pb.Blogr {
	return &pb.Blogr{
		Id:       data.Id,
		Authorid: data.Authorid,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func (s *Server) ReadBlog(ctx context.Context, req *pb.Blogid) (*pb.Blogr, error) {
	var data BlogItem
	m := map[string]interface{}{}
	iter := Session.Query("select * from allblog where id=?", req.Id).Iter()
	for iter.MapScan(m) {
		data.Id = m["id"].(string)
		data.Authorid = m["authorid"].(string)
		data.Title = m["title"].(string)
		data.Content = m["content"].(string)
	}
	fmt.Println(data)
	return jsonToBlog(&data), nil
}