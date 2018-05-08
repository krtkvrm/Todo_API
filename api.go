package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// List is to store Todo List in Map
var List ListDataType = make(map[string][]Tasks)

// CreateList to Create Todo List
func CreateList(w http.ResponseWriter, r *http.Request) {
	list, ok := r.URL.Query()["list"]
	if !ok || len(list) < 1 {
		fmt.Fprintln(w, "URL Queries missing")
		return
	}
	// Check if List Name Already Exist
	if List.IfExist(list[0]) {
		fmt.Fprintf(w, "List Already Exists")
	} else {
		// Creating list
		List[list[0]] = nil
		fmt.Fprintln(w, "List Made")
	}
	return
}

// Add Task to a List
func Add(w http.ResponseWriter, r *http.Request) {
	var newtask Tasks
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newtask)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	if List.IfExist(newtask.TaskList) {
		List.AddTask(newtask)
		fmt.Fprintln(w, "Task Added")
	} else {
		fmt.Fprintln(w, "List Doesn't Exist")
	}
	return
}

// GetList all the tasks of a List
func GetList(w http.ResponseWriter, r *http.Request) {
	inp, ok := r.URL.Query()["list"]
	if !ok || len(inp) < 1 {
		fmt.Fprintln(w, "URL Query 'list' missing")
		return
	}
	encoder := json.NewEncoder(w)
	listName := inp[0]
	if List.IfExist(listName) {
		encoder.Encode(List[listName])
	} else {
		fmt.Fprintln(w, "List Doesn't Exist")
	}
	return
}

// UpdateTask is to delete a particular task in Todo List
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	inp, ok := r.URL.Query()["list"]
	status, ok := r.URL.Query()["status"]
	if !ok || len(inp) < 1 || len(id) < 0 || len(status) < 0 {
		fmt.Fprintln(w, "URL Query 'list' missing")
		return
	}
	listName := inp[0]
	if List.IfExist(listName) {
		index := List.IndexOfTask(id[0], listName)
		if index == -1 {
			// if element is not found
			fmt.Fprintln(w, "Task not Found")
		} else {
			List.UpdateTaskAt(listName, status[0], index)
			fmt.Fprintln(w, "Task Updated")
		}
	} else {
		fmt.Fprintln(w, "List Doesn't Exist")
	}
	return
}

// DeleteTask is to Update a Particular Task Status
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	inp, ok := r.URL.Query()["list"]
	if !ok || len(inp) < 1 || len(id) < 0 {
		fmt.Fprintln(w, "URL Query 'list' missing")
		return
	}
	listName := inp[0]
	if List.IfExist(listName) {
		index := List.IndexOfTask(id[0], listName)
		if index == -1 {
			// if element is not found
			fmt.Fprintln(w, "Task not Found")
		} else {
			List.DeleteTaskAt(listName, index)
			fmt.Fprintln(w, "Task Deleted")
		}
	} else {
		fmt.Fprintln(w, "List Doesn't Exist")
	}
	return
}

// DeleteList to delete a particular Todo List
func DeleteList(w http.ResponseWriter, r *http.Request) {
	newlist, ok := r.URL.Query()["list"]
	if !ok || len(newlist) < 1 {
		fmt.Fprintln(w, "URL Query 'list' missing")
		return
	}
	// Check if List Name Already Exist
	if List.IfExist(newlist[0]) {
		delete(List, newlist[0])
		fmt.Fprintln(w, "List Deleted")
	} else {
		fmt.Fprintln(w, "List Doesn't Exist")
	}
	return
}
