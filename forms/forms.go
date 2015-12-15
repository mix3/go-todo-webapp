package forms

import (
	"net/http"

	"github.com/mholt/binding"
)

type CreateForm struct {
	Text string
}

func (cf *CreateForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.Text: binding.Field{
			Form:     "text",
			Required: true,
		},
	}
}

type SwitchForm struct {
	Id int
}

func (sf *SwitchForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&sf.Id: binding.Field{
			Form:     "id",
			Required: true,
		},
	}
}

type DeleteForm struct {
	Id []int
}

func (df *DeleteForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&df.Id: binding.Field{
			Form:     "id",
			Required: true,
		},
	}
}
