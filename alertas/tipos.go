package alertas

import (
    "fmt"
)

type AlertaBase struct {
    CheckName string `json:"_check_name"`
    Level string `json:"_level"`
    Host string `json:"host"`
}


func (alerta AlertaBase) String() string{
    return fmt.Sprintf(`
Nombre: %s
Nivel: %s
Hostname: %s
    `, alerta.CheckName, alerta.Level, alerta.Host) 
}


type AlertaDisco struct {
    CheckName string `json:"_check_name"`
    Level string `json:"_level"`
    Host string `json:"host"`
    Path string `json:"path"`
    UsedPercent string `json:"used_percent"`
}


func (alerta AlertaDisco) String() string{
    return fmt.Sprintf(`
Nombre: %s
Nivel: %s
Hostname: %s

Particion: %s    -    Uso: %s
`, alerta.CheckName, alerta.Level, alerta.Host, alerta.Path, alerta.UsedPercent) 
}

