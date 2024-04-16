
# build the app.
FROM golang:1.22.2
ADD . /.
WORKDIR /app
EXPOSE 8080
CMD ["go", "get", "-u", "github.com/gin-gonic/gin"]
CMD ["go", "run", "/main.go"]
