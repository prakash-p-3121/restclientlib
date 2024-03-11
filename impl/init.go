package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"net/http"
	"sync"
)

// Concurrent HashMap
var statusCodeToErrMap sync.Map

func init() {
	statusCodeToErrMap.Store(http.StatusBadRequest, errorlib.NewBadReqError)
	statusCodeToErrMap.Store(http.StatusConflict, errorlib.NewConflictError)
	statusCodeToErrMap.Store(http.StatusNotFound, errorlib.NewNotFoundError)
	statusCodeToErrMap.Store(http.StatusInternalServerError, errorlib.NewInternalServerError)
}
