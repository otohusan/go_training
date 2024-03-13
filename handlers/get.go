package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) ReturnUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.ReturnUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserList(c *gin.Context) {
	userList, err := h.userService.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userList)
}
