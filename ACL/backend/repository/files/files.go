package files
import (
		//"fmt"
		"context"
		"database/sql"
	
		"github.com/pucsd2020-pp/ACL/backend/driver"
		"github.com/pucsd2020-pp/ACL/backend/model"
	)
type filesRepository struct {
	conn *sql.DB
}

func NewFilesRepository(conn *sql.DB) *filesRepository {
	return &filesRepository{conn: conn}
}

func (files *filesRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	fil := obj.(model.Files)
	result, err := driver.Create(files.conn, &fil)
	if nil != err {

		return 0, err
	}
	id, _ := result.LastInsertId()
	fil.Fid = id
	res,err:=driver.GetPath(files.conn,id)
	var i int64
	i=1
	if i==fil.Ftype{
		driver.CreateFile(res)
	}else{
		driver.CreateDirectory(res)
	}
	return id, nil
}

func (files *filesRepository) GetFilesByParent(cntx context.Context, pid int64) ([] interface{}, error) {
	obj := new(model.Files)
	return driver.GetFilesByParent(files.conn, obj, pid)	
	
}
func (files *filesRepository) Delete(cntx context.Context, id int64) (interface {},error){
	obj := &model.Files{Fid: id}
	path,_:=driver.GetPath(files.conn,id)
	driver.RemoveFile(path)
	return driver.DeleteById(files.conn, obj, id)
}