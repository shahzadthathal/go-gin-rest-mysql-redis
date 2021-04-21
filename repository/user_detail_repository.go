package repository

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
	"log"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// GetUserDetailByID ...
func GetUserDetailByID(id int64) (model.MUserDetail, error) {
	db := configuration.DB

	var userDtl model.MUserDetail

	result, err := db.Query("select d.*,m.* from m_user_detail d JOIN m_user AS m ON d.user_id = m.id where d.id = ?", id)
	if err != nil {
		// print stack trace
		log.Println("Error query user detail: " + err.Error())
		return userDtl, err
	}

	// defer db.Close()

	for result.Next() {

		err := result.Scan(&userDtl.ID, &userDtl.Address, &userDtl.DOB, &userDtl.POB, &userDtl.Phone,
			&userDtl.Email, &userDtl.UserID, &userDtl.MUser.ID, &userDtl.MUser.Password,
			&userDtl.MUser.UserName, &userDtl.MUser.AccountExpired, &userDtl.MUser.AccountLocked,
			&userDtl.MUser.CredentialsExpired, &userDtl.MUser.Enabled)
		if err != nil {
			return userDtl, err
		}
	}

	return userDtl, nil
}

// GetAllUserDetail ...
func GetAllUserDetail() ([]model.MUserDetail, error) {
	db := configuration.DB
	// defer db.Close()

	var userDtl model.MUserDetail
	var userDtls []model.MUserDetail

	rows, err := db.Query("select d.*,m.* from m_user_detail d JOIN m_user AS m ON d.user_id = m.id")
	if err != nil {
		return userDtls, err
	}

	// defer db.Close()

	var isError bool
	for rows.Next() {
		if err := rows.Scan(&userDtl.ID, &userDtl.Address, &userDtl.DOB, &userDtl.POB, &userDtl.Phone,
			&userDtl.Email, &userDtl.UserID, &userDtl.MUser.ID, &userDtl.MUser.Password,
			&userDtl.MUser.UserName, &userDtl.MUser.AccountExpired, &userDtl.MUser.AccountLocked,
			&userDtl.MUser.CredentialsExpired, &userDtl.MUser.Enabled); err != nil {
			isError = true
			break
		} else {
			userDtls = append(userDtls, userDtl)
		}
	}

	if isError {
		return userDtls, err
	}

	return userDtls, nil
}

// CreateUserDetail ...
func CreateUserDetail(usr model.MUserDetail) (model.MUserDetail, error) {
	db := configuration.DB
	defer db.Close()

	var userDtl model.MUserDetail

	crt, err := db.Prepare("insert into m_user_detail (address, dob, pob, phone, email, user_id) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Panic(err)
		return userDtl, err
	}

	res, err := crt.Exec(usr.Address, usr.DOB, usr.POB, usr.Phone, usr.Email, usr.UserID)
	if err != nil {
		log.Panic(err)
		return userDtl, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return userDtl, err
	}

	userDtl.ID = int64(rowID)

	// find user detail by id
	resval, err := GetUserDetailByID(userDtl.ID)
	if err != nil {
		log.Panic(err)
		return userDtl, err
	}

	return resval, nil
}

// UpdateUserDetail ...
func UpdateUserDetail(usr model.MUserDetail) (model.MUserDetail, error) {
	db := configuration.DB
	defer db.Close()

	var userDtl model.MUserDetail

	crt, err := db.Prepare("update m_user_detail set address =?, dob =?, pob =?, phone =?, email =?, user_id =? where id=?")
	if err != nil {
		return userDtl, err
	}
	_, queryError := crt.Exec(usr.Address, usr.DOB, usr.POB, usr.Phone, usr.Email, usr.UserID, usr.ID)
	if queryError != nil {
		return userDtl, err
	}

	// find user detail by id
	res, err := GetUserDetailByID(usr.ID)
	if err != nil {
		return userDtl, err
	}

	return res, nil
}

// DeleteUserDetailByID ...
func DeleteUserDetailByID(id int64) error {
	db := configuration.DB
	defer db.Close()

	crt, err := db.Prepare("delete from m_user_detail where id=?")
	if err != nil {
		return err
	}
	_, queryError := crt.Exec(id)
	if queryError != nil {
		return err
	}

	return nil
}
