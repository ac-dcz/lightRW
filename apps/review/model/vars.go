package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	AuditIng  = 10
	AuditSuc  = 20
	AuditFail = 30
)
