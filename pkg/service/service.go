package service

import (
	medapp "github.com/mnogohoddovochka/med-app"
	"github.com/mnogohoddovochka/med-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorisation interface {
	CreateDoctor(doctor medapp.Doctor) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type DoctorList interface {
	GetAll() ([]medapp.Doctor, error)
	GetById(id int) (medapp.Doctor, error)
}

type PatientList interface {
	CreatePatient(input medapp.Patient) (int, error)
	GetAll() ([]medapp.Patient, error)
	GetById(id int) (medapp.Patient, error)
	UpdatePatient(id int, input medapp.UpdPatient) error
	DeletePatient(id int) error
}

type VisitList interface {
	CreateVisit(input medapp.Visit) (int, error)
	GetAll() ([]medapp.VisitOutput, error)
	GetById(id int) (medapp.VisitOutput, error)
	UpdateVisit(id int, input medapp.UpdVisit) error
	DeleteVisit(id int) error
}

type Service struct {
	Authorisation
	DoctorList
	PatientList
	VisitList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repo.Authorisation),
		DoctorList:    NewDoctorsListService(repo.DoctorList),
		PatientList:   NewPatientsListService(repo.PatientList),
		VisitList:     NewVisitsListPostgres(repo.VisitList),
	}
}
