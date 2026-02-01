package model

type ServerResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type StudentDetailResponse struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	GroupName  string `json:"group_name"`
	MajorName  string `json:"major_name,omitempty"`
	SchoolName string `json:"school_name,omitempty"`
}

type ScheduleResponse struct {
	GroupName     string `json:"group_name"`
	Subject       string `json:"subject"`
	ProfessorName string `json:"professor_name,omitempty"`
	TimeSlot      string `json:"time_slot"`
}

type Attendance struct {
	StudentID int    `json:"student_id"`
	SubjectID int    `json:"subject_id"`
	VisitDay  string `json:"visit_day"`
	Visited   bool   `json:"visited"`
}
