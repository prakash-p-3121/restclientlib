package restclientlib

import "github.com/prakash-p-3121/restclientlib/impl"

func NewRestClientFactory() RestClient {
	return &impl.RestClientImpl{}
}
