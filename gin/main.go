package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	return str1 + str2
}
func main() {
	r := gin.Default()
	//自定义模版函数，必须在loadhtmlglob前面
	// r.SetFuncMap(template.FuncMap{
	// 	"UnixToTime": UnixToTime,
	// 	"Println":    Println,
	// })
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//加载templates中的模版文件，使用不同目录下的名称相同的模版,注意:一定要放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录,第一个参数表示路由,第二个参数表示映射的目录
	r.Static("/static", "./static")
	//前台路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"msg":   "我是msg",
			"hobby": []string{"吃饭", "唱歌", "写代码"},
			"newList": []interface{}{
				&Article{
					Title:   "新闻标题1111",
					Content: "新闻详情1111",
				},
				&Article{
					Title:   "新闻标题2222",
					Content: "新闻详情2222",
				},
			},
			"date": 1701328649,
		})
	})

	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"Title": "新闻页面",
			"News":  news,
		})
	})
	r.Run()
}
