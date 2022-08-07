package base

//go:generate stringer -type=ErrCode --linecomment
type ErrCode int

const (
	Success ErrCode = 0 // Operation succeed

	InputError            ErrCode = 1000 // Input Error
	OpenSqlError          ErrCode = 1001 // Error open Mysql database
	ConnectSqlError       ErrCode = 1002 // Cannot connect to mysql database
	InsertError           ErrCode = 1003 // Insert data error
	QueryError            ErrCode = 1004 // Query db error
	WrongLoginError       ErrCode = 1005 // Wrong username or passwd
	AccountExistError     ErrCode = 1006 // Account is already registered
	PasswdUpdateError     ErrCode = 1007 // Update password failed
	InsertProfileError    ErrCode = 1008 // insert into profile error: maybe account already registered
	CreateCollectionError ErrCode = 1009 // Create collection error
	UserDoNotExist        ErrCode = 1010 // user do not exist : could not get this user in database

	AuthFailed      ErrCode = 1101 // Permission denied, lack token
	AuthFormatError ErrCode = 1102 // The auth format in the request header is incorrect
	InvalidToken    ErrCode = 1103 // The token has expired or is invalid or could not parse with claims
	GenTokenError   ErrCode = 1104 // Generate token error: Sign Token Failed
	LackTokenError  ErrCode = 1105 // Lack token in request header
	UserTokenError  ErrCode = 1105 // Wrong account or password, maybe the password has been changed
	EnrollFail      ErrCode = 1106 // Register to Fabric failed
	ResetEmailError ErrCode = 1107 // sending reset email error

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

	UserIDNotExist         ErrCode = 3000 // user do not exist
	UserProfileUpdateError ErrCode = 3001 // Update profile failed
	BalanceNotEnough       ErrCode = 3002 // Sorry, your credit is running low
	UpdateBalanceError     ErrCode = 3003 //Transfer failed
	UpdateBlindBoxError    ErrCode = 3004 //Transfer failed

	EditItemError ErrCode = 4001
	GetItemError  ErrCode = 4002

	EmptyInput        ErrCode = 1401 // Invalid Input: one of the parameters is empty
	EmailFormatError  ErrCode = 1402 // Email format invalid or not related to ust email
	PasswdLengthError ErrCode = 1403 // Passwd not in valid length: must be 8-20
	PasswdFormatError ErrCode = 1404 // Passwd not in valid format: must have at least three types of special characters, upper case letters, lower case letters and numbers

	FileTypeError ErrCode = 1501 // Not in format file type, recommended: jpg png jpeg gif jfif webp mp3 flac mp4 avi
	FileSizeError ErrCode = 1502 // File too large or could not parse and obtain file size

	SigVerifyError ErrCode = 1601 // failed to verify signature
	SigNotFound    ErrCode = 1602 // Lack signature in parameter field
	SigCountError  ErrCode = 1603 // Not enough signature related to URL

	PermissionDenied ErrCode = 1701 // only admin accounts could edit act and owner to edit his collection/nft/transfer

	FabricInvokeError ErrCode = 1801 // invoke fabric failed: check backend log or fabric status

	TransferToError     ErrCode = 1901 // User in transfer to field doesn't exist
	TransferAuthError   ErrCode = 1902 // User in transfer from field doesn't match the item creater
	TransferOwnerChange ErrCode = 1902 // Only the item with the same owner and creater could be used transfer

)
