# Aplicación de Gestión de Tareas

Esta es una aplicación de gestión de tareas construida con React, Next.js y un backend en Golang. Permite crear, leer, actualizar y eliminar tareas.

## Requisitos previos

- Docker: Asegúrate de tener Docker instalado en tu máquina. Puedes descargarlo desde [https://www.docker.com/get-started](https://www.docker.com/get-started).

## Instalación

1. Clona este repositorio en tu máquina local:

```sh
git clone git@github.com:Julianlg752/to-do.git
```
2. Navega al directorio del proyecto:
```sh
cd to-do
```
3. Compila el frontend
```sh
cd frontend && npm run build
```
4. Construye y ejecuta los contenedores de Docker:

```sh
cd .. && docker-compose up -d
```

Esto creará y ejecutará los contenedores de Docker para el frontend (React) y el backend (Golang).

4. Accede a la aplicación:

Abre tu navegador web y visita [http://localhost:3000](http://localhost:3000) para acceder a la aplicación de gestión de tareas.

5. Usuario de Pruebas
En el login usa este usuario de pruebas para ingresar al dashboard
```
usuario: admin
contraseña: admin_to_do
```
## Uso

Una vez que la aplicación esté en funcionamiento, podrás realizar las siguientes acciones:

- **Crear una tarea**: En la página de inicio, ingresa un título y una descripción para la nueva tarea y haz clic en el botón "Agregar tarea". La tarea se agregará a la lista de tareas.
- **Editar una tarea**: Para editar una tarea existente, haz clic en el botón "Editar" junto a la tarea que deseas modificar. Podrás cambiar el título y la descripción de la tarea. Luego, haz clic en el botón "Actualizar" para guardar los cambios.
- **Eliminar una tarea**: Para eliminar una tarea, haz clic en el botón "Eliminar" junto a la tarea que deseas eliminar. La tarea se eliminará de la lista.
- **Marcar una tarea como completada**: Cada tarea tiene un checkbox que puedes marcar para indicar que la tarea está completada. Esta acción se actualizará en el backend.

## Backend en Golang y arquitectura hexagonal

El backend de esta aplicación está desarrollado en Golang y sigue una arquitectura hexagonal. Esto significa que se ha separado la lógica de negocio del resto de la aplicación, lo que permite una mayor modularidad, reutilización y testabilidad del código.

## To-Do
- Pruebas Unitarias
- Implementacion Paginación

## Contribuir

Si deseas contribuir a este proyecto, puedes seguir estos pasos:

1. Haz un fork de este repositorio.
2. Crea una nueva rama para tu contribución: `git checkout -b mi-nueva-funcionalidad`.
3. Realiza tus cambios y realiza commits: `git commit -m "Agrega mi nueva funcionalidad"`.
4. Haz push a tu rama: `git push origin mi-nueva-funcionalidad`.
5. Abre una Pull Request en este repositorio.

¡Agradecemos tu contribución!

## Licencia

Este proyecto está bajo la Licencia MIT. Para más información, consulta el archivo [LICENSE](LICENSE).
