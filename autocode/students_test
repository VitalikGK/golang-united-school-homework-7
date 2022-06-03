package coverage

import (
	"os"
	"fmt"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T){
	var p People
	p = append(p, Person{firstName: "firstName", lastName: "Test_lastName", birthDay: time.Now() })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: time.Now().Add(time.Minute * 2) })
	// fmt.Println("p= ",p)
	// fmt.Println("p[0]= ",p[0])
	// fmt.Println("len p= ",len(p))
	err := p.Len()
	if err != 2 {
		t.Error("Ожидается 2, возвращает ", err)
	}

}

func TestLess(t *testing.T){
	var p People
	p = append(p, Person{firstName: "firstNam", lastName: "Test_lastName", birthDay: time.Now().Add(time.Minute *2) })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: time.Now() })

	
	err := p.Less(0,1)

	if err != true{
		t.Error("Ожидается true, возвращает ", err)
	}
}

func TestSwap(t *testing.T){
	var p People
	p = append(p, Person{firstName: "firstNam", lastName: "Test_lastName", birthDay: time.Now().Add(time.Minute *2) })
	p = append(p, Person{firstName: "firstName1", lastName: "Test_lastName1", birthDay: time.Now() })

	p0 := p[0]

	p.Swap(0,1)

	fmt.Println("Элемент 0 ", p[0])

	if p0 == p[0]{

		t.Error("Не поменялись элементы ", p[0])
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestNewMatrix(t *testing.T){

	//var m0(Matrix)

	_, err := New("10 20 30 40 50 60 70 80 90")
	if err != nil {
		t.Error("Ошибка")
	}

//	fmt.Println("m0 ", m0)
}