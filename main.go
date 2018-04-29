package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/hashmup/QuestionBankAPI/src/application"
	"github.com/hashmup/QuestionBankAPI/src/config"
	"github.com/hashmup/QuestionBankAPI/src/infrastructure/persistence"
	"github.com/hashmup/QuestionBankAPI/src/interfaces/class"
	"github.com/hashmup/QuestionBankAPI/src/interfaces/folder"
	"github.com/hashmup/QuestionBankAPI/src/interfaces/question"
	"github.com/hashmup/QuestionBankAPI/src/interfaces/school"
	"github.com/hashmup/QuestionBankAPI/src/interfaces/session"
	"github.com/hashmup/QuestionBankAPI/src/interfaces/user"
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

	classRepo := persistence.NewClassRepository(dbConn)
	schoolRepo := persistence.NewSchoolRepository(dbConn)
	schoolService := application.NewSchoolService(schoolRepo, classRepo)
	schoolDependency := &school.Dependency{
		SchoolService: schoolService,
	}
	r = school.MakeSchoolHandler(schoolDependency, r)

	userRepo := persistence.NewUserRepository(dbConn)
	userService := application.NewUserService(userRepo)
	userDependency := &user.Dependency{
		UserService: userService,
	}

	r = user.MakeUserHandler(userDependency, r)

	sessionRepo := persistence.NewSessionRepository(dbConn, redisConn)
	sessionService := application.NewSessionService(sessionRepo, userRepo)
	sessionDependency := &session.Dependency{
		SessionService: sessionService,
	}

	r = session.MakeSessionHandler(sessionDependency, r)

	classService := application.NewClassService(classRepo)
	classDependency := &class.Dependency{
		ClassService:   classService,
		SessionService: sessionService,
	}

	r = class.MakeClassHandler(classDependency, r)

	folderRepo := persistence.NewFolderRepository(dbConn)
	folderService := application.NewFolderService(folderRepo)
	folderDependency := &folder.Dependency{
		FolderService:  folderService,
		SessionService: sessionService,
	}

	r = folder.MakeFolderHandler(folderDependency, r)

	questionRepo := persistence.NewQuestionRepository(dbConn)
	questionService := application.NewQuestionService(questionRepo)
	questionDependency := &question.Dependency{
		QuestionService: questionService,
		SessionService:  sessionService,
	}
	r = question.MakeQuestionHandler(questionDependency, r)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	h := chain.Then(r)

	return h
}
