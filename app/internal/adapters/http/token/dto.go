package tokenclienthttp

type tokenClientDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type blockToken struct {
	TokenName string `json:"token_name"`
}

type allClients struct {
	Clients []string `json:"clients"`
}
