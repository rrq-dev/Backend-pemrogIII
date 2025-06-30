package model

type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type Payload struct {
	User string `json:"user"`
	Role string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type LoginResponse struct {
	Message string `json:"message" example:"Login success"`
	Token   string `json:"token" example:"v4.public.xxxxxx"`
}

type MahasiswaRequest struct {
	Nama       string       `bson:"nama" json:"nama" example:"Al Rumi"`
	NPM        int          `bson:"npm" json:"npm" example:"1234567890"`
	Prodi      string       `bson:"prodi" json:"prodi" example:"Teknik Informatika"`
	Fakultas   string       `bson:"fakultas" json:"fakultas" example:"Sekolah Vokasi"`
	Alamat     Alamat       `bson:"alamat" json:"alamat"`
	Minat      []string     `bson:"minat" json:"minat" example:"Programming, Data Science"`
	MataKuliah []MataKuliah `bson:"mata_kuliah" json:"mata_kuliah"`
}