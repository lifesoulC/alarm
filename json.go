package main

type Req struct {
	Content []AlarmMsg `json:"alarmmsg"`
	Hoster  Host       `json:"host"`
}

type AlarmMsg struct {
	MsgName  string `json:"alarmname"`
	MsgValue string `json:"alarmvalue"`
}

type Host struct {
	SrcIP    string `json:"srcip"`
	HostName string `json:"hostname"`
	Email    string `json:"destemail"`
	Grade    int    `json:"grade"`
}

type Resperr struct {
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
}

type SenderType struct {
	User     string `json:"user"`
	Password string `json:"password"`
	MailHost string `json:"host"`
}
