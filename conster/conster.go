package conster

import "fmt"

func ConstInsideFunction() string {
	const queryInside = `this is a const query this query the table %s`
	return fmt.Sprintf(queryInside, "inside")
}

const queryOutside = `this is a const query this query the table %s`
func ConstOutsideFunction() string {
	return fmt.Sprintf(queryOutside, "outside")
}

