package model

//Comment -
type Comment struct {
	ID     int64
	UserID int64
	PostID int64
	Title  string
	Body   string
}
