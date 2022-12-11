package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"time"
)

/*
性能对比： 10w条数据

==================序列化JSON==================
序列化耗时： 34.2005ms
写文件： 91.2864ms
==================序列化YAML==================
序列化耗时： 1.0374918s
写文件： 91.5277ms
==================序列化ProtoBuf==================
序列化耗时： 6.1438ms
写文件： 1.0555ms

==================反序列化JSON==================
反序列化JSON时间: 97.3399ms
==================反序列化Yaml==================
反序列化Yaml时间: 791.4535ms
==================反序列化ProtoBuf==================
反序列化ProtoBuf时间: 10.108ms
*/
func main() {
	// 10万条数据
	// json yaml protobuf 分别保存，记录序列化耗时
	count := 100000
	persons := make([]*apis.PersonalInformation, 0, count)

	for i := 0; i < count; i++ {
		persons = append(persons, &apis.PersonalInformation{
			Name:   "Hud",
			Sex:    "男",
			Tall:   1.81,
			Weight: 81,
			Age:    24,
		})
	}
	{
		fmt.Println("==================序列化JSON==================")
		startTime := time.Now()
		data, err := json.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishJsonMarshalTime := time.Now()
		err = ioutil.WriteFile("F:/Learning-JobAccess-Camp/chapter08/04.think/data/data.json", data, 0777)
		if err != nil {
			log.Fatal(err)
		}
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishJsonMarshalTime.Sub(startTime))
		fmt.Println("写文件：", finishWriteFileTime.Sub(finishJsonMarshalTime))
	}
	{
		fmt.Println("==================序列化YAML==================")
		startTime := time.Now()
		data, err := yaml.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishYamlMarshalTime := time.Now()
		err = ioutil.WriteFile("F:/Learning-JobAccess-Camp/chapter08/04.think/data/data.yaml", data, 0777)
		if err != nil {
			log.Fatal(err)
		}
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishYamlMarshalTime.Sub(startTime))
		fmt.Println("写文件：", finishWriteFileTime.Sub(finishYamlMarshalTime))
	}
	{
		fmt.Println("==================序列化ProtoBuf==================")
		pLister := &apis.PersonalInformationList{
			Items: persons,
		}
		startTime := time.Now()
		data, err := proto.Marshal(pLister)
		if err != nil {
			log.Fatal(err)
		}
		finishYamlMarshalTime := time.Now()
		err = ioutil.WriteFile("F:/Learning-JobAccess-Camp/chapter08/04.think/data/data.protobuf", data, 0777)
		if err != nil {
			log.Fatal(err)
		}
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时：", finishYamlMarshalTime.Sub(startTime))
		fmt.Println("写文件：", finishWriteFileTime.Sub(finishYamlMarshalTime))
	}

	//    反序列化JSON时间: 95.6069ms
	{
		fmt.Println("==================反序列化JSON==================")
		startTime := time.Now()
		data, err := ioutil.ReadFile("F:/Learning-JobAccess-Camp/chapter08/04.think/data/data.json")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &persons)
		if err != nil {
			log.Fatal(err)
		}
		finishTime := time.Now()
		fmt.Println("反序列化JSON时间:", finishTime.Sub(startTime))
		//	}

		//反序列化Yaml时间: 803.5403ms
		{
			fmt.Println("==================反序列化Yaml==================")
			startTime := time.Now()
			data, err := ioutil.ReadFile("F:/Learning-JobAccess-Camp/chapter08/04.think/data/data.yaml")
			if err != nil {
				log.Fatal(err)
			}
			err = yaml.Unmarshal(data, &persons)
			if err != nil {
				log.Fatal(err)
			}
			finishTime := time.Now()
			fmt.Println("反序列化Yaml时间:", finishTime.Sub(startTime))
		}
		//反序列化ProtoBuf时间: 20.8213ms
		{
			pLister := &apis.PersonalInformationList{}
			fmt.Println("==================反序列化ProtoBuf==================")
			startTime := time.Now()
			data, err := ioutil.ReadFile("F:/Learning-JobAccess-Camp/chapter08/04.think/data/data.protobuf")
			if err != nil {
				log.Fatal(err)
			}
			err = proto.Unmarshal(data, pLister)
			if err != nil {
				log.Fatal(err)
			}
			finishTime := time.Now()
			fmt.Println("反序列化ProtoBuf时间:", finishTime.Sub(startTime))
		}

	}
}
