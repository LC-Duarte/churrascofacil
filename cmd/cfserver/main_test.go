package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/LC-Duarte/churrascofacil/cmd/cfserver" // Importe o pacote principal da aplicação
	"github.com/stretchr/testify/assert"
	// Importe o pacote principal da aplicação
)

func TestListCarnesEndpoint(t *testing.T) {
	app := main.CreateApp()

	req := httptest.NewRequest(http.MethodGet, "/carnes", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expected := []string{"Frango", "Boi", "Porco", "Carneiro"}
	actual := make([]string, 0)
	err = json.NewDecoder(resp.Body).Decode(&actual)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestListCortesEndpoint(t *testing.T) {
	app := main.CreateApp()

	req := httptest.NewRequest(http.MethodGet, "/cortes?tipo=Frango", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expected := []string{"Peito", "Asa", "Coxa", "Sobrecoxa"}
	actual := make([]string, 0)
	err = json.NewDecoder(resp.Body).Decode(&actual)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestMain(m *testing.M) {
	// Configuração antes de executar os testes, se necessário
	// ...

	// Executa os testes
	m.Run()
}
