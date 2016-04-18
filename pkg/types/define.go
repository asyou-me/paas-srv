package types

const (
	// 版本信息
	ApiVersion = "v0.0.1"
	Info       = ""
	// 默认集群编号
	DefaultRegion = "shanghai"
)

var (
	Err403 = Event{
		Code: 403,
		Type: "err",
	}
	Err404 = Event{
		Code: 404,
		Type: "err",
	}
)
