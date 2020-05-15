package repository

import (
	"context"
)

type IRepository interface {
	GetByID(context.Context, int64) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, int64) (interface {},error)
	GetAll(context.Context) ([]interface{}, error)
}
type JRepository interface {
	IsValidUser(context.Context, interface{}) (interface{}, error)
}
type UGRepository interface{
	AddUserGroup(context.Context, interface{}) ( error)
	GetUsersByGroup(context.Context,int64) ([]interface{}, error)
	Delete2(context.Context, int64,int64) error
}

type FRepository interface {
	Create(context.Context, interface{}) (interface{}, error)
	GetFilesByParent(context.Context, int64) ([]interface{}, error)
	Delete(context.Context, int64) (interface {},error)
}
type UFRepository interface {
	Create(context.Context, interface{}) (interface{}, error)
	GetAll(context.Context) ([]interface{}, error)
	Delete2(context.Context, int64,int64) error

}
type ARepository interface{
	IsAccess(context.Context, int64,int64) (int64, error)
}



type Repository struct {
}

func (repo *Repository) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}

func (repo *Repository) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (repo *Repository) Delete(cntx context.Context, id int64) (obj interface {}, err error) {
	return
}

func (repo *Repository) GetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}

func (repo *Repository) AddUserGroup(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) GetUsersByGroup(cntx context.Context,gid int64) (obj []interface{}, err error) {
	return
}

func (repo1 *Repository) IsValidUser(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) DeleteGroup(cntx context.Context, gid int64) ( err error) {
	return
}

func (repo *Repository) Delete2(cntx context.Context,id int64,gid int64) ( err error) {
	return
}

func (repo *Repository) GetFilesByParent(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}

func (repo *Repository) IsAccess(cntx context.Context, id int64,fid int64) (res int64, err error) {
	return
}