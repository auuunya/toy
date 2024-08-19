package model

const t_todo = `
CREATE TABLE IF NOT EXISTS t_todo(
	id int primary key auto_increment,
	title varchar(255) not null,
	content text,
	is_success boolean DEFAULT '0',
	data timestamp,
	user_id int
);
`
const t_user = `
CREATE TABLE IF NOT EXISTS t_user(
	id int primary key auto_increment,
	username varchar(64) not null,
	password varchar(125) not null,
	email varchar(125)
);
`

var Tables = [...]string{t_todo, t_user}

func CreateTable() error {
	tables := Tables
	for _, table := range tables {
		_, err := Db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil
}
