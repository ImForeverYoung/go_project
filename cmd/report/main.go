package main

import (
	storage "HW_5/internal/storage"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/postgres" // Default fallback
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	store := storage.NewStorage(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("--- REPORT START ---")

	// 1. Verify Student Details (including Major and School)
	fmt.Println("\n[1] Testing GetStudent (ID=1)...")
	student, err := store.GetStudent(ctx, "1")
	if err != nil {
		log.Printf("Error getting student: %v\n", err)
	} else {
		fmt.Printf("Student: %s\nMajor: %s\nSchool: %s\n", student.FullName, student.MajorName, student.SchoolName)
	}

	// 2. Verify Schedules (including Professor)
	fmt.Println("\n[2] Testing GetAllSchedule...")
	schedules, err := store.GetAllSchedule(ctx)
	if err != nil {
		log.Printf("Error getting schedules: %v\n", err)
	} else {
		for _, s := range schedules {
			fmt.Printf("Group: %s | Subject: %s | Prof: %s | Time: %s\n", s.GroupName, s.Subject, s.ProfessorName, s.TimeSlot)
		}
	}

	// 3. Verify Group Analytics (GROUP BY)
	fmt.Println("\n[3] Testing GetGroupStudentCounts (GROUP BY)...")
	groupStats, err := store.GetGroupStudentCounts(ctx)
	if err != nil {
		log.Printf("Error getting group stats: %v\n", err)
	} else {
		for _, stat := range groupStats {
			fmt.Printf("Group: %s | Count: %d\n", stat.GroupName, stat.StudentCount)
		}
	}

	// 4. Verify Attendance Analytics (HAVING)
	fmt.Println("\n[4] Testing GetStudentsWithAbsences (HAVING missed >= 1)...")
	absentStudents, err := store.GetStudentsWithAbsences(ctx, 1)
	if err != nil {
		log.Printf("Error getting absent students: %v\n", err)
	} else {
		for _, stat := range absentStudents {
			fmt.Printf("Student: %s | Missed: %d\n", stat.StudentName, stat.MissedClasses)
		}
	}

	fmt.Println("\n--- REPORT END ---")
}
