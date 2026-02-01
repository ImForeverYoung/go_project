package storage

import (
	"HW_5/internal/model"
	"context"
	"fmt"
	"os"
)

// GetGroupStudentCounts returns the number of students in each group using GROUP BY
func (s *Storage) GetGroupStudentCounts(ctx context.Context) ([]model.GroupStat, error) {
	query := `
		SELECT g.group_name, COUNT(s.student_id)
		FROM groups g
		LEFT JOIN students s ON g.group_id = s.group_id
		GROUP BY g.group_name
		ORDER BY g.group_name;
	`
	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetGroupStudentCounts failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var results []model.GroupStat
	for rows.Next() {
		var stat model.GroupStat
		if err := rows.Scan(&stat.GroupName, &stat.StudentCount); err != nil {
			return nil, err
		}
		results = append(results, stat)
	}
	return results, nil
}

// GetStudentsWithAbsences returns students who missed more than minAbsences using HAVING
func (s *Storage) GetStudentsWithAbsences(ctx context.Context, minAbsences int) ([]model.StudentAttendanceStat, error) {
	query := `
		SELECT s.student_name, COUNT(a.id) as missed_count
		FROM students s
		JOIN attendance a ON s.student_id = a.student_id
		WHERE a.visited = FALSE
		GROUP BY s.student_name
		HAVING COUNT(a.id) >= $1;
	`
	rows, err := s.conn.Query(ctx, query, minAbsences)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetStudentsWithAbsences failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var results []model.StudentAttendanceStat
	for rows.Next() {
		var stat model.StudentAttendanceStat
		if err := rows.Scan(&stat.StudentName, &stat.MissedClasses); err != nil {
			return nil, err
		}
		results = append(results, stat)
	}
	return results, nil
}
