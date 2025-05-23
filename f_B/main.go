package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"encoding/gob"
	"encoding/json"

	"xius.com/xfgc/application"
	"xius.com/xfgc/common"
	"xius.com/xfgc/postgres"
	"xius.com/xfgc/xfgc"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

var config postgres.Config

func init() {
	file, err := os.Open("db.config")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}

	if config.Port <= 0 {
		fmt.Println("Invalid Http Port: ", config.Port)
		os.Exit(0)
	}

	if len(config.PGConnStr) == 0 {
		fmt.Println("Postgres DB String Not Found: ")
		os.Exit(0)
	}

	if config.PGPoolSize <= 0 {
		fmt.Println("Invalid Postgres DB Pool PGPoolSize: ", config.PGPoolSize)
		os.Exit(0)
	}

	if config.MaxIdleDuration <= 0 {
		fmt.Println("Invalid Postgres DB MaxIdleDuration: ", config.MaxIdleDuration)
		os.Exit(0)
	}

}

func main() {

	fmt.Printf("Starting XIUS 5GC Web Server %v+ %d -- %s | %s\n", time.Now(), 1, time.Now().Format(time.RFC3339), time.Now().UTC().Format("2006-01-02 15:04:05"))

	common.InitConfig()
	common.InitLogger(true, true, true)

	//postgres.CreatePool( "xfgc", "host=localhost port=5432 user=postgres password=123456 dbname=FiveGCore sslmode=disable", 5, 40, false);
	postgres.CreatePool(config)
	//postgres.CreatePool("xfgc", "host=192.168.149.72 port=5432 user=fgcuser password=fgcpass dbname=xius5gc sslmode=disable", 5, 40, false)

	gob.Register(xfgc.AppUser{})

	router := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(CORSMiddleware())

	xfgc.RegisterCounter_FSubscriber()
	xfgc.InitModule_FSubscriber(router)

	xfgc.RegisterCounter_GnbACLIPAddres()
	xfgc.InitModule_GnbACLIPAddres(router)

	xfgc.RegisterCounter_GnbAclPlmn()
	xfgc.InitModule_GnbAclPlmn(router)

	xfgc.RegisterCounter_PlmnTac()
	xfgc.InitModule_PlmnTac(router)

	xfgc.RegisterCounter_UserType()
	xfgc.InitModule_UserType(router)

	xfgc.RegisterCounter_AccountStatus()
	xfgc.InitModule_AccountStatus(router)

	xfgc.RegisterCounter_AppUser()
	xfgc.InitModule_AppUser(router)

	xfgc.RegisterCounter_Dnn()
	xfgc.InitModule_Dnn(router)

	xfgc.RegisterCounter_IMSIRange()
	xfgc.InitModule_IMSIRange(router)

	application.InitModule_Auth(router)

	router.Static("/assets", "./assets/")
	router.LoadHTMLGlob("html/*")

	router.GET("/", func(c *gin.Context) {

		session := sessions.Default(c)
		x := session.Get("user")

		//fmt.Printf("x=%+v\n", x);

		if x != nil {
			c.HTML(http.StatusOK, "home.html", gin.H{
				"title": "Main website",
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Main website",
			})
		}
	})

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/logout", func(c *gin.Context) {

		session := sessions.Default(c)
		session.Set("user", nil)
		session.Save()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBind(&json); err != nil {
			q := url.Values{}
			q.Set("error", err.Error())
			location := url.URL{Path: "/", RawQuery: q.Encode()}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		user := application.AppWebLogin(json.User, json.Password)

		//if json.User == "admin" {
		if user != nil {

			session := sessions.Default(c)
			session.Set("user", user)
			session.Save()

			//fmt.Printf("user found login succesful\n");

			q := url.Values{}
			location := url.URL{Path: "/home", RawQuery: q.Encode()}
			c.Redirect(http.StatusFound, location.RequestURI())
		} else {
			q := url.Values{}
			q.Set("error", "Invalid-Credentials")

			location := url.URL{Path: "/", RawQuery: q.Encode()}
			c.Redirect(http.StatusFound, location.RequestURI())
		}

		return
	})

	//router.Run(":8480")
	port := strconv.Itoa(config.Port)
	router.Run(":" + port)

	fmt.Printf("Init Completed %v+\n", time.Now())
	for {
		time.Sleep(time.Second)
	}
}
