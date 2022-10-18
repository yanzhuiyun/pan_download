package mysql

import (
	"fmt"
	"pandownload/settings"
	"testing"
)

func TestGetfiles(t *testing.T) {
	settings.Init()
	Init()
	fmt.Println(Getfiles("zhangsan"))
}
