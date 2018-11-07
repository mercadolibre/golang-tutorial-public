// Interface
type Receptor interface {
    ProcesarMensaje(int, string) string
}

// Receptor de Alta Prioridad
type ReceptorAltaPrioridad struct{
    siguiente Receptor
}

func (rap ReceptorAltaPrioridad) ProcesarMensaje(prioridad int, mensaje string) string {
    if prioridad >= 5 {
       return "Procesando mensaje de alta prioridad: " + mensaje
    }

    if rap.siguiente != nil {
       return rap.siguiente.ProcesarMensaje(prioridad, mensaje)
    }

    return ""
}

// Receptor de Baja Prioridad
type ReceptorBajaPrioridad struct{
    siguiente Receptor
}

func (rbp ReceptorBajaPrioridad) ProcesarMensaje(prioridad int, mensaje string) string {
    if prioridad < 5 {
       return "Procesando mensaje de baja prioridad: " + mensaje
    }

    if rbp.siguiente != nil {
       return rbp.siguiente.ProcesarMensaje(prioridad, mensaje)
    }

    return ""
}


func main () {
  manejadores := ReceptorBajaPrioridad {
    siguiente: ReceptorAltaPrioridad {},
  }
  fmt.Println(manejadores.ProcesarMensaje(4, "Mensaje 1 - Prioridad 4"))
  fmt.Println(manejadores.ProcesarMensaje(5, "Mensaje 2 - Prioridad 5"))
  fmt.Println(manejadores.ProcesarMensaje(10, "Mensaje 3 - Prioridad 10"))
}