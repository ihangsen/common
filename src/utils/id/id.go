package id

import (
	"math/rand"
	"strings"
	"time"
)

func Merge(index uint8, id uint64) uint64 {
	return id<<8 + uint64(index)
}

func Split(key uint64) (index uint8, id uint64) {
	return uint8(key & 0xFF), key >> 8
}

const (
	base    = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 进制的包含字符string类型
	pad     = "0"
	decimal = 32 // 进制长度
	length  = 6  // code最小长度
)

var baseMap = func() map[byte]uint64 {
	baseArr := []byte(base)          // 字符串进制转换为byte数组
	baseRev := make(map[byte]uint64) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = uint64(k)
	}
	return baseRev
}()

func ToCode(id uint64) string {
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % decimal
		id = id / decimal
		res += string(base[mod])
	}
	resLen := len(res)
	if resLen < length {
		res += pad
		for i := 0; i < length-resLen-1; i++ {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			res += string(base[rand.Intn(decimal)])
		}
	}
	return res
}

func ToId(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)

	// 查找补位字符的位置
	isPad := strings.Index(code, pad)
	if isPad != -1 {
		lenCode = isPad
	}

	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) != pad {
			index := baseMap[code[i]]
			b := uint64(1)
			for j := 0; j < r; j++ {
				b *= decimal
			}
			// pow 类型为 float64 类型转换太麻烦所以自己循环实现pow的功能
			//res += float64(index) * math.Pow(float64(32)float64(2))
			res += index * b
			r++
		}
	}
	return res
}
