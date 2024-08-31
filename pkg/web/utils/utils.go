package utils

import (
	"fmt"

	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUIDFromParam(c *gin.Context, key ...string) (string, error) {
	paramKey := "id"
	if len(key) > 0 {
		paramKey = key[0]
	}
	id := c.Param(paramKey)
	if _, err := uuid.Parse(id); err != nil {
		return "", &webErrors.BadRequest{
			Message: fmt.Sprintf("invalid uuid: %v", err),
		}
	}
	return id, nil
}
