package model

type Group struct {
	Id            int64  `json:"gid,omitempty" key:"primary" autoincr:"1" column:"gid"`
	Gname 		  string `json:"gname" column:"gname"` 

}



func (group *Group) Table() string {
	return "groups_"
}

func (group *Group) String() string {
	return Stringify(group)
}
