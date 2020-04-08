package util

import (
	"easyweb/lib/erro"
	"easyweb/lib/htt"
)

func Validate(params map[string]interface{}, ctx htt.HttpCtx, errors map[string]erro.Error) int{
	for key, err := range errors{
		if params[key] == nil {
			ctx.Send(err.ToMap())
			break
		}
	}

	return 0
}