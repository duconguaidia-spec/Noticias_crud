# noticias_crud

API REST para la gestión de noticias desarrollada en Go con conexión a PostgreSQL.

## Especificaciones Técnicas

### Tecnologías Implementadas

* [Golang](https://go.dev/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [PostgreSQL](https://www.postgresql.org/)
* [lib/pq](https://github.com/lib/pq)

## Ejecución del Proyecto

```bash
# 1. Clonar el repositorio
git clone https://github.com/tu_usuario/noticias_crud

# 2. Moverse a la carpeta del repositorio
cd noticias_crud

# 3. Instalar dependencias
go mod tidy

# 4. Configurar las variables en config/db.go con los datos de conexión

# 5. Ejecutar el proyecto
go run main.go
```

## Endpoints

### Noticias

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/noticias` | Obtiene todas las noticias |
| GET | `/noticias/{id}` | Obtiene una noticia por ID |
| POST | `/noticias` | Crea una nueva noticia |
| PUT | `/noticias/{id}` | Actualiza una noticia por ID |
| DELETE | `/noticias/{id}` | Elimina una noticia por ID |

## Ejemplos de Uso

### GET /noticias
```bash
localhost:8083/noticias
```

### GET /noticias/{id}
```bash
localhost:8083/noticias/1
```

### POST /noticias
```bash
localhost:8083/noticias
{
    "titulo": "Precio del kilo de café pergamino sube un 12% en el Huila",
    "tipo": "Articulo",
    "cuerpo": "El mercado cafetero del departamento del Huila registró un incremento del 12% en el precio del kilo de café pergamino durante el primer trimestre del año, según datos de la Cooperativa de Caficultores del sur del país.",
    "url_video": null,
    "imagen_destacada": "https://cdn.agrocampo.co/img/noticia-precio-cafe.jpg",
    "fuente": "Federación Nacional de Cafeteros",
    "fecha_noticia": "2025-03-20T00:00:00Z",
    "id_usuario": 1,
    "acceso_limitado": false,
    "estado": "Publicado",
    "activo": true,
    "fecha_creacion": "2026-04-13T14:09:56.209884Z",
    "fecha_modificacion": "2026-04-13T14:09:56.209884Z"
}
```

### PUT /noticias/{id}
```bash
localhost:8083/noticias/4 
{
    "id": 4,
    "titulo": "Precio del kilo de café pergamino sube un 12% en el Huila",
    "tipo": "Articulo",
    "cuerpo": "El mercado cafetero del departamento del Huila registró un incremento del 12% en el precio del kilo de café pergamino durante el primer trimestre del año, según datos de la Cooperativa de Caficultores del sur del país.",
    "url_video": null,
    "imagen_destacada": "https://cdn.agrocampo.co/img/noticia-precio-cafe.jpg",
    "fuente": "Federación Nacional de Cafeteros",
    "fecha_noticia": "2025-03-20T00:00:00Z",
    "id_usuario": 1,
    "acceso_limitado": True,
    "estado": "Publicado",
    "activo": true,
    "fecha_creacion": "2026-04-13T14:09:56.209884Z",
    "fecha_modificacion": "2026-04-13T14:09:56.209884Z"
}
```

### DELETE /noticias/{id}
```bash
localhost:8083/noticias/1
```
