// Go doens't support inheritance, but rather composition and promotion

package main 

import(
	"fmt"
)


type Employee struct {
	Name 	string
	ID 		string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}


type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	
	// Do business logic
	return []Employee{}

}

func main() {
	m := Manager{
		Employee: Employee {
			Name: 	"Bob_Bobson", 
			ID: 	"12345",
		},
	
		Reports: []Employee{},
	}
	
	fmt.Println(m.ID)
	fmt.Println(m.Description())
}



