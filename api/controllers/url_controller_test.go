package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/helpers"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"errors"
)

var URL_TEST = models.URL{
	Address: "http://www.example.com.br",
	Alias: "wilson",
}

var URL_CONTROLLER_TEST URLController

func init() {
	gin.SetMode(gin.ReleaseMode)

	items := models.URLS{
		Urls: []models.URL{
			models.URL{UUID: "1", Alias: "wilson1", Address: "www.example1.com.br"},
			models.URL{UUID: "2", Alias: "wilson2", Address: "www.example2.com.br"},
			models.URL{UUID: "3", Alias: "wilson3", Address: "www.example3.com.br"},
		},
	}

	URL_CONTROLLER_TEST = URLController{&URLRepositoryMock{items}}
}

type URLRepositoryMock struct {
	items models.URLS
}

func (repository URLRepositoryMock) GetAllRanking(q url.Values) models.URLS {
	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)
	repository.items.Meta.Pagination = pagination
	return repository.items
}

func (repository URLRepositoryMock) GetByAlias(alias string) models.URL {
	for i, item := range repository.items.Urls {
		if item.Alias == alias {
			return repository.items.Urls[i]
		}
	}

	return models.URL{}
}

func (repository URLRepositoryMock) Create(url *models.URL) error {
	for _, item := range repository.items.Urls {
		if item.Alias == url.Alias {
			return errors.New("existe")
		}
	}
	
	return nil
}

func (repository URLRepositoryMock) CountRows() int {
	return len(repository.items.Urls)
}

func (repository URLRepositoryMock) UpdateView(url *models.URL) error {
	return nil
}

func TestCadastroURLValida(t *testing.T) {
	url, _ := json.Marshal(URL_TEST)
	reader := strings.NewReader(string(url))
	req, _ := http.NewRequest("POST", "/shorten", reader)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		URL_CONTROLLER_TEST.Create(c)
	})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Esperado %v, recebido %d", http.StatusCreated, w.Code)
	}
}

func TestCadastroURLAliasJaExiste(t *testing.T) {
	url := URL_TEST
	url.Alias = "wilson1"

	urlJson, _ := json.Marshal(url)
	reader := strings.NewReader(string(urlJson))
	req, _ := http.NewRequest("POST", "/shorten", reader)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		URL_CONTROLLER_TEST.Create(c)
	})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusConflict {
		t.Errorf("Esperado %v, recebido %d", http.StatusCreated, w.Code)
	}
}

func TestCadastroURLSemAddress(t *testing.T) {
	url := URL_TEST
	url.Address = ""

	urlJson, _ := json.Marshal(url)
	reader := strings.NewReader(string(urlJson))
	req, _ := http.NewRequest("POST", "/shorten", reader)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		URL_CONTROLLER_TEST.Create(c)
	})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Esperado %v, recebido %d", http.StatusCreated, w.Code)
	}
}

func TestBuscaURLPorAlias(t *testing.T) {
	req, _ := http.NewRequest("GET", "/details/wilson1", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/details/:alias", URL_CONTROLLER_TEST.GetByAlias)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado %v, recebido %d", http.StatusOK, w.Code)
	}
}

func TestBuscaURLPorAliasQueNaoExiste(t *testing.T) {
	req, _ := http.NewRequest("GET", "/details/wilson4", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/details/:alias", URL_CONTROLLER_TEST.GetByAlias)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Esperado %v, recebido %d", http.StatusOK, w.Code)
	}
}

func TestBuscaTodasURLSRanking(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ranking", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/ranking", URL_CONTROLLER_TEST.GetAllRanking)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado %v, recebido %d", http.StatusOK, w.Code)
	}
}