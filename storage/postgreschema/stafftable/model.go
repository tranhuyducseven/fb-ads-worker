package stafftable

type Staff struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	RoleName     string `json:"roleName"`
	LarkId       string `json:"larkId"`
	PrivatePhone string `json:"privatePhone"`
}
