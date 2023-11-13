package categorytable

type Category struct {
	Id            string `json:"id"`
	ParentId      string `json:"parentId"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Order         string `json:"order"`
	ShowHome      string `json:"showHome"`
	ShowHomeOrder string `json:"showHomeOrder"`
	PrivateId     string `json:"privateId"`
	Status        string `json:"status"`
	Image         string `json:"image"`
}
