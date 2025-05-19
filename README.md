# API_Go

//Có thể sửa port 3308 -> port trong máy mình   (tag mysql là đang lấy bảng mới nhất  -> có thể mysql:8.1.0 ...  -> xem trên dockerhub)

docker ps

docker run --name theanh-mysql -e MYSQL_ROOT_PASSWORD=tta1301 -p 3308:3306 -d mysql


truy vấn thì nên đánh index -> trên các cột có where -> hay dùng where để lọc như status 


dùng explain
select  * from ... where status = 'Doing' 

Gin Framework 

Thêm vào go.mod 

```
"github.com/gin-gonic/gin"
```

main.go  -> API đầu tiên


``
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
``



<h1>Install GORM </h1>
go get -u gorm.io/gorm

Cài driver với mysql

go get -u gorm.io/driver/mysql
