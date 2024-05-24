package customer

import "testing"

func TestPerson(t *testing.T) {
	person := Person{Name: "John", Age: 30}
	if person.Name != "John" {
		t.Errorf("person.Name = %s; want John", person.Name)
	}
	if person.Age != 30 {
		t.Errorf("person.Age = %d; want 30", person.Age)
	}
}
