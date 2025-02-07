package app

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/xorima/slogger"
)

type App struct {
	router chi.Router
	log    *slog.Logger
	todos  []string
	mu     sync.Mutex
}

func NewApp() *App {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	app := &App{
		router: router,
		log:    slogger.NewLogger(slogger.NewLoggerOpts("example-app", "todo-api")),
		todos:  []string{},
	}

	app.routes()
	return app
}

func (a *App) Run() error {
	a.log.Info("server started on 8080")
	return http.ListenAndServe(":8080", a.router)
}

func (a *App) routes() {
	a.router.Route("/api/v1/todo", func(r chi.Router) {
		r.Post("/", a.addTodo)
		r.Get("/", a.getTodos)
	})
}

func (a *App) addTodo(w http.ResponseWriter, r *http.Request) {
	var todo string
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	a.todos = append(a.todos, todo)
	a.mu.Unlock()

	a.log.Info("Added new todo", slog.String("todo", todo))
	w.WriteHeader(http.StatusCreated)
}

func (a *App) getTodos(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if err := json.NewEncoder(w).Encode(a.todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
