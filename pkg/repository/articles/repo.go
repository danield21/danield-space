package articles

import (
	"time"

	"github.com/danield21/danield-space/pkg/repository"
	"github.com/danield21/danield-space/pkg/repository/categories"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const entity = "Articles"

//GetAll gets all articles written for this website.
func GetAll(c context.Context, limit int) (articles []Article, err error) {
	q := datastore.NewQuery(entity).Order("Publish").Limit(limit)
	_, err = q.GetAll(c, &articles)
	return
}

//GetAllByType gets all articles of the same type.
func GetAllByType(ctx context.Context, category string, limit int) (articles []Article, err error) {
	cat, key, err := categories.Get(ctx, category)
	if err != nil {
		return
	}
	q := datastore.NewQuery(entity).Filter("Publish <", time.Now()).Order("Publish").Ancestor(key).Limit(limit)
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

//GetMapKeyedByTypes gets a map of articles with the key being the article type.
//Map returns an array of article with the same type limited by Limit.
func GetMapKeyedByTypes(ctx context.Context, Limit int) (articleMap map[string][]Article, err error) {
	articleMap = make(map[string][]Article)

	var types []string
	types, err = GetTypes(ctx)
	if err != nil {
		return
	}

	for _, t := range types {
		articles, aErr := GetAllByType(ctx, t, Limit)

		if aErr != nil {
			err = aErr
			return
		}

		articleMap[t] = articles
	}
	return
}

//GetTypes gets a list of article types that are in the database
func GetTypes(ctx context.Context) (types []string, err error) {
	var typesStruct []map[string]string
	q := datastore.NewQuery(entity).Project("Type").Distinct()
	_, err = q.GetAll(ctx, &types)

	for _, t := range typesStruct {
		types = append(types, t["Type"])
	}

	return
}

func Set(ctx context.Context, article Article) (err error) {
	oldArticle, key, dErr := Get(ctx, article.Category.Url, article.Url)

	if dErr != nil {
		key = datastore.NewIncompleteKey(ctx, entity, nil)
		article.DataElement = repository.WithNew("site")
	} else {
		article.DataElement = repository.WithOld(oldArticle.DataElement, "site")
	}

	_, err = datastore.Put(ctx, key, &article)

	return
}
