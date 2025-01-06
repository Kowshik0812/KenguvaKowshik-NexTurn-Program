package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Salary     int    `json:"salary"`
	Department string `json:"department"`
}

const (
	HR    = "HR"
	IT    = "IT"
	Admin = "Admin"
)

var employees = []Employee{
	{ID: 1, Name: "Amul", Age: 30, Salary: 30000, Department: HR},
	{ID: 2, Name: "Sujith", Age: 25, Salary: 40000, Department: IT},
	{ID: 3, Name: "Chandu", Age: 35, Salary: 50000, Department: Admin},
	{ID: 4, Name: "Ajith", Age: 28, Salary: 55000, Department: HR},
	{ID: 5, Name: "Kailash", Age: 22, Salary: 60000, Department: IT},
}

func addEmployee(id int, name string, age int, department string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	if department != HR && department != IT && department != Admin {
		return errors.New("invalid department")
	}

	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

func searchEmployeeByID(id int) (Employee, error) {
	for _, emp := range employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

func searchEmployeeByName(name string) (Employee, error) {
	for _, emp := range employees {
		if strings.EqualFold(emp.Name, name) {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

func listEmployeesByDepartment(department string) []Employee {
	var deptEmployees []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			deptEmployees = append(deptEmployees, emp)
		}
	}
	return deptEmployees
}

func countEmployeesByDepartment(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func addEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil || age <= 0 {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	department := r.FormValue("department")

	err = addEmployee(id, name, age, department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Employee added successfully"))
}

func searchEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	name := r.URL.Query().Get("name")

	var emp Employee
	var err error

	if id != 0 {
		emp, err = searchEmployeeByID(id)
	} else if name != "" {
		emp, err = searchEmployeeByName(name)
	} else {
		http.Error(w, "Please provide an ID or Name to search", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func listEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	department := r.URL.Query().Get("department")
	if department == "" {
		http.Error(w, "Department is required", http.StatusBadRequest)
		return
	}

	deptEmployees := listEmployeesByDepartment(department)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deptEmployees)
}

func countEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	department := r.URL.Query().Get("department")
	if department == "" {
		http.Error(w, "Department is required", http.StatusBadRequest)
		return
	}

	count := countEmployeesByDepartment(department)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Count: %d", count)))
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

func main() {

	http.HandleFunc("/add", addEmployeeHandler)
	http.HandleFunc("/search", searchEmployeeHandler)
	http.HandleFunc("/list", listEmployeesHandler)
	http.HandleFunc("/count", countEmployeesHandler)
	http.HandleFunc("/employee", GetAllEmployees)

	fmt.Println("Server running on 8090")
	http.ListenAndServe(":8090", nil)
}
