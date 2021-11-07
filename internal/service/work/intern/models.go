package intern

import (
	"fmt"
)

var idIncrementer uint64

var allInterns = []Intern{
	{Name: "Иван Охлобыстин", UniqueKey: getNextInternId(), InternshipID: 1},
	{Name: "Кристина Асмус", UniqueKey: getNextInternId(), InternshipID: 1},
	{Name: "Аглая Тарасова", UniqueKey: getNextInternId(), InternshipID: 2},
	{Name: "Светлана Пермякова", UniqueKey: getNextInternId(), InternshipID: 2},
	{Name: "Илья Глинников", UniqueKey: getNextInternId(), InternshipID: 2},
	{Name: "Дмитрий Шаракоис", UniqueKey: getNextInternId(), InternshipID: 2},
	{Name: "Александр Ильин", UniqueKey: getNextInternId(), InternshipID: 3},
	{Name: "Один Ланд Байрон", UniqueKey: getNextInternId(), InternshipID: 3},
	{Name: "Вадим Демчог", UniqueKey: getNextInternId(), InternshipID: 3},
	{Name: "Азамат Мусагалиев", UniqueKey: getNextInternId(), InternshipID: 3},
}

var allInternships = []Internship{
	{Name: "Хирургия", InternshipID: 1},
	{Name: "Кардиология", InternshipID: 1},
	{Name: "Отолорингология", InternshipID: 2},
}

type Intern struct {
	Name         string
	UniqueKey    uint64
	InternshipID uint64
}

type Internship struct {
	InternshipID uint64
	Name         string
}

func NewIntern(name string, internshipId uint64) *Intern {
	return &Intern{
		Name:         name,
		InternshipID: internshipId,
	}
}

func (i Intern) String() string {
	return fmt.Sprintf("Name: %s\nUnique key: %d\nInternship ID: %d", i.Name, i.UniqueKey, i.InternshipID)
}

func getNextInternId() uint64 {
	idIncrementer += 1
	return idIncrementer
}
