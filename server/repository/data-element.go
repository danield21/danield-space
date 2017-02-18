package repository

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

//DataElement contains basic informtion on data that has been retrieved by a database
type DataElement struct {
	Key        *datastore.Key `datastore:"-"`
	CreatedOn  time.Time
	CreatedBy  string
	ModifiedOn time.Time
	ModifiedBy string
}

func WithOld(changer string, oldE DataElement) DataElement {
	return DataElement{
		Key:        oldE.Key,
		CreatedBy:  oldE.CreatedBy,
		CreatedOn:  oldE.CreatedOn,
		ModifiedBy: changer,
		ModifiedOn: time.Now(),
	}
}

func WithNew(creator string) DataElement {
	return DataElement{
		Key:        nil,
		CreatedBy:  creator,
		CreatedOn:  time.Now(),
		ModifiedBy: creator,
		ModifiedOn: time.Now(),
	}
}

type contextKey string

const personKey = contextKey("person")
const defaultPerson = "unknown"

func WithPerson(ctx context.Context) string {
	if iPerson := ctx.Value(personKey); iPerson != nil {
		return defaultPerson
	} else if sPerson, ok := iPerson.(string); !ok {
		return defaultPerson
	} else {
		return sPerson
	}
}

func StorePerson(ctx context.Context, person string) context.Context {
	return context.WithValue(ctx, personKey, person)
}
