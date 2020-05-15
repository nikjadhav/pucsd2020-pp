package model

type UserGroup struct {
	Id            int64  `json:"id"  column:"id"`
	Gid			  int64  `json:"gid"  column:"gid"`



}



func (usergroup *UserGroup) Table() string {
	return "usergroup"
}

func (usergroup *UserGroup) String() string {
	return Stringify(usergroup)
}
