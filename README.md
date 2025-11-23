1) Clonar el repositorio

   git clone https://github.com/Carlos-Daniel-2025/Sistema-de-Control-Escolar-API-REST-.git

2) Entrar a la carpeta del Proyecto

   Escuela

3) Instalar dependencias

   go mod tidy

4) Ejecutar el servidor

   go run .

5) El servidor se debe de iniciar en: http://localhost:3030
 ________________________________________________________
   

   DESCRIPCIÓN DE RUTAS PRINCIPALES DEL API Y SU PROPOSITO
   

ESTUDIANTES (/api/students)

Método	    //     Ruta	         //           Descripción

POST	      /api/students	            //Crear un estudiante

GET	        /api/students	            //Consultar todos los estudiantes

GET	        /api/students/{id}	      //Consultar estudiante por ID

PUT	        /api/students/{id}	      //Actualizar estudiante

DELETE	    /api/students/{id}	      //Eliminar estudiante
__________________________________________________________

MATERIAS (/api/subjects)

Método	   //      Ruta           //          Descripción

POST	      /api/subjects	            //Crear una materia

GET	        /api/subjects/{id}	      //Consultar materia por ID

PUT	        /api/subjects/{id}	      //Actualizar materia

DELETE	    /api/subjects/{id}	      //Eliminar materia
__________________________________________________________

CALIFICACIONES (/api/grades)

Método	   //     Ruta	              //                   Descripción

POST	      /api/grades	                          //Crear una calificación

PUT	        /api/grades/{id}	                    //Actualizar calificación

DELETE	    /api/grades/{id}	                    //Eliminar calificación

GET	        /api/grades/student/{student_id}	    //Ver calificaciones de un estudiante
__________________________________________________________

//cURL DE POSTMAN (EJEMPLOS DE PETICIONES)

//EJEMPLO DE DELETE STUDENT 2

curl --location --request DELETE 'http://localhost:3030/api/students/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Ximena Diaz Sanchez",
    "group": "3A",
    "email": "ximenadisa45@gmail.com"
}'

//EJEMPLO DE POST EN MATERIAS

curl --location 'http://localhost:3030/api/subjects' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Fisica"
}'

//EJEMPLOS DE GET EN GRADES STUDENT 2

curl --location --request GET 'http://localhost:3030/api/grades/student/2' \
--header 'Content-Type: application/json' \
--data '{
    "student_id": 3,
    "subject_id": 1,
    "grade": 7.0
}'
___________________________________________________________


HERRAMIENTAS USADAS PARA ESTE PROYECTO

1. Go 1.25
2. Gorilla Mux
3. GORM
4. SQLite
