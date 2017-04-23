package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/helpers"
	_"github.com/satori/go.uuid"
)

type URL struct {
	UUID    string `json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4(); unique"`
	Serial  int    `json:"-" sql:"auto_increment; primary_key; unique"`
	Address string `json:"address" sql:"type:varchar(500); not null;" valid:"required~Endereço é obrigatório"`
	Alias   string `json:"alias" sql:"type:varchar(30); unique"`
	View    int    `json:"view" sql:"type:integer;"`
}

type URLS struct {
	Urls []URL        `json:"urls"`
	Meta helpers.Meta `json:"meta"`
}

//Verifica se o objetivo está vazio
func (url URL) IsEmpty() bool {
	return url == URL{}
}

//Verifica se o alias está vazio
func (url URL) AliasIsEmpty() bool {
	return url.Alias == ""
}

// Valida a estrutura pelas tags do Govalidator
func (url URL) Validate() []string {
	var errors []string
	if _, err := govalidator.ValidateStruct(url); err != nil {
		errs := err.(govalidator.Errors).Errors()
		for _, element := range errs {
			errors = append(errors, element.Error())
		}
	}

	return errors
}