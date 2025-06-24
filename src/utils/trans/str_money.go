package trans

import (
	"fmt"
)

func PayAliMoney(money uint32) string {
	return fmt.Sprintf("%.2f", float32(money)/float32(1000))
}

func PayWxMoney(money uint32) *int64 {
	u := int64(money / 10)
	return &u
}
