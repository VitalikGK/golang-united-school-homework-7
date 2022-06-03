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

	

	m, err := New("10 20 30 40 50 60 70 80 90")
	if err != nil {
		t.Error("Ошибка")
	}

	var cols0 [][]int
	for i :=0; i < 9; i++{
		cols0 = append(cols0, []int{(i+1)*10})
	} 

	cols := m.Cols()



	b := compareSlices(cols, cols0)
	if b != true {
		t.Error("Ошибка, слайс не верный")
	}

	s := []int{10,20,30,40,50,60,70,80,90}

	var rows0 [][]int
		rows0 = append(rows0, s)

	rows := m.Rows()

	b0 := compareSlicesR(rows, rows0)
	if b0 != true {
		t.Error("Ошибка, слайс не верный")
	}

	 fmt.Println("Rows ", rows)


}

func compareSlices(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
	return false
	}
	
	for i := 0; i < len(s1); i++ {
	if s1[i][0] != s2[i][0] {
	return false
	}
	}
	
	return true
   }

   func compareSlicesR(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
	return false
	}
	
	for i := 0; i < len(s1); i++ {
	if s1[0][i] != s2[0][i] {
	return false
	}
	}
	
	return true
   }