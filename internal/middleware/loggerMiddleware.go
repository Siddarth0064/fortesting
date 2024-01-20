package middleware

import (
	"context"
	"hospetal/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Middleware struct {
	a *auth.Auth
}

func NewMiddleware(a *auth.Auth) (Middleware, error) {
	return Middleware{a: a}, nil
}

type key string

const TraceIdKey key = "1"

func (m Middleware) Log() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceId := uuid.NewString()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, TraceIdKey, traceId)
		req := c.Request.WithContext(ctx)
		c.Request = req

		log.Info().Str("traceId", traceId).Msg("in log file")
		defer log.Logger.Info().Msg("request processing complete")
		c.Next()

	}

}
