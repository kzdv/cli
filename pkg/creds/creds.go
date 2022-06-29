package creds

type Creds struct {
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}
