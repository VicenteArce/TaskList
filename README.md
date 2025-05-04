# citiapscap

## Requisitos y Ejecución
Como requisitos se tiene MongoDB, Golang y Node.js.

1. Clonar el repositorio  
Primero clona el repositorio con:  
```bash
git clone https://github.com/VicenteArce/TaskList-CitiApsCap
```

2. Ejecución del proyecto
Primero, cree una base de datos en MongoDB cuyo nombre sea: "tododatabase" y una colección llamada "tasks". El nombre de estas se encuentra hardcodeada dentro del controlador de las tareas, pero se podría haber dejado el nombre de la base de datos como una variable de entorno.  
Para ejecutar el backend, abra la carpeta [backend](./backend) y ejecute el backend con el siguiente comando:  
```bash
go run main.go
```  
En este caso no será necesario instalar dependencias, ya que estas ya vienen insladas en los archivos correspondientes.  

Para la ejecución del frontend, abra la carpeta [frontend](./frontend) y ejecute los siguientes comandos:  
```bash
npm install
npm run dev
```
