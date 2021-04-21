package repository

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
	"JWT_REST_Gin_MySQL/util"
	"errors"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// GetUserByID ...
func GetUserByID(id int64) (model.MUser, error) {
	db := configuration.DB

	var user model.MUser

	result, err := db.Query("select id, user_name, password, account_expired, account_locked, credentials_expired, enabled from m_user where id = ?", id)
	if err != nil {
		// print stack trace
		log.Println("Error query user: " + err.Error())
		return user, err
	}

	for result.Next() {
		err := result.Scan(&user.ID, &user.UserName, &user.Password, &user.AccountExpired, &user.AccountLocked, &user.CredentialsExpired, &user.Enabled)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

// GetUserLogin ...
func GetUserLogin(username string, password string) (model.MUser, error) {

	var mUser model.MUser
	var err error

	// find by user
	mUser, err = GetUserByUsername(username)
	if err != nil {
		return mUser, err
	}

	if (model.MUser{} == mUser) {
		return mUser, errors.New("bad credential")
	}

	var retVal bool = util.CheckPasswordHash(password, mUser.Password)
	if retVal == false {
		return mUser, errors.New("wrong password")
	}

	return mUser, nil
}

// GetUserByUsername ...
func GetUserByUsername(username string) (model.MUser, error) {
	db := configuration.DB

	var mUser model.MUser
	result, err := db.Query("select id, user_name, password, account_expired, account_locked, credentials_expired, enabled from m_user where user_name = ?", username)
	if err != nil {
		return mUser, err
	}

	for result.Next() {
		err := result.Scan(&mUser.ID, &mUser.UserName, &mUser.Password, &mUser.AccountExpired, &mUser.AccountLocked, &mUser.CredentialsExpired, &mUser.Enabled)
		if err != nil {
			return mUser, err
		}
	}

	return mUser, nil
}

// GetUserAll ...
func GetUserAll() ([]model.MUser, error) {
	db := configuration.DB

	var mUser model.MUser
	var mUsers []model.MUser

	rows, err := db.Query("select id, user_name, password, account_expired, account_locked, credentials_expired, enabled from m_user")
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return mUsers, err
	}

	for rows.Next() {
		if err := rows.Scan(&mUser.ID, &mUser.UserName, &mUser.Password, &mUser.AccountExpired,
			&mUser.AccountLocked, &mUser.CredentialsExpired, &mUser.Enabled); err != nil {
			return mUsers, err
		}
		mUsers = append(mUsers, mUser)
	}

	return mUsers, nil
}

// CreateUser ...
func CreateUser(mUser model.MUser) (model.MUser, error) {
	db := configuration.DB

	var err error

	hash, _ := util.HashPassword(mUser.Password, bcrypt.DefaultCost)
	mUser.Password = hash

	crt, err := db.Prepare("insert into m_user (user_name, password, account_expired, account_locked, credentials_expired, enabled) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Panic(err)
		return mUser, err
	}

	res, err := crt.Exec(mUser.UserName, mUser.Password, mUser.AccountExpired,
		mUser.AccountLocked, mUser.CredentialsExpired, mUser.Enabled)
	if err != nil {
		log.Panic(err)
		return mUser, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return mUser, err
	}

	mUser.ID = int64(rowID)

	// find user by id
	resval, err := GetUserByID(mUser.ID)
	if err != nil {
		log.Panic(err)
		return mUser, err
	}

	return resval, nil
}

// UpdateUser ...
func UpdateUser(mUser model.MUser) (model.MUser, error) {
	db := configuration.DB

	var err error

	hash, _ := util.HashPassword(mUser.Password, bcrypt.DefaultCost)
	mUser.Password = hash

	crt, err := db.Prepare("update m_user set user_name =?, password =?, account_expired =?, account_locked =?, credentials_expired =?, enabled =? where id=?")
	if err != nil {
		return mUser, err
	}
	_, queryError := crt.Exec(mUser.ID, mUser.UserName, mUser.Password, mUser.AccountExpired,
		mUser.AccountLocked, mUser.CredentialsExpired, mUser.Enabled)
	if queryError != nil {
		return mUser, err
	}

	// find user by id
	res, err := GetUserByID(mUser.ID)
	if err != nil {
		return mUser, err
	}

	return res, nil
}

// DeleteUserByID ...
func DeleteUserByID(id int64) error {
	db := configuration.DB

	res, err := GetUserByID(id)
	if err != nil {
		return err
	}

	s := strconv.FormatInt(res.ID, 10)
	if (model.MUser{} == res) {
		return errors.New("no record value with id: %v" + s)
	}

	crt, err := db.Prepare("delete from m_user where id=?")
	if err != nil {
		return err
	}
	_, queryError := crt.Exec(id)
	if queryError != nil {
		return err
	}

	return nil
}
