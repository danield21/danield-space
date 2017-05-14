package articles

import (
	"time"

	"github.com/danield21/danield-space/server/models"
	"github.com/danield21/danield-space/server/repository/categories"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const entity = "Articles"

//Get gets a single article with the same type and key.
//Returns a error if there is no match.
func Get(ctx context.Context, cat *categories.Category, url string) (*Article, error) {
	var articles []*Article
	var keys []*datastore.Key
	var err error

	if cat == nil {
		return nil, categories.ErrNilCategory
	} else if cat.Key == nil {
		cat, err = categories.Get(ctx, cat.URL)
		if err != nil {
			return nil, err
		}
	}

	q := datastore.NewQuery(entity).Filter("URL =", url).Ancestor(cat.Key).Limit(1)

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
func GetAll(ctx context.Context, limit int) ([]*Article, error) {
	var (
		cat      categories.Category
		articles []*Article
	)

	q := datastore.NewQuery(entity).Order("PublishDate").Limit(limit)
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
func GetAllByCategory(ctx context.Context, cat *categories.Category, limit int) ([]*Article, error) {
	var (
		err      error
		articles []*Article
	)

	if cat == nil {
		return nil, categories.ErrNilCategory
	} else if cat.Key == nil {
		cat, err = categories.Get(ctx, cat.URL)
		if err != nil {
			return nil, err
		}
	}

	q := datastore.NewQuery(entity).Filter("PublishDate <", time.Now()).Order("PublishDate").Ancestor(cat.Key).Limit(limit)
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
func GetMapKeyedByCategory(ctx context.Context, Limit int) (map[*categories.Category][]*Article, error) {
	articleMap := make(map[*categories.Category][]*Article)

	cats, err := categories.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, cat := range cats {
		articles, err := GetAllByCategory(ctx, cat, Limit)

		if err != nil || len(articles) == 0 {
			continue
		}

		articleMap[cat] = articles
	}
	return articleMap, nil
}

func Set(ctx context.Context, article *Article) error {
	if article == nil {
		return ErrNilArticle
	}

	oldArticle, err := Get(ctx, article.Category, article.URL)

	if err != nil {
		cat, _ := categories.Get(ctx, article.Category.URL)
		article.DataElement = models.WithNew(models.WithPerson(ctx))
		article.Key = datastore.NewIncompleteKey(ctx, entity, cat.Key)
	} else {
		article.DataElement = models.WithOld(models.WithPerson(ctx), oldArticle.DataElement)
	}

	article.Key, err = datastore.Put(ctx, article.Key, article)

	return err
}
