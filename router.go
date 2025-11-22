package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	//METODOS PARA ESTUDIANTES

	r.HandleFunc("/api/students", CreateStudent).Methods("POST") //Crear estudiante
	r.HandleFunc("/api/students", GetStudents).Methods("GET") //Obtener todos los estudiantes
	r.HandleFunc("/api/students/{id}", GetStudent).Methods("GET") //Obtener un estudiante
	r.HandleFunc("/api/students/{id}", UpdateStudent).Methods("PUT") //Actualizar estudiante
	r.HandleFunc("/api/students/{id}", DeleteStudent).Methods("DELETE") //Eliminar estudiante

	//METODOS PARA MATERIAS

	r.HandleFunc("/api/subjects", CreateSubject).Methods("POST") //Crear materia
	r.HandleFunc("/api/subjects/{id}", GetSubject).Methods("GET") //Crear materia
	r.HandleFunc("/api/subjects/{id}", UpdateSubject).Methods("PUT") //Actualizar materia
	r.HandleFunc("/api/subjects/{id}", DeleteSubject).Methods("DELETE") //Eliminar materia

	//METODOS PARA CALIFICACIONES

	r.HandleFunc("/api/grades", CreateGrade).Methods("POST") //Crear calificación
	r.HandleFunc("/api/grades/{id}", UpdateGrade).Methods("PUT") //Actualizar calificación
	r.HandleFunc("/api/grades/{id}", DeleteGrade).Methods("DELETE") //Eliminar calificación
	r.HandleFunc("/api/grades/student/{student_id}", GetGradesByStudent).Methods("GET") //Obtener calificación de una materia de un estudiante

	return r
}
