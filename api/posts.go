package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Post struct {
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Blog        string     `json:"blog"`
	BlogID      int64      `json:"blog_id"`
	Published   string     `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
}

// Link        string     `json:"link"`
// Hash string `json:"hash"`
// Saved       bool       `json:"saved"`

type Posts struct {
	Posts []Post `json:"posts"`
}

type PostRequest struct {
	Blogs []int `json:"blogs"`
}

func GetPostsByBlog(blogID int) Posts {
	pr := PostRequest{Blogs: []int{blogID}}

	return GetPosts(pr)
}

func GetPosts(reqPost PostRequest) Posts {
	rJSON, err := json.Marshal(reqPost)
	if err != nil {
		panic(err)
	}
	// fmt.Println(rJSON)
	client := &http.Client{}

	payload := strings.NewReader(string(rJSON))

	req, err := http.NewRequest("GET", "https://api.dev-blogs.tech/api/posts", payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("blogio-key", "LALA")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		panic(err)
	}
	return posts
}
