package parsed

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/pkg/parsed/gozeroApiParser"
)

func ParseGoZeroApiByFile(filePath string) (*spec.ApiSpec, error) {
	api, err := gozeroApiParser.Parse(filePath)
	if err != nil {
		fmtd.Error(err)
		return nil, err
	}
	if err := api.Validate(); err != nil {
		fmtd.Error(err)
		return nil, err
	}
	return api, nil
}
