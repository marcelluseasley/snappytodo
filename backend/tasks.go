package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/matoous/go-nanoid/v2"
)

const (
	JSONFilePath = "./tasks.json"
)

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type CreateTaskBody struct {
	Name string `json:"name"`
}

type UpdateTaskBody struct {
	Name *string `json:"name"`
	Done *bool   `json:"done"`
}

func tasks(c *fiber.Ctx) error {
	tasks := []Task{}
	jsonFile, err := os.ReadFile(JSONFilePath)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}
	err = json.Unmarshal(jsonFile, &tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}
	return c.JSON(tasks)
}

func createTask(c *fiber.Ctx) error {
	body := CreateTaskBody{}

	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	jsonFile, err := os.ReadFile(JSONFilePath)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	tasks := []Task{}
	err = json.Unmarshal(jsonFile, &tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	id, err := gonanoid.New()
	if err != nil {
		id = strconv.Itoa(len(tasks))
	}
	newTask := Task{
		Name: body.Name,
		Done: false,
		Id:   id,
	}

	tasks = append(tasks, newTask)
	j, err := json.Marshal(tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	err = os.WriteFile(JSONFilePath, j, 0755)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	c.SendStatus(fiber.StatusCreated)

	return c.JSON(newTask)
}

func updateTask(c *fiber.Ctx) error {
	body := UpdateTaskBody{}

	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	jsonFile, err := os.ReadFile(JSONFilePath)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	tasks := []Task{}
	err = json.Unmarshal(jsonFile, &tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	taskId := c.Params("taskId")
	for i, task := range tasks {
		if taskId == task.Id {
			if body.Name != nil {
				task.Name = *body.Name
			}
			if body.Done != nil {
				task.Done = *body.Done
			}
			tasks[i] = task
		}
	}

	j, err := json.Marshal(tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	err = os.WriteFile(JSONFilePath, j, 0755)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteTask(c *fiber.Ctx) error {

	jsonFile, err := os.ReadFile(JSONFilePath)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	tasks := []Task{}
	err = json.Unmarshal(jsonFile, &tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	taskId := c.Params("taskId")

	prevLen := len(tasks)

	for i, task := range tasks {
		if taskId == task.Id {
			tasks = append(tasks[:i],tasks[i+1:]...)
		}
	}

	currLen := len(tasks)

	if currLen == prevLen {
		return c.SendStatus(fiber.StatusNotFound)
	}

	j, err := json.Marshal(tasks)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	err = os.WriteFile(JSONFilePath, j, 0755)
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"code":    code,
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}