# AGENTS.md

Guía para agentes de IA (y personas) que trabajen en este repositorio.

## Proyecto

LANShare es una carpeta compartida que se sincroniza entre dispositivos de una misma red local, de forma P2P, sin servidores ni nube. El alcance y las funcionalidades están en [README.md](README.md). Leelo antes de proponer cambios.

**Restricciones de diseño que no se negocian:**
- **Solo LAN.** No agregar sincronización por internet, relays, servidores de descubrimiento externos ni dependencias de servicios en la nube.
- **Sin servidor central.** Todos los dispositivos corren el mismo binario y son pares.
- **Nunca perder datos.** Ante una duda (conflicto, error de transferencia, hash que no coincide), conservar las copias y no sobrescribir ni borrar.
- **Archivos sin modificar.** Nada de compresión, recodificación ni cambios de metadatos.
- **Emparejamiento diferido.** Por decisión del dueño del repo, la primera versión no implementa emparejamiento ni control de acceso: se asume una red doméstica de confianza. No dar por hecho que existe autenticación entre pares.

## Estado actual

Existe el esqueleto del módulo: `cmd/lanshare` y los paquetes de `internal/` (ver README.md), sin lógica todavía. Las demás decisiones de implementación (librerías, protocolo, formato del índice, interfaces entre paquetes) **todavía no están tomadas**. No las inventes ni las des por hechas: proponelas y esperá confirmación antes de implementarlas.

`internal/reconcile` no debe hacer I/O (red ni disco): recibe logs y devuelve decisiones. `controller` es quien obtiene los datos de `transfer` y `storage` y se los pasa.

## Stack

- **Lenguaje:** Go. Preferir la biblioteca estándar; antes de sumar una dependencia externa, justificar por qué hace falta.
- **Plataformas objetivo:** Windows, macOS y Linux (escritorio). El código tiene que ser portable: usar `path/filepath` para rutas y no asumir separadores, mayúsculas/minúsculas ni permisos de un sistema operativo en particular.

## Convenciones de código

- Formatear con `gofmt` (o `goimports`) y pasar `go vet` sin advertencias.
- Seguir las convenciones idiomáticas de Go: errores como valores, envolverlos con contexto (`fmt.Errorf("...: %w", err)`) y no usar `panic` para errores esperables.
- Toda goroutine tiene que poder terminar: usar `context.Context` para cancelación y evitar fugas.
- Proteger el estado compartido (mutex o canales) y correr los tests con `-race`.
- Identificadores y nombres de paquetes en inglés; comentarios y documentación en español.

## Comandos

Una vez que exista el módulo de Go:

```sh
go build ./...
go test -race ./...
go vet ./...
gofmt -l .
```

Antes de dar un cambio por terminado, los cuatro tienen que pasar limpios.

## Tests

- Toda lógica de sincronización (comparación de índices, conflictos, borrados) necesita tests unitarios. Es el corazón del proyecto y donde se pierden datos si hay bugs.
- Los tests no deben depender de la red real ni de mDNS. Usar directorios temporales (`t.TempDir()`) e interfaces que permitan simular pares.

## Git

- Mensajes de commit en español, en infinitivo y concisos (por ejemplo: `Agregar detección de cambios en la carpeta`).
- No hacer commit ni push sin que lo pida el dueño del repo.
- No commitear binarios, carpetas de prueba ni archivos de configuración local.

## Documentación

Si un cambio modifica el alcance, las funcionalidades o el funcionamiento general, actualizar el README.md en el mismo cambio.
