package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	xtime "github.com/gogoods/x/time"
	"github.com/gogoods/x/file"
)

var start_time = time.Now()

func AddRoutesForGin(s *gin.Engine) {
	s.GET("/ok", func(c *gin.Context) {
		c.String(200, "ok")
	})

	s.GET("/up", func(c *gin.Context) {
		m := map[string]interface{}{
			"up":     xtime.Format(start_time),
			"uptime": time.Since(start_time).String(),
		}
		c.JSON(200, &m)
	})

	s.GET("/workdir", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("%s\n", file.SelfDir()))
	})

	s.GET("/myip", func(c *gin.Context) {
		c.String(200, c.ClientIP())
	})
}

func AddRoutes() {
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		m := map[string]interface{}{
			"up":     start_time,
			"uptime": time.Since(start_time),
		}
		bytes, _ := json.Marshal(&m)
		w.Write(bytes)
	})

	http.HandleFunc("/workdir", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s\n", file.SelfDir())))
	})

	http.HandleFunc("/myip", func(w http.ResponseWriter, r *http.Request) {
		ip := GetRemoteIp(r)

		w.Write([]byte(ip))
	})
}

func GetRemoteIp(r *http.Request) string {
	ip := r.Header.Get("Repost-Real-IP") //获取转发前的真实ip

	if len(ip) < 7 {
		ip = r.Header.Get("X-Real-IP")
	}

	if len(ip) < 7 {
		ip = r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip = (strings.Split(ip, ":"))[0]
		}
	}
	return ip
}
