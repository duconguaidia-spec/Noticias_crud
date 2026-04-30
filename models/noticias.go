package models
import "time"
type Noticia struct {
    ID *int `json:"id"`
    Titulo *string `json:"titulo"`
    Tipo *string `json:"tipo"`
    Cuerpo *string `json:"cuerpo"`
    UrlVideo *string `json:"url_video"`
    ImagenDestacada *string `json:"imagen_destacada"`
    Fuente *string `json:"fuente"`
    FechaNoticia *string `json:"fecha_noticia"`       
    IDUsuario *int `json:"id_usuario"`
    AccesoLimitado *bool `json:"acceso_limitado"`
    Estado *string `json:"estado"`
    Activo *bool `json:"activo"`
    FechaCreacion *time.Time `json:"fecha_creacion"`     
    FechaModificacion *time.Time `json:"fecha_modificacion"`  
}