package self_errors

const (
	//通用
	ErrInvalidParams = 1001
	ErrLackOfRequiredParams = 1002
	ErrServiceNotAvailable = 1003
	ErrRecordNotFound = 1004

	//season模块
	ErrSeasonAlreadyExist = 1101


)

var Messages = map[int64]string{
	//通用
	ErrInvalidParams:"参数错误",
	ErrLackOfRequiredParams:"缺少必填参数",
	ErrServiceNotAvailable:"系统繁忙,请稍后重试",
	ErrRecordNotFound:"记录不存在",

	//配置模块
	ErrSeasonAlreadyExist:"season已存在，请勿重复添加",

}