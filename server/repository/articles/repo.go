package articles

import (
	"time"

	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/categories"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const entity = "Articles"

//GetAll gets all articles written for this website.
func GetAll(ctx context.Context, limit int) (articles []Article, err error) {
	var keys []*datastore.Key
	q := datastore.NewQuery(entity).Order("PublishDate").Limit(limit)
	keys, err = q.GetAll(ctx, &articles)

	if err != nil {
		return
	}

	for i, key := range keys {
		datastore.Get(ctx, key.Parent(), &articles[i].Category)
	}

	return
}

//GetAllByCategory gets all articles of the same category.
func GetAllByCategory(ctx context.Context, category string, limit int) (articles []Article, err error) {
	cat, key, err := categories.Get(ctx, category)
	if err != nil {
		return
	}
	q := datastore.NewQuery(entity).Filter("PublishDate <", time.Now()).Order("PublishDate").Ancestor(key).Limit(limit)
	_, err = q.GetAll(ctx, &articles)
	if err != nil {
		return
	}
	for i := range articles {
		articles[i].Category = cat
	}
	return
}

//Get gets a single article with the same type and key.
//Returns a error if there is no match.
func Get(ctx context.Context, category, url string) (article Article, key *datastore.Key, err error) {
	var articles []Article
	var keys []*datastore.Key

	cat, catkey, err := categories.Get(ctx, category)

	q := datastore.NewQuery(entity).Filter("Url =", url).Ancestor(catkey).Limit(1)

	keys, err = q.GetAll(ctx, &articles)
	if err != nil {
		return
	}

	if len(articles) == 0 {
		err = ErrNoArticle
		return
	}

	key = keys[0]
	article = articles[0]
	article.Category = cat

	return
}

//GetMapKeyedByCategory gets a map of articles with the key being the article type.
//Map returns an array of article with the same type limited by Limit.
func GetMapKeyedByCategory(ctx context.Context, Limit int) (articleMap map[categories.Category][]Article, err error) {
	articleMap = make(map[categories.Category][]Article)

	var cats []categories.Category
	cats, err = categories.GetAll(ctx)
	if err != nil {
		return
	}

	for _, cat := range cats {
		articles, aErr := GetAllByCategory(ctx, cat.Url, Limit)

		if aErr != nil {
			err = aErr
			return
		}

		if len(articles) == 0 {
			continue
		}

		articleMap[cat] = articles
	}
	return
}

func Set(ctx context.Context, article Article) (err error) {
	oldArticle, key, dErr := Get(ctx, article.Category.Url, article.Url)

	if dErr != nil {
		_, catKey, _ := categories.Get(ctx, article.Category.Url)
		key = datastore.NewIncompleteKey(ctx, entity, catKey)
		article.DataElement = repository.WithNew("site")
	} else {
		article.DataElement = repository.WithOld(oldArticle.DataElement, "site")
	}

	_, err = datastore.Put(ctx, key, &article)

	return
}
