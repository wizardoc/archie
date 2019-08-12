package robust

type ArchieError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var CANNOT_FIND_ORGANIZATION = ArchieError{1001, "Cannot find organization"}
var CONNOT_CREATE_ORGANIZATION = ArchieError{1002, "Could't create new organization"}

// Account
var REGISTER_FAILURE = ArchieError{1003, "Register fail"}
var REGISTER_EXIST_USER = ArchieError{1004, "The user is exist"}
var LOGIN_USER_DOES_NOT_EXIST = ArchieError{1005, "The user does not exist"}

// JWT
var JWT_DOES_NOT_EXIST = ArchieError{4001, "Jwt does not exist"}
var JWT_PARSE_ERROR = ArchieError{4002, "Cannot parse jwt"}
var JWT_NOT_ALLOWED = ArchieError{4003, "The jwt is not allowed"}

// DB
var CREATE_DATA_FAILURE = ArchieError{3001, "Create data to db failure"}
