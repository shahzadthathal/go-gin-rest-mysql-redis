package repository

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
	"errors"
	"log"
	"strconv"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// GetPostByID ...
func GetPostByID(id int64) (model.MPost, error) {
	db := configuration.DB

	var post model.MPost

	result, err := db.Query("select id, title, description, status from posts where id = ?", id)
	if err != nil {
		// print stack trace
		log.Println("Error query post: " + err.Error())
		return post, err
	}

	for result.Next() {
		err := result.Scan(&post.ID, &post.Title, &post.Description, &post.Status)
		if err != nil {
			return post, err
		}
	}

	return post, nil
}

// GetPostAll ...
func GetPostAll() ([]model.MPost, error) {
	db := configuration.DB

	var mPost model.MPost
	var mPosts []model.MPost

	rows, err := db.Query("select id, title, description, status from posts")
	if err != nil {
		log.Println("Error query post: " + err.Error())
		return mPosts, err
	}

	for rows.Next() {
		if err := rows.Scan(&mPost.ID, &mPost.Title, &mPost.Description, &mPost.Status); err != nil {
			return mPosts, err
		}
		mPosts = append(mPosts, mPost)
	}

	return mPosts, nil
}

// CreatePost ...
func CreatePost(mPost model.MPost) (model.MPost, error) {
	db := configuration.DB

	var err error

	crt, err := db.Prepare("insert into posts (title, description, status) values (?, ?, ?)")
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	res, err := crt.Exec(mPost.Title, mPost.Description, mPost.Status)
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	mPost.ID = int64(rowID)

	// find post by id
	resval, err := GetPostByID(mPost.ID)
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	return resval, nil
}

// UpdatePost ...
func UpdatePost(mPost model.MPost) (model.MPost, error) {
	db := configuration.DB

	var err error

	crt, err := db.Prepare("update posts set title =?, description =?, status =? where id=?")
	if err != nil {
		return mPost, err
	}
	_, queryError := crt.Exec(mPost.ID, mPost.Title, mPost.Description, mPost.Status)
	if queryError != nil {
		return mPost, err
	}

	// find post by id
	res, err := GetPostByID(mPost.ID)
	if err != nil {
		return mPost, err
	}

	return res, nil
}

// DeletePostByID ...
func DeletePostByID(id int64) error {
	db := configuration.DB

	res, err := GetPostByID(id)
	if err != nil {
		return err
	}

	s := strconv.FormatInt(res.ID, 10)
	if (model.MPost{} == res) {
		return errors.New("no record value with id: %v" + s)
	}

	crt, err := db.Prepare("delete from posts where id=?")
	if err != nil {
		return err
	}
	_, queryError := crt.Exec(id)
	if queryError != nil {
		return err
	}

	return nil
}
