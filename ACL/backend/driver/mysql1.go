package driver

import (
	"os"
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"reflect"
	"strings"
	"github.com/pucsd2020-pp/ACL/backend/model"
	_ "github.com/go-sql-driver/mysql"

)

func GetUsersByGroup(conn *sql.DB, object model.IModel, gid int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}
		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT id")
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE gid= ?")
	
	query := queryBuffer.String()
	row, err := conn.Query(query, gid)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	objects := make([]interface{}, 0)
	defer row.Close()
	for row.Next() {
		user_id:=0
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(&user_id)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
		objects = append(objects,user_id)

	}

	return objects, nil
}

func Delete2(conn *sql.DB, object model.IModel, id int64,gid int64) (error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		fmt.Println("field",field)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE "+columns[0]+"=? and "+columns[1]+"= ? ")
	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return err
	}

	defer stmt.Close()
	_,err = stmt.Exec(id,gid)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	return err
}
func GetFilesByParent(conn *sql.DB, object model.IModel, pid int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)
	var row *sql.Rows
	var err error
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	if(pid == 0){
	queryBuffer.WriteString(" WHERE parent IS NULL")
	}else {
		queryBuffer.WriteString(" WHERE parent = ?")
	}	

	query := queryBuffer.String()

	if (pid == 0){ 
	row, err = conn.Query(query)
	}else{
		row, err = conn.Query(query,pid)
	}
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		val:=new(model.Files)	
		err = row.Scan(&val.Fid,&val.Fname,&val.Parent,&val.Ftype,&val.Owner)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}		
		objects = append(objects,val)
	}
	return objects, nil
}
func CreateFile(path string){
	f, err := os.Create("/"+path)
	if err != nil {
        log.Printf("Error: Create File: %s\n", err.Error())
	}
	defer f.Close()
}
func CreateDirectory(path string){
	_, err := os.Stat("/"+path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("/"+path,0777)
		if errDir != nil {
			log.Printf("Error: Create Directory: %s\n", err.Error())
		}
 
	}
}

func GetPath(conn *sql.DB,fid int64) (string,error){
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("WITH RECURSIVE file_path (fid, fname, path) AS")
	queryBuffer.WriteString(" (")
	queryBuffer.WriteString(" SELECT fid, fname, fname as path")
	queryBuffer.WriteString(" FROM filesystem")
	queryBuffer.WriteString(" WHERE parent is NULL")
	queryBuffer.WriteString(" UNION ALL")
	queryBuffer.WriteString(" SELECT c.fid, c.fname, CONCAT(cp.path, '/', c.fname)")
	queryBuffer.WriteString(" FROM file_path AS cp JOIN filesystem AS c ")
	queryBuffer.WriteString(" ON cp.fid = c.parent ")
	queryBuffer.WriteString(" )")
	queryBuffer.WriteString(" SELECT path FROM file_path where fid= ? ")
	query := queryBuffer.String()

	var return_res string
	return_res="Path Not Found"
	row, err := conn.Query(query,fid)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return return_res, err
	}
	defer row.Close()
	var scanval string
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return return_res, err
		}
		
		err := row.Scan(&scanval)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return return_res, err
		}else{
			return scanval,err
		}
		
	}
	return return_res,err
}
func RemoveFile(path string){
	err := os.RemoveAll("/"+path)
    if err != nil {
        fmt.Println(err);
        return
    }
       fmt.Println("File /"+path+ " successfully deleted")
}
func IsAccess(conn *sql.DB, object model.IModel, id int64,fid int64) (int64,error){
	var queryBuffer bytes.Buffer
	var _id,_fid string
	_id=strconv.FormatInt(id, 10)
	_fid=strconv.FormatInt(fid, 10)
	fmt.Println("int to string",_id,_fid)
	queryBuffer.WriteString("select  case  when count(*)>0 ")
	queryBuffer.WriteString("then 1 else 0 end as access from ")
	queryBuffer.WriteString("(select gid from usergroup  where id="+_id+" and gid in(select gid from ")
	queryBuffer.WriteString("groupfilesystem  where fid="+_fid+" and ptype=2)) as s ")
	queryBuffer.WriteString("union select  case  when count(*)>0 then 1 else 0  end as res ")
	queryBuffer.WriteString("from userfilesystem where id="+_id+" ")
	queryBuffer.WriteString("and fid="+_fid+" and ptype=2 ")
	queryBuffer.WriteString("union select case when count(*)>0 then 1 else 0 end as own from filesystem where owner="+_id+" and fid="+_fid+" ")
	queryBuffer.WriteString("union select case when count(*)>0 then 1 else 0 end as admin from user_detail where id="+_id+" and rtype=1")
	query := queryBuffer.String()
	fmt.Println("q",query)
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return 0, err
	}
	defer row.Close()
	var res int64
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return 0, err
		}
		var res int64
		err = row.Scan(&res)
		if(res==1){
			return res,nil
		}
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return 0, err
		}
		}
	return res, nil
}
