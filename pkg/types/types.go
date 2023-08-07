package types

type Book struct {
	UserName string `json:"username"`
	UserID int `json:"userid"`
	BookID    int    `json:"bookid"`
	BookName  string `json:"book_name"`
	RequestID int `json:"requestid"` 
	Publisher string `json:"publisher"`
	ISBN      string `json:"isbn"`
	Edition   int    `json:"edition"`
	Quantity  int    `json:"quantity"`

}

type HashedPassword struct {
	Salt string
	Hash string
}
type UserReg struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password  string `json:"password"`
	PasswordC string `json:"passwordc"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}


type UserInfo struct {
	UserName string `json:"username"`
}


type BookID struct {
    BookId int `json:"book_id,string"`	
}

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
	JWT_Key     string `yaml:"JWTSECRETKEY"`
}