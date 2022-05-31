package prototype

type address struct {
	street string
	city   string
}

type person struct {
	name    string
	address *address
}
