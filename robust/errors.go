package robust

type ArchieError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (ae ArchieError) Error() string {
	return ae.Msg
}

var CANNOT_FIND_ORGANIZATION = ArchieError{1001, "Cannot find organization"}
var CONNOT_CREATE_ORGANIZATION = ArchieError{1002, "Could't create new organization"}

// Account
var REGISTER_FAILURE = ArchieError{1003, "Register fail"}
var REGISTER_EXIST_USER = ArchieError{1004, "The user is exist"}
var LOGIN_USER_DOES_NOT_EXIST = ArchieError{1005, "The user does not exist"}
var LOGIN_PASSWORD_NOT_VALID = ArchieError{1006, "The password not valid"}
var EMAIL_DOSE_EXIST = ArchieError{1007, "The email does exist"}
var REMOVE_PERMISSION = ArchieError{1008, "Remove permission error"}
var REMOVE_ORG_FAILURE = ArchieError{1009, "Remove organization failure"}
var CANNOT_FIND_USER = ArchieError{1010, "Cannot find this user"}
var CONNOT_FOLLOW_YOURSELF = ArchieError{1011, "Cannot follow ur self"}
var REPEAT_FOLLOW_USER = ArchieError{1012, "Repeat follow"}
var CANNOT_UPDATE_USERINFO = ArchieError{1013, "Cannot update userInfo"}
var INVALID_EMAIL_CODE = ArchieError{1014, "Email verify code is wrong"}
var SEND_VERIFY_CODE_FAILURE = ArchieError{1015, "Send verify code failure"}
var MISSING_PARAMS = ArchieError{1016, "Missing params"}
var NO_VALID_EMAIL = ArchieError{1017, "The email has not been verify"}
var REPEAT_EMAIL = ArchieError{1018, "Repeat email"}
var USER_DOSE_NOT_EXIST = ArchieError{1019, "The user does not exist"}
var ORIGIN_PASSWORD_FAILURE = ArchieError{1020, "The origin password is wrong"}
var REPEAT_PASSWORD = ArchieError{1021, "The new password is equal to origin password"}
var EMAIL_IS_REQUIRED = ArchieError{1022, "The email is required"}
var UPDATE_LOGIN_TIME_ERROR = ArchieError{1023, "Cannot update the time of login"}
var USER_ALREADY_EXIST = ArchieError{1024, "The user does already exist"}

// Messages
var MESSAGE_SIGNAL_NOT_EXIST = ArchieError{6001, "The signal of message operator is not exist"}
var MESSAGE_SEND_FAILURE = ArchieError{6002, "Send message failure"}
var MESSAGE_CANNOT_FIND_TO = ArchieError{6003, "The user does not exist"}
var MESSAGE_SEND_TO_YOURSELF = ArchieError{6004, "Cannot send message to yourself"}

// Organization
var ORGANIZATION_FIND_EMPTY = ArchieError{2008, "Cannot find organizations"}
var ORGANIZATION_CREATE_FAILURE = ArchieError{2009, "Cannot create organization"}
var ORGANIZATION_INVITE_YOURSELF = ArchieError{2010, "Cannot invite yourself"}
var ORGANIZATION_INVITE_EXIST = ArchieError{2011, "The user does exist in the organization"}
var ORGANIZATION_INVITE_ERROR = ArchieError{2012, "Invite Error"}

// DB
var CREATE_DATA_FAILURE = ArchieError{3001, "Create data to db failure"}
var DOUBLE_KEY = ArchieError{3002, "The key is duplicate"}
var DB_WRITE_FAILURE = ArchieError{3003, "DB write failure"}
var DB_READ_FAILURE = ArchieError{3004, "DB read failure"}
var DB_UPDATE_FAILURE = ArchieError{3005, "DB update failure"}

// JWT
var JWT_DOES_NOT_EXIST = ArchieError{4001, "Jwt does not exist"}
var JWT_PARSE_ERROR = ArchieError{4002, "Cannot parse jwt"}
var JWT_NOT_ALLOWED = ArchieError{4003, "The jwt is not allowed"}
var JWT_CANNOT_PARSE_CLAIMS = ArchieError{4004, "Cannot parse the claims"}

// IO
var CANNOT_FIND_FILE = ArchieError{5001, "Connot find this file"}

// Validation
var INVALID_PARAMS = ArchieError{10001, "Invalid params"}
var INVALID_PERMISSION = ArchieError{10002, "Invalid permission"}
