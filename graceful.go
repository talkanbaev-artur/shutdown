package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

//Wait is warapper for graceful shutdown. It waits for OS signals or for external Context cancel call
//
// Warning! It uses the global zap Sugered Logger and requires it as a dependency!
// Make sure you configure the global logger with zap.ReplaceGlobals()
func Wait(ctx context.Context, cancel context.CancelFunc, shutdown *Shutdown) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		oscall := <-c
		zap.S().Infow("Catched shutdown sygnal", "Signal", oscall)
		cancel()
	}()

	zap.S().Infow("Initialised authentication server", "ConfigStatus", "OK", "ServiceStatus", "OK")
	<-ctx.Done()
	errs := shutdown.Close() // initialise the shutdown process and print all errors
	if len(errs) != 0 {
		for _, v := range errs {
			zap.S().Errorw("Shutdown error", "Error msg", v)
		}
	}
}
