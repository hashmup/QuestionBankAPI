package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/config"
	"github.com/justinas/alice"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	h := initHandlers()
	http.ListenAndServe(":9000", h)
}

func initHandlers() http.Handler {
	// Routing
	r := chi.NewRouter()

	// middleware chain
	chain := alice.New(
	// loggingMiddleware,
	// middleware.AccessControl,
	// middleware.Authenticator,
	)

	// DB connection
	var dbConfig config.DBConfig
	err := envconfig.Process("qb", &dbConfig)
	if err != nil {
		panic(err)
	}

	dbConn, err := config.NewDBConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	// Redis Connection
	var redisConfig config.RedisConfig
	err = envconfig.Process("qb", &redisConfig)
	if err != nil {
		panic(err)
	}
	redisConn, err := config.NewRedisConnection(redisConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", dbConn)
	fmt.Printf("%#v\n", redisConn)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	h := chain.Then(r)

	return h
}
