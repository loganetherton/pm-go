package utils

import (
	"fmt"
	"github.com/loganetherton/pm-go/cmd/blog"
	"github.com/loganetherton/pm-go/cmd/cache"
)

// GenericsWithStructs is an example of how to use generics with structs
func GenericsWithStructs() {
	category := blog.Category{
		ID:   1,
		Name: "Blah",
	}

	categoryCache := cache.New[blog.Category]()
	categoryCache.Set(category.ID, category)

	categories := make([]blog.Category, 0)
	categories = append(categories, category)
	post := blog.Post{
		ID:         1,
		Categories: categories,
	}
	postCache := cache.New[blog.Post]()
	postCache.Set(post.ID, post)

	if categoryCache.Get(1) == category && postCache.Get(1).ID == post.ID {
		fmt.Println("Generics with structs works")
	}
}
