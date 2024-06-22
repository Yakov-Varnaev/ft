package web

import "github.com/gin-gonic/gin"

func GetDataFromBody[T any](c *gin.Context) (*T, error) {
	var data T
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return nil, &ValidationError{Message: err.Error()}
	}
	return &data, nil
}
