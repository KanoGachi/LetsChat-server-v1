package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kanogachi/gin/ginday01/dbconn"
)

type msgBody struct {
	Sid      int    `json:"sid" db:"sid"`
	Rid      int    `json:"rid" db:"rid"`
	Msg      string `json:"msg" db:"msg"`
	UnixTime int64  `json:"time" db:"time"`
}

type userInfo struct {
	Uid    int    `json:"uid" db:"uid"`
	Name   string `json:"name" db:"name"`
	Avasrc string `json:"avasrc" db:"avasrc"`
}

func RegisterRouter(router *gin.Engine) {
	{
		// 获取对话消息记录
		router.GET("/conv", func(c *gin.Context) {
			sid, _ := strconv.Atoi(c.Query("sid"))
			rid, _ := strconv.Atoi(c.Query("rid"))
			log.Printf("conversation request s:%v,r:%v\n", sid, rid)
			var msgData []msgBody
			sqlStr := "select sid,rid,msg from conv where (sid,rid) in ((?,?),(?,?)) order by time desc"
			err := dbconn.Database.Select(&msgData, sqlStr, sid, rid, rid, sid)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "query failed",
				})
			}
			c.JSON(http.StatusOK, msgData)
		})
		// 处理发送的新消息
		router.POST("/conv", func(c *gin.Context) {
			var curMsg msgBody
			err := c.BindJSON(&curMsg)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "handle data failed",
				})
			}
			log.Println(curMsg.Sid, " sending to ", curMsg.Rid, ":", curMsg.Msg)
			sqlStr := "insert into conv (sid,rid,msg) values (?,?,?)"
			_, err = dbconn.Database.Exec(sqlStr, curMsg.Sid, curMsg.Rid, curMsg.Msg)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "insert failed",
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}
	{
		// 获取登录用户信息
		router.GET("/user", func(c *gin.Context) {
			passport := c.Query("passport")
			var user userInfo
			sqlStr := "select uid, name, avasrc from user where passport = ?"
			err := dbconn.Database.Get(&user, sqlStr, passport)
			if err != nil || user.Uid == 0 {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "query failed",
				})
			}
			c.JSON(http.StatusOK, user)
		})
	}
	{
		// 获取朋友信息
		router.GET("/friends", func(c *gin.Context) {
			uid := c.Query("uid")
			var fList []userInfo
			log.Println("uid ", uid, " requesting friend list")
			sqlStr := "select u.uid, u.name, u.avasrc from (select fid from friend where uid = ?) t inner join user u on t.fid = u.uid"
			err := dbconn.Database.Select(&fList, sqlStr, uid)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "query failed",
				})
			}
			c.JSON(http.StatusOK, fList)
		})
	}
}
