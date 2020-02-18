package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error
var tpl *template.Template

type user struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	Password  []byte
}

/*
CREATE TABLE `users` (
	`Id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(255) NOT NULL,
	`first_name` VARCHAR(255) NOT NULL,
	`last_name` VARCHAR(255) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`Id`)
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;

CREATE TABLE `userroles` (
	`Id` INT(11) NOT NULL AUTO_INCREMENT,
	`UserId` BIGINT(20) NOT NULL,
	`Role` VARCHAR(50) NOT NULL,
	PRIMARY KEY (`Id`),
	INDEX `FK_UserId_Users_Id` (`UserId`),
	CONSTRAINT `userroles_ibfk_1` FOREIGN KEY (`UserId`) REFERENCES `users` (`Id`) ON UPDATE CASCADE ON DELETE CASCADE
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;

*/

func init() {
	dbUser := "admin"
	dbPass := "admin"
	dbName := "dev.golang.org"
	db, err = sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	defer db.Close()
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/userForm", userForm)
	http.HandleFunc("/createUsers", createUsers)
	http.HandleFunc("/editUsers", editUsers)
	http.HandleFunc("/deleteUsers", deleteUsers)
	http.HandleFunc("/updateUsers", updateUsers)
	log.Println("Server is up on 8090 port")
	log.Fatalln(http.ListenAndServe(":8090", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	rows, e := db.Query(
		`SELECT *
		FROM users;
		`)
	checkErr(e)

	users := make([]user, 0)
	for rows.Next() {
		usr := user{}
		rows.Scan(&usr.ID, &usr.Username, &usr.FirstName, &usr.LastName, &usr.Password)
		users = append(users, usr)
	}
	log.Println(users)
	tpl.ExecuteTemplate(w, "index.html", users)
}

func userForm(w http.ResponseWriter, req *http.Request) {
	err = tpl.ExecuteTemplate(w, "userForm.html", nil)
	checkErr(err)
}

func createUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		usr := user{}
		usr.Username = req.FormValue("username")
		usr.FirstName = req.FormValue("firstName")
		usr.LastName = req.FormValue("lastName")
		bPass, e := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)
		checkErr(e)
		usr.Password = bPass
		y, er := db.Exec(
			"INSERT INTO users (username, first_name, last_name, password) VALUES (?, ?, ?, ?)",
			usr.Username,
			usr.FirstName,
			usr.LastName,
			usr.Password,
		)
		checkErr(er)

		var userId, err = y.LastInsertId()
		checkErr(err)

		// insert user role
		_, er1 := db.Exec(
			"INSERT INTO userroles (userId, Role) VALUES (?, ?)",
			userId,
			"Admin",
		)
		checkErr(er1)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}

func editUsers(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		http.Error(res, "Please send ID", http.StatusBadRequest)
	}
	log.Println(id)
	rows, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	checkErr(err)
	usr := user{}
	for rows.Next() {
		rows.Scan(&usr.ID, &usr.Username, &usr.FirstName, &usr.LastName, &usr.Password)
	}
	log.Println(usr)
	tpl.ExecuteTemplate(res, "editUser.html", usr)
}

func deleteUsers(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		http.Error(res, "Please send ID", http.StatusBadRequest)
	}
	_, er := db.Exec("DELETE FROM users WHERE id = ?", id)
	checkErr(er)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func updateUsers(w http.ResponseWriter, req *http.Request) {
	_, er := db.Exec(
		"UPDATE users SET username = ?, first_name = ?, last_name = ? WHERE id = ? ",
		req.FormValue("username"),
		req.FormValue("firstName"),
		req.FormValue("lastName"),
		req.FormValue("id"),
	)
	checkErr(er)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
