package storage

import (
	"HW_5/internal/model"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn *pgx.Conn
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) GetStudent(ctx context.Context, id string) (model.StudentDetailResponse, error) {
	var fullName string
	var gender string
	var birthDate string
	var groupName string
	var majorName string
	var schoolName string

	query := `
		SELECT s.student_name, s.gender, s.birth_date::text, g.group_name, m.major_name, sc.school_name
		FROM students s
		JOIN groups g ON s.group_id = g.group_id
		JOIN majors m ON g.major_id = m.major_id
		JOIN schools sc ON m.school_id = sc.school_id
		WHERE s.student_id=$1`

	err := s.conn.QueryRow(ctx, query, id).Scan(&fullName, &gender, &birthDate, &groupName, &majorName, &schoolName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return model.StudentDetailResponse{}, err
	}

	idInt, _ := strconv.Atoi(id)

	return model.StudentDetailResponse{
		ID:         idInt,
		FullName:   fullName,
		Gender:     gender,
		BirthDate:  birthDate,
		GroupName:  groupName,
		MajorName:  majorName,
		SchoolName: schoolName,
	}, nil
}

func (s *Storage) GetAllSchedule(ctx context.Context) ([]model.ScheduleResponse, error) {
	results := []model.ScheduleResponse{}

	query := `
		SELECT g.group_name, sub.subject_name, p.name, sch.time_slot 
		FROM schedules sch 
		JOIN groups g ON sch.group_id = g.group_id 
		JOIN subjects sub ON sch.subject_id = sub.subject_id
		LEFT JOIN professors p ON sub.professor_id = p.id`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var groupName string
		var subject string
		var professorName *string // Handle nullable professor
		var timeSlot string

		err := rows.Scan(&groupName, &subject, &professorName, &timeSlot)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			continue
		}

		profName := ""
		if professorName != nil {
			profName = *professorName
		}

		schedule := model.ScheduleResponse{
			GroupName:     groupName,
			Subject:       subject,
			ProfessorName: profName,
			TimeSlot:      timeSlot,
		}
		results = append(results, schedule)
	}
	return results, nil
}

func (s *Storage) GetGroupSchedule(ctx context.Context, id string) ([]model.ScheduleResponse, error) {
	results := []model.ScheduleResponse{}

	query := `
		SELECT g.group_name, sub.subject_name, p.name, sch.time_slot 
		FROM schedules sch 
		JOIN groups g ON sch.group_id = g.group_id 
		JOIN subjects sub ON sch.subject_id = sub.subject_id
		LEFT JOIN professors p ON sub.professor_id = p.id
		WHERE g.group_id=$1`

	rows, err := s.conn.Query(ctx, query, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var groupName string
		var subject string
		var professorName *string
		var timeSlot string

		err := rows.Scan(&groupName, &subject, &professorName, &timeSlot)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			continue
		}

		profName := ""
		if professorName != nil {
			profName = *professorName
		}

		schedule := model.ScheduleResponse{
			GroupName:     groupName,
			Subject:       subject,
			ProfessorName: profName,
			TimeSlot:      timeSlot,
		}
		results = append(results, schedule)
	}
	return results, nil
}

func (s *Storage) MarkAttendance(ctx context.Context, request model.Attendance) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	queryI := `INSERT INTO attendance (student_id, subject_id, visit_day, visited)
VALUES ($1, $2, $3, $4) RETURNING id;`

	row := s.conn.QueryRow(ctx, queryI,
		request.StudentID,
		request.SubjectID,
		request.VisitDay,
		request.Visited,
	)
	err = row.Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return 0, err
	}

	return id, nil

}

func (s *Storage) GetAttendanceBySubjectId(ctx context.Context, id string) ([]model.Attendance, error) {
	results := []model.Attendance{}

	rows, err := s.conn.Query(ctx, `select student_id, visit_day::text, visited
			 from attendance
			 where subject_id=$1`, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var studentId int
		var visitedDay string
		var visited bool

		rows.Scan(&studentId, &visitedDay, &visited)
		idInt, _ := strconv.Atoi(id)
		attendance := model.Attendance{
			StudentID: studentId,
			SubjectID: idInt,
			VisitDay:  visitedDay,
			Visited:   visited,
		}
		results = append(results, attendance)
	}
	return results, nil
}

func (s *Storage) GetAttendanceByStudentId(ctx context.Context, id string) ([]model.Attendance, error) {
	results := []model.Attendance{}

	rows, err := s.conn.Query(ctx, `select subject_id, visit_day::text, visited
			 from attendance
			 where student_id=$1`, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subjectId int
		var visitedDay string
		var visited bool

		rows.Scan(&subjectId, &visitedDay, &visited)
		idInt, _ := strconv.Atoi(id)
		attendance := model.Attendance{
			StudentID: idInt,
			SubjectID: subjectId,
			VisitDay:  visitedDay,
			Visited:   visited,
		}
		results = append(results, attendance)
	}
	return results, nil
}
