package alertas

import (
    "fmt"
)


func FormatearAlertaBase(alerta map[string]any) string{
    return fmt.Sprintf(`
Nombre: %s
Nivel: %s
Hostname: %s
    `, alerta["_check_name"], alerta["_level"], alerta["host"]) 
}

func FormatearAlertaDisco(alerta map[string]any) string{
    return fmt.Sprintf(`
Nombre: %s
Nivel: %s
Hostname: %s

Particion: %s    -    Uso: %s
`, alerta["_check_name"], alerta["_level"], alerta["host"], alerta["path"], alerta["used_percent"]) 
}

func FormatearAlertaTemperatura(alerta map[string]any) string {
    temperatura, existe := alerta["temp1"]
    if !existe {
        temperatura = alerta["temp2"]
    }
    return fmt.Sprintf(`
Nombre: %s
Nivel: %s
Hostname: %s

Temperatura: %.2f 
`, alerta["_check_name"], alerta["_level"], alerta["host"], temperatura) 
}

var Formateadores = map[string]func(map[string]any) string{
    "Temperatura": FormatearAlertaTemperatura,
    "Disco": FormatearAlertaDisco,
}
