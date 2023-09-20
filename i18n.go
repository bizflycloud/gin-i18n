package i18n

import (
	"github.com/gin-gonic/gin"
)

// NewI18n ...
func NewI18n(opts ...Option) GinI18n {
	// init ins
	ins := &GinI18nImpl{}

	// set ins property from opts
	for _, opt := range opts {
		opt(ins)
	}

	// 	if Bundle isn't constructed then assign it from default
	if ins.Bundle == nil {
		ins.SetBundle(defaultBundleConfig)
	}

	// if getLngHandler isn't constructed then assign it from default
	if ins.getLngHandler == nil {
		ins.getLngHandler = defaultGetLngHandler
	}

	return ins
}

// Localize ...
func Localize(opts ...Option) gin.HandlerFunc {
	atI18n := NewI18n(opts...)
	return func(context *gin.Context) {
		context.Set("i18n", atI18n)
		atI18n.SetCurrentContext(context)
	}
}

/*
GetMessage get the i18n message

	 param is one of these type: messageID, *i18n.LocalizeConfig
	 Example:
		GetMessage(context, "hello") // messageID is hello
		GetMessage(context, &i18n.LocalizeConfig{
				MessageID: "welcomeWithName",
				TemplateData: map[string]string{
					"name": context.Param("name"),
				},
		})
*/
func GetMessage(context *gin.Context, param interface{}) (string, error) {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.GetMessage(param)
}

/*
MustGetMessage get the i18n message without error handling

	  param is one of these type: messageID, *i18n.LocalizeConfig
	  Example:
		MustGetMessage(context, "hello") // messageID is hello
		MustGetMessage(context, &i18n.LocalizeConfig{
				MessageID: "welcomeWithName",
				TemplateData: map[string]string{
					"name": context.Param("name"),
				},
		})
*/
func MustGetMessage(context *gin.Context, param interface{}) string {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.MustGetMessage(param)
}
