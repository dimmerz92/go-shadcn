package internal

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/dimmerz92/go-shadcn/pkg/lib"
	"github.com/dimmerz92/go-shadcn/pkg/middleware"
)

func DevServer(args []string) {
	cmd := flag.NewFlagSet("dev", flag.ExitOnError)
	port := cmd.Int("p", 8000, "use a port between 1024 and 65535")
	cmd.Parse(args)

	if *port < 1024 || *port > 65535 {
		fmt.Printf("%d out of range, use a port between 1024 and 65535", *port)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/themes/output.css")
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		lib.Render(w, r, http.StatusOK, Index())
	})

	loggedMux := middleware.Chain(mux, middleware.Logger)

	fmt.Printf("listening on port %d\n", *port)
	fmt.Println("go to localhost:8001 (or the proxy you specified) for hot reloading")
	http.ListenAndServe(fmt.Sprintf(":%d", *port), loggedMux)
}
