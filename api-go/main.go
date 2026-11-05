package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
    cargarEnv(".env")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status": "ok"}`))
	})
http.HandleFunc("/login", login)
	http.HandleFunc("/qr", requiereToken(recibirMatriz))

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 

func recibirMatriz(w http.ResponseWriter, r *http.Request) {
	if permitirNavegador(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Metodo no permitido"}`))
		return
	}

	var body struct {
		Matriz [][]float64 `json:"matriz"`
	}


	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "JSON INVALIDO}`))
		return
	}


	if err := validar(body.Matriz); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}


	factorQ, factorR := factorizar(body.Matriz)
	analisis, err := pedirAnalisis(factorQ, factorR)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
w.Header().Set("Content-Type", "application/json")
encoder := json.NewEncoder(w)
encoder.SetIndent("", "  ")
encoder.Encode(map[string]any{
	"analisis": analisis,
})


}

// VAlIDAR LA MATRIZ 
func validar(a [][]float64) error {
	//VALIDAR SI LOS DATOS DE LA MATRIZ EXISTEN
	if len(a) == 0 || len(a[0]) == 0 {
		return fmt.Errorf("la matriz no puede estar vacia")
	}
	columnas := len(a[0])
	// RECORRIDO PARA VALIDAR QUE LA CANTIDAD DE MATRIZES SEAN LA MISMAS COLUMNAS Y FILAS
	for i := 1; i < len(a); i++ {
		if len(a[i]) != columnas {
			return fmt.Errorf("la matriz debe tener el mismo numero de columnas en todas las filas")
		}
	}
	return nil
}

// CONEXION CON LA API DE NODE OBTENCION DE RESULTADOS
func pedirAnalisis(factorQ, factorR [][]float64) (map[string]any, error) {

	playload, err := json.Marshal(map[string]any{
		"Q": factorQ,
		"R": factorR,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		"http://localhost:3000/operaciones",
		"application/json",
		bytes.NewReader(playload),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cuerpo, err := io.ReadAll(resp.Body)

	if err != nil{
		return nil, err
	}

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("node respondio %s", resp.Status)
	}

	var nodo struct {
	    Analisis map[string]any `json:"analisis"`
	}
	if err := json.Unmarshal(cuerpo, &nodo); err != nil {
		return nil, err
	}
	return nodo.Analisis, nil
	
}


func permitirNavegador(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

func cargarEnv(ruta string) {
	datos, err := os.ReadFile(ruta)
	if err != nil {
		log.Println("no se pudo leer .env:", err)
		return
	}
	for _, linea := range strings.Split(string(datos), "\n") {
		linea = strings.TrimSpace(linea)
		if linea == "" || strings.HasPrefix(linea, "#") {
			continue
		}
		clave, valor, ok := strings.Cut(linea, "=")
		if !ok {
			continue
		}
		os.Setenv(strings.TrimSpace(clave), strings.TrimSpace(valor))
	}
}