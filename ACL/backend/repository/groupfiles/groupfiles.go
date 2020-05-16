package groupfiles
import (
	"context"
	"database/sql"
	//"fmt"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type groupfilesRepository struct {
	conn *sql.DB
}

func NewGroupFilesRepository(conn *sql.DB) *groupfilesRepository {
	return &groupfilesRepository{conn: conn}
}

func (groupfiles *groupfilesRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.GroupFiles)
	return driver.GetById(groupfiles.conn, obj, id)
}

func (groupfiles *groupfilesRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.GroupFiles)
	result, err := driver.Create(groupfiles.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (groupfiles *groupfilesRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.GroupFiles)
	err := driver.UpdateById(groupfiles.conn, &usr)
	return obj, err
}

func (groupfiles *groupfilesRepository) Delete2(cntx context.Context, gid int64,fid int64) (error){
	obj := &model.GroupFiles{Gid: gid,Fid: fid}
	return driver.Delete2(groupfiles.conn, obj,gid,fid)
}

func (groupfiles *groupfilesRepository) GetAll(cntx context.Context) ([]interface{}, error){ 
	obj := &model.GroupFiles{}

	return driver.GetAll(groupfiles.conn, obj, 0, 0)
}


