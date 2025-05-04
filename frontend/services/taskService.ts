import httpCommon from "~/http-common";

// Función para crear una nueva tarea
const create = (data: { [key: string]: any }) => {
  return httpCommon.post("/tasks/", data);
};


// Función para obtener todas las tareas
const getAll = () => {
  return httpCommon.get("/tasks/");
};


// Función para modificar una tarea 
const update = (id: string, data: { [key: string]: any }) => {
  return httpCommon.put(`/tasks/${id}`, data);
};

// Función para marcar como completada una tarea
const complete = (id: string) => {
  return httpCommon.put(`/tasks/${id}/complete`);
}

// Función para eliminar una tarea
const remove = (id: string) => {
  return httpCommon.delete(`/tasks/${id}`);
};

export default {
  create,
  getAll,
  update,
  complete,
  remove
};