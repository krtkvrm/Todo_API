package main

import (
	"strconv"
)

// Tasks is struct defined to store Task Information
type Tasks struct {
	TaskList    string
	TaskID      int
	TaskName    string
	TaskDetails string
	TaskStatus  string
}

// ListDataType is for Function Reciever
type ListDataType map[string][]Tasks

// IndexOfTask is to find location of Task using TaskID in ListDataType
func (list ListDataType) IndexOfTask(id, name string) int {
	index := -1
	for i, e := range list[name] {
		if strconv.Itoa(e.TaskID) == id {
			index = i
			break
		}
	}
	return index
}

// IfExist will check if List Exist
func (list ListDataType) IfExist(name string) bool {
	// Check if List Name Already Exist
	if _, ok := list[name]; ok {
		return true
	}
	return false
}

// AddTask to existing List
func (list ListDataType) AddTask(newTask Tasks) {
	newTask.TaskID = len(List[newTask.TaskList])
	newTask.TaskStatus = "To Comptele"
	list[newTask.TaskList] = append(list[newTask.TaskList], newTask)
}

// DeleteTaskAt is to delete task from a List
func (list ListDataType) DeleteTaskAt(listname string, index int) {
	listTask := list[listname]
	listTask = append(listTask[:index], listTask[index+1:]...)
	list[listname] = listTask
}

// UpdateTaskAt will update a particular task's status from "To Complete"
func (list ListDataType) UpdateTaskAt(listname string, newStatus string, index int) {
	list[listname][index].TaskStatus = newStatus
}