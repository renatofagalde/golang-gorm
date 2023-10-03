package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/renatofagalde/gorm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
)

func main() {
	db, err := gorm.Open(sqlite.Open("pessoa.db"), &gorm.Config{})
	if err != nil {
		panic("erro ao acessar o banco de dados")
	}
	db.AutoMigrate(&model.Pessoa{})

	for i := 1; i < 10; i++ {
		db.Create(&model.Pessoa{
			Model:       gorm.Model{},
			Nome:        faker.Name(),
			Idade:       int8(rand.Intn(50-10) + 10),
			PhoneNumber: faker.Phonenumber(),
		})
	}

	var pessoa model.Pessoa
	db.Last(&pessoa, 9)
	fmt.Println("Pessoa ID 9 é: ", pessoa)

	db.Model(&pessoa).Update("Nome", "Senna")

	//var pessoaFilter model.Pessoa
	var pessoas = []model.Pessoa{}
	db.Find(&pessoas).Rows()

	for _, p := range pessoas {
		fmt.Println(
			fmt.Sprintf("%d nome -> %s, criado na data: %s", p.ID, p.Nome, p.CreatedAt))

	}

}
