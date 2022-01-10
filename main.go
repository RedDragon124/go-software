package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
	"github.com/RedDragon124/go-software/pkg/myname"
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	fmt.Println("")
	fmt.Println(myname.returnManas())
}
