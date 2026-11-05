package main

import (
	"math"
	"testing"
)

func TestEjemploConocido(t *testing.T) {
	q, r := factorizar([][]float64{{0, 1}, {1, 0}})
	if diferencia(q, [][]float64{{0, 1}, {1, 0}}) != 0 || diferencia(r, [][]float64{{1, 0}, {0, 1}}) != 0 {
		t.Fatalf("Q=%v R=%v", q, r)
	}
}

func TestReconstruyeLaMatriz(t *testing.T) {
	a := [][]float64{{1, 2}, {3, 4}, {5, 6}}
	q, r := factorizar(a)
	producto := multiplicar(q, r)
	if diferencia(producto, a) > 1e-8 {
		t.Fatalf("QR no reconstruye la matriz: %v", producto)
	}
}

func multiplicar(a, b [][]float64) [][]float64 {
	filas := len(a)
	medio := len(b)
	columnas := len(b[0])
	c := make([][]float64, filas)
	for i := 0; i < filas; i++ {
		c[i] = make([]float64, columnas)
		for k := 0; k < medio; k++ {
			for j := 0; j < columnas; j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func diferencia(a, b [][]float64) float64 {
	mayor := 0.0
	for i := range a {
		for j := range a[i] {
			if d := math.Abs(a[i][j] - b[i][j]); d > mayor {
				mayor = d
			}
		}
	}
	return mayor
}