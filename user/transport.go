package user

import (
	"net/http"
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"demo/log"
	"github.com/gorilla/mux"
)
func MakeHandler(bs Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(log.Logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	registerHandler := kithttp.NewServer(
		makeRegisterEndpoints(bs),
		decodeRequest,
		encodeResponse,
		opts...,
	)

	loginHandler := kithttp.NewServer(
		makeLoginEndpoints(bs),
		decodeRequest,
		encodeResponse,
		opts...,
	)

	changePasswordHandler := kithttp.NewServer(
		makeChangePasswordEndpoints(bs),
		decodeRequest,
		encodeResponse,
		opts...,
	)

	deleteHandler := kithttp.NewServer(
		makeDeleteEndpoints(bs),
		decodeRequest,
		encodeResponse,
		opts...,
	)

	r:=mux.NewRouter()
	r.Handle("/users/register",registerHandler).Methods("GET")
	r.Handle("/users/login",loginHandler).Methods("GET")
	r.Handle("/users/changepassword",changePasswordHandler).Methods("GET")
	r.Handle("/users/delete",deleteHandler).Methods("GET")
	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return &userRequest{
		username:r.FormValue("username"),
		password:r.FormValue("password"),
		newPass:r.FormValue("newpass"),
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	rsp := response.(*userResponse)
	content, err:=json.Marshal(rsp)
	if err!=nil {
		return err
	}
	w.Write(content)
	return nil
}
