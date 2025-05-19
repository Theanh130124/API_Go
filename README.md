# API_Go

//Có thể sửa port 3308 -> port trong máy mình   (tag mysql là đang lấy bảng mới nhất  -> có thể mysql:8.1.0 ...  -> xem trên dockerhub)

docker ps

docker run --name theanh-mysql -e MYSQL_ROOT_PASSWORD=tta1301 -p 3308:3306 -d mysql


truy vấn thì nên đánh index -> trên các cột có where -> hay dùng where để lọc như status 


dùng explain
select  * from ... where status = 'Doing' 