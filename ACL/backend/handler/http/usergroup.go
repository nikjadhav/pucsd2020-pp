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
	"github.com/pucsd2020-pp/ACL/backend/repository/usergroup"

)

type UserGroup struct {
	handler.HTTPHandler
	repo repository.UGRepository
}

func NewUserGroupHandler(conn *sql.DB) *UserGroup {
	return &UserGroup{
		repo: usergroup.NewUserGroupRepository(conn),
	}
}

func (usergroup *UserGroup) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup", Func: usergroup.GetByIDGroup},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "usergroup", Func: usergroup.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup/{gid}", Func: usergroup.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "usergroup/{id}/{gid}", Func: usergroup.Delete},
	}
}


func (usergroup *UserGroup) Create(w http.ResponseWriter, r *http.Request) {
	var grp model.UserGroup
	err := json.NewDecoder(r.Body).Decode(&grp)
	for { 
		if nil != err {
			break
		}

		err = usergroup.repo.AddUserGroup(r.Context(),grp)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (usergroup *UserGroup) GetByID(w http.ResponseWriter, r *http.Request) {

	gid, err := strconv.ParseInt(chi.URLParam(r, "gid"), 10, 64)
	grp, err := usergroup.repo.GetUsersByGroup(r.Context(), gid)
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (usergroup *UserGroup) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id,err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	gid,err1:=  strconv.ParseInt(chi.URLParam(r, "gid"), 10, 64)
	for {
		if nil != err {
			break
		}
		if nil != err1 {
			break
		}
		err = usergroup.repo.Delete2(r.Context(),id,gid)
		if nil != err {
			break
		}
		payload = "user deleted from group successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

