package model

type Project struct{
	Id	int64  `json:"pid,omitempty" key:"primary" autoincr:"1" column:"pid"`
        Pname   string `json:"pname" column:"pname"`
	Pmanager string `json:"pmanager" column:"pmanager"`

}

func (project *Project) Table() string{
	return "project"
}

func (project *Project) String() string{
	return Stringify(project)
}
