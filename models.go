package main

type Student struct {
	StudentID uint    `gorm:"primaryKey;autoIncrement" json:"student_id"`
	Name      string  `json:"name"`
	Group     string  `json:"group"`
	Email     string  `json:"email"`
	Grades    []Grade `json:"grades,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}

type Subject struct {
	SubjectID uint    `gorm:"primaryKey;autoIncrement" json:"subject_id"`
	Name      string  `json:"name"`
	Grades    []Grade `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}

type Grade struct {
	GradeID   uint    `gorm:"primaryKey;autoIncrement" json:"grade_id"`
	StudentID uint    `json:"student_id"`
	SubjectID uint    `json:"subject_id"`
	Grade     float64 `json:"grade"`
}
