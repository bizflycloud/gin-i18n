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

	// 	if bundle isn't constructed then assign it from default
	if ins.bundle == nil {
		ins.SetBundle(defaultBundleConfig)
	}

	// if getLngHandler isn't constructed then assign it from default
	if ins.getLngHandler == nil {
		ins.getLngHandler = defaultGetLngHandler
	}

	return ins
}

func CloneGinI18n(ins GinI18n) GinI18n {
	if ins == nil {
		return nil
	}
	castedIns, ok := ins.(*GinI18nImpl)
	if !ok {
		return ins
	}
	newIns := &GinI18nImpl{
		bundle:          castedIns.bundle,
		currentContext:  castedIns.currentContext,
		localizerByLng:  castedIns.localizerByLng,
		defaultLanguage: castedIns.defaultLanguage,
		getLngHandler:   castedIns.getLngHandler,
	}
	return newIns
}

// Localize ...
func Localize(opts ...Option) gin.HandlerFunc {
	atI18n := NewI18n(opts...)
	return func(context *gin.Context) {
		newAtI18n := CloneGinI18n(atI18n)
		newAtI18n.SetCurrentContext(context)
		context.Set("i18n", newAtI18n)
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
