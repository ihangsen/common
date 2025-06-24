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
	SystemErr         string
	PermissionDenied  string
	RefreshValid      string
	NotLogin          string
	TooEarly          string
	TokenValid        string
	LoginErr          string
	AccountErr        string
	AccountNotExist   string
	NoAuditPass       string
	AuditIngErr       string
	VerifyErr         string
	VerifyCodeExpired string
	SendValid         string
	VerifyCodeErr     string
	SendFailed        string
	RegisterFailed    string
	IDCardErr         string
	IDCardFrontalErr  string
	IDCardBackErr     string
	GetNumFailed      string
	CurrentPwdErr     string
	PwdInvalid        string
	UploadCategoryErr string
	UploadFailed      string
	DeleteFailed      string
	PwdFormatErr      string
	UserFormatErr     string
	TimeErr           string
	PayWayErr         string
	PayErr            string
	NotFilled         string
	SmsCategoryErr    string
	NonExistentInfo   string
	OperationFailure  string
	Executed          string
	ParamErr          string
	StatusErr         string
	BalanceErr        string
	BlackErr          string
	LimitErr          string
	MemberErr         string
	SuperAdmin        string
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
