package main

import (
	"l/gin/news"
	"l/gin/router"
)

func main() {
	r := router.New()
	a := news.New()
	go a.CollectNews()
	r.Run(":3000")
}
