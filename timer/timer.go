package timer

import (
	"fmt"
	"time"
)

func ShowTime()  {
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
}
