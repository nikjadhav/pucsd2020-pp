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
	"github.com/pucsd2020-pp/ACL/backend/repository/groupfiles"
)

type GroupFiles struct {
	handler.HTTPHandler
	repo repository.UFRepository
}

func NewGroupFilesHandler(conn *sql.DB) *GroupFiles {
	return &GroupFiles{
		repo: groupfiles.NewGroupFilesRepository(conn),
	}
}

func (groupfiles *GroupFiles) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groupfiles/{id}", Func: groupfiles.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "groupfiles", Func: groupfiles.Create},
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "groupfiles/{id}", Func: groupfiles.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "groupfiles/{gid}/{fid}", Func: groupfiles.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groupfiles", Func: groupfiles.GetAll},
		
	}
}

/*
func (groupfiles *GroupFiles) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = groupfiles.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
*/
func (groupfiles *GroupFiles) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.GroupFiles
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = groupfiles.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
/*
func (groupfiles *GroupFiles) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.GroupFiles{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in groupfiles id for tracking update
		usr.UpdatedBy = 0

		iUsr, err = groupfiles.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.GroupFiles)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
*/
func (groupfiles *GroupFiles) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	gid, err := strconv.ParseInt(chi.URLParam(r, "gid"), 10, 64)
	fid, err1 := strconv.ParseInt(chi.URLParam(r, "fid"), 10, 64)
	for {
		if nil != err {
			break
		}
		if nil != err1 {
			break
		}
		err = groupfiles.repo.Delete2(r.Context(), gid,fid)
		if nil != err {
			break
		}
		payload = "file permission delete from group successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (groupfiles *GroupFiles) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := groupfiles.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}





