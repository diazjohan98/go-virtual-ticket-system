# High-Concurrency Virtual Queue System (Backend) 🚀

Este es el motor de alta concurrencia diseñado en **Go (Golang)** que maneja el flujo de usuarios e inventario para la venta masiva de boletos en eventos de alta demanda. Utiliza una arquitectura robusta para mitigar la sobrecarga de la base de datos principal mediante un sistema de colas asíncronas.

> 🖥️ **¿Buscas el Frontend?** Puedes ver la interfaz de usuario en el siguiente repositorio: [react-virtual-ticket-system](https://github.com/diazjohan98/react-virtual-ticket-system.git)

## 🧠 Arquitectura y Patrones

El proyecto está construido bajo los principios de **Clean Architecture** (Arquitectura Limpia), garantizando el desacoplamiento total entre la lógica de negocio, los frameworks y los motores de bases de datos.

- **Dominio (Domain):** Entidades puras e interfaces (contratos) del sistema.
- **Casos de Uso (Usecases):** Orquestación de la lógica de negocio (ej. validación e ingreso a la fila).
- **Infraestructura (Infrastructure):** Implementaciones concretas de las bases de datos (MySQL y Redis) y el Worker asíncrono.
- **Delivery (HTTP):** Manejo de peticiones y respuestas mediante handlers desacoplados.

## 🛠️ Tecnologías Utilizadas

- **Go (Golang):** Lenguaje principal, aprovechando las _Goroutines_ y _Contexts_ para concurrencia eficiente.
- **Redis:** Utilizado como buffer de alta velocidad (In-Memory) para gestionar la fila virtual mediante estructuras FIFO.
- **MySQL:** Base de datos relacional para la persistencia firme y el control estricto del inventario final de boletos.
- **Docker & Docker Compose:** Containerización de la infraestructura local para garantizar un entorno idéntico a producción.

## 🚀 Cómo Ejecutar el Backend

### Prerrequisitos

- Go (v1.20 o superior)
- Docker Desktop

### Pasos

1. Clonar el repositorio.
2. Crear un archivo `.env` en la raíz basado en el archivo `.env.example`.
3. Levantar la infraestructura de bases de datos con Docker:
   ```bash
   docker-compose up -d
   ```
4. Sincronizar las dependencias de Go:

```
Bash
go mod tidy
```

5. Arrancar el servidor HTTP y el Worker asíncrono:

```
Bash
go run cmd/api/main.go
```

El servidor estará escuchando de forma segura en http://localhost:8080.

## 👨‍💻 Autor

**Johan Sebastian Vasquez Diaz**  
_Desarrollador Fullstack | Tecnólogo en Sistemas de Información_

- **LinkedIn:** [Johan Sebastian Vasquez Diaz](https://www.linkedin.com/in/johan98vdiaz/)
- **GitHub:** [@diazjohan98](https://github.com/diazjohan98)
