package main

import ("fmt"
		"encoding/json"
		"net/http"
		"github.com/gorilla/mux"
		"io/ioutil"
)

type Employee struct{
	ID string `json:"id"`
	Name string `json::"name"`
	Salary float64 `json:"salary"`
}

var emps []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	
	json.NewEncoder(w).Encode(emps)
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id := vars["id"]

	
	for _, emp := range emps {
		if emp.ID == id {
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(emp)
			return
		}
	}

	
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Employee not found"))
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	
	var newEmployee Employee
	err = json.Unmarshal(body, &newEmployee)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	
	emps = append(emps, newEmployee)

	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emps)
}


func main(){
	emps = []Employee{
		{ID: "E001", Name: "John Doe", Salary: 50000.0},
		{ID: "E002", Name: "Jane Smith", Salary: 60000.0},
	}

	router := mux.NewRouter()

	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", getEmployeeByID).Methods("GET")
	router.HandleFunc("/employees", addEmployee).Methods("POST")


	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", router)
}

