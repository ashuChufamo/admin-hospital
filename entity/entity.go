package entity

import "time"

// Category represents Food Menu Category
type Profile struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null;unique"`
	// UserName      string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:varchar(255);"`
	Email       string `gorm:"type:varchar(255);"`
	Phone       string `gorm:"type:varchar(255);"`
	Address     string `gorm:"type:varchar(255);"`
	Image       string `gorm:"type:varchar(255);"`
	Sex         string `gorm:"type:varchar(255);"`
	Role        string `gorm:"type:varchar(255);"`
	BirthDate   time.Time
	Description string
}
type Doctor struct {
	ID         uint    `gorm:"not null"`
	Profile    Profile `gorm:"ForeignKey:ID"`
	Uuid       uint
	Department string `gorm:"type:varchar(255);not null"`

	//DoctorHistory []DoctorHistory
	//PetientHistory []PetientHistory
	Prescription []Prescription `gorm:"ForeignKey:DoctorId"`
	Diagnosis    []Diagnosis    `gorm:"ForeignKey:DoctorId"`
	Appointment  []Appointment  `gorm:"ForeignKey:DoctorId"`
}
type Appointment struct {
	ID           uint
	PatientId    uint   `gorm:"not null"`
	PatientUname string `gorm:"type:varchar(255);not null"`
	DoctorId     uint   `gorm:"not null"`
	Date         time.Time
}
type Petient struct {
	ID      uint `gorm:"not null"`
	Uuid    uint
	Profile Profile `gorm:"ForeignKey:Uuid"`

	BloodGroup   string `gorm:"type:varchar(255);not null"`
	Age          int
	Prescription []Prescription `gorm:"ForeignKey:PatientId"`
	Diagnosis    []Diagnosis    `gorm:"ForeignKey:PatientId"`
	Appointment  []Appointment  `gorm:"ForeignKey:PatientId"`
	Request      []Request      `gorm:"ForeignKey:PatientId"`
}
type Pharmacist struct {
	ID           uint
	Uuid         uint
	Profile      Profile
	Medicine     []Medicine     `gorm:"ForeignKey:AddedBy"`
	Prescription []Prescription `gorm:"ForeignKey:PhrmacistId"`
}
type Laboratorist struct {
	ID        uint
	Uuid      uint
	Profile   Profile
	Diagnosis []Diagnosis `gorm:"ForeignKey:LaboratoristId"`
}

type Admin struct {
	ID      uint
	Uuid    uint
	Profile Profile `gorm:"ForeignKey:Uuid"`
	// Appointment []Appointment `gorm:"many2many:admin_appointment"`
	Request []Request `gorm:"ForeignKey:AdminId"`
}

type Request struct {
	ID            uint
	PatientId     uint
	DoctorId      uint
	ApproveStatus string `gorm:"type:varchar(255);"`
	ApprovedBy    uint
}

type Prescription struct {
	ID           uint
	PatientId    uint `gorm:"not null"`
	DoctorId     uint `gorm:"not null"`
	PhrmacistId  uint
	Date         time.Time
	MedicineName string `gorm:"type:varchar(255);"`
	Description  string `gorm:"type:varchar(255);"`
	GivenStatus  string `gorm:"type:varchar(255);"`
}

type Medicine struct {
	ID           uint
	CategoryName string `gorm:"type:varchar(255);not null"`
	MedicineName string `gorm:"type:varchar(255);not null"`
	ExpiredAt    time.Time
	Amount       uint
	AddedBy      uint
}

type Diagnosis struct {
	ID             uint
	PatientId      uint `gorm:"not null"`
	DoctorId       uint `gorm:"not null"`
	LaboratoristId uint
	Date           time.Time
	Reponse        string `gorm:"type:varchar(255);"`
	Description    string `gorm:"type:varchar(255);"`
}

type Error struct {
	Code    int
	Message string
}
