package repositories

import (
	"github.com/wilsontamarozzi/bemobi-hire-me/api/models"
	"net/url"
	"testing"
)

var URL_TEST = models.URL{
	Address: "http://example.com.br",
	Alias:   "wilson",
}

var URL_REPOSITORY_TEST = NewURLRepository()

func TestCadastroURLComAlias(t *testing.T) {
	url := URL_TEST
	err := URL_REPOSITORY_TEST.Create(&url)

	if err != nil {
		t.Errorf("Não era esperado o erro: %s", err.Error())
	}

	if url.UUID == "" {
		t.Error("Era esperado o ID da URL")
	}

	if url.Alias == "" {
		t.Error("Era esperado o ALIAS da URL")
	}
}

func TestCadastroURLSemAlias(t *testing.T) {
	url := URL_TEST
	url.Alias = ""
	err := URL_REPOSITORY_TEST.Create(&url)

	if err != nil {
		t.Errorf("Não era esperado o erro: %s", err.Error())
	}

	if url.UUID == "" {
		t.Error("Era esperado o ID da URL")
	}

	if url.Alias == "" {
		t.Error("Era esperado o ALIAS da URL")
	}
}

func TestBuscaURLPorAlias(t *testing.T) {
	newUrl := URL_TEST
	newUrl.Alias = "teste-busca-por-alias"

	errCreate := URL_REPOSITORY_TEST.Create(&newUrl)
	if errCreate != nil {
		t.Errorf("Não era esperado o erro: %s", errCreate.Error())
	}

	url := URL_REPOSITORY_TEST.GetByAlias(newUrl.Alias)
	if url.UUID == "" {
		t.Error("Era esperado o ID da URL")
	}

	if !url.IsEmpty() {
		t.Errorf("Esperado registro da URL")
		t.Errorf("Recebido vazio")
	}
}

func TestListagemDoRanking(t *testing.T) {
	url1 := URL_TEST
	url2.Alias = "alias1"
	url2 := URL_TEST
	url2.Alias = "alias2"

	err1 := URL_REPOSITORY_TEST.Create(&url1)
	err2 := URL_REPOSITORY_TEST.Create(&url2)

	if err1 != nil || err2 != nil {
		t.Errorf("Não era esperado o erro1: %s", err1)
		t.Errorf("Não era esperado o erro2: %s", err2)
	}

	params := url.Values{}
	params.Add("page", "1")
	params.Add("per_page", "50")
	urls := URL_REPOSITORY_TEST.GetAllRanking(params)
	if len(urls) <= 0 {
		t.Error("Esperado maior que 0")
		t.Error("Recebido %d", len(url2))
	}
}
