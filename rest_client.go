package restclientlib

type RestClient interface {
	Post()
	Get()
	Put()
	Patch()
	Delete()
}
