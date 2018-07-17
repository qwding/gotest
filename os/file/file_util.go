package file

import (
	"fmt"
	"os"
)

//makedir 不是智能递归创建的
func MakeDir(name) error {
	err := os.Mkdir(name, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
