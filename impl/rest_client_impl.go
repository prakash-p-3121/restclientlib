package impl

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/prakash-p-3121/errorlib"
	"strconv"
)

type RestClientImpl struct {
}

func (client *RestClientImpl) buildAppError(resp *resty.Response) errorlib.AppError {
	if resp.IsSuccess() {
		return nil
	}
	statusCode := resp.StatusCode()
	factoryFunc, ok := statusCodeToErrMap.Load(statusCode)
	if !ok {
		newErr := errors.New("unhandled httpStatus code =" + strconv.Itoa(statusCode))
		return errorlib.NewInternalServerError(newErr.Error())
	}
	appErrFactoryFunc, ok := factoryFunc.(func(string) errorlib.AppError)
	if !ok {
		newErr := errors.New("cannot type cast to Factory function func(string) errorlib.AppError")
		return errorlib.NewInternalServerError(newErr.Error())
	}

	return appErrFactoryFunc(string(resp.Body()))
}

func (client *RestClientImpl) Post(url string, request, resultPtr interface{}) errorlib.AppError {
	restyClient := resty.New()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(resultPtr).Post(url)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return client.buildAppError(resp)
}

func (client *RestClientImpl) Get(url string, resultPtr interface{}) errorlib.AppError {
	restyClient := resty.New()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(resultPtr).Get(url)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return client.buildAppError(resp)
}

func (client *RestClientImpl) Put(url string, request, resultPtr interface{}) errorlib.AppError {
	restyClient := resty.New()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(resultPtr).Put(url)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return client.buildAppError(resp)
}

func (client *RestClientImpl) Patch(url string, request, resultPtr interface{}) errorlib.AppError {
	restyClient := resty.New()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(resultPtr).Patch(url)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return client.buildAppError(resp)
}

func (client *RestClientImpl) Delete(url string, resultPtr interface{}) errorlib.AppError {
	restyClient := resty.New()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(resultPtr).Delete(url)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return client.buildAppError(resp)
}
