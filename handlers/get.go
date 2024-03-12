package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetUserList(c *gin.Context) {
	id := c.Param("id")
	userList, err := h.userService.ReturnUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userList)
}
