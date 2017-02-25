package account

//Default is the default account
//Be sure to change it
var Default = Account{Username: "Root", Hashword: nil, Super: true}

func init() {
	Default.Password([]byte("ThisIsAVerySimplePassword!"))
}
