package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Función para enviar respuestas en formato JSON
func sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//DATOS DEL ESTUDIANTE

//Función para la creacion de un estudiante
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student

	//Convertimos el JSON que envía el usuario en una estructura Student
	json.NewDecoder(r.Body).Decode(&student)

	//Guardamos los datos ingrsados en la base de datos
	DB.Create(&student)

	//Mostramos el estudiante creado
	sendJSON(w, student)
}

//Función para obtener todos los datos de los estudiantes
func GetStudents(w http.ResponseWriter, r *http.Request) {
	var students []Student

	DB.Find(&students)

	sendJSON(w, students)
}

//Función para obtener un estudiante por su ID
func GetStudent(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	var student Student
	DB.First(&student, id)

	sendJSON(w, student)
}

//Función para actualizar a un estudiante
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	var student Student
	DB.First(&student, id)

	//Variable para recibir datos nuevos
	var data Student
	json.NewDecoder(r.Body).Decode(&data)

	//Reemplazamos los datos por los nuevos
	student.Name = data.Name
	student.Group = data.Group
	student.Email = data.Email

	DB.Save(&student)

	sendJSON(w, student)
}

//Función para eliminar a un estudiante
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	DB.Delete(&Student{}, id)

	w.Write([]byte("El estudiante ha sido eliminado con éxito"))
}

//DATOS DE LAS MATERIAS

//Función para crear una materia
func CreateSubject(w http.ResponseWriter, r *http.Request) {
	var subject Subject
	json.NewDecoder(r.Body).Decode(&subject)

	DB.Create(&subject)

	sendJSON(w, subject)
}

//Función para obtener una  materia por ID
func GetSubject(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	var subject Subject
	DB.First(&subject, id)

	sendJSON(w, subject)
}

//Función para actualizar una materia
func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	var subject Subject
	DB.First(&subject, id)

	var data Subject
	json.NewDecoder(r.Body).Decode(&data)

	subject.Name = data.Name

	DB.Save(&subject)

	sendJSON(w, subject)
}

//Función para eliminar una materia
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	DB.Delete(&Subject{}, id)

	w.Write([]byte("La materia ha sido eliminada con éxito"))
}

//DATOS DE LAS CALIFICACIONES

//Función para crear una calificación
func CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade Grade
	json.NewDecoder(r.Body).Decode(&grade)

	DB.Create(&grade)

	sendJSON(w, grade)
}

//Función para actualizar una calificación
func UpdateGrade(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	var grade Grade
	DB.First(&grade, id)

	var data Grade
	json.NewDecoder(r.Body).Decode(&data)

	grade.Grade = data.Grade
	grade.StudentID = data.StudentID
	grade.SubjectID = data.SubjectID

	DB.Save(&grade)

	sendJSON(w, grade)
}

//Función para eliminar una calificación
func DeleteGrade(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTexto)

	DB.Delete(&Grade{}, id)

	w.Write([]byte("La calificación ha sido eliminada con éxito"))
}

//Función para obtener todas las calificaciones de un estudiante por ID
func GetGradesByStudent(w http.ResponseWriter, r *http.Request) {
	idTexto := mux.Vars(r)["student_id"]
	studentID, _ := strconv.Atoi(idTexto)

	var grades []Grade
	DB.Where("student_id = ?", studentID).Find(&grades)

	sendJSON(w, grades)
}
