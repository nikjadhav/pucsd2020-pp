package model

type UserFiles	 struct {
	Id           int64  `json:"id" column:"id"`
	Fid		  int64  `json:"fid" column:"fid"`
	Ptype		  int64  `json:"ptype" column:"ptype"`

}



func (userfiles *UserFiles) Table() string {
	return "userfilesystem"
}

func (userfiles *UserFiles) String() string {
	return Stringify(userfiles)
}
