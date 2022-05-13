package base

//go:generate stringer -type=ErrCode --linecomment
type ErrCode int

const (
	Success ErrCode = 100 // Operation succeed

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

)
