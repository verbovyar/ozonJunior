package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pkgs/internal/repositories/interfaces"
	"pkgs/internal/validation"
)

const (
	queryParamKey = "key"
)

type Server struct {
	Mux        *http.ServeMux
	Repository interfaces.Repository
}

func New(repository interfaces.Repository) *Server {
	mux := http.NewServeMux()

	server := Server{Mux: mux, Repository: repository}

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			server.Read(writer, request)
		case http.MethodPost:
			server.Create(writer, request)
		case http.MethodPut:
			server.Update(writer, request)
		case http.MethodDelete:
			server.Delete(writer, request)
		default:
			fmt.Printf("unsupported method: [%s]", request.Method)
		}
	})

	return &server
}

func (s *Server) Create(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}

	type data struct {
		Key   string
		Value string
	}
	var unmarshalledData data

	if err = json.Unmarshal(body, &unmarshalledData); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if validation.CreateValidate(unmarshalledData.Key, unmarshalledData.Value, writer) == false {
		return
	}

	s.Repository.Create(unmarshalledData.Key, unmarshalledData.Value)
}

func (s *Server) Read(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get(queryParamKey)

	if validation.ReadValidate(key, writer) == false {
		return
	}

	value := s.Repository.Read(key)

	_, err := writer.Write([]byte(value))
	if err != nil {
		log.Printf("error while writing body, err: [%s]", err.Error())
		return
	}
}

func (s *Server) Update(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}

	type data struct {
		Key   string
		Value string
	}
	var unmarshalledData data

	if err = json.Unmarshal(body, &unmarshalledData); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if validation.UpdateValidation(unmarshalledData.Key, unmarshalledData.Value, writer) == false {
		return
	}

	s.Repository.Update(unmarshalledData.Key, unmarshalledData.Value)
}

func (s *Server) Delete(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get(queryParamKey)

	if validation.DeleteValidate(key, writer) == false {
		return
	}

	s.Repository.Delete(key)
}
