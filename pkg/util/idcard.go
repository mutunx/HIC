package util

// 通过身份证前17位获取第18位校验位
func getValidPos(card string) string {

	card = card[0:17]
	var cardNumber []int32
	for _, v := range card {
		cardNumber = append(cardNumber, v-48)
	}
	// 17位加权位
	power := []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var sum int32
	// 计算位数*加权位总和
	for i, v := range cardNumber {
		sum += v * power[i]
	}
	// 总和模11获取校验位
	y := sum % 11

	// 获取对应的校验位
	validPos := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	return validPos[y]
}
