package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/challenges/night_challenge/rest-api/handler"
	"github.com/pucsd2020-pp/challenges/night_challenge/rest-api/model"
	"github.com/pucsd2020-pp/challenges/night_challenge/rest-api/repository"
	"github.com/pucsd2020-pp/challenges/night_challenge/rest-api/repository/project"
)

type Project struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewProjectHandler(conn *sql.DB) *Project {
	return &Project{
		repo: project.NewProjectRepository(conn),
	}
}

func (project *Project) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "project/{id}", Func: project.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "project", Func: project.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "project/{id}", Func: project.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "project/{id}", Func: project.Delete},
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user", Func: user.GetAll},
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user/count", Func: user.Count},
	}
}

func (project *Project) GetByID(w http.ResponseWriter, r *http.Request) {
	var sub interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		sub, err = project.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, sub, http.StatusOK, err)
}

func (project *Project) Create(w http.ResponseWriter, r *http.Request) {
	var sub model.Project
	err := json.NewDecoder(r.Body).Decode(&sub)
	for {
		if nil != err {
			break
		}

		_, err = project.repo.Create(r.Context(), sub)
		break
	}
	handler.WriteJSONResponse(w, r, sub, http.StatusOK, err)
}

func (project *Project) Update(w http.ResponseWriter, r *http.Request) {
	var isub interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	sub := model.Project{}
	err := json.NewDecoder(r.Body).Decode(&sub)
	for {
		if nil != err {
			break
		}
		sub.Id = id
		if nil != err {
			break
		}

		// set logged in user id for tracking update
		//sub.UpdatedBy = 0

		isub, err = project.repo.Update(r.Context(),sub)
		if nil != err {
			break
		}
		sub = isub.(model.Project)
		break
	}

	handler.WriteJSONResponse(w, r, sub, http.StatusOK, err)
}

func (project *Project) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = project.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Project deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}
/*
func (user *User) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := user.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (user *User) Count(w http.ResponseWriter, r *http.Request) {
        usrs, err := user.repo.Count(r.Context())
        handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}*/
