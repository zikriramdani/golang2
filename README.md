## Golang Login

go mod init chapter-c32
go get -u

### Salah satu dari packages tersebut adalah jwt-go, yang digunakan untuk keperluan JWT.
go get -u github.com/dgrijalva/jwt-go@v3.2.0
go get -u github.com/novalagung/gubrak/v2

### Jalankan aplikasi, lalu test menggunakan curl.

curl -X POST --user zikri:zikri123 http://localhost:8080/login
curl -X GET \
    --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzg2NTI0NzksImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwiVXNlcm5hbWUiOiJub3ZhbCIsIkVtYWlsIjoidGVycGFsbXVyYWhAZ21haWwuY29tIiwiR3JvdXAiOiJhZG1pbiJ9.JREJgUAYs2R5zuquqyX8tk23QlajVVe19susm6JsZq8" \
    http://localhost:8080/index

### go run main.go

### http://localhost:8080