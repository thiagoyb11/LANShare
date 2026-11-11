// Package reconcile compara el log local con el de otro dispositivo y decide,
// archivo por archivo, qué hacer: descargar, borrar, marcar un conflicto o nada.
//
// No hace operaciones de red ni de disco: recibe datos y devuelve decisiones,
// para poder testearlo de forma aislada. Se llama reconcile y no sync para no
// chocar con el paquete sync de la biblioteca estándar.
package reconcile
