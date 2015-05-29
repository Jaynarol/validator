package validator

/*
 * :field - e.Field | User
 * :tag - e.Tag | min
 * :param - e.Param | 10
 * :value - e.Value | jaynarol
 */
var MsgEn = map[string]string{
	"required":     `The :field field is required.`,
	"len":          `The :field must size be :param.`,
	"min":          `The :field must be at least :param.`,
	"max":          `The :field must not be greater than :param.`,
	"eq":           `The :field is invalid.`,
	"ne":           `The :field is invalid.`,
	"lt":           `The :field must be less than :param.`,
	"lte":          `The :field must be less than or equal :param.`,
	"gt":           `The :field must be greater than :param.`,
	"gte":          `The :field must be greater than or equal :param.`,
	"eqfield":      `The :field and :param must match.`,
	"nefield":      `The :field and :param must not match.`,
	"gtefield":     `The :field and must be greater than or equal :param.`,
	"gtfield":      `The :field and must be greater than :param .`,
	"ltefield":     `The :field and must be less than or equal :param.`,
	"ltfield":      `The :field and must be less than :param.`,
	"alpha":        `The :field must only contain letters.`,
	"alphanum":     `The :field must only contain letters and numbers.`,
	"numeric":      `The :field must be a number.`,
	"number":       `The :field must be a positive number.`,
	"hexadecimal":  `The :field format Hexadecimal is invalid. (ex: 65ADF9998EBB)`,
	"hexcolor":     `The :field format Hex Color is invalid.(ex: #0F2389)`,
	"rgb":          `The :field format RGB Color is invalid. (ex: rgb(230, 255, 0))`,
	"rgba":         `The :field format RGBA Color is invalid. (ex: rgba(200, 54, 54, 0.5))`,
	"hsl":          `The :field format HSL Color is invalid. (ex: hsl(120, 100%, 50%))`,
	"hsla":         `The :field format HSLA Color is invalid. (ex: hsla(0, 100%, 50%, 0.4))`,
	"email":        `The :field must be a valid email address.`,
	"url":          `The :field format URL is invalid.`,
	"uri":          `The :field format URI is invalid.`,
	"base64":       `The :field format Base64 is invalid.`,
	"contains":     `The :field is invalid.`,
	"containsany":  `The :field is invalid.`,
	"containsrune": `The :field is invalid.`,
	"excludes":     `The :field is invalid.`,
	"excludesall":  `The :field is invalid.`,
	"excludesrune": `The :field is invalid.`,
}