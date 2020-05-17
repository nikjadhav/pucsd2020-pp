package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/sessions"
	"fmt"
	"os"
	"errors"
)
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type IHTTPHandler interface {
	GetHTTPHandler() []*HTTPHandler
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}
type HTTPHandler struct {
	Authenticated bool
	Method        string
	Path          string
	Func          func(http.ResponseWriter, *http.Request)
}

type response struct {
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (hdlr *HTTPHandler) GetHTTPHandler() []HTTPHandler {
	return []HTTPHandler{}
}

func (hdlr *HTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	return
}

func WriteJSONResponse(w http.ResponseWriter,
	r *http.Request,
	payload interface{},
	code int,
	err error) {
	resp := &response{
		Status: code,
		Data:   payload,
	}

	if nil != err {
		resp.Message = err.Error()
	}

	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}
func SetSession(w http.ResponseWriter,
	r *http.Request,id int64){
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["session"] = id
	fmt.Println("sess",session.Values["session"])
	err = session.Save(r, w)
	fmt.Println("session saving issue",err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}
func GetSession(w http.ResponseWriter,
	r *http.Request,) (int,error){
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0,err
	}
	if (session.Values["session"]!=nil){

	return 1,nil
	}
	err = errors.New("Unauthorize Access")
	return 0,err
}

func ClearSession(w http.ResponseWriter,
	r *http.Request,){
		session, err := store.Get(r,"session-name")
		session.Options.MaxAge = -1
		err = session.Save(r, w)
		if err != nil {
			fmt.Println("Failed to delete session")
		}
		fmt.Println("Delete Session")
		
	}
	