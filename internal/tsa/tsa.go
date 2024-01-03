package tsa

import (
	"fmt"
	"net/http"

	"github.com/bzhtux/tsa/models"
	"github.com/gin-gonic/gin"
)

// Get all HTTP Status Codes
func (h *BaseHandler) GetAllHttpStatusCodes(c *gin.Context) {
	var codes []models.HttpStatusCode

	res := h.db.Find(&codes)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "internal server error",
			"data":    nil,
		})
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "204",
			"message": "no content",
			"data":    "{}",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "ok",
			"data":    codes,
		})
	}
}

// Get one HTTP Status Code
func (h *BaseHandler) GetOneHttpStatusCode(c *gin.Context) {
	var code models.HttpStatusCode
	codeID := c.Params.ByName("codename")

	res := h.db.Where("Code = ?", codeID).First(&code)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "bad request",
			"data":    nil,
		})
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "404",
			"message": "not found",
			"data":    nil,
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "ok",
			"data":    code,
		})
	}
}

func (h *BaseHandler) GetIndex(c *gin.Context) {
	var codes []models.HttpStatusCode

	res := h.db.Find(&codes)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "internal server error",
			"data":    nil,
		})
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "204",
			"message": "no content",
			"data":    "{}",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"status":  "200",
			"message": "ok",
			"data":    codes,
			"title":   "Home",
		})
	}
}

func (h *BaseHandler) DisplayCode(c *gin.Context) {
	var code models.HttpStatusCode
	codeID := c.Params.ByName("codename")

	res := h.db.Where("Code = ?", codeID).First(&code)
	if res.RowsAffected == 0 {
		c.Redirect(http.StatusTemporaryRedirect, "/static/404.html")
	} else {
		c.HTML(http.StatusOK, "code.html", gin.H{
			"status":  "200",
			"message": "ok",
			"data":    code,
			"title":   fmt.Sprint(code.Code) + " : " + string(code.Name),
		})
	}
}
