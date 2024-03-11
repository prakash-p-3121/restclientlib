package restclientlib

import "github.com/prakash-p-3121/errorlib"

type RestClient interface {
	Post(url string, request, resultPtr interface{}) errorlib.AppError
	Get(url string, resultPtr interface{}) errorlib.AppError
	Put(url string, request, resultPtr interface{}) errorlib.AppError
	Patch(url string, request, resultPtr interface{}) errorlib.AppError
	Delete(url string, resultPtr interface{}) errorlib.AppError
}
