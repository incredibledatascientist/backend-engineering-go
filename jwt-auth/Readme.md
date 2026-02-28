- cmd> go mod init projname
- go get github.com/gin-gonic/gin

internal/
-handler/ ← HTTP layer
-service/ ← Business logic
-auth/ ← JWT + hashing logic
-repository/ ← DB access
-domain/ ← Entities

// File server
r.PathPrefix("/static/"). // 1. Match all URLs starting with /static/
Handler(http.StripPrefix( // 2. Remove "/static" from the URL
"/static",
http.FileServer(http.Dir("internal/static")), // 3. Serve files from this folder
))

req: /static/login.html
Strip "/static" → becomes "/login.html"
FileServer serves → internal/static/login.html

\d movie : desc of the table
