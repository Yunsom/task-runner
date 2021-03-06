package http

import (
	"net/http"
	"os"

	"github.com/mylxsw/task-runner/config"
	"github.com/mylxsw/task-runner/console"
	"github.com/mylxsw/task-runner/http/handler"
	"github.com/mylxsw/task-runner/log"
)

// StartHTTPServer start an http server instance serving for api request
func StartHTTPServer() {

	// print welcome message
	http.HandleFunc("/", handler.Home)
	// check the server status
	http.HandleFunc("/status", handler.Status)
	// push task to task queue
	http.HandleFunc("/push", handler.PushTask)
	// create new task queue
	http.HandleFunc("/queue", handler.NewQueue)

	runtime := config.GetRuntime()
	log.Debug("http listening on %s", console.ColorfulText(console.TextCyan, runtime.Config.HTTP.ListenAddr))
	if err := http.ListenAndServe(runtime.Config.HTTP.ListenAddr, nil); err != nil {
		log.Error("failed listening http on %s: %v", runtime.Config.HTTP.ListenAddr, err)
		os.Exit(2)
	}
}
