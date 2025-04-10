package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/task"
)

func AddTaskHandler(taskService *task.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var request task.AddOneRequest
		if err := c.ReadJson(&request); err != nil {
			log.Errorf("failed to read json: %v", err)
			return err
		}

		response, err := taskService.AddOne(c.Ctx(), log, request)
		if err != nil {
			log.Errorf("failed to add one task: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type getTasksByInspectorIdVars struct {
	InspectorId int `path:"inspector_id"`
}

func GetTasksByInspectorId(taskService *task.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars getTasksByInspectorIdVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		response, err := taskService.GetByInspectorId(c.Ctx(), log, vars.InspectorId)
		if err != nil {
			log.Errorf("failed to get tasks by inspector id: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type updateTaskStatusVars struct {
	TaskId string `path:"task_id"`
}

func UpdateTaskStatusHandler(taskService *task.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars updateTaskStatusVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		var request task.UpdateStatusRequest
		if err := c.ReadJson(&request); err != nil {
			log.Errorf("failed to read json: %v", err)
			return err
		}

		err := taskService.UpdateStatus(c.Ctx(), log, vars.TaskId, request)
		if err != nil {
			log.Errorf("failed to update task status: %v", err)
			return err
		}

		c.Write(http.StatusOK)

		return nil
	}
}
