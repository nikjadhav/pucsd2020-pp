package model
import (
	"database/sql"
)
type NullInt64 struct {
    sql.NullInt64
}
type Files struct {
	Fid            int64  `json:"fid,omitempty" key:"primary" autoincr:"1" column:"fid"`
	Fname 		  string `json:"fname" column:"fname"` 
	Parent 		  NullInt64  `json:"parent" column:"parent"` 
	Ftype	      int64  `json:"ftype" column:"ftype"`
	Owner		  NullInt64  `json:"owner" column:"owner"`
}


func (files *Files) Table() string {
	return "filesystem"
}

func (files *Files) String() string {
	return Stringify(files)
}
