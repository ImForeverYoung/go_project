package model

type GroupStat struct {
	GroupName    string `json:"group_name"`
	StudentCount int    `json:"student_count"`
}

type StudentAttendanceStat struct {
	StudentName   string `json:"student_name"`
	MissedClasses int    `json:"missed_classes"`
}
