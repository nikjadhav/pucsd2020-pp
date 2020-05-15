package http
import (
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/ACL/backend/handler"
	"github.com/pucsd2020-pp/ACL/backend/model"
	"github.com/pucsd2020-pp/ACL/backend/repository"
	"github.com/pucsd2020-pp/ACL/backend/repository/files"

)

type Files struct {
	handler.HTTPHandler
	repo repository.FRepository
}

func NewFilesHandler(conn *sql.DB) *Files {
	return &Files{
		repo: files.NewFilesRepository(conn),
	}
}

func (files *Files) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "files", Func: files.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "files/{pid}", Func: files.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "files/{id}", Func: files.Delete},
	}
}
func (files *Files) Create(w http.ResponseWriter, r *http.Request) {
	var fil model.Files
	err := json.NewDecoder(r.Body).Decode(&fil)
	for { 
		if nil != err {
			break
		}
		_, err = files.repo.Create(r.Context(),fil)
		break
	}
	handler.WriteJSONResponse(w, r, fil, http.StatusOK,err)
}

func (files *Files) GetByID(w http.ResponseWriter, r *http.Request) {
	var fil interface{}
	pid, err := strconv.ParseInt(chi.URLParam(r, "pid"), 10, 64)
	for {
		if nil != err {
			break
		}

		fil, err = files.repo.GetFilesByParent(r.Context(), pid)
		break
	}
	fmt.Println("fil",fil)
	handler.WriteJSONResponse(w, r, fil, http.StatusOK, err)
}
func (files *Files) Delete(w http.ResponseWriter, r *http.Request) {
		var payload string
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		for {
			if nil != err {
				break
			}
	
			_,err = files.repo.Delete(r.Context(), id)
			if nil != err {
				break
			}
			payload = "file deleted successfully"
			break
		}
	
		handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
	}