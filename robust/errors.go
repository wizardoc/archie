package robust

type ArchieError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var CANNOT_FIND_ORGANIZATION = ArchieError{1001, "cannot find organization"}
var CONNOT_CREATE_ORGANIZATION = ArchieError{1002, "could't create new organization"}

// Account
var REGISTER_FAILURE = ArchieError{1003, "register fail"}
var REGISTER_EXIST_USER = ArchieError{1004, "the user is exist"}

// JWT
var JWT_DOES_NOT_EXIST = ArchieError{4001, "jwt does not exist"}
var JWT_PARSE_ERROR = ArchieError{4002, "cannot parse jwt"}
var JWT_NOT_ALLOWED = ArchieError{4003, "the jwt is not allowed"}

// DB
var CREATE_DATA_FAILURE = ArchieError{3001, "create data to db failure"}
