// Interface Mediador
type Mediador interface {
    MostrarMensaje(Usuario, string)
}

// Mediador Concreto
type ChatRoom struct{}

func (cr *ChatRoom) MostrarMensaje(usuario Usuario, mensaje string) {
    fmt.Printf("El mensaje de %s es: %s\n", usuario.GetNombre(), mensaje)
}

// Interface Colega
type Usuario interface {
    EnviarMensaje(string)
    GetNombre() string
}

// Colega Concreto
type UsuarioChat struct {
    nombre   string
    mediador Mediador
}

func (u *UsuarioChat) GetNombre() string {
    return u.nombre
}

func (u *UsuarioChat) EnviarMensaje(mensaje string) {
    u.mediador.MostrarMensaje(u, mensaje)
}

func main(){
	
	mediador := &ChatRoom{}

	usuarioA := &UsuarioChat{"Daniel", mediador}
	usuarioB := &UsuarioChat{"Pedro", mediador}

	usuarioA.EnviarMensaje("Hola como estas?")
	usuarioB.EnviarMensaje("Muy bien y vos?")
}