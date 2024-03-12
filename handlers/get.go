package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetUserList(c *gin.Context) {
	userList, err := h.userService.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userList)
}
