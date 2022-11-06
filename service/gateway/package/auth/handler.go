package auth

import (
	"github.com/elisfromkirov/service/service/gateway/package/gtw"
	"github.com/elisfromkirov/service/service/request_registry/package/util"
	"net/http"
	"net/url"
)

const (
	UserKey     string = "user"
	PasswordKey string = "password"
)

type Handler struct {
}

func NewHandler(config *gtw.Config) *Handler {
	return &Handler{}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	info := ParseQuery(request.URL.Query())

	access := Access{Access:  info.User != "user" && info.Password != "password"}

	buffer, err := MarshalAccess(&access)
	if err != nil {
		util.LogError(err)
		return
	}

	_, err = writer.Write([]byte(buffer))
	if err != nil {
		util.LogError(err)
		return
	}
}

func ParseQuery(values url.Values) AuthInfo {
	return AuthInfo{
		User: values.Get(UserKey),
		Password: values.Get(PasswordKey),
	}
}
