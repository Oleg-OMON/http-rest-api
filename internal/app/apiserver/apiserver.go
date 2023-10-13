package apiserver

import (
	"fmt"
	"net/http"

	"github.com/Oleg-OMON/http-rest-api.git/internal/app/models"
	"github.com/Oleg-OMON/http-rest-api.git/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// структура ApiServera
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// инициализациия  сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// функция запуска сервера
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Сервер запушен")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handlehallo())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handlehallo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := s.store.DB.Query("Select * FROM players")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		players := []models.Player{}

		for rows.Next() {
			pl := models.Player{}
			err := rows.Scan(&pl.Player_id, &pl.First_name, &pl.Last_name, &pl.Nickname, &pl.Citizenship, &pl.Dob, &pl.Role)
			if err != nil {
				fmt.Println("хуня 7 а не 6")
				continue
			}
			players = append(players, pl)
		}
		for _, p := range players {
			fmt.Println(p)
		}

	}
}
