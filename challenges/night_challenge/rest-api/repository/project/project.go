package project

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/challenges/night_challenge/rest-api/driver"
	"github.com/pucsd2020-pp/challenges/night_challenge/rest-api/model"
)

type projectRepository struct {
	conn *sql.DB
}

func NewProjectRepository(conn *sql.DB) *projectRepository {
	return &projectRepository{conn: conn}
}

func (project *projectRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Project)
	return driver.GetById(project.conn, obj, id)
}

func (project *projectRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	sub := obj.(model.Project)
	result, err := driver.Create(project.conn, &sub)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	sub.Id = id
	return id, nil
}

func (project *projectRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	sub := obj.(model.Project)
	err := driver.UpdateById(project.conn, &sub)
	return obj, err
}

func (project *projectRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Project{Id: id}
	return driver.SoftDeleteById(project.conn,obj,id)
}
/*
func (user *userRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.User{}
	return driver.GetAll(user.conn, obj, 0, 0)
}

func (user *userRepository) Count(cntx context.Context) (interface{}, error) {
        obj := &model.User{}
        return driver.Count(user.conn, obj)
}
*/
