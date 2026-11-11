// Package controller coordina los componentes de LANShare: decide cuándo
// sincronizar (al iniciar, ante un cambio local o un aviso de otro dispositivo),
// pide los datos a transfer y storage, se los pasa a reconcile y ejecuta las
// decisiones resultantes. También ofrece los servicios que usa el front end.
package controller
