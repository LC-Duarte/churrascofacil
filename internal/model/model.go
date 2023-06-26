package model

import (
	"encoding/json"
	"fmt"

	"github.com/LC-Duarte/churrascofacil/internal/loader"
)

type Cortes struct {
	Cortes []Corte `json:"cortes"`
}

type Corte struct {
	ID   int      `json:"id"`
	Name string   `json:"Nome"`
	Type string   `json:"tipo"`
	PP   float32  `json:"pp"`
	Cat  []string `json:"cat"`
}

func (cortes *Cortes) Load(path string) error {
	byteValue, err := loader.Load(path)
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal([]byte(byteValue), cortes)
	return err

}

func (corte *Corte) IsMain() bool {
	for _, s := range corte.Cat {
		if "M" == s {
			return true
		}
	}
	return false
}
func (corte *Corte) IsAppetizer() bool {
	for _, s := range corte.Cat {
		if "A" == s {
			return true
		}
	}
	return false
}
func (corte *Corte) IsSec() bool {
	for _, s := range corte.Cat {
		if "S" == s {
			return true
		}
	}
	return false
}
