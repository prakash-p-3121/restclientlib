package restclientlib

import "github.com/prakash-p-3121/restclientlib/impl"

func NewRestClient() RestClient {
	return &impl.RestClientImpl{}
}
