package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jaxfu/ape/engine/core/bus"
	"github.com/jaxfu/ape/engine/core/server/api"
)

// TODO: wip server get/set routes
type Server struct {
	Config     ServerConfig
	FileServer http.Handler
	Api        api.Api
}

type ServerConfig struct {
	FullUrl   string
	BaseUrl   string
	Port      uint
	ClientDir string
}

func NewServer(
	url string,
	port uint,
	clientDirFp string,
	bus *bus.Bus,
) (*Server, error) {
	return &Server{
		Config: ServerConfig{
			FullUrl: fmt.Sprintf("%s:%d", url, port),
			BaseUrl: url,
			Port:    port,
		},
		FileServer: http.FileServer(http.Dir(clientDirFp)),
		Api:        api.NewApi(bus),
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	router := http.ServeMux{}
	createComp := http.HandlerFunc(s.Api.CreateComponent)
	router.Handle(
		"POST /api/components",
		logRequest(createComp),
	)

	// objectHandler := http.HandlerFunc(func(
	// 	w http.ResponseWriter,
	// 	r *http.Request,
	// ) {
	// 	data, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		fmt.Printf("io.ReadAll: %+v\n", err)
	// 	}
	// 	defer r.Body.Close()
	//
	// 	switch r.Method {
	// 	case "POST":
	// 		req := components.Object{}
	// 		if err := json.Unmarshal(data, &req); err != nil {
	// 			fmt.Printf("error unmarshalling: %+v\n", err)
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			return
	// 		}
	//
	// 		w.WriteHeader(http.StatusCreated)
	// 		return
	//
	// 	case "GET":
	// 		// objs, err := s.Store.GetObjects()
	// 		// if err != nil {
	// 		// 	fmt.Printf("Store.GetObject: %+v\n", err)
	// 		// 	w.WriteHeader(http.StatusInternalServerError)
	// 		// 	return
	// 		// }
	//
	// 		// fmt.Printf("%+v\n", objs)
	//
	// 		// marshalled, err := json.Marshal(objs)
	// 		// if err != nil {
	// 		// 	fmt.Printf("json.Marshal: %+v\n", err)
	// 		// 	w.WriteHeader(http.StatusInternalServerError)
	// 		// 	return
	// 		// }
	// 		//
	// 		// w.Write(marshalled)
	// 		return
	//
	// 	case "OPTIONS":
	// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Refresh-Token")
	// 		w.Header().Set("Access-Control-Expose-Headers", "Authorization, X-Refresh-Token")
	// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
	// 		w.WriteHeader(http.StatusNoContent)
	//
	// 	default:
	// 		fmt.Printf("unsupported method '%s'\n", r.Method)
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// })
	// router.Handle("GET /api/object", handleCors(logRequest(objectHandler)))

	// mainHandler := http.HandlerFunc(func(
	// 	w http.ResponseWriter,
	// 	r *http.Request,
	// ) {
	// 	s.FileServer.ServeHTTP(w, r)
	// })
	// router.Handle("/", handleCors(logRequest(mainHandler)))

	router.Handle("/api/health", handleCors(logRequest(healthCheck())))
	server := http.Server{
		Addr:    s.Config.FullUrl,
		Handler: &router,
	}

	go func() {
		fmt.Printf("\u2713 Server Open On %s\n\n", s.Config.FullUrl)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}

func healthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func handleCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Refresh-Token")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization, X-Refresh-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)
	})
}
