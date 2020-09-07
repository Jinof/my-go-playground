package typer

import "fmt"

func TestTypeInside(data string) string {
	type reqInside struct {
		Data string
	}

	req := &reqInside{Data: data}

	return fmt.Sprintf("data is %s", req.Data)
}

type reqOutside struct {
	Data string
}
func TestTypeOutside(data string) string {

	req := &reqOutside{Data: data}

	return fmt.Sprintf("data is %s", req.Data)
}
