package validator

// controllers层的状态封装
type ResCode int64

const (
	CodeInvalidToken  = 9090
	CodeAuthForbidden = 403
)
const (
	CodeSuccess ResCode = 200 + iota
	CodeInvalidParm
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeSqlProblem
	CodeUserCannotBeEmpty

	CodeMenuExist
	CodeCannotBeEmpty
	// jwt
	CodeNeedLogin
	CodeRepeat
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParm:       "请求参数错误",
	CodeUserExist:         "用户名已存在",
	CodeUserNotExist:      "用户名不存在",
	CodeInvalidPassword:   "用户名或者密码错误",
	CodeUserCannotBeEmpty: "用户与密码不能为空",
	CodeServerBusy:        "服务繁忙",
	CodeSqlProblem:        "请检查sql语句",

	CodeMenuExist:     "菜单重复",
	CodeCannotBeEmpty: "不能为空",
	// jwt
	CodeNeedLogin:     "需要登录",
	CodeRepeat:        "重复存在",
	CodeInvalidToken:  "无效token",
	CodeAuthForbidden: "权限拒绝",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
