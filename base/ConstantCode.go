package base

//go:generate stringer -type=ErrCode --linecomment
type ErrCode int

const (
	Success ErrCode = 0 // Operation succeed

	InputError        ErrCode = 1000 // Input Error
	OpenSqlError      ErrCode = 1001 // Error open Mysql database
	ConnectSqlError   ErrCode = 1002 // Cannot connect to mysql database
	InsertError       ErrCode = 1003 // Insert data error
	QueryError        ErrCode = 1004 // Query db error
	WrongLoginError   ErrCode = 1005 // Wrong username or passwd
	AccountExistError ErrCode = 1006 // Account is already registered
	PasswdUpdateError ErrCode = 1007 // Update password failed

	AuthFailed      ErrCode = 1101 // Permission denied, lack token
	AuthFormatError ErrCode = 1102 // The auth format in the request header is incorrect
	InvalidToken    ErrCode = 1103 // The token has expired or is invalid or could not parse with claims
	GenTokenError   ErrCode = 1104 // Generate token error: Sign Token Failed
	LackTokenError  ErrCode = 1105 // Lack token in request header
	UserTokenError  ErrCode = 1105 // Wrong account or password, maybe the password has been changed

	GetPersonError        ErrCode = 1201 // Can not get user info from database
	StoreEmailTokenError  ErrCode = 1202 // Store Email Token Error
	ActivateEmailError    ErrCode = 1203 // Activate email token failed
	SendEmailError        ErrCode = 1204 // Error sending activation email to the user
	TokenNotActivated     ErrCode = 1205 // Token not activated, please rerun activation email
	TokenInvalidError     ErrCode = 1206 // Token invalid: Overtime
	TokenNotExist         ErrCode = 1207 // Could not find this token in database
	TokenAlreadyActivated ErrCode = 1208 // Token has already been activated

	WrongVerifyCode ErrCode = 1301 // The verify code is invalid or expired

	ServerError ErrCode = 2000 // Server Error

	UserIDNotExist	ErrCode = 3000
	UserProfileUpdateError ErrCode = 3001 // Update profile failed
	BalanceNotEnough		ErrCode = 3002// Sorry, your credit is running low
	UpdateBalanceError		ErrCode = 3003//Transfer failed

	EditItemError		ErrCode = 4001
	GetItemError		ErrCode = 4002


)
