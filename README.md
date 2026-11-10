# LANShare

Carpeta compartida entre dispositivos de una misma red local (LAN), sin servidores, sin nube y sin límites de peso ni pérdida de calidad.

Cada dispositivo tiene una carpeta en su disco. Todo lo que se agrega, modifica o borra en ella aparece automáticamente en los demás dispositivos conectados a la misma red. Los archivos viajan directo entre dispositivos (P2P) y solo existen como copias locales.

## Motivación

Muchas plataformas para pasar archivos entre dispositivos (mensajería, nubes gratuitas, etc.) imponen límites de tamaño, comprimen imágenes (A vos te hablo, WhatsApp) y videos o requieren subir los datos a un servidor de terceros. LANShare busca resolver el caso más común: **mover archivos entre mis propios dispositivos, en mi casa u oficina**, de forma rápida, privada y sin pérdida de calidad.

## Alcance

### Incluido
- Sincronización de una carpeta compartida entre dispositivos de la **misma red local**.
- Funcionamiento **100% P2P**: no hay servidor central ni almacenamiento en la nube.
- Archivos de **cualquier tipo y tamaño**, transferidos sin modificar.
- Aplicación de escritorio para **Windows, macOS y Linux**.

### Fuera de alcance (por ahora)
- Sincronización a través de internet (fuera de la LAN).
- Aplicaciones móviles (Android / iOS).
- Historial de versiones de archivos.
- Múltiples carpetas compartidas o permisos por usuario.

## Funcionalidades básicas

1. **Descubrimiento automático**: los dispositivos con LANShare se encuentran solos en la red (mDNS), sin configurar IPs.
2. **Emparejamiento**: la primera vez, dos dispositivos se vinculan con un código de confirmación. Solo se sincroniza con dispositivos emparejados; el resto de la red no tiene acceso.
3. **Detección de cambios**: la app vigila la carpeta compartida y detecta archivos nuevos, modificados, renombrados y borrados.
4. **Sincronización**:
   - Archivos nuevos o modificados se copian al resto de los dispositivos.
   - Los borrados se propagan (un dispositivo que estaba apagado no "resucita" archivos borrados).
   - La integridad de cada archivo se verifica mediante hash.
5. **Conflictos**: si un mismo archivo se modifica en dos dispositivos antes de sincronizarse, se conservan ambas versiones (`archivo.conflicto-<dispositivo>.ext`). Nunca se pierden datos.
6. **Transferencias reanudables**: los archivos grandes se transfieren por bloques; si la conexión se corta, se retoma desde donde quedó.
7. **Interfaz mínima**: ver dispositivos conectados, estado de sincronización y emparejar nuevos dispositivos.

## Cómo funciona (visión general)

```
 Dispositivo A                                Dispositivo B
┌──────────────────┐                        ┌──────────────────┐
│ Carpeta          │                        │ Carpeta          │
│   ▲   │ watcher  │      mDNS discovery    │   ▲   │ watcher  │
│   │   ▼          │ ◄────────────────────► │   │   ▼          │
│ Índice local     │                        │ Índice local     │
│ (hash, versión,  │  intercambio de índice │ (hash, versión,  │
│  borrados)       │ ◄────────────────────► │  borrados)       │
│                  │  transferencia de      │                  │
│ Servidor local   │  archivos por bloques  │ Servidor local   │
└──────────────────┘ ◄────────────────────► └──────────────────┘
```

1. Cada dispositivo se anuncia en la LAN y descubre a los demás.
2. Los dispositivos emparejados intercambian sus índices (lista de archivos con su hash y versión).
3. Cada uno compara y pide los archivos que le faltan o que tiene desactualizados.
4. Ante cada cambio local, se notifica a los demás para volver a sincronizar.

## Tecnologías

- **Lenguaje**: Go
- Los detalles de implementación (librerías, protocolo, formato del índice) se definirán más adelante.

## Estado

🚧 En etapa inicial de diseño.

## Licencia

Distribuido bajo la licencia Apache 2.0. Ver [LICENSE](LICENSE).
