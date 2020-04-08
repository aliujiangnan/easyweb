package erro

import (
	"strconv"
)

type Error struct{
	id 		int
	msg		string 
}

func NewError(id int, msg string)Error{
	return Error{id, msg}
}

func (this Error)ToString() string{
	return `{"id":` + strconv.Itoa(this.id) + `"msg":"` + this.msg + `"}`
}

func (this Error)ToMap() map[string]interface{}{
	return map[string]interface{}{
		"id":this.id,
		"msg":this.msg,
	}
}