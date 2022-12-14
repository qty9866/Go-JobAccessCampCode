package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectDB() *gorm.DB {
	connection, err := gorm.Open(mysql.Open("root:199866@tcp(127.0.0.1:3306)/go_db"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("数据库连接成功")
	return connection
}

func CreateNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Id:     6,
		Name:   "SunXX",
		Sex:    "女",
		Tall:   1.63,
		Weight: 63,
		Age:    49,
	})
	err := resp.Error
	if err != nil {
		fmt.Println("创建SunXX对象失败")
	}
	return nil
}

// 用于准确查询
func ormSelect(conn *gorm.DB) {
	result := make([]*PersonalInformation, 0)
	resp := conn.Where(&PersonalInformation{Age: 24}).Find(&result) //将结果放在find函数的参数里
	if resp.Error != nil {
		fmt.Println("查找失败", resp.Error)
		return
	}
	fmt.Println("查询结果:")
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", string(data))
}

// 用于不准确查询
func ormSearchWithInaccurateQuery(conn *gorm.DB) {
	result := make([]*PersonalInformation, 0)
	resp := conn.Where("tall>? and age > ?", 1.80, 25).Find(&result)
	if resp.Error != nil {
		fmt.Println("查询错误", resp.Error)
	}
	data, _ := json.Marshal(result)
	fmt.Println("不精准查询结果:")
	fmt.Printf("%+v\n", string(data))
}

// 全覆盖 全量更新 必须存在主键
func updateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:     1,
		Name:   "hud",
		Sex:    "男",
		Tall:   1.81,
		Weight: 81,
		Age:    24,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新失败！", err)
		return err
	}
	fmt.Println("更新成功！")
	return nil
}

// 部分更新模式
func updateDataSelectedFields(conn *gorm.DB) error {
	// 从数据库获取已有的数据
	var person = ormSelectByID(conn, 1)

	// 对部分数据进行修改
	person.Name = "Hud"
	person.Tall = 1.80

	// 只修改名字、身高
	resp := conn.Model(person).Select("NAME", "TALL").Updates(person)
	if resp.Error != nil {
		return resp.Error
	}
	fmt.Println("数据更新成功")
	return nil
}

// 根据ID查询
func ormSelectByID(conn *gorm.DB, id int64) *PersonalInformation {
	result := &PersonalInformation{}
	resp := conn.Where(&PersonalInformation{Id: id}).Find(&result) //将结果放在find函数的参数里
	if resp.Error != nil {
		fmt.Println("查找失败", resp.Error)
		return nil
	}
	fmt.Println("查询结果:")
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", string(data))
	return result
}

// 删除数据(DELETE) 删除数据时，必须通过”主键“来删除，批量删除操作可以通过特殊语法完成
func deleteData(conn *gorm.DB) error {
	p := &PersonalInformation{Id: 6}
	resp := conn.Delete(p)
	if resp.Error != nil {
		return resp.Error
	}
	fmt.Println("数据删除成功！")
	return nil
}

func main() {
	conn := connectDB()
	fmt.Println(conn)
	if err := CreateNewPerson(conn); err != nil {
		// todo handle error
	}
	ormSelect(conn)
	ormSearchWithInaccurateQuery(conn)

	err := updateExistingPerson(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = updateDataSelectedFields(conn)
	if err != nil {
		fmt.Println("更新失敗", err)
	}
	err = updateDataSelectedFields(conn)
	if err != nil {
		fmt.Println("更新失败！", err)
	}
	err = deleteData(conn)
	if err != nil {
		fmt.Println("删除数据错误")
	}
}
