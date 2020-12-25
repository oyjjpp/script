package nocopy

import (
	"log"
	"testing"
)

func TestNoCopy(t *testing.T) {
	// userinfo := UserInfo{Name: "ouyangjun", Address: "北京"}
	per := Person{
		User: UserInfo{Name: "ouyangjun", Address: "北京"},
	}
	log.Println(&per)
}
