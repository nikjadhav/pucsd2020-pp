package model

type UserFiles	 struct {
	Id           int64  `json:"id" key:"primary" column:"id"`
	Fid		  int64  `json:"fid" key:"primary" column:"fid"`
	Ptype		  int64  `json:"ptype" column:"ptype"`

}



func (userfiles *UserFiles) Table() string {
	return "userfilesystem"
}

func (userfiles *UserFiles) String() string {
	return Stringify(userfiles)
}
