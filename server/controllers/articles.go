package controllers

import (
	"errors"
	"strings"
	"time"
)

//Article contains information about articles written on this website.
type Article struct {
	DataElement
	Type     string
	Key      string
	Title    string
	Abstract string
}

//Path returns the path for a article.
func (a Article) Path() string {
	return "/" + a.Type + "/" + a.Key
}

//Heading returns a heading for the article using the Type and Title.
func (a Article) Heading() string {
	return strings.Title(a.Type + ": " + a.Title)
}

//ArticleController is a controller for Article.
type ArticleController struct {
}

//GetAll gets all articles written for this website.
//TODO create database connection to get this.
func (c ArticleController) GetAll() []Article {
	return []Article{
		Article{
			DataElement: DataElement{
				UUID:       "1",
				CreatedBy:  "Daniel J Dominguez",
				CreatedOn:  time.Now(),
				ModifiedBy: "Daniel J Dominguez",
				ModifiedOn: time.Now(),
			},
			Type:     "blog",
			Key:      "ttc-web-development",
			Title:    "Things to Consider - Web Development",
			Abstract: "Brings up things to consider, like accessibility, progressive enchancements, and performance.",
		},
		Article{
			DataElement: DataElement{
				UUID:       "2",
				CreatedBy:  "Daniel J Dominguez",
				CreatedOn:  time.Now(),
				ModifiedBy: "Daniel J Dominguez",
				ModifiedOn: time.Now(),
			},
			Type:     "project",
			Key:      "this-website",
			Title:    "This Website",
			Abstract: "Built on the idea of progressive enchancements.",
		},
		Article{
			DataElement: DataElement{
				UUID:       "3",
				CreatedBy:  "Daniel J Dominguez",
				CreatedOn:  time.Now(),
				ModifiedBy: "Daniel J Dominguez",
				ModifiedOn: time.Now(),
			},
			Type:     "project",
			Key:      "swau-app",
			Title:    "SWAU App",
			Abstract: "First major project.",
		},
	}
}

//GetType gets all articles of the same type.
func (c ArticleController) GetType(Type string) []Article {
	var articles []Article

	for _, a := range c.GetAll() {
		if a.Type == Type {
			articles = append(articles, a)
		}
	}

	return articles
}

//Get gets a single article with the same type and key.
//Returns a error if there is no match.
func (c ArticleController) Get(Type, Key string) (article Article, err error) {
	err = errors.New("Unable to find article with type/key pair")

	for _, a := range c.GetAll() {
		if a.Type == Type && a.Key == Key {
			article = a
			err = nil
		}
	}

	return
}

//GetMapKeyedByTypes gets a map of articles with the key being the article type.
//Map returns an array of article with the same type limited by Limit.
func (c ArticleController) GetMapKeyedByTypes(Limit int) (articleMap map[string][]Article) {
	articleMap = make(map[string][]Article)
	for _, article := range c.GetAll() {
		if len(articleMap[article.Type]) <= Limit {
			articleMap[article.Type] = append(articleMap[article.Type], article)
		}
	}

	return
}
