// Command lanshare es el daemon de LANShare: vigila la carpeta compartida y la
// sincroniza con los demás dispositivos de la red local.
//
// Por ahora solo arranca y espera a que se lo detenga (Ctrl+C o SIGTERM).
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "lanshare:", err)
		os.Exit(1)
	}
}

// run contiene la lógica de arranque. Está separada de main para poder
// devolver errores en lugar de cortar el programa en cualquier punto.
func run() error {
	// El contexto se cancela cuando llega Ctrl+C (os.Interrupt) o SIGTERM.
	// Todos los componentes van a recibir este contexto para poder terminar
	// de forma ordenada.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	fmt.Println("LANShare iniciado. Presioná Ctrl+C para salir.")
	<-ctx.Done()
	fmt.Println("LANShare detenido.")
	return nil
}
