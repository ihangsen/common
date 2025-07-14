package i18n

import "sync"

type Language string

const (
	Cn Language = "cn"
	En Language = "en"
)

var (
	Get = chinese
)

var once sync.Once

type content struct {
	SystemErr        string
	PermissionDenied string
	RefreshValid     string
	NotLogin         string
	TooEarly         string
	TokenValid       string
	UserFormatErr    string
	PwdFormatErr     string
	TimeErr          string
}

func Init(language Language) {
	once.Do(func() {
		switch language {
		case Cn:
			Get = chinese
		case En:
			Get = english
		default:
			Get = chinese
		}
	})
}
