package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"Learning-JobAccess-Camp/pkg/apis"
	goBMI "github.com/armstrongli/go-bmi"
)

func main() {
	input := &inputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("F:/Learning-JobAccess-Camp/chapter08/03.format/data/Hud.self.information.json")

	for {
		pi := input.GetInput()
		err := records.savePersonalInformation(pi)
		if err != nil {
			fmt.Println("main:保存错误", err)
		}

		fr, err := calc.FatRate(pi)
		if err != nil {
			fmt.Println("main:体脂率计算错误:", err)
		}

		rk.inputRecord(pi.Name, fr)

		rankResult, _ := rk.getRank(pi.Name)
		fmt.Println("排名结果:", rankResult)
	}

}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))

	personalInformation := apis.PersonalInformation{}
	err = json.Unmarshal(data, &personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("开始计算体脂信息：", personalInformation)
	bmi, _ := goBMI.BMI(float64(personalInformation.Weight), float64(personalInformation.Tall)) // todo handle error
	fmt.Printf("%s 的 BMI是：%v\n", personalInformation.Name, bmi)
	fatRate := goBMI.CalcFatRate(bmi, int(personalInformation.Age), personalInformation.Sex)
	fmt.Printf("%s 的体脂率是：%v\n", personalInformation.Name, fatRate)
}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	// b := make([]byte, 10)
	// var n int
	_, err = file.Write(data)
	fmt.Println(err)
}
