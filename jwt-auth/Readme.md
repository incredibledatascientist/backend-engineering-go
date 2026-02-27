- cmd> go mod init projname
- go get github.com/gin-gonic/gin

internal/
-handler/ ← HTTP layer
-service/ ← Business logic
-auth/ ← JWT + hashing logic
-repository/ ← DB access
-domain/ ← Entities
