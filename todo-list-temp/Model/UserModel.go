package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []*User

func GetAllUser() ([]*User, error) {
	users := []*User{}
	rows, err := Db.Query("select id,username,password,email from t_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := new(User)
		_ = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users, nil
}

func GetUser(username string) (*User, error) {
	user := new(User)
	err := Db.QueryRow("select id,username,password,email from t_user where username=?", username).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func CheckPass(email, password string) (*User, error) {
	user := new(User)
	err := Db.QueryRow("select id,username,password,email from t_user where email=?", email).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	if SetMd5(password) != user.Password {
		return nil, err
	}
	return user, nil
}
func (user *User) CreateUser() error {
	stmt, err := Db.Prepare("insert into t_user (username,password,email) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Username, SetMd5(user.Password), user.Email)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
func (user *User) UpdateUser() error {
	stmt, err := Db.Prepare("update t_user set username=?,password=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Username, SetMd5(user.Password), user.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
func (user *User) DeleteUser() error {
	stmt, err := Db.Prepare("delete from t_user where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
