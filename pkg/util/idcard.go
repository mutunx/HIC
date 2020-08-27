package util

import (
	"fmt"
	"regexp"
	"strconv"
)

type Info struct {
	Id         string
	Name       string
	Birthday   string
	Province   string
	City       string
	Area       string
	Sex        int
	ValidState int
	RealState  int
	Error      string
}

// 通过身份证前17位获取第18位校验位
func getValidPos(id string) string {

	id = id[0:17]
	var idNumber []int32
	for _, v := range id {
		idNumber = append(idNumber, v-48)
	}
	// 17位加权位
	power := []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var sum int32
	// 计算位数*加权位总和
	for i, v := range idNumber {
		sum += v * power[i]
	}
	// 总和模11获取校验位
	y := sum % 11

	// 获取对应的校验位
	validPos := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	return validPos[y]
}

func ValidIDCard(id string) (err error) {
	hasLetter := regexp.MustCompile("[a-zA-WY-Z]|[\\W]").MatchString
	if len(id) != 18 {
		err = fmt.Errorf("系统只支持处理18位身份证")
		return
	}
	if hasLetter(id) {
		err = fmt.Errorf("身份证不合法")
		return
	}

	return nil
}

func ValidIdValidPos(id string) bool {
	return id[len(id)-1:] == getValidPos(id)
}

func GetProvince(id string) string {
	return id[0:2]
}

func GetCity(id string) string {
	return id[0:4]
}

func GetArea(id string) string {
	return id[0:6]
}

func GetBirthday(id string) string {
	return id[6:12]
}

func GetSex(id string) int {
	sex, err := strconv.Atoi(id[17:18])
	if err != nil {
		_ = fmt.Errorf("string %s to int error ", id[17:17])
	}
	return sex % 2
}

func GetInfo(id string, realSate int) (info *Info) {
	err := ValidIDCard(id)
	if err != nil {
		return &Info{
			Id:         id,
			Birthday:   "",
			Province:   "",
			City:       "",
			Area:       "",
			Sex:        0,
			ValidState: 0,
			RealState:  realSate,
			Error:      err.Error(),
		}

	}
	var validState int
	if ValidIdValidPos(id) {
		validState = 1
	}
	return &Info{
		Id:         id,
		Birthday:   GetBirthday(id),
		Province:   GetProvince(id),
		City:       GetCity(id),
		Area:       GetArea(id),
		Sex:        GetSex(id),
		ValidState: validState,
		RealState:  realSate,
		Error:      "",
	}
}
