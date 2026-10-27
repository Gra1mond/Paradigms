package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/solve", FindHandler)

	fmt.Println("Сервер запущен: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}

func FindSqrt(a, b, c float64) []float64 {
	if a == 0 {
		fmt.Println("а не должно быть равно 0")
		return nil
	}
	d := b*b - 4*a*c
	if d < 0 {
		fmt.Println("действительных корней нет")
		return nil
	}

	ts := []float64{(-b + math.Sqrt(d)) / (2 * a)}
	if d > 0 {
		ts = append(ts, (-b-math.Sqrt(d))/(2*a))
	}
	return ts
}

type Equation struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Используйте POST", http.StatusMethodNotAllowed)
		return
	}

	var eq Equation

	if err := json.NewDecoder(r.Body).Decode(&eq); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	if eq.A == 0 {
		http.Error(w, "a не должно быть равно 0", http.StatusBadRequest)
		return
	}

	ts := FindSqrt(eq.A, eq.B, eq.C)
	roots := []float64{}

	for _, t := range ts {
		if t > 0 {
			x := math.Sqrt(t)
			roots = append(roots, x, -x)
		} else if t == 0 {
			roots = append(roots, 0)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roots)
}
