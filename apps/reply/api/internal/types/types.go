// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ProposeReplyReq struct {
	ReviewId uint64 `json:"review_id" validate:"required`
	StoreId  uint64 `json:"store_id" validate:"required`
	Sku      string `json:"sku" validate:"required`
	Content  string `json:"content" validate:"required`
	HasImage uint8  `json:"has_image" validate:"required`
	ImageCDN string `json:"image_cdn,optional"`
}

type ProposeReplyResp struct {
	ReplyId uint64 `json:"reply_id"`
	Status  uint8  `json:"status"`
}

type ReplyByMidResp struct {
	Infos []ReplyInfo `json:"infos"`
}

type ReplyByReplyIdReq struct {
	ReplyId uint64 `json:"reply_id" validate:"required`
}

type ReplyByReplyIdResp struct {
	Info ReplyInfo `json:"info"`
}

type ReplyByReviewIdReq struct {
	ReviewId uint64 `json:"review_id" validate:"required`
}

type ReplyByReviewIdResp struct {
	Infos []ReplyInfo `json:"infos"`
}

type ReplyBySSkuReq struct {
	StoreId uint64 `json:"store_id" validate:"required"`
	Sku     string `json:"sku" validate:"required`
}

type ReplyBySSkuResp struct {
	Infos []ReplyInfo `json:"infos"`
}

type ReplyInfo struct {
	Mid      uint64 `json:"mid"`
	ReviewId uint64 `json:"review_id"`
	StoreId  uint64 `json:"store_id"`
	Sku      string `json:"sku"`
	Content  string `json:"content"`
	HasImage uint32 `json:"has_image"`
	ImageCDN string `json:"image_cdn"`
	ReplyId  uint64 `json:"reply_id"`
	Status   uint8  `json:"status"`
	CreateAt string `json:"create_at"`
	IsDel    uint8  `json:"is_del"`
}
