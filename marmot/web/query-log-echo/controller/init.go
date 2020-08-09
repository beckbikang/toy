package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Uni      *ut.UniversalTranslator
	ValidateBiz *validator.Validate
	Trans ut.Translator
)


func init(){
	z := zh.New()
	Uni = ut.New(z,z)
	Trans, _ = Uni.GetTranslator("zh")
	ValidateBiz = validator.New()
	zht.RegisterDefaultTranslations(ValidateBiz, Trans)
}