package usergroup
import (
	"fmt"
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type usergroupRepository struct {
	conn *sql.DB
}


func NewUserGroupRepository(conn *sql.DB) *usergroupRepository {
	return &usergroupRepository{conn: conn}
}


func (usergroup *usergroupRepository) AddUserGroup(cntx context.Context, obj interface{}) (error) {
	grp := obj.(model.UserGroup)
	result, err := driver.Create(usergroup.conn, &grp)
	if nil != err {
		return  err
	}
	fmt.Println("result",result)
	return nil
}

func (usergroup *usergroupRepository) GetUsersByGroup(cntx context.Context, gid int64) ([] interface{}, error) {
	obj := &model.UserGroup{}
	return driver.GetUsersByGroup(usergroup.conn, obj, gid)
}

func (usergroup *usergroupRepository) Delete2(cntx context.Context,id int64,gid int64) error {
	obj := &model.UserGroup{Id: id,Gid:gid}
	return driver.Delete2(usergroup.conn,obj,id,gid)
}
