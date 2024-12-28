package codes

type Code uint32

const (
	Ok                  Code = 200
	InternalError       Code = 500
	UnKnown             Code = 400
	InvalidParams       Code = 100
	NotFoundMetaData    Code = 101
	InvalidToken        Code = 102
	InvalidTokenPayLoad Code = 103
	BuildTokenError     Code = 104
	RateLimit           Code = 105
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

// for goods
const (
	SkuAlreadyExists Code = 30000 + iota
	SkuNotRegistry
)

// for store
const (
	StoreNotRegistry Code = 40000 + iota
	GoodsNotFound
)

// for order
const (
	OrderNotFound Code = 50000 + iota
	OrderPayed
	OrderExpire
	InvalidStoreIdAndSku
	StockNotEnough
)

// for review
const (
	InvalidOrderStatus Code = 60000 + iota
	SkuNotInOrder
	NotFoundReview
)

// for reply
const (
	ReviewNotPassAudit Code = 70000 + iota
	InvalidSSku
	NotFoundReply
)
