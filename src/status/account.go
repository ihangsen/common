package status

// 系统统一返回code
const (
	OK               = uint8(1) // 请求成功
	SystemError      = uint8(2) // 服务器繁忙,请稍后再试
	TokenValid       = uint8(3) // Token 过期
	RefreshValid     = uint8(4) // 账户已在其他设备登录
	TooEarly         = uint8(5) // 刷新太早
	NotLogin         = uint8(6) // 未登录
	PermissionDenied = uint8(7) // 没有权限
	CustomCode       = uint8(8) // 自定义异常
	ParamCode        = uint8(9) // 请求参数异常
)

// 用户状态
const (
	// todo 修改router
	ACCOUNT_NORMAL = iota + uint8(1) //账户正常
	//ACCOUNT_ENABLED                   		//账户不可用
	//ACCOUNT_NON_LOCKED                        //账户锁定
	//ACCOUNT_NON_EXPIRED                       //账户过期
	//CREDENTIALS_NON_EXPIRED                   //凭证过期
)
