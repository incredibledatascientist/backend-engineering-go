# Multiple Files with the Same Package:
----------------------------------
-> I have api.go and main.go files with the same main package

- If we build and then run, it works because Go compiles all the files at once.
- But when we run main.go, it only runs the current file. To run all files, we have to use [go run .] | [go run *.go]


## ServerConfig:
-----------------
-> ServerConfig helps configure multiple fields so that our NewAPIServer looks cleaner, and we can later change the fields without needing to modify the passed arguments.

Before:
    - func NewAPIServer(addr string, readTimeout, writeTimeout time.Time) *APIServer {}

After:
    func NewAPIServer(cfg ServerConfig) *APIServer {
        return &APIServer{
            Addr:         cfg.Addr,
            ReadTimeout:  cfg.ReadTimeout,
            WriteTimeout: cfg.WriteTimeout,
            IdleTimeout:  cfg.IdleTimeout,
        }
    }
