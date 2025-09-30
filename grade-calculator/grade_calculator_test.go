package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 71, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 62, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 20, Assignment)
	gradeCalculator.AddGrade("exam 1", 14, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 24, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// AI was used for help in writing the TestGradeTypeString test function
func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gt       GradeType
		expected string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		actual := test.gt.String()
		if actual != test.expected {
			t.Errorf("Expected GradeType(%d).String() to return '%s', got '%s'", test.gt, test.expected, actual)
		}
	}
}

func TestZeroLenInComputeAvg(t *testing.T) {
	result := computeAverage([]Grade{})
	if result != 0 {
		t.Errorf("Dividing by 0 should return 0 to avoid error, got %d", result)
	}
}

func TestPassedClassPass(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam", 100, Exam)
	gradeCalculator.AddGrade("essay", 100, Essay)

	if !gradeCalculator.passed {
		t.Errorf("Expected passed to be true, got false")
	}
}

func TestPassedClassFail(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 0, Assignment)
	gradeCalculator.AddGrade("exam", 40, Exam)
	gradeCalculator.AddGrade("essay", 20, Essay)

	if gradeCalculator.passed {
		t.Errorf("Expected passed to be false, got true")
	}
}

