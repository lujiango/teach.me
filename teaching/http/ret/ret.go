package ret

import (
	"io"
	"net/http"
)

const (
	LOCATION_IS_EMPTY  = "{ret : 400100, msg : 'location is empty.'}"
	USER_IS_EXIST      = "{ret : 400200, msg : 'user has exist. '}"
	USER_CHECK_FAILED  = "{ret : 400201, msg : 'user check failed.'}"
	USER_TOKEN_NULL    = "{ret : 400202, msg : 'user token is null.'}"
	USER_TOKEN_ILLEGAL = "{ret : 400203, msg : 'user token is illegal.'}"
	USER_TOKEN_EXPIRE  = "{ret : 400204, msg : 'user token is expire.'}"
)

func HanderError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, err.Error())
}
