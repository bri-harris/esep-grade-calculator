package esepunittests

type GradeCalculator struct {
	grades     []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})

}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentType := filterTypes(gc.grades, Assignment)
	examType := filterTypes(gc.grades, Exam)
	essayType := filterTypes(gc.grades, Essay)

	assignment_average := computeAverage(assignmentType)
	exam_average := computeAverage(examType)
	essay_average := computeAverage(essayType)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	for _, g := range grades {
		sum += g.Grade
	}

	if len(grades) == 0 {
		return 0
	}

	return sum / len(grades)
}

func filterTypes(grades []Grade, gradeType GradeType) []Grade {
	desiredType := make([]Grade, 0)
	for _, g := range grades {
		if g.Type == gradeType {
			desiredType = append(desiredType, g)
		}
	}
	return desiredType
}
