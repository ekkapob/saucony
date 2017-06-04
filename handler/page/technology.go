package page

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type TechTmpl struct {
	model.Tmpl
	Technologies []string
}

func Technology(w http.ResponseWriter, r *http.Request) {
	tmpl := helper.InitTemplate(w, r)
	t := handler.BaseTemplate("technology.tmpl", nil)
	techTmpl := TechTmpl{
		Tmpl:         tmpl,
		Technologies: []string{"all"},
	}
	t.ExecuteTemplate(w, "main", techTmpl)
}
