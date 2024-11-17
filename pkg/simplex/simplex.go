package simplex

import (
	"fmt"
	glpk "github.com/lukpank/go-glpk/glpk"
)

type SimplexOptimizer struct{}

func (s *SimplexOptimizer) Minimize(coefficients []float64, constraints [][]float64, bounds []float64, variableBounds [][2]float64) ([]float64, float64, error) {
	return Minimize(coefficients, constraints, bounds, variableBounds)
}

func Minimize(coefficients []float64, constraints [][]float64, bounds []float64, variableBounds [][2]float64) ([]float64, float64, error) {
	lp := glpk.New()
	defer lp.Delete()

	numVars := len(coefficients)
	numConstraints := len(constraints)

	lp.SetObjDir(glpk.MIN)
	lp.AddCols(numVars)
	for j := 1; j <= numVars; j++ {
		lb, ub := variableBounds[j-1][0], variableBounds[j-1][1]
		lp.SetColBnds(j, glpk.DB, lb, ub)
		lp.SetObjCoef(j, coefficients[j-1])
	}

	lp.AddRows(numConstraints)
	for i := 1; i <= numConstraints; i++ {
		lp.SetRowBnds(i, glpk.UP, 0, bounds[i-1])
		indices := make([]int32, numVars)
		values := make([]float64, numVars)
		for j := 0; j < numVars; j++ {
			indices[j] = int32(j + 1)
			values[j] = constraints[i-1][j]
		}
		lp.SetMatRow(i, indices, values)
	}

	err := lp.Simplex(nil)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка при решении задачи: %v", err)
	}

	status := lp.Status()
	if status != glpk.OPT {
		return nil, 0, fmt.Errorf("не удалось найти оптимальное решение")
	}

	result := make([]float64, numVars)
	for j := 1; j <= numVars; j++ {
		result[j-1] = lp.ColPrim(j)
	}
	objective := lp.ObjVal()

	return result, objective, nil
}
