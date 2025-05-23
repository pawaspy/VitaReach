// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"context"
)

type Querier interface {
	AddAppointmentNotes(ctx context.Context, arg AddAppointmentNotesParams) (Appointment, error)
	CheckDoctorEmailExists(ctx context.Context, email string) (bool, error)
	CheckDoctorUsernameExists(ctx context.Context, username string) (bool, error)
	CheckPatientEmailExists(ctx context.Context, email string) (bool, error)
	CheckPatientUsernameExists(ctx context.Context, username string) (bool, error)
	CreateAppointment(ctx context.Context, arg CreateAppointmentParams) (Appointment, error)
	CreateDoctor(ctx context.Context, arg CreateDoctorParams) (Doctor, error)
	CreatePatient(ctx context.Context, arg CreatePatientParams) (Patient, error)
	CreatePrescription(ctx context.Context, arg CreatePrescriptionParams) (Prescription, error)
	DeleteAppointment(ctx context.Context, id int64) error
	DeleteDoctor(ctx context.Context, username string) error
	DeletePatient(ctx context.Context, username string) error
	DeletePrescription(ctx context.Context, appointmentID int64) error
	GetAppointmentById(ctx context.Context, id int64) (Appointment, error)
	GetDoctorByEmail(ctx context.Context, email string) (Doctor, error)
	GetDoctorByUsername(ctx context.Context, username string) (Doctor, error)
	GetPatientByEmail(ctx context.Context, email string) (Patient, error)
	GetPatientByUsername(ctx context.Context, username string) (Patient, error)
	GetPrescription(ctx context.Context, appointmentID int64) (Prescription, error)
	ListCompletedPatientAppointments(ctx context.Context, patientUsername string) ([]Appointment, error)
	ListDoctorAppointments(ctx context.Context, doctorUsername string) ([]Appointment, error)
	ListDoctors(ctx context.Context, arg ListDoctorsParams) ([]Doctor, error)
	ListDoctorsBySpecialization(ctx context.Context, arg ListDoctorsBySpecializationParams) ([]Doctor, error)
	ListPatientAppointments(ctx context.Context, patientUsername string) ([]Appointment, error)
	ListPatients(ctx context.Context, arg ListPatientsParams) ([]Patient, error)
	ListTodayDoctorAppointments(ctx context.Context, doctorUsername string) ([]Appointment, error)
	ListTodayPatientAppointments(ctx context.Context, patientUsername string) ([]Appointment, error)
	ListUpcomingDoctorAppointments(ctx context.Context, doctorUsername string) ([]Appointment, error)
	ListUpcomingPatientAppointments(ctx context.Context, patientUsername string) ([]Appointment, error)
	UpdateAppointmentStatus(ctx context.Context, arg UpdateAppointmentStatusParams) (Appointment, error)
	UpdateDoctorPassword(ctx context.Context, arg UpdateDoctorPasswordParams) error
	UpdateDoctorProfile(ctx context.Context, arg UpdateDoctorProfileParams) (Doctor, error)
	UpdateFeedback(ctx context.Context, arg UpdateFeedbackParams) (Prescription, error)
	UpdateOnlineStatus(ctx context.Context, arg UpdateOnlineStatusParams) (Appointment, error)
	UpdatePatientPassword(ctx context.Context, arg UpdatePatientPasswordParams) error
	UpdatePatientProfile(ctx context.Context, arg UpdatePatientProfileParams) (Patient, error)
	UpdatePrescription(ctx context.Context, arg UpdatePrescriptionParams) (Prescription, error)
}

var _ Querier = (*Queries)(nil)
