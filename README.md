# 🎟️ Virtual Queue System (Sistema de Fila Virtual)

Un sistema backend de alta concurrencia diseñado para manejar la venta masiva de entradas (ej. festivales de música) sin bloquear la base de datos, implementando una sala de espera en tiempo real.

## 🚀 Tecnologías

- **Backend:** Go (Golang)
- **Base de Datos Relacional:** MySQL
- **Caché / Cola de mensajes:** Redis
- **Comunicación en Tiempo Real:** Server-Sent Events (SSE)
- **Frontend:** React (Próximamente)
- **Arquitectura:** Clean Architecture

## 🧠 Arquitectura y Flujo

1. **El Muro (Redis):** Cuando miles de usuarios intentan comprar al mismo tiempo, el servidor HTTP de Go no consulta a MySQL. Encola a los usuarios en Redis, respondiendo en milisegundos.
2. **Workers (Goroutines):** Un _pool_ de workers en Go procesa la cola de Redis de forma asíncrona, controlando el flujo exacto de peticiones que llegan a MySQL para evitar sobrecargas y ventas dobles (_overselling_).
3. **Sala de Espera (SSE):** El cliente de React mantiene una conexión unidireccional con Go a través de Server-Sent Events, recibiendo actualizaciones en tiempo real sobre su turno en la fila.

## 📁 Estructura del Proyecto

El proyecto sigue los principios de **Clean Architecture** para mantener el dominio completamente aislado de los frameworks y la infraestructura (bases de datos, colas, protocolos HTTP).

## 👨‍💻 Autor

**Johan Sebastian Vasquez Diaz**  
_Desarrollador Fullstack | Tecnólogo en Sistemas de Información_

- **LinkedIn:** [Johan Sebastian Vasquez Diaz](https://www.linkedin.com/in/johan98vdiaz/)
- **GitHub:** [@diazjohan98](https://github.com/diazjohan98)
