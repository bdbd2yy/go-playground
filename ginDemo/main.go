// you can find the tutorial here: https://geektutu.com/post/quick-go-gin.html
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func main() {
    r := gin.Default()
    r.GET("/", func (c *gin.Context)  {
        c.String(http.StatusOK, "who are you")
    })
    // dynamic routine
    r.GET("/user/:name", func(ctx *gin.Context) {
        name := ctx.Param("name")
        ctx.String(http.StatusOK, "hello %s", name)
    })
    r.GET("/users", func(ctx *gin.Context) {
        name := ctx.Query("name")
        // set the default value if the key does not exist
        role := ctx.DefaultQuery("role", "teacher")
        ctx.String(http.StatusOK, "%s is a %s", name, role)
    })

    // POST
    r.POST("/form", func(ctx *gin.Context) {
        username := ctx.PostForm("username")
        password := ctx.DefaultPostForm("password", "000000")

        ctx.JSON(http.StatusOK, gin.H{
            "username": username,
            "password": password,
        })
    })

    // combine the post and query
    r.POST("/posts", func(ctx *gin.Context) {
        id := ctx.Query("id")
        page := ctx.DefaultQuery("page", "0")
        username := ctx.PostForm("username")
        password := ctx.DefaultPostForm("password", "000000")

        ctx.JSON(http.StatusOK, gin.H{
            "id": id,
            "page": page,
            "username": username,
            "password": password,
        })
    })

    // map
    r.POST("/post", func(ctx *gin.Context) {
        ids := ctx.QueryMap("ids")
        names := ctx.PostFormMap("names")

        ctx.JSON(http.StatusOK, gin.H{
            "ids": ids,
            "names": names,
        })
    })

    // redirect
    r.GET("/redirect", func(ctx *gin.Context) {
        ctx.Redirect(http.StatusPermanentRedirect, "/index")
    })

    r.GET("/goindex", func(ctx *gin.Context) {
        ctx.Request.URL.Path = "/"
        r.HandleContext(ctx)
    })

    // group routes
    defaultHandler := func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "path": ctx.FullPath(),
        })
    }

    // v1 := r.Group("/v1", defaultHandler)
    // if you add a default handler to the group, then all the child routes will execute it.
    // it is not recommended to add the handler that actually respond the request.
    v1 := r.Group("/v1")
    {
        v1.GET("/posts", defaultHandler)
        v1.GET("/series", defaultHandler)
    }

    v2 := r.Group("/v2")
    {
        v2.GET("/posts", defaultHandler)
        v2.GET("/series", defaultHandler)
    }

    // upload a file
    r.POST("/upload1", func(ctx *gin.Context) {
        file, _ := ctx.FormFile("file")
        // ctx.SaveUploadedFile(file, dst)
        ctx.String(http.StatusOK, "%s uploaded",  file.Filename)
    })

    // upload multiple files
    r.POST("/upload2", func(ctx *gin.Context) {
        form, _ := ctx.MultipartForm()
        // the "upload[]" is just front-end filed name
        // files's type is pointer array
        files := form.File["upload[]"]

        for _, file := range files {
            log.Println(file.Filename)
        }

        ctx.String(http.StatusOK, "%d files uploaded!", len(files))
    })

    // html template
    // you can define a struct within a funciton for temporary internal use
    type student struct {
        Name string
        Age int8
    }

    r.LoadHTMLGlob("templates/*")

    stu1 := &student{"bdbd", 20}
    stu2 := &student{"wzy", 19}

    r.GET("/arr", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "arr.tmpl", gin.H{
            "title": "Gin",
            "stuArr": [2]*student{stu1, stu2},
        })
    })

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.Run("127.0.0.1:8081")

    // apply to a single route
    // r.GET("/benchmark", MyBenchLogger(), benchendpoint)
    // autherized := r.Group("/")
    // autherized.Use(AuthRequied())
}

// customize middleware
func Logger() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        t := time.Now()
        ctx.Set("bdbd", "1011")
        ctx.Next()
        latency := time.Since(t)
        log.Print(latency)
    }
}
