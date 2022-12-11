package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// MySQL: dataSourceName参数: <username>:<password>@tcp(<ip>:<port>)/<database name>
	db, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	err = TestDbConnection(err, db)
	if err != nil {
		log.Fatal(err)
	}
	queryAllData(err, db)
	err = insertData(db)
	if err != nil {
		log.Fatal(err)
	}
	queryAllData(err, db)
}

func insertData(db *sql.DB) (err error) {
	_, err = db.Exec(fmt.Sprintf("INSERT INTO personal_information(NAME,SEX,TALL,WEIGHT,AGE) values ('%s','%s','%f','%f','%d')",
		"QTY",
		"男",
		1.81,
		81.0,
		24,
	))
	if err != nil {
		log.Fatal("插入数据失败", err)
		return
	}
	return nil
}

// 查询数据
func queryAllData(err error, db *sql.DB) {
	rows, err := db.Query("SELECT NAME,SEX,TALL,WEIGHT,AGE FROM personal_information")
	if err != nil {
		log.Fatal(err)
	}

	// 构造结构体
	list := &apis.PersonalInformationList{}
	// 只要有下一行 就循环
	for rows.Next() {
		var name string
		var sex string
		var tall, weight float64
		var age int
		err := rows.Scan(&name, &sex, &tall, &weight, &age)
		if err != nil {
			return
		}
		fmt.Printf("name:%v\tage: %v\t\n", name, age)
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name:   name,
			Sex:    sex,
			Tall:   float32(tall),
			Weight: float32(weight),
			Age:    int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func TestDbConnection(err error, db *sql.DB) error {
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:199866@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		log.Fatalf("打开数据库错误:%v", err)
	}
	return db, err
}
