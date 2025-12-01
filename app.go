package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
)

//go:embed swagger-ui
var content embed.FS

func logReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Infof("New request to: '%s %s'", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func newApp() http.Handler {
	Logger.Info("Init the backend")

	router := http.NewServeMux()
	// Root
	router.HandleFunc("/", getHomeHandler)

	// Collection endpoints
	router.HandleFunc("/api/cats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			makeHandlerFunc(createCat).ServeHTTP(w, r)
			return
		}
		if r.Method == http.MethodGet {
			makeHandlerFunc(listCats).ServeHTTP(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	// Item endpoints (GET, DELETE)
	router.HandleFunc("/api/cats/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			makeHandlerFunc(getCat).ServeHTTP(w, r)
		case http.MethodDelete:
			makeHandlerFunc(deleteCat).ServeHTTP(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fsys, _ := fs.Sub(content, "swagger-ui")
	router.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(http.FS(fsys))))

	return logReq(router)
}

// Simpler way to handle requests
type ServiceFunc func(*http.Request) (int, any)

// Wraps the ServiceFunc to make a http.HandlerFunc with panic handling and JSON response encoding
func makeHandlerFunc(svcFunc ServiceFunc) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {

		code, body := func(req *http.Request) (code int, body any) {
			// General panic/error handler to keep the server up
			defer func() {
				if recov := recover(); recov != nil {
					Logger.Error("Recovering from a panic: ", recov)
					// Using the named return values
					code = http.StatusInternalServerError
					body = http.StatusText(code)
				}
			}()
			return svcFunc(req)
		}(req)

		// Single response
		res.Header().Set("content-type", "application/json")
		res.WriteHeader(code)
		json.NewEncoder(res).Encode(body)
	}
}
