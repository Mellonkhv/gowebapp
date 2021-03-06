package ui

import(
	"fmt"
	"github.com/mellonkhv/gowebapp/model"
	"net/http"
	"net"
	"time"
)

type Config struct {
	Assets http.FileSystem
}

func Start(cfg Config, m *model.Model, listener net.Listener){
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/", indexHandler(m))
	go server.Serve(listener)
}

const indexHTML  = `
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="utf-8">
		<title>Simple Go Web App</title>
	</head>
	<body>
		<div id='root'></div>
	</body
</html>
`

func indexHandler(m *model.Model) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,indexHTML)
	})
}