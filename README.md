# Go-shutdown

A small utility, which migrated from project to project, and now extracted to its own lib.

Basically it is a stack, which is passed to various functions, which require the execution of `defer` clause as a cleanup before shutdown.

### Example

`db.go`

```golang
func ConnectDatabase(shutdown *shutdown.Shutdown) (*sql.DB, error){
    //add the deferred close function to the shutdown struct
    db, _ := sql.Open("postgres", psqlInfo)

    shutdown.Add(db.Close)
}
```

`main.go`

```golang

shutdown := shutdown.NewShutdown()

repo, _ := db.ConnectDatabase(shutdown)
email, _ := utils.NewEmailSender(shutdown)

// at the end after graceful shutdown code
<-ctx.Done()
errs := shutdown.Close() 
if len(errs) != 0 {
	for _, v := range errs {
		logger.Infow("Shutdown error", "Error msg", v)
	}
}
```