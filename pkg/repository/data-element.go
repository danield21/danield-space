package repository

import "time"

//DataElement contains basic informtion on data that has been retrieved by a database
type DataElement struct {
	CreatedOn  time.Time
	CreatedBy  string
	ModifiedOn time.Time
	ModifiedBy string
}

func WithOld(oldE DataElement, changer string) DataElement {
	return DataElement{
		CreatedBy:  oldE.CreatedBy,
		CreatedOn:  oldE.CreatedOn,
		ModifiedBy: changer,
		ModifiedOn: time.Now(),
	}
}

func WithNew(creator string) DataElement {
	return DataElement{
		CreatedBy:  creator,
		CreatedOn:  time.Now(),
		ModifiedBy: creator,
		ModifiedOn: time.Now(),
	}
}
