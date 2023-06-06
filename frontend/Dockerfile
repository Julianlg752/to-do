# Establecer la imagen base
FROM node:16-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de dependencias
COPY . .

# Instalar las dependencias del proyecto
RUN npm install


# Exponer el puerto utilizado por la aplicación
EXPOSE 3000

# Compilar el proyecto Next.js
RUN npm run build
CMD [ "npm", "run", "start" ]

