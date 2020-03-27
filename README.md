## User registration Go & postgreSQL

Simple user wallet registration using golang and postgresql database to store the data

## Installation

```bash
go get github.com/Neufal777/NSpace
```

## Usage

```golang

func main() {

	registerUser("NAME", "SURNAME", "USERNAME", 0)

	u := user{}

	log.Println("Connecting to SQL...")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully connected!")

	rows, _ := db.Query("SELECT id, name, surname, nickname, balance FROM users")

	for rows.Next() {
		rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Nickname, &u.balance)
		log.Printf("|%v|%v|%v|%v|%v  ", u.Id, u.Name, u.Surname, u.Nickname, u.balance)
	}

}

```

