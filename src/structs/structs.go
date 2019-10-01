package structs

type DataClient struct {
	Page      string `json:"page"`
	DelCookie bool   `json:"delCookie"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type SigninData struct {
	HTML   string
	MsgErr string
	Title  string
}

type Users struct {
	Login    string
	Password string
	Token    string
}

type Data struct {
	HTML    string
	Token   string
	ReLogin bool
}

type TMPL struct {
	Auth  bool
	Title string
}
