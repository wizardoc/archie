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

// Organization
var ORGANIZATION_FIND_EMPTY = ArchieError{2008, "Cannot find organizations"}

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
