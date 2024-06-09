package customer

import "testing"

func TestPersonAge(t *testing.T) {
	person := Person{Name: "John", Age: 30}
	if person.Age != 30 {
		t.Errorf("person.Age = %d; want 30", person.Age)
	}
}

func TestPersonName(t *testing.T) {
	person := Person{Name: "John", Age: 30}
	if person.Name != "JohnX" {
		t.Errorf("person.Name = %s; want John", person.Name)
	}
}
