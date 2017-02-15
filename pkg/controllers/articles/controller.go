package articles

import (
	"github.com/danield21/danield-space/pkg/controllers"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const entity = "Articles"

//GetAll gets all articles written for this website.
func GetAll(c context.Context, limit int) (articles []Article, err error) {
	q := datastore.NewQuery(entity).Order("ModifiedOn").Limit(limit)
	_, err = q.GetAll(c, &articles)
	return
}

//GetAllByType gets all articles of the same type.
func GetAllByType(c context.Context, Type string, limit int) (articles []Article, err error) {
	q := datastore.NewQuery(entity).Filter("Type =", Type).Order("ModifiedOn").Limit(limit)
	_, err = q.GetAll(c, &articles)
	return
}

//Get gets a single article with the same type and key.
//Returns a error if there is no match.
func Get(c context.Context, Type, Key string) (article Article, key *datastore.Key, err error) {
	var articles []Article
	var keys []*datastore.Key

	q := datastore.NewQuery(entity).Filter("Type =", Type).Filter("Key =", Key).Limit(1)

	keys, err = q.GetAll(c, &articles)
	if err != nil {
		return
	}

	if len(articles) == 0 {
		err = ErrNoArticle
		return
	}

	key = keys[0]
	article = articles[0]

	return
}

//GetMapKeyedByTypes gets a map of articles with the key being the article type.
//Map returns an array of article with the same type limited by Limit.
func GetMapKeyedByTypes(c context.Context, Limit int) (articleMap map[string][]Article, err error) {
	articleMap = make(map[string][]Article)

	var types []string
	types, err = GetTypes(c)
	if err != nil {
		return
	}

	for _, t := range types {
		articles, aErr := GetAllByType(c, t, Limit)

		if aErr != nil {
			err = aErr
			return
		}

		articleMap[t] = articles
	}
	return
}

//GetTypes gets a list of article types that are in the database
func GetTypes(c context.Context) (types []string, err error) {
	var typesStruct []map[string]string
	q := datastore.NewQuery(entity).Project("Type").Distinct()
	_, err = q.GetAll(c, &types)

	for _, t := range typesStruct {
		types = append(types, t["Type"])
	}

	return
}

func Set(c context.Context, article Article) (err error) {
	oldArticle, key, dErr := Get(c, article.Type, article.Key)

	if dErr != nil {
		key = datastore.NewIncompleteKey(c, entity, nil)
		article.DataElement = controllers.WithNew("site")
	} else {
		article.DataElement = controllers.WithOld(oldArticle.DataElement, "site")
	}

	_, err = datastore.Put(c, key, &article)

	return
}
