package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"regexp"
	"strings"
)

// 加密
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logs.Error(err.Error())
	}
	return string(hash)
}

// 密码验证
func ComparePwd(hashPwd string, plainPwd []byte) bool {
	logs.Info(hashPwd)
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logs.Error(err.Error())
		return false
	}

	return true
}

// 判断array contain item
func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// [a] -> a -> a
// [a b c] -> a b c -> a,b,c
func Convert(array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func IntToBool(num int8) bool {
	if num > 0 {
		return true
	}
	return false
}

func ReturnQ(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += ",?"
	}
	return str[1:]
}

func GetKeys(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetValues(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	values := make([]string, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetSum(m []int) int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	sum := 0
	for _, v := range m {
		sum = sum + v
	}
	return sum
}

func RegJsonData(Data []byte) []byte {

	reg := regexp.MustCompile("([a-zA-Z]\\w*):")
	regStr := reg.ReplaceAllString(string(Data), `"$1":`)
	//单引号转装双引号
	newStr := strings.Replace(regStr, `'`, `"`, -1)

	return []byte(newStr)
}
