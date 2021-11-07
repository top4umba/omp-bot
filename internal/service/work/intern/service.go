package intern

import (
	"fmt"
)

type Service interface {
	Describe(internID uint64) (*Intern, error)
	List(cursor uint64, limit uint64) ([]Intern, error)
	Create(Intern) (uint64, error)
	Update(internID uint64, intern Intern) error
	Remove(internID uint64) (bool, error)
}

type InternService struct{}

func NewService() *InternService {
	return &InternService{}
}

func (s *InternService) List() []Intern {
	return allInterns
}

func (s *InternService) Describe(internID uint64) (*Intern, error) {
	idx, intern := getIntern(internID)
	if idx >= 0 {
		return &intern, nil
	}
	return nil, fmt.Errorf("cannot find intern with id %d", internID)
}

func (s *InternService) Create(intern Intern) (uint64, error) {
	intern.UniqueKey = getNextInternId()
	allInterns = append(allInterns, intern)
	return intern.UniqueKey, nil
}

func (s *InternService) Update(internID uint64, intern Intern) error {
	idx, _ := getIntern(internID)
	if idx < 0 {
		return fmt.Errorf("cannot find intern with id %d", internID)
	}
	updatedIntern := allInterns[idx]
	updatedIntern.Name = intern.Name
	updatedIntern.InternshipID = intern.InternshipID
	allInterns[idx] = updatedIntern
	return nil
}

func (s *InternService) Remove(internID uint64) (bool, error) {
	idx, _ := getIntern(internID)
	if idx >= 0 {
		allInterns[idx] = allInterns[len(allInterns)-1]
		allInterns = allInterns[:len(allInterns)-1]
		return true, nil
	}
	return false, nil
}

func getIntern(internID uint64) (int, Intern) {
	var res Intern
	for idx, intern := range allInterns {
		if intern.UniqueKey == internID {
			res = intern
			return idx, res
		}
	}
	return -1, res
}
