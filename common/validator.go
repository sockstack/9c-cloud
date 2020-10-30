package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
)

var (
	e = en.New()
	u = ut.New(e, e)
)

type Validator struct {
	err error
	message string
	rule Rule
	query IQuery
	trans ut.Translator
}

type IQuery interface {
	Rule() Rule
}

type Rule map[string]map[string]string

func NewValidator() *Validator {

	trans, _ := u.GetTranslator("en")
	engine := binding.Validator.Engine()
	en2.RegisterDefaultTranslations(engine.(*validator.Validate), trans)
	return &Validator{trans: trans}
}

func (v *Validator) Json(ctx *gin.Context, query IQuery) {
	v.validate(ctx.ShouldBindJSON(query)).unwrap()
}

func (v *Validator) Validate(query IQuery) *Validator {
	v.query = query
	return v
}

func (v *Validator) validate(err error) *Validator {
	errs := err.(validator.ValidationErrors)
	rules := v.query.Rule()
	for _, e := range errs {
		v.err = errors.New(e.Translate(v.trans))
		if m, ok := rules[e.Field()]; ok {
			if msg, ok := m[e.ActualTag()]; ok {
				v.err = errors.New(msg)
			}
		}
	}
	return v
}

func (v *Validator) unwrap() {
	if v.err != nil {
		panic(v.err)
	}
}
