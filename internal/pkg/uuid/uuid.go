package uuid

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

var (
	chars   = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	letters = []rune("abcdefghijklmnopqrstuvwxyz")
)

func GetUuid(prefix string) string {
	u1 := ksuid.New()
	if prefix == "" {
		prefix = fmt.Sprintf("%s%s", string(letters[rand.Intn(len(letters))]), randString(4))
	}
	return strings.ToLower(prefix + randString(4) + u1.String()[4:])
}

func randString(lenNum int) string {
	str := strings.Builder{}
	length := len(chars)
	rand.Seed(time.Now().UnixNano()) // 重新播种，否则值不会变
	for i := 0; i < lenNum; i++ {
		str.WriteString(chars[rand.Intn(length)])
	}
	return str.String()
}
