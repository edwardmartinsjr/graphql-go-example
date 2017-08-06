package resolvers

import (
	"log"

	"github.com/graphql-go-example/conf"
	"github.com/graphql-go-example/model"
)

//InsertPost -
func InsertPost(post *model.Post) error {
	strsql := `
		INSERT INTO posts(user_id, title, body)
		VALUES (?, ?, ?)
	`

	res, err := conf.DB.Exec(strsql, post.UserID, post.Title, post.Body)
	if err != nil {
		log.Printf("[posts] Error INSERT: [%s] \nError: [%s]\n", strsql, err.Error())
		return err
	}

	post.ID, _ = res.LastInsertId()
	return nil

}

//RemovePostByID -
func RemovePostByID(id int) error {
	_, err := conf.DB.Exec("DELETE FROM posts WHERE id=?", id)
	return err
}

//GetPostByID -
func GetPostByID(id int64) (*model.Post, error) {
	var (
		userID      int64
		title, body string
	)
	err := conf.DB.QueryRow(`
		SELECT user_id, title, body
		FROM posts
		WHERE id=?
	`, id).Scan(&userID, &title, &body)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:     id,
		UserID: userID,
		Title:  title,
		Body:   body,
	}, nil
}

//GetPostByIDAndUser -
func GetPostByIDAndUser(id, userID int64) (*model.Post, error) {
	var title, body string
	err := conf.DB.QueryRow(`
		SELECT title, body
		FROM posts
		WHERE id=?
		AND user_id=?
	`, id, userID).Scan(&title, &body)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:     id,
		UserID: userID,
		Title:  title,
		Body:   body,
	}, nil
}

//GetPostsForUser -
func GetPostsForUser(id int64) ([]*model.Post, error) {
	rows, err := conf.DB.Query(`
		SELECT p.id, p.title, p.body
		FROM posts AS p
		WHERE p.user_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		posts       = []*model.Post{}
		pid         int
		title, body string
	)
	for rows.Next() {
		if err = rows.Scan(&pid, &title, &body); err != nil {
			return nil, err
		}
		posts = append(posts, &model.Post{ID: id, UserID: id, Title: title, Body: body})
	}
	return posts, nil
}
