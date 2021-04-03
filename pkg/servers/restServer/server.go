package restServer

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/keepcalmist/grpcFibonacci/pkg/fibonacci"
	"github.com/keepcalmist/grpcFibonacci/pkg/servers"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type server struct {
	serv *http.Server
	fibo fibonacci.FiboService
}

func New(fibo fibonacci.FiboService) servers.Server {
	return &server{
		serv: &http.Server{},
		fibo: fibo,
	}
}

func (s *server) Run(conf *viper.Viper, quit chan bool) {
	s.serv.Addr = ":" + conf.GetString("REST_SERVER_ADDRESS")
	s.serv.Handler = s.initRouter()
	go func() {
		if err := s.serv.ListenAndServe(); err != nil {
			log.Println("Server error: ", err)
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.serv.Shutdown(ctx)
}

func (s *server) initRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/calculate", s.Get).Methods(http.MethodPost)
	return r
}
