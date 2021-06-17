package basecontroller

import (
	"fmt"
)

//Controller  controller基类
type Controller struct{}

func init() {
	fmt.Println("base controller init")
}
