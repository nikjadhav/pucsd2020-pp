package userfiles
import (
	"context"
	"database/sql"
	//"fmt"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type userfilesRepository struct {
	conn *sql.DB
}

func NewUserFilesRepository(conn *sql.DB) *userfilesRepository {
	return &userfilesRepository{conn: conn}
}

func (userfiles *userfilesRepository) IsAccess(cntx context.Context, id int64,fid int64) (int64, error) {
	obj := new(model.UserFiles)
	return driver.IsAccess(userfiles.conn, obj, id,fid)
}

func (userfiles *userfilesRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserFiles)
	
	result, err := driver.Create(userfiles.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}
func (userfiles *userfilesRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserFiles)
	err := driver.UpdateById(userfiles.conn, &usr)
	return obj, err
}

func (userfiles *userfilesRepository) Delete2(cntx context.Context, id int64,fid int64) (error){
	obj := &model.UserFiles{Id: id,Fid: fid}
	return driver.Delete2(userfiles.conn, obj, id,fid)
}

func (userfiles *userfilesRepository) GetAll(cntx context.Context) ([]interface{}, error){ 
	obj := &model.UserFiles{}

	return driver.GetAll(userfiles.conn, obj, 0, 0)
}


