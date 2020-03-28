## User registration Go & postgreSQL

Simple user wallet registration using golang and postgresql database to store the data

## Installation

```bash
go get github.com/Neufal777/NSpace
```

## Usage

```golang

	//Introduce value in the database
	registerUser("NAME", "SURNAME", "USERNAME", 0)

	//show results
	showUsers()

```

## Output

```bash
2020/03/28 10:09:05 |0|NAME|SURNAME|USERNAME|BALANCE  
2020/03/28 10:09:05 |1|naoufal|dahouli|neufal79|324  
2020/03/28 10:09:05 |3|Jorsh|Mclean|jmc89|12  
```
