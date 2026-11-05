package main

import "math"

func factorizar(a [][]float64) ([][]float64, [][]float64) {
	r := copiar(a)
	filas := len(r)
	columnas := len(r[0])
	q := indentidad(filas)

	pasos := columnas
	if filas < pasos {
		pasos = filas
	}

	for k := 0; k < pasos; k++ {
		largo := filas - k

		cola := 0.0
		for i := 1; i < largo; i++ {
			cola = math.Hypot(cola, r[k+i][k])
		}
		x0 := r[k][k]
		if cola == 0 && x0 >= 0 {
			continue
		}

		norma := math.Hypot(x0, cola)
		beta := -norma
		if x0 < 0 {
			beta = norma
		}
		tau := (beta - x0) / beta
		denominador := x0 - beta

		v := make([]float64, largo)
		v[0] = 1
		for i := 1; i < largo; i++ {
			v[i] = r[k+i][k] / denominador
		}

		// La misma reflexión se aplica a R y a Q.
		for j := k; j < columnas; j++ {
			punto := 0.0
			for i := 0; i < largo; i++ {
				punto += v[i] * r[k+i][j]
			}
			factor := tau * punto
			for i := 0; i < largo; i++ {
				r[k+i][j] -= factor * v[i]
			}
		}
		for j := 0; j < filas; j++ {
			punto := 0.0
			for i := 0; i < largo; i++ {
				punto += v[i] * q[k+i][j]
			}
			factor := tau * punto
			for i := 0; i < largo; i++ {
				q[k+i][j] -= factor * v[i]
			}
		}
	}

	q = transponer(q)
    diagonalPositiva(q, r)
	limpiar(q)
    limpiar(r)
    return q, r
}

func transponer(a [][]float64) [][]float64 {
	filas := len(a)
	columnas := len(a[0])
	t := make([][]float64, columnas)
	for j := 0; j < columnas; j++ {
		t[j] = make([]float64, filas)
		for i := 0; i < filas; i++ {
			t[j][i] = a[i][j]
		}
	}
	return t
}

func indentidad(n int) [][]float64 {
	q := make([][]float64, n)
	for i := range q {
		q[i] = make([]float64, n)
		q[i][i] = 1
	}
	return q
}

func copiar(a [][]float64) [][]float64 {
	copia := make([][]float64, len(a))
	for i := range a {
		copia[i] = make([]float64, len(a[i]))
		copy(copia[i], a[i])
	}
	return copia
}

func diagonalPositiva(q, r [][]float64) {
	filas := len(q)
	columnas := len(r[0])
	limite := filas
	if columnas < limite {
		limite = columnas
	}

	for k := 0; k < limite; k++ {
		if r[k][k] >= 0 {
			continue
		}
		for i := 0; i < filas; i++ {
			q[i][k] = -q[i][k]
		}
		for j := k; j < columnas; j++ {
			r[k][j] = -r[k][j]
		}
	}
}

func limpiar(m [][]float64) {
	for i := range m {
		for j := range m[i] {
			if math.Abs(m[i][j]) < 1e-12 {
				m[i][j] = 0
			}
		}
	}
}