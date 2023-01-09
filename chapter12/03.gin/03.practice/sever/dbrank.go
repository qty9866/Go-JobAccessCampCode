package main

import (
	"Learning-JobAccess-Camp/chapter12/02.practice/frinterface"
	"Learning-JobAccess-Camp/pkg/apis"
	"fmt"
	"gorm.io/gorm"
	"log"
)

var _ frinterface.ServeInterface = &dbRank{}

var _ frinterface.RankInitInterface = &dbRank{}

func NewDbRank(conn *gorm.DB, embedRank frinterface.ServeInterface) frinterface.ServeInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedRank == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbRank{
		conn:      conn,
		embedRank: embedRank,
	}
}

type dbRank struct {
	conn *gorm.DB
	// 在下面调用的过程中，这个embedRank都是一个空的，所以要将更新/修改的数据嵌入到内存中去
	embedRank frinterface.ServeInterface
}

func (d *dbRank) Init() error {
	output := make([]*apis.PersonalInformation, 0)
	// 查找所有数据  1=1是True
	resp := d.conn.Find(&output)
	if err := resp.Error; err != nil {
		fmt.Println("初始化导出数据出错", err)
		return err
	}
	// 初始化embedRank
	for _, item := range output {
		if _, err := d.embedRank.UpdatePersonalInformation(item); err != nil {
			log.Fatalf("初始化%s时失败:%v", item.Name, err)
		}
	}
	return nil
}

func (d dbRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	resp := d.conn.Create(pi)
	err := resp.Error
	if err != nil {
		fmt.Printf("创建%s对象失败:%v\n", pi.Name, err)
	}
	fmt.Printf("创建%s对象成功\n", pi.Name)
	err = d.embedRank.RegisterPersonalInformation(pi)
	if err != nil {
		// todo handle error
	}
	return nil
}

func (d dbRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	resp := d.conn.Updates(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s失败:%v\n", pi.Name, err)
		return nil, err
	}
	fmt.Printf("更新%s成功\n", pi.Name)
	return d.embedRank.UpdatePersonalInformation(pi)
}

func (d dbRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	return d.embedRank.GetFatRate(name)
}

func (d dbRank) GetTop() ([]*apis.PersonalRank, error) {
	return d.embedRank.GetTop()
}
