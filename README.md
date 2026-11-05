# Análisis de matrices

El proyecto tiene tres procesos. Hay que abrir uno por terminal, en este orden:

1. **Node** calcula máximo, mínimo, promedio, suma y si alguna matriz es diagonal.
2. **Go** recibe la matriz, calcula la factorización QR y le pide ese análisis a Node.
3. **El formulario** envía la matriz a Go y muestra la respuesta.

Node tiene que estar corriendo antes de usar Go. Si no, Go no puede completar la petición y responde un error.

## 1. API Node

```bash
cd dctestgo/api-node
npm start
```

Queda en `http://localhost:3000`. No hace falta instalar dependencias.

## 2. API Go

En otra terminal:

```bash
cd dctestgo/api-go
go run .
```

Queda en `http://localhost:8080`.

## 3. Formulario

En una tercera terminal, la primera vez instala los paquetes:

```bash
cd dctestgo/front/frontData
npm install
npm run dev
```

Vite imprime la dirección local, normalmente `http://localhost:5173`. Ábrela, llena las seis casillas y pulsa **Analizar**.

por ultimo se mostrara un cuadro de text input en una matriz donde debera ingresar los valores

Las variables de entorno pueden ser cualqueir usaurio que asignes para realizar las pruebas