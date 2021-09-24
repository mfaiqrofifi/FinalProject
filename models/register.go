package models

type Register struct {
	Nama     string `json:"nama"`
	Umur     int    `json:"umur"`
	Alamat   string `json:"alamat"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
