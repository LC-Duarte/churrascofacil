package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"os"

	docs "github.com/LC-Duarte/churrascofacil/docs" // Importa os arquivos de documentação gerados pelo Swagger

	"github.com/LC-Duarte/churrascofacil/internal/model"
	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
)

var cortes model.Cortes

func loadConfig(path string) map[string]interface{} {
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result

}

func CreateApp() *fiber.App {
	docs.SwaggerInfo.BasePath = "/"
	app := fiber.New()
	//app.Get("/", home)
	// Rota para a documentação Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",

		/*// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",*/
	}))

	// Rota para listar os tipos de carne
	app.Get("/carnes", listCarnes)

	// Rota para listar os cortes de churrasco
	app.Get("/cortes", listCortes)
	app.Get("/calcular", calcularQuantidade)
	return app
}

// @title Churrasco Fácil API
// @description API para listar e calcular quantidades de carne para um churrasco fácil.
func main() {

	cortes.Load("../../data/load.json")
	fmt.Println(cortes)

	app := CreateApp()

	// Rota para calcular a quantidade de carne
	//app.Get("/calcular", calcularQuantidade)

	err := app.Listen(":8080")

	if err != nil {
		fmt.Println("The port is already in use!")
		os.Exit(1)
	}

}

// @Summary Listar tipos de carne
// @Description Retorna uma lista de tipos de carne disponíveis
// @Produce json
// @Success 200 {object} []string
// @Router /carnes [get]
func listCarnes(c *fiber.Ctx) error {
	carnes := []string{"Frango", "Boi", "Porco", "Carneiro"}
	return c.JSON(carnes)
}

// @Summary Listar cortes de churrasco
// @Description Retorna uma lista de cortes de churrasco para um determinado tipo de carne
// @Param tipo query string true "Tipo de carne (Frango, Boi, Porco, Carneiro)"
// @Produce json
// @Success 200 {object} []string
// @Router /cortes [get]
func listCortes(c *fiber.Ctx) error {
	tipo := c.Query("tipo")
	if tipo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tipo de carne inválido"})
	}

	cortes := make([]string, 0)

	switch tipo {
	case "Frango":
		cortes = []string{"Peito", "Asa", "Coxa", "Sobrecoxa"}
	case "Boi":
		cortes = []string{"Picanha", "Maminha", "Alcatra", "Cupim", "Assado de Tira", "Fraldinha", "Bananinha", "Entranha", "Bife Ancho", "Entrecote", "Costela", "Prime Rib", "Short Rib"}
	case "Porco":
		cortes = []string{"Costela", "Linguiça", "Pernil", "Lombo"}
	case "Carneiro":
		cortes = []string{"Paleta", "Carré", "Perna"}
	case "Todos":
		cortes = []string{"Peito", "Asa", "Coxa", "Sobrecoxa", "Picanha", "Maminha", "Alcatra", "Cupim",
			"Assado de Tira", "Fraldinha", "Bananinha", "Entranha", "Bife Ancho", "Entrecote", "Costela",
			"Prime Rib", "Short Rib", "Costela", "Linguiça", "Pernil", "Lombo",
			"Paleta", "Carré", "Perna"}

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tipo de carne inválido"})
	}

	return c.JSON(cortes)
}

// @Summary Calcular quantidade de carne
// @Description Calcula a quantidade de carne necessária para um churrasco
// @Param pessoas query int true "Quantidade de pessoas"
// @Produce json
// @Success 200 {object} map[string]float64
// @Router /calcular [get]
func calcularQuantidade(c *fiber.Ctx) error {

	quantidades := make(map[string]interface{})
	switch pessoas := float32(c.QueryInt("pessoas", -1)); {
	case pessoas == 1:
		quantidades["Principal"] = calcMain(pessoas)
		break
	case pessoas > 1 && pessoas < 5:
		quantidades["Principal"] = calcMain(pessoas)
		quantidades["Acompanhamentos"] = calcAppetizer(pessoas)
		break
	case pessoas >= 5 && pessoas < 10:
		quantidades["Principal"] = calcMain(pessoas)
		quantidades["Secundaria"] = calcSec(pessoas)
		quantidades["Aperitivo"] = calcAppetizer(pessoas)
		break
	case pessoas >= 10:
		quantidades["Principal"] = calcMain(pessoas / 2)
		quantidades["Principal_2"] = quantidades["Principal"]
		quantidades["Secundaria"] = calcSec(pessoas)
		quantidades["Aperitivo"] = calcAppetizer(pessoas)
		break
	default:
		break
	}
	sugestion := make(map[string]string)
	for k, value := range quantidades {
		mv := value.(map[string]string)
		s_name, s_value := pick(mv)
		sugestion[k] = fmt.Sprintf("%s  de %s", s_value, s_name)
	}
	quantidades["sugestão"] = sugestion

	return c.JSON(quantidades)
}

func calcMain(pessoas float32) map[string]string {
	mainCarnes := make(map[string]string)
	for _, corte := range cortes.Cortes {
		if corte.IsMain() {
			mainCarnes[corte.Name] = fmt.Sprintf("%.2fKg", pessoas*corte.PP)
		}
	}
	return mainCarnes

}

func calcAppetizer(pessoas float32) map[string]string {
	appetizers := make(map[string]string)
	for _, corte := range cortes.Cortes {
		if corte.IsAppetizer() {
			appetizers[corte.Name] = fmt.Sprintf("%.2fKg", pessoas*corte.PP)
		}
	}
	return appetizers
}

func calcSec(pessoas float32) map[string]string {
	secs := make(map[string]string)
	for _, corte := range cortes.Cortes {
		if corte.IsSec() {
			secs[corte.Name] = fmt.Sprintf("%.2fKg", pessoas*corte.PP)
		}
	}
	return secs
}

func pick(m map[string]string) (string, string) {
	X := rand.Intn(len(m))
	i := 0
	for k, v := range m {
		if i == X {
			return k, v
		}
		i++
	}

	return "", ""
}
