package exp

import "fmt"

type Context interface {
	GetHost() string
	GetUsername() string
	GetPassword() string
	Log(line string)
}

type ContextImpl struct {
	Host     string
	Username string
	Password string
}

func NewContext(host string, username string, password string) Context {
	return &ContextImpl{
		Host:     host,
		Username: username,
		Password: password,
	}
}

func (_this *ContextImpl) GetHost() string {
	if _this.Host == "" {
		fmt.Printf("    - 请输入光猫地址 (默认: 192.168.1.1) : ")
		_, _ = fmt.Scanln(&_this.Host)
		if _this.Host == "" {
			_this.Host = "192.168.1.1"
		}
	}
	return _this.Host
}

func (_this *ContextImpl) GetUsername() string {
	if _this.Username == "" {
		fmt.Printf("    - 请输入普通管理员账号 (默认: useradmin) : ")
		_, _ = fmt.Scanln(&_this.Username)
		if _this.Username == "" {
			_this.Username = "useradmin"
		}
	}
	return _this.Username
}

func (_this *ContextImpl) GetPassword() string {
	if _this.Password == "" {
		fmt.Printf("    - 请输入普通管理员密码 : ")
		_, _ = fmt.Scanln(&_this.Password)
	}
	return _this.Password
}

func (_this *ContextImpl) Log(line string) {
	fmt.Printf("    - %s\n", line)
}

type Result struct {
	Username string
	Password string
	Config   []byte
}

type Exp struct {
	Name   string
	Handle func(context Context) (*Result, error)
}
