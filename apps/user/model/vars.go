package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	NormalUser byte = 0b0001
	MerChant   byte = 0b0010
	Admin      byte = 0b0100
)

const (
	StatusOk uint64 = 0
	StatusNo uint64 = 1
)
