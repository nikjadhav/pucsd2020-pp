package model

type GroupFiles	 struct {
	Gid           int64  `json:"gid" column:"gid"`
	Fid		  int64  `json:"fid" column:"fid"`
	Ptype		  int64  `json:"ptype" column:"ptype"`

}



func (groupfiles *GroupFiles) Table() string {
	return "groupfilesystem"
}

func (groupfiles *GroupFiles) String() string {
	return Stringify(groupfiles)
}
