package middleware

import (
	"cms/helper"
	"cms/model/request"
	"context"
	"github.com/gin-gonic/gin"
)

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Token")
		apiKey := c.GetHeader("apiKey")
		valueApiKey := helper.ReadENV("", "API_KEY")
		if header == "" || apiKey == "" {
			c.AbortWithStatusJSON(400, &request.CustomError{
				Message: "No Header or apiKey",
				Code:    3991,
			})
		} else if apiKey != valueApiKey {
			c.AbortWithStatusJSON(400, &request.CustomError{
				Message: "api key is wrong",
				Code:    3999,
			})
		}
		ctx := context.WithValue(c.Request.Context(), "Token", header )
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
