package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/task"
)

func AddTaskHandler(taskService *task.Service) router.Handler {
	return func(c router.Context) error {
		var request task.AddOneRequest
		if err := c.ReadJson(&request); err != nil {
			return fmt.Errorf("failed to read json: %w", err)
		}

		response, err := taskService.AddOne(c.Ctx(), c.Log(), request)
		if err != nil {
			return fmt.Errorf("failed to add one task: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type getTasksByInspectorIdVars struct {
	InspectorId int `path:"inspector_id"`
}

func GetTasksByInspectorId(taskService *task.Service) router.Handler {
	return func(c router.Context) error {
		var vars getTasksByInspectorIdVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		response, err := taskService.GetByInspectorId(c.Ctx(), c.Log(), vars.InspectorId)
		if err != nil {
			return fmt.Errorf("failed to get tasks by inspector id: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type updateTaskStatusVars struct {
	TaskId string `path:"task_id"`
}

func UpdateTaskStatusHandler(taskService *task.Service) router.Handler {
	return func(c router.Context) error {
		var vars updateTaskStatusVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		var request task.UpdateStatusRequest
		if err := c.ReadJson(&request); err != nil {
			return fmt.Errorf("failed to read json: %w", err)
		}

		err := taskService.UpdateStatus(c.Ctx(), c.Log(), vars.TaskId, request)
		if err != nil {
			return fmt.Errorf("failed to update task status: %w", err)
		}

		c.Write(http.StatusOK)

		return nil
	}
}
