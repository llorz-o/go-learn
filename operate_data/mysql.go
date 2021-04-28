package operate_data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	open, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_2_sql")
	if err != nil {
		fmt.Println("sqlx.Open err:", err)
		return
	}
	Db = open
}

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func mysqlInsert() {
	r, err := Db.Exec(`insert into person(username, sex, email)values(?,?,?)`, "jojo", "man", "jojo@gmail.com")
	if err != nil {
		fmt.Println("exec failed:", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("insert failed:", err)
		return
	}
	fmt.Println("insert succ:", id)
}

func mysqlSelect() {
	var person []Person
	err := Db.Select(&person, `select user_id, username, sex, email from person where user_id=?`, 2)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("select succ:", person)
}

func mysqlUpdate() {
	res, err := Db.Exec(`update person set username=? where user_id=?`, "new name", 2)
	if err != nil {
		fmt.Println("update failed", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("res rows affected failed", err)
		return
	}
	fmt.Println("update success:", row)
}

func mysqlDelete() {
	exec, err := Db.Exec(`delete from person where user_id=?`, 3)
	if err != nil {
		fmt.Println("delete failed", err)
		return
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("row effected failed", err)
		return
	}
	fmt.Println("delete success", affected)
}

func FuncMySQL() {
	defer Db.Close()
	//mysqlInsert()
	//mysqlSelect()
	//mysqlUpdate()
	mysqlDelete()
}
