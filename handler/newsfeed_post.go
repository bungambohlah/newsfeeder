package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newsfeeder/platform/newsfeed"
	"github.com/aidarkhanov/nanoid"
)

type newfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Id: nanoid.New(),
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
