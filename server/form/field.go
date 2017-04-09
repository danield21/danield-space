package form

type Field struct {
	Error  error
	Values []string
}

func (f Field) Get() string {
	if len(f.Values) == 0 {
		return ""
	}
	return f.Values[0]
}

func (f Field) Has(value string) bool {
	for _, v := range f.Values {
		if v == value {
			return true
		}
	}
	return false
}
