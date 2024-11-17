package simplex

import (
	"errors"
	"fmt"
	glpk "github.com/lukpank/go-glpk/glpk"
)

// SimplexOptimizer реализует линейный оптимизатор
type SimplexOptimizer struct{}

// NewSimplexOptimizer создает новый экземпляр SimplexOptimizer
func NewSimplexOptimizer() *SimplexOptimizer {
	return &SimplexOptimizer{}
}

// Minimize решает задачу линейного программирования с помощью Simplex-метода
func (s *SimplexOptimizer) Minimize(
	coefficients []float64,
	constraints [][]float64,
	bounds []float64,
	variableBounds [][2]float64,
) ([]float64, float64, error) {
	// Проверка входных данных
	if err := validateInput(coefficients, constraints, bounds, variableBounds); err != nil {
		return nil, 0, err
	}

	// Создаем LP модель
	lp := glpk.New()
	defer lp.Delete()

	numVars := len(coefficients)
	numConstraints := len(constraints)

	// Устанавливаем направление оптимизации (минимизация)
	lp.SetObjDir(glpk.MIN)

	// Добавляем переменные и задаем их границы
	lp.AddCols(numVars)
	for j := 1; j <= numVars; j++ {
		lb, ub := variableBounds[j-1][0], variableBounds[j-1][1]
		lp.SetColBnds(j, glpk.DB, lb, ub)
		lp.SetObjCoef(j, coefficients[j-1])
	}

	// Добавляем ограничения
	lp.AddRows(numConstraints)
	for i := 1; i <= numConstraints; i++ {
		lp.SetRowBnds(i, glpk.UP, 0, bounds[i-1]) // Верхняя граница
		indices := make([]int32, numVars)
		values := make([]float64, numVars)
		for j := 0; j < numVars; j++ {
			indices[j] = int32(j + 1)
			values[j] = constraints[i-1][j]
		}
		lp.SetMatRow(i, indices, values)
	}

	// Запускаем Simplex-метод
	err := lp.Simplex(nil)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка при решении задачи: %v", err)
	}

	// Проверяем статус решения
	status := lp.Status()
	if status != glpk.OPT {
		return nil, 0, fmt.Errorf("решение не найдено, статус: %v", status)
	}

	// Извлекаем результаты
	result := make([]float64, numVars)
	for j := 1; j <= numVars; j++ {
		result[j-1] = lp.ColPrim(j)
	}
	objective := lp.ObjVal()

	return result, objective, nil
}

// validateInput проверяет входные данные на корректность
func validateInput(
	coefficients []float64,
	constraints [][]float64,
	bounds []float64,
	variableBounds [][2]float64,
) error {
	numVars := len(coefficients)

	// Проверяем количество ограничений
	if len(bounds) != len(constraints) {
		return errors.New("несоответствие между количеством ограничений и границ")
	}

	// Проверяем размер переменных и коэффициентов
	for _, constraint := range constraints {
		if len(constraint) != numVars {
			return errors.New("несоответствие между количеством переменных и коэффициентов в ограничениях")
		}
	}

	// Проверяем размер границ переменных
	if len(variableBounds) != numVars {
		return errors.New("несоответствие между количеством переменных и их границами")
	}

	return nil
}
