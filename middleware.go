package main

import "github.com/gin-gonic/gin"

func RequestIDMiddleWare(ctx *gin.Context) {
	h := ctx.Request.Header
	if _, ok := h["request_id"]; !ok {
		h["request_id"] = []string{GenRequestId()}
	}

	ctx.Next()
}
