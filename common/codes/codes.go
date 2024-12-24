package codes

type Code uint32

const (
	Ok            Code = 200
	InternalError Code = 500
	UnKnown       Code = 400
)

// for verify code
const (
	GenCodeFast Code = 10000 + iota
	GenCodeMany
)

// for user
const (
	VerifyCodeErr Code = 20000 + iota
	TelAlreadyExists
	TelNotRegistry
	PassWordError
)
