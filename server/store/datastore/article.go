package datastore

import (
	"context"
	"errors"
	"time"

	"github.com/danield21/danield-space/server/store"
	"google.golang.org/appengine/datastore"
)

const articleEntity = "Articles"

type Article struct {
	Category store.CategoryRepository
}

//ErrNoArticle appears when the article requested does not exist in the database
var ErrNoArticle = errors.New("unable to find article with category/url pair")

//ErrNilArticle appears when item parameter is nil
var ErrNilArticle = errors.New("err was nil")

//Get gets a single article with the same type and key.
//Returns a error if there is no match.
func (a Article) Get(ctx context.Context, cat *store.Category, url string) (*store.Article, error) {
	var articles []*store.Article
	var keys []*datastore.Key
	var err error

	if cat == nil {
		return nil, ErrNilCategory
	} else if cat.Key == nil {
		cat, err = a.Category.Get(ctx, cat.URL)
		if err != nil {
			return nil, err
		}
	}

	q := datastore.NewQuery(articleEntity).Filter("URL =", url).Ancestor(cat.Key).Limit(1)

	keys, err = q.GetAll(ctx, &articles)
	if err != nil {
		return nil, err
	}

	if len(articles) == 0 {
		return nil, ErrNoArticle
	}

	article := articles[0]
	article.Category = cat
	article.Key = keys[0]

	return article, nil
}

//GetAll gets all articles written for this website.
func (Article) GetAll(ctx context.Context, limit int) ([]*store.Article, error) {
	var (
		cat      store.Category
		articles []*store.Article
	)

	q := datastore.NewQuery(articleEntity).Order("PublishDate").Limit(limit)
	keys, err := q.GetAll(ctx, &articles)

	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		articles[i].Key = key
		err := datastore.Get(ctx, key.Parent(), &cat)
		if err != nil {
			return nil, err
		}
		articles[i].Category = &cat
	}

	return articles, nil
}

//GetAllByCategory gets all articles of the same category.
func (a Article) GetAllByCategory(ctx context.Context, cat *store.Category, limit int) ([]*store.Article, error) {
	var (
		err      error
		articles []*store.Article
	)

	if cat == nil {
		return nil, ErrNilCategory
	} else if cat.Key == nil {
		cat, err = a.Category.Get(ctx, cat.URL)
		if err != nil {
			return nil, err
		}
	}

	q := datastore.NewQuery(articleEntity).Filter("PublishDate <", time.Now()).Order("PublishDate").Ancestor(cat.Key).Limit(limit)
	keys, err := q.GetAll(ctx, &articles)
	if err != nil {
		return nil, err
	}
	for i := range articles {
		articles[i].Category = cat
		articles[i].Key = keys[i]
	}
	return articles, nil
}

//GetMapKeyedByCategory gets a map of articles with the key being the article type.
//Map returns an array of article with the same type limited by Limit.
func (a Article) GetMapKeyedByCategory(ctx context.Context, Limit int) (map[*store.Category][]*store.Article, error) {
	articleMap := make(map[*store.Category][]*store.Article)

	cats, err := a.Category.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, cat := range cats {
		articles, err := a.GetAllByCategory(ctx, cat, Limit)

		if err != nil || len(articles) == 0 {
			continue
		}

		articleMap[cat] = articles
	}
	return articleMap, nil
}

func (a Article) Set(ctx context.Context, article *store.Article) error {
	if article == nil {
		return ErrNilArticle
	}

	oldArticle, err := a.Get(ctx, article.Category, article.URL)

	if err != nil {
		cat, _ := a.Category.Get(ctx, article.Category.URL)
		article.DataElement = store.WithNew(store.WithPerson(ctx))
		article.Key = datastore.NewIncompleteKey(ctx, articleEntity, cat.Key)
	} else {
		article.DataElement = store.WithOld(store.WithPerson(ctx), oldArticle.DataElement)
	}

	article.Key, err = datastore.Put(ctx, article.Key, article)

	return err
}
