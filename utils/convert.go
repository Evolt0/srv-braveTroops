package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryInt(ctx *gin.Context, query string) (int, error) {
	number := ctx.Query(query)
	return strconv.Atoi(number)
}

func QueryInt64(ctx *gin.Context, query string) (int64, error) {
	number := ctx.Query(query)
	return strconv.ParseInt(number, 10, 64)
}

func QueryUInt64(ctx *gin.Context, query string) (uint64, error) {
	number := ctx.Query(query)
	return strconv.ParseUint(number, 10, 64)
}
