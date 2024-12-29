package middleware

import (
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/time/rate"
	"net/http"
)

type RateLimitMiddleware struct {
	Limiter *rate.Limiter
}

func NewRateLimitMiddleware(Limiter *rate.Limiter) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		Limiter: Limiter,
	}
}

func (m *RateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !m.Limiter.Allow() {
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.RateLimit, "rate limit"))
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
