package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func Upgrader(c *gin.Context) {
	// uid := c.Request.Header.Get("UID")
	// if uid == "" {
	// 	log.Println("no have uid")
	// 	ans := gin.H{
	// 		"err_msg": "no have uid"}
	// 	c.JSON(http.StatusBadRequest, ans)
	// 	return
	// }

	// Upgrade connection
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		return
	}
	if err := epoller.Add(conn); err != nil {
		log.Printf("Failed to add connection %v", err)
		conn.Close()
	}
}
