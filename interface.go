package i18n

import (
	"context"
)

// GinI18n ...
type GinI18n interface {
	GetMessage(param interface{}) (string, error)
	MustGetMessage(param interface{}) string
	SetCurrentContext(ctx context.Context)
	SetBundle(cfg *BundleCfg)
	SetGetLngHandler(handler GetLngHandler)
}
