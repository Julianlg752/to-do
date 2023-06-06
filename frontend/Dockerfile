# Establecer la imagen base
FROM node:14-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de dependencias
COPY package.json ./

# Instalar las dependencias del proyecto
RUN yarn install --frozen-lockfile

# Copiar el código fuente al contenedor
COPY . .

# Compilar el proyecto Next.js
RUN yarn build

# Exponer el puerto utilizado por la aplicación
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["npm", "start"]
