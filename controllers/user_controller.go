package controllers

import (
	"encoding/json"
	"gestfin-apis/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetUserDataByToken(c *fiber.Ctx) error {
	if session, err := ValidateTokenSession(c); err == nil {
		var user = models.User{ID: uint(session.UserID)}
		var dataUser models.DataUser

		if err, dataUser = user.FindByID(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Usuário não encontrado!")
		}

		return c.Status(fiber.StatusOK).JSON(dataUser)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON("Você não tem autorização!")
	}
}

func UpdateUserByColumn(c *fiber.Ctx) error {
	if _, err := ValidateTokenSession(c); err == nil {
		body := c.Body()
		list := []string{"NAME", "EMAIL", "DATE_BIRTH", "TEL", "CELL", "GENDER"}
		var u models.User
		var user struct {
			NewValue string `json:"new_value"`
			Column   string `json:"column"`
		}

		if err := json.Unmarshal(body, &user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Estrutura de dados incorreta!")
		}

		if indexOf(list, user.Column) == -1 {
			return c.Status(fiber.StatusBadRequest).JSON("Informe o campo que deseja editar.")
		}

		if err = u.UpdateByColumn(user.Column, user.NewValue); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Ocorreu um erro ao tentar atualizar o campo.")
		}

		return c.Status(fiber.StatusOK).JSON("Dados atualizados com sucesso!")
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON("Você não tem acesso a essa rota ")
	}
}

func indexOf(slice []string, value string) int {
	for i, item := range slice {
		if item == strings.ToUpper(value) {
			return i
		}
	}
	return -1
}
