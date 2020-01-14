package robust

import "gopkg.in/go-playground/validator.v9"

/**
* 封装 Validation，返回一个 ArchieError，统一返回接口
 */
var validate *validator.Validate

type Validation struct {
	Target interface{}
}

func (v *Validation) Valid() error {
	if err := validate.Struct(v.Target); err != nil {
		return ArchieError{Code: 4000, Msg: err.Error()}
	}
	return nil
}

func init() {
	validate = validator.New()
}
