package http

import (
//	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/ACL/backend/handler"
	"github.com/pucsd2020-pp/ACL/backend/model"
	"github.com/pucsd2020-pp/ACL/backend/repository"
	"github.com/pucsd2020-pp/ACL/backend/repository/group"

)

type Group struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewGroupHandler(conn *sql.DB) *Group {
	return &Group{
		repo: group.NewGroupRepository(conn),
	}
}

func (group *Group) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "group/{id}", Func: group.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "group", Func: group.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "group", Func: group.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "group/{gid}", Func: group.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "group/{id}", Func: group.Update},
	}
}

func (group *Group) GetByID(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		grp, err = group.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (group *Group) Create(w http.ResponseWriter, r *http.Request) {
	var grp model.Group
	err := json.NewDecoder(r.Body).Decode(&grp)
	for { 
		if nil != err {
			break
		}

		_, err = group.repo.Create(r.Context(),grp)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (group *Group) GetAll(w http.ResponseWriter, r *http.Request) {
	grps, err := group.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, grps, http.StatusOK, err)
}

func (group *Group) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	//var grp interface{}
	gid, err := strconv.ParseInt(chi.URLParam(r, "gid"), 10, 64)
	for {
		if nil != err {
			break
		}

		_,err = group.repo.Delete(r.Context(), gid)
		if nil != err {
			break
		}
		payload = "Group deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}
func (group *Group) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.Group{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}
		iUsr, err = group.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.Group)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

