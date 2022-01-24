package person

type RegisterRequestFormat struct {
	Nama     string `json:"name"`
	HP       string `json:"hp"`
	Umur     int    `json:"umur"`
	Password string `json:"password"`
}
