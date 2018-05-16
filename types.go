package main

//	подключение к БД
const (
	//	Файл настроек
	USERS_FILE_NAME = "users.conf"
	//	роутинг
)

//	структура одного поста
type settingsStruct struct {
	ListeningIp                 string
	ListeningPort               string
	EmailSmtpServerIp           string
	EmailSmtpServerPort         string
	EmailSmtpLogin              string
	EmailSmtpPassword           string
	EmailFromName               string
	EmailFromMail               string
	EmailTo                     string
	EmailCopy                   string
	EmailShadow                 string
	EmailSubjectAdm             string
	EmailTextBeforeUsernameAdm  string
	EmailTextAfterUsernameAdm   string
	EmailSubjectUser            string
	EmailTextBeforeUsernameUser string
	EmailTextAfterUsernameUser  string
}

type cloudUsersStruct struct {
	UserName  string
	UserEmail string
}

//	структура одного поста
type postStruct struct {
	PostId     string   `json:"_id"`
	PostTitle  string   `json:"posttitle"`
	CreateDate string   `json:"createdate"`
	PostTags   []string `json:"posttags"`
	PostBody   string   `json:"postbody"`
	Picture    string   `json:"picture"`
}

//	структура одного поста
type postStructForSend struct {
	PostTitle  string   `json:"posttitle"`
	CreateDate string   `json:"createdate"`
	PostTags   []string `json:"posttags"`
	PostBody   string   `json:"postbody"`
	Picture    string   `json:"picture"`
}

//	структура одной новости
type newsStruct struct {
	Author  string `json:"author"`
	Date    string `json:"date"`
	Picture string `json:"picture"`
	Text    string `json:"text"`
	Title   string `json:"title"`
}

//	проверочная структура
type testStruct struct {
	Text   string
	Number int
}

//	структура пользователей
type usersStruct struct {
	Login      string //	логин
	Password   string //	пароль
	Name       string //	имя
	SecGroup   string //	группа безопасности
	Properties string //	доп. информация
}
