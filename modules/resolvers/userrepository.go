package resolvers

import (
	"log"

	"github.com/graphql-go-example/conf"
	"github.com/graphql-go-example/model"
)

//InsertUser -
func InsertUser(user *model.User) error {
	strsql := `
		INSERT INTO users(email)
		VALUES (?)
	`

	res, err := conf.DB.Exec(strsql, user.Email)
	if err != nil {
		log.Printf("[users] Error INSERT: [%s] \nError: [%s]\n", strsql, err.Error())
		return err
	}

	user.ID, _ = res.LastInsertId()
	return nil
}

//GetUserByID -
func GetUserByID(id int64) (*model.User, error) {
	var email string
	err := conf.DB.QueryRow("SELECT email FROM users WHERE id=?", id).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    id,
		Email: email,
	}, nil
}

//RemoveUserByID -
func RemoveUserByID(id int) error {
	_, err := conf.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}

//Follow -
func Follow(followerID, followeeID int) error {
	_, err := conf.DB.Exec(`
		INSERT INTO followers(follower_id, followee_id)
		VALUES (?, ?)
	`, followerID, followeeID)
	return err
}

//Unfollow -
func Unfollow(followerID, followeeID int) error {
	_, err := conf.DB.Exec(`
		DELETE FROM followers
		WHERE follower_id=?
		AND followee_id=?
	`, followerID, followeeID)
	return err
}

//GetFollowerByIDAndUser -
func GetFollowerByIDAndUser(id int64, userID int64) (*model.User, error) {
	var email string
	err := conf.DB.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.follower_id
		AND f.follower_id=?
		AND f.followee_id=?
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    id,
		Email: email,
	}, nil
}

//GetFollowersForUser -
func GetFollowersForUser(id int64) ([]*model.User, error) {
	rows, err := conf.DB.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.followee_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*model.User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &model.User{ID: id, Email: email})
	}
	return users, nil
}

//GetFolloweeByIDAndUser -
func GetFolloweeByIDAndUser(id int64, userID int64) (*model.User, error) {
	var email string
	err := conf.DB.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.followee_id
		AND f.followee_id=?
		AND f.follower_id=?
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    id,
		Email: email,
	}, nil
}

//GetFolloweesForUser -
func GetFolloweesForUser(id int64) ([]*model.User, error) {
	rows, err := conf.DB.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.follower_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*model.User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &model.User{ID: id, Email: email})
	}
	return users, nil
}
