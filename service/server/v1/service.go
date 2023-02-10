package v1

import "github.com/Carbohz/gtfs-viewer/service/server"

var _ server.CommonService = (*CommonServiceImpl)(nil)

type CommonServiceImpl struct {
}

func NewService() *CommonServiceImpl {
	svc := &CommonServiceImpl{}
	return svc
}

func (c CommonServiceImpl) SomeUsefulMethodInFututre() {
	panic("implement me")
}
