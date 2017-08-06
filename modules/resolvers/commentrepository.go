package resolvers

import (
	"log"

	"github.com/graphql-go-example/conf"
	"github.com/graphql-go-example/model"
)

//InsertComment -
func InsertComment(comment *model.Comment) error {
	strsql := `
		INSERT INTO comments(user_id, post_id, title, body)
		VALUES (?, ?, ?, ?)`

	res, err := conf.DB.Exec(strsql, comment.UserID, comment.PostID, comment.Title, comment.Body)
	if err != nil {
		log.Printf("[comments] Error INSERT: [%s] \nError: [%s]\n", strsql, err.Error())
		return err
	}

	comment.ID, _ = res.LastInsertId()
	return nil
}

//RemoveCommentByID -
func RemoveCommentByID(id int) error {
	_, err := conf.DB.Exec("DELETE FROM comments WHERE id=?", id)
	return err
}

//GetCommentByIDAndPost -
func GetCommentByIDAndPost(id int64, postID int64) (*model.Comment, error) {
	var (
		userID      int64
		title, body string
	)
	err := conf.DB.QueryRow(`
		SELECT user_id, title, body
		FROM posts
		WHERE id=?
		AND post_id=?
	`, id, postID).Scan(&userID, &title, &body)
	if err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:     id,
		UserID: userID,
		PostID: postID,
		Title:  title,
		Body:   body,
	}, nil
}

//GetCommentsForPost -
func GetCommentsForPost(id int64) ([]*model.Comment, error) {
	rows, err := conf.DB.Query(`
		SELECT c.id, c.user_id, c.title, c.body
		FROM comments AS c
		WHERE c.post_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		comments    = []*model.Comment{}
		cid, userID int64
		title, body string
	)
	for rows.Next() {
		if err = rows.Scan(&cid, &userID, &title, &body); err != nil {
			return nil, err
		}
		comments = append(comments, &model.Comment{
			ID:     cid,
			UserID: userID,
			PostID: id,
			Title:  title,
			Body:   body,
		})
	}
	return comments, nil
}
