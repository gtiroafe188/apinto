package httplog

import (
	"reflect"

	"github.com/eolinker/eosc"
)

//Register 注册httplog驱动工厂
func Register(register eosc.IExtenderDriverRegister) {
	register.RegisterExtenderDriver("log_httplog", NewFactory())
}

type factory struct {
}

//NewFactory 创建httplog驱动工厂
func NewFactory() eosc.IExtenderDriverFactory {
	return &factory{}
}

//Create 创建httplog驱动
func (f *factory) Create(profession string, name string, label string, desc string, params map[string]string) (eosc.IExtenderDriver, error) {
	return &driver{
		profession: profession,
		name:       name,
		label:      label,
		desc:       desc,
		driver:     driverName,
		configType: reflect.TypeOf((*DriverConfig)(nil)),
		params:     params,
	}, nil
}
