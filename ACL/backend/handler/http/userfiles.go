package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	//"fmt"

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/ACL/backend/handler"
	"github.com/pucsd2020-pp/ACL/backend/model"
	"github.com/pucsd2020-pp/ACL/backend/repository"
	"github.com/pucsd2020-pp/ACL/backend/repository/userfiles"
)

type UserFiles struct {
	handler.HTTPHandler
	repo repository.UFRepository
	repo1 repository.ARepository
}

func NewUserFilesHandler(conn *sql.DB) *UserFiles {
	return &UserFiles{
		repo: userfiles.NewUserFilesRepository(conn),
		repo1: userfiles.NewUserFilesRepository(conn),
	}
}

func (userfiles *UserFiles) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userfiles/{id}", Func: userfiles.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userfiles", Func: userfiles.Create},
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "userfiles/{id}", Func: userfiles.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userfiles/{id}/{fid}", Func: userfiles.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userfiles", Func: userfiles.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userfiles/{id}/{fid}", Func: userfiles.GetByID},
		
	}
}


func (userfiles *UserFiles) GetByID(w http.ResponseWriter, r *http.Request) {
	//var usr interface{}
	var res int64
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	fid, err1 := strconv.ParseInt(chi.URLParam(r, "fid"), 10, 64)
	for {
		if nil != err {
			break
		}
		if nil != err1 {
			break
		}
		res,err = userfiles.repo1.IsAccess(r.Context(),id,fid)
		break
	}
	handler.WriteJSONResponse(w, r,res, http.StatusOK, err)
}

func (userfiles *UserFiles) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserFiles
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = userfiles.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
/*
func (userfiles *UserFiles) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserFiles{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in userfiles id for tracking update
		usr.UpdatedBy = 0

		iUsr, err = userfiles.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserFiles)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
*/
func (userfiles *UserFiles) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	//var res interface {}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	fid, err1 := strconv.ParseInt(chi.URLParam(r, "fid"), 10, 64)
	for {
		if nil != err {
			break
		}
		if nil != err1 {
			break
		}
		err = userfiles.repo.Delete2(r.Context(), id,fid)
		if nil != err {
			break
		}
		payload = " file permission delete from user successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}
func (userfiles *UserFiles) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := userfiles.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}





