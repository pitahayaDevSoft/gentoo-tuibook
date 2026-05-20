# Disks

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Disks/de "Handbuch:X86/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Disks/tr "Handbook:X86/Installation/Disks/tr (0% translated)")
- español
- [français](/wiki/Handbook:X86/Installation/Disks/fr "Handbook:X86/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Disks/it "Handbook:X86/Installation/Disks (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Disks/hu "Handbook:X86/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Disks/pl "Handbook:X86/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Disks/pt-br "Manual:X86/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Disks/ta "கையேடு:X86/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Disks/zh-cn "手册：X86/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Disks/ja "ハンドブック:X86/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Disks/ko "Handbook:X86/Installation/Disks/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86 "Handbook:X86")[Installation](/wiki/Handbook:X86/Full/Installation "Handbook:X86/Full/Installation")[About the installation](/wiki/Handbook:X86/Installation/About/es "Handbook:X86/Installation/About/es")[Choosing the media](/wiki/Handbook:X86/Installation/Media/es "Handbook:X86/Installation/Media/es")[Configuring the network](/wiki/Handbook:X86/Installation/Networking/es "Handbook:X86/Installation/Networking/es")Preparing the disks[The stage file](/wiki/Handbook:X86/Installation/Stage/es "Handbook:X86/Installation/Stage/es")[Installing base system](/wiki/Handbook:X86/Installation/Base/es "Handbook:X86/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:X86/Installation/Kernel/es "Handbook:X86/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:X86/Installation/System/es "Handbook:X86/Installation/System/es")[Installing tools](/wiki/Handbook:X86/Installation/Tools/es "Handbook:X86/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:X86/Installation/Bootloader/es "Handbook:X86/Installation/Bootloader/es")[Finalizing](/wiki/Handbook:X86/Installation/Finalizing/es "Handbook:X86/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:X86/Full/Working "Handbook:X86/Full/Working")[Portage introduction](/wiki/Handbook:X86/Working/Portage/es "Handbook:X86/Working/Portage/es")[USE flags](/wiki/Handbook:X86/Working/USE/es "Handbook:X86/Working/USE/es")[Portage features](/wiki/Handbook:X86/Working/Features/es "Handbook:X86/Working/Features/es")[Initscript system](/wiki/Handbook:X86/Working/Initscripts/es "Handbook:X86/Working/Initscripts/es")[Environment variables](/wiki/Handbook:X86/Working/EnvVar/es "Handbook:X86/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:X86/Full/Portage "Handbook:X86/Full/Portage")[Files and directories](/wiki/Handbook:X86/Portage/Files/es "Handbook:X86/Portage/Files/es")[Variables](/wiki/Handbook:X86/Portage/Variables/es "Handbook:X86/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:X86/Portage/Branches/es "Handbook:X86/Portage/Branches/es")[Additional tools](/wiki/Handbook:X86/Portage/Tools/es "Handbook:X86/Portage/Tools/es")[Custom package repository](/wiki/Handbook:X86/Portage/CustomTree/es "Handbook:X86/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:X86/Portage/Advanced/es "Handbook:X86/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:X86/Full/Networking "Handbook:X86/Full/Networking")[Getting started](/wiki/Handbook:X86/Networking/Introduction/es "Handbook:X86/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:X86/Networking/Advanced/es "Handbook:X86/Networking/Advanced/es")[Modular networking](/wiki/Handbook:X86/Networking/Modular/es "Handbook:X86/Networking/Modular/es")[Wireless](/wiki/Handbook:X86/Networking/Wireless/es "Handbook:X86/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:X86/Networking/Extending/es "Handbook:X86/Networking/Extending/es")[Dynamic management](/wiki/Handbook:X86/Networking/Dynamic/es "Handbook:X86/Networking/Dynamic/es")

## Contents

- [1Introducción a los dispositivos de bloque](#Introducci.C3.B3n_a_los_dispositivos_de_bloque)
  - [1.1Dispositivos de bloque](#Dispositivos_de_bloque)
  - [1.2Tablas de particionamiento](#Tablas_de_particionamiento)
    - [1.2.1Tabla de Particiones GUID (GPT)](#Tabla_de_Particiones_GUID_.28GPT.29)
    - [1.2.2Registro maestro de arranque (MBR) o sector de arranque DOS](#Registro_maestro_de_arranque_.28MBR.29_o_sector_de_arranque_DOS)
  - [1.3Almacenamiento avanzado](#Almacenamiento_avanzado)
  - [1.4Esquema de particionamiento por defecto](#Esquema_de_particionamiento_por_defecto)
- [2Diseñar un esquema de particionamiento](#Dise.C3.B1ar_un_esquema_de_particionamiento)
  - [2.1¿Cuántas particiones y de qué tamaño?](#.C2.BFCu.C3.A1ntas_particiones_y_de_qu.C3.A9_tama.C3.B1o.3F)
  - [2.2¿Qué decir sobre el espacio de intercambio?](#.C2.BFQu.C3.A9_decir_sobre_el_espacio_de_intercambio.3F)
    - [2.2.1¿Qué es la Partición del Sistema EFI (ESP)?](#.C2.BFQu.C3.A9_es_la_Partici.C3.B3n_del_Sistema_EFI_.28ESP.29.3F)
  - [2.3¿Qué es la partición de arranque BIOS?](#.C2.BFQu.C3.A9_es_la_partici.C3.B3n_de_arranque_BIOS.3F)
- [3Partición del disco con GPT para UEFI](#Partici.C3.B3n_del_disco_con_GPT_para_UEFI)
  - [3.1Examinar el esquema de particionamiento actual](#Examinar_el_esquema_de_particionamiento_actual)
  - [3.2Creando una nueva etiqueta de disco / eliminando todas las particiones](#Creando_una_nueva_etiqueta_de_disco_.2F_eliminando_todas_las_particiones)
  - [3.3Creando la partición del sistema EFI (ESP)](#Creando_la_partici.C3.B3n_del_sistema_EFI_.28ESP.29)
  - [3.4Crear la partición de intercambio](#Crear_la_partici.C3.B3n_de_intercambio)
  - [3.5Crear la partición raíz](#Crear_la_partici.C3.B3n_ra.C3.ADz)
  - [3.6Almacenar la tabla de particiones](#Almacenar_la_tabla_de_particiones)
- [4Particionado del disco con MBR para arranque BIOS / heredado](#Particionado_del_disco_con_MBR_para_arranque_BIOS_.2F_heredado)
  - [4.1Ver el diseño de particionado actual](#Ver_el_dise.C3.B1o_de_particionado_actual)
  - [4.2Creando una nueva etiqueta de disco / eliminando todas las particiones](#Creando_una_nueva_etiqueta_de_disco_.2F_eliminando_todas_las_particiones_2)
  - [4.3Creando la partición de arranque](#Creando_la_partici.C3.B3n_de_arranque)
  - [4.4Creando la partición de intercambio](#Creando_la_partici.C3.B3n_de_intercambio)
  - [4.5Creando la partición raíz](#Creando_la_partici.C3.B3n_ra.C3.ADz)
  - [4.6Guardar el diseño del particionado](#Guardar_el_dise.C3.B1o_del_particionado)
- [5Crear los sistemas de archivos](#Crear_los_sistemas_de_archivos)
  - [5.1Introducción](#Introducci.C3.B3n)
  - [5.2Sistemas de archivos](#Sistemas_de_archivos)
  - [5.3Creación de un sistema de archivos en una partición](#Creaci.C3.B3n_de_un_sistema_de_archivos_en_una_partici.C3.B3n)
    - [5.3.1Sistema de archivos de partición de arranque BIOS heredado](#Sistema_de_archivos_de_partici.C3.B3n_de_arranque_BIOS_heredado)
    - [5.3.2Particiones ext4 pequeñas](#Particiones_ext4_peque.C3.B1as)
  - [5.4Activar la partición de intercambio](#Activar_la_partici.C3.B3n_de_intercambio)
- [6Montar la partición raíz](#Montar_la_partici.C3.B3n_ra.C3.ADz)

## Introducción a los dispositivos de bloque

### Dispositivos de bloque

Examinaremos de forma detallada los aspectos de Gentoo Linux así como Linux en general que tengan que ver con discos, incluyendo dispositivos de bloques, particiones y sistemas de archivos de Linux. Una vez familiarizados con las entrañas de los discos y sistemas de archivos, podemos establecer las particiones y sistemas de archivos para la instalación.

Para empezar, veamos los dispositivos de bloque. Los discos SCSI y Serial ATA (SATA) aparecen ambos etiquetados entre los dispositivos gestionados como /dev/sda, /dev/sdb, /dev/sdc, etc. En los equipos mas modernos, los discos de estado sólido NVMe sobre PCI Express tienen nombres de dispositivo como /dev/nvme0n1, /dev/nvme0n2, etc..

La siguiente tabla ayudará a los lectores a saber dónde encontrar un tipo concreto de dispositivo de bloque en su sistema:

Tipo de dispositivoNombre del dispositivo por defectoNotas y consideraciones
IDE, SATA, SAS, SCSI, o USB flash/dev/sdaSe puede encontrar desde aproximadamente 2007 hasta la actualidad, es quizá el mas comunmente usado en Linux. Estos tipos de dispositivos pueden se conectados

via bus [SATA](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") como almacenamiento de bloque. Como ejemplo, la primera partición en el primer dispositivo SATA es nombrada /dev/sda1.

NVM Express (NVMe)/dev/nvme0n1Lo último en tecnología de estado sólido, los discos [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") son conectados al bus PCI Express y tienen la mejor velocidad de transferencia de bloques del mercado. Equipos desde alrededor de 2014 hasya la actualidad pueden tener soporte para el hardware NVMe. La primera partición en el primer dispositivo NVMe en nombrada /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0Los dispositivos [embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard"), tarjetas SD, y otros tipos de tarjetas de memoria pueden ser útiles para el almacenamiento de datos. No obstante, es posible que muchos sistemas no permitan el arranque desde este tipo de dispositivos. Se segiere **no** usar estos dispositivos para instalaciones Linux activas; más bien considere usarlos para transferir archivos, que es la idea de diseño típica. Alternativamente, este tipo de almacenamiento podría resultar útil para copias de seguridad de archivos o instantáneas a corto plazo.
VirtIO/dev/vdaLos dispositivos [virtualizados](/wiki/Virtualization "Virtualization") solo se encuentran dentro de un emulador virtual [QEMU](/wiki/QEMU "QEMU"). El controlador `virtio_blk` proporciona este acceso al espacio de disco asignado en el anfitrión para esta máquina virtual. Dicho esto, es una excelente manera de probar Gentoo dentro de una máquina virtual.

Los dispositivos de bloque mencionados anteriormente representan una interfaz abstracta de disco. Las aplicaciones pueden hacer uso de estas interfaces para interactúar con el disco duro de la máquina sin tener que saber el tipo de unidad que tiene: SATA, SCSI, o cualquier otra. La aplicación puede simplemente dirigirse al almacenamiento en el disco como a una serie de bloques contiguos de acceso aleatorio de 4096-bytes (4K).

### Tablas de particionamiento

Aunque teóricamente es posible utilizar un disco de forma directa, sin particionar, para albergar la instalación Linux (por ejemplo cuando se crea un RAID btrfs), en la práctica, esto casi nunca ocurre. En su lugar, los dispositivos de bloque se dividen en partes más pequeñas y manejables. En los sistemas **x86** éstas se llaman particiones. Actualmente hay dos tecnologías estándar de particionamiento en uso: MBR (a veces también llamado etiqueta de disco DOS) y GPT; estos están vinculados a los dos tipos de procesos de arranque: arranque BIOS heredado y UEFI.

#### Tabla de Particiones GUID (GPT)

La configuración de la _Tabla de particiones GUID (GPT)_ (también llamada etiqueta de disco GPT) utiliza identificadores de 64 bits para las particiones. La ubicación en la que almacena la información de la partición es mucho mayor que los 512 bytes de la tabla de particiones MBR (etiqueta de disco DOS), lo que significa que prácticamente no hay límite en el número de particiones para un disco GPT. Además, el tamaño máximo de partición es mucho mayor (casi 8 ZiB, sí, zebibytes).

Cuando se utiliza UEFI (en lugar de BIOS) como interfaz de software de sistema entre el sistema operativo y el firmware, se requiere el uso de GPT ya que se podrían producir problemas de compatibilidad si se utiliza una etiqueta de disco DOS.

GPT también utiliza sumas de comprobación y redundancia. Realiza sumas de comprobación CRC32 para detectar errores en la cabecera y el las tablas de particiones y dispone de una copia de respaldo GPT al final del disco. Esta copia de respaldo puede utilizarse para recuperarse en caso de que se produzcan daños en la GPT primaria que se almacena al comienzo del disco.

**Importante**

Hay algunas advertencias con respecto a GPT:

- El uso de GPT en un ordenador basado en BIOS funciona, pero el usuario no podrá realizar un arranque dual con un sistema operativo Microsoft Windows, ya que Microsoft Windows se niega a arrancar desde una partición GPT cuando está en modo BIOS.
- Algunos firmware de placa base con errores (antiguos) configurados para arrancar en modo BIOS/CSM/legacy también pueden tener problemas con el arranque desde discos etiquetados con GPT.

#### Registro maestro de arranque (MBR) o sector de arranque DOS

El sector de arranque _[Registro de arranque maestro](https://en.wikipedia.org/wiki/Master_boot_record "wikipedia:Master boot record")_ (también llamado sector de arranque de DOS, etiqueta de disco de DOS y, más recientemente, en oposición a las configuraciones GPT/UEFI, arranque de BIOS heredado) fue el primero en introducido en 1983 con PC DOS 2.x. MBR utiliza identificadores de 32 bits para el sector de inicio y la longitud de las particiones, y admite tres tipos de particiones: primaria, extendida y lógica. Las particiones primarias tienen su información almacenada en el propio registro de arranque maestro: una ubicación muy pequeña (generalmente 512 bytes) al comienzo de un disco. Debido a este pequeño espacio, solo se admiten cuatro particiones primarias (por ejemplo, /dev/sda1 a /dev/sda4).

Para admitir más particiones, una de las particiones primarias en el MBR se puede marcar como una partición "extendida". Esta partición puede contener particiones lógicas adicionales (particiones dentro de una partición).

**Importante**

Aunque todavía son compatibles con la mayoría de los fabricantes de placas base, los sectores de arranque MBR y sus limitaciones de particionamiento asociadas se consideran heredados. A menos que trabaje con hardware anterior a 2010, es mejor particionar un disco con [tabla de particiones GUID](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table"). Los lectores que deban continuar con este tipo de configuración deben ser conscientes de lo siguiente:

- La mayoría de las placas base posteriores a 2010 consideran el uso de sectores de arranque MBR como un modo de arranque heredado (compatible, pero no ideal).
- Debido al uso de identificadores de 32 bits, las tablas de particiones en el MBR no pueden abordar un espacio de almacenamiento que tenga un tamaño superior a 2 TiB.
- A menos que se cree una partición extendida, MBR admite un máximo de cuatro particiones.
- Esta configuración no proporciona un sector de arranque de respaldo, por lo que si algo sobrescribe la tabla de particiones, se perderá toda la información de las particiones.

Dicho esto, es posible que el MBR y el arranque BIOS heredado aún se utilicen en entornos de nube virtualizados como AWS.

Los autores del manual recomiendan utilizar [GPT](#GPT) siempre que sea posible para realizar una instalación de Gentoo.

### Almacenamiento avanzado

El medio de arranque oficial de Gentoo proporciona soporte para [Logical Volume Manager (LVM)](/wiki/LVM/es "LVM/es"). LVM puede combinar "volúmenes físicos", como particiones o discos, en "grupos de volúmenes". Los grupos de volúmenes son más flexibles que las particiones y se pueden utilizar para definir grupos RAID o cachés en SSD rápidos para HD lentos. Aunque el uso no está cubierto en el manual, LVM es totalmente compatible con Gentoo.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/es "LVM/es")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption "Rootfs encryption")

### Esquema de particionamiento por defecto

A lo largo del resto del manual, discutiremos y explicaremos dos casos:

1. Firmware UEFI con disco de tabla de particiones GUID (GPT).
2. Firmware MBR DOS/BIOS obsoleto con un disco de tabla de particiones MBR.

Si bien es posible mezclar y combinar tipos de arranque con cierto firmware de la placa base, la combinación va más allá de la intención del manual. Como se indicó anteriormente, se recomienda encarecidamente que las instalaciones en hardware moderno utilicen el arranque UEFI con un disco con etiqueta GPT.

El siguiente esquema de particiones se utilizará como un diseño de ejemplo simple.

**Importante**

La primera fila de la siguiente tabla contiene información exclusiva para _**ya sea**_ una etiqueta de disco GPT _**o**_ una etiqueta de disco MBR DOS/BIOS heredado. En caso de duda, continúe con GPT, ya que las máquinas **x86** fabricadas después del año 2010 generalmente admiten firmware UEFI y sector de arranque GPT.

Partición
Sistema de archivos
Tamaño
Descripción
/dev/sda1fat32 Sistema de archivos requerido para la partición del sistema EFI, que siempre está asociado con una etiqueta de disco GPT.1 GB
Detalles de la partición del sistema EFI. _Aplicable al firmware del sistema que admite una implementación UEFI. Este suele ser el caso de los sistemas fabricados desde alrededor del año 2010 hasta el presente._ext4 Sistema de archivos recomendado para la partición de inicio de una tabla de particiones MBR, que se utiliza junto con firmware anterior limitado a la etiqueta de disco DOS/BIOS heredado.Detalles de la partición de arranque MBR DOS/BIOS heredado. _Aplicable al firmware de máquina BIOS heredado. Los sistemas de este tipo normalmente se fabricaban <u>antes</u> del año 2010 y, en general, han dejado de producirse._/dev/sda2Intercambio de Linux
Tamaño de RAM \* 2
Detalles de la partición de intercambio.
/dev/sda3xfs
Resto del disco El _**perfil**_ seleccionado, _**particiones**_ adicionales (opcional) y el _**propósito**_ del sistema agrega complejidades para dimensionar apropiadamente los sistemas de archivos raíz, por lo tanto los autores del Manual no pueden ofrecer una sugerencia única para la partición del sistema de archivos raíz.</br></br> Cuando Gentoo es el único sistema operativo que usa el disco, seleccionar el resto del disco es la opción más segura y sugerida.Detalles de la partición raíz.

Si esta información es suficiente, el lector avanzado puede saltar directamente al paricionamiento.

Tanto fdisk como parted son utilidades de partición incluidas en los entornos de imágenes live oficiales de Gentoo. fdisk es bien conocido, estable y maneja discos MBR y GPT. parted fue una de las primeras utilidades de administración de dispositivos de bloques de Linux que admitió particiones GPT. Puede usarse como alternativa a fdisk si el lector lo prefiere; sin embargo, el manual solo proporcionará instrucciones para fdisk, ya que está comúnmente disponible en la mayoría de los entornos Linux.

Antes de pasar a las instrucciones de creación, el primer conjunto de secciones describirán con mas detalle cómo pueden crearse esquemas de particionamiento y mencionan algunos problemas comunes.

## Diseñar un esquema de particionamiento

### ¿Cuántas particiones y de qué tamaño?

El diseño de la distribución de la partición del disco depende en gran medida de lo que se pida al sistema y de los sistemas de archivos aplicados al dispositivo. Si hay muchos usuarios, se recomienda tener /home en una partición separada, lo que aumentará la seguridad y facilitará las copias de seguridad y otros tipos de mantenimiento. Si se está instalando Gentoo para funcionar como un servidor de correo, entonces /var debería ser una partición separada ya que todos los correos se almacenan dentro del directorio /var. Los servidores de juegos pueden tener una partición /opt separada, ya que la mayoría del software del servidor de juegos está instalado allí. El motivo de estas recomendaciones es similar al directorio /home: seguridad, copias de seguridad y mantenimiento.

En la mayoría de situaciones dentro de Gentoo, /usr y /var deberían mantenerse relativamente grandes en lo que a tamaño se refiere. /usr alberga la mayoría de aplicaciones disponibles y las fuentes del núcleo Linux (dentro de /usr/src). Por defecto, /var alberga el repositorio de ebuilds de Gentoo ebuild (localizado en /var/db/repos/gentoo), el cual, dependiendo del sistema de ficheros, normalmente ocupa cerda de 650 MiB de espacio en disco. Esta estimación de espacio _excluye_ los directorios /var/cache/distfiles y /var/cache/binpkgs, los cuales gradualmente se llenarán con ficheros fuente y (opcionalmente) paquetes binarios conforme se van añadiendo al sistema.

Cuántas particiones y como son de grandes depende mayoritariamente de considerar o no las compensaciones y la elección de la mejor opción para cada caso. Tener particiones o volúmenes separados tiene las siguientes ventajas:

- Puede elegir el mejor sistema de archivos para cada partición o volumen.
- El sistema entero no puede quedarse sin espacio si una herramienta fallara y escribiera datos continuamente en una partición o volumen.
- Si es el caso, el tiempo dedicado a las comprobaciones de integridad de los sistemas de archivos se reduce ya que las éstas pueden ser hechas en paralelo (aunque esta mejora se realiza más con varios discos que con varias particiones).
- Se puede mejorar la seguridad montando algunas particiones o volúmenes en modo solo lectura, `nosuid` (los bits setuid son ignorados), `noexec` (los bits de ejecución son ignorados), etc.

Sin embargo, tener múltiples particiones tiene también ciertas desventajas:

- Si no se configura correctamente, el sistema puede tener mucho espacio libre en una partición y poco espacio libre en otra.
- Una partición separada para /usr/ puede requerir que el administrador arranque con un initramfs para montar la partición antes de que comiencen otros guiones de arranque. Dado que la generación y mantenimiento de un initramfs está más allá del alcance de este manual, **recomendamos que los recién llegados no usen una partición separada para /usr/**.
- También hay un límite de 15 particiones para SCSI y SATA a menos que el disco utilice etiquetas GPT.

**Nota**

Las instalaciones en las que se desee utilizar systemd como el sistema para inicio y servicios deben tener el directorio /usr disponible en el momento del inicio, bien como parte del sistema de archivos o montado a través de initramfs.

### ¿Qué decir sobre el espacio de intercambio?

Recomendaciones para el tamaño del espacio de intercambio
Tamaño de la RAM¿Soporte para suspender?¿Soporte para hibernación?
2 GB o menos2 \* RAM3 \* RAM
2 a 8 GBTanto como RAM2 \* RAM
8 a 64 GB8 GB mínimo, 16 máximo1.5 \* RAM
64 GB o mas8 GB mínimo¡Hibernación no recomendada! La hibernación _no_ se recomienda para sistemas con gran cantidad de memoria. Si bien es posible, se debe escribir todo el contenido de la memoria en el disco para que la hibernación sea exitosa. Escribir decenas de gigabytes (¡o más!) en el disco puede llevar una cantidad considerable de tiempo, especialmente cuando se utilizan discos rotatorios. En ese caso es mejor suspender.

No existe un valor perfecto para el tamaño del espacio de intercambio. El propósito del espacio es proporcionar almacenamiento en disco al núcleo cuando la memoria interna dinámica (RAM) está bajo presión. Un espacio de intercambio permite que el núcleo mueva páginas de memoria a las que no es probable que se acceda pronto al disco (intercambio o salida de página), lo que liberará memoria en RAM para la tarea actual. Por supuesto, si las páginas intercambiadas en el disco se necesitan repentinamente, deberán volver a colocarse en la memoria (entrada de página), lo que llevará mucho más tiempo que leer desde la RAM (ya que los discos son muy lentos en comparación con la memoria interna).

Cuando un sistema no va a ejecutar aplicaciones con uso intensivo de memoria o tiene mucha RAM disponible, probablemente no necesite mucho espacio de intercambio. Sin embargo, tenga en cuenta que, en caso de hibernación, el espacio de intercambio se utiliza para almacenar _todo el contenido de la memoria_ (probablemente en sistemas de escritorio y portátiles mas que en sistemas de servidor). Si el sistema requiere soporte para la hibernación, entonces se necesita un espacio de intercambio mayor o igual a la cantidad de memoria necesaria.

Como regla general para cantidades de RAM menores a 4 GB, se recomienda que el tamaño del espacio de intercambio sea el doble de la memoria interna (RAM). Para sistemas con varios discos duros, es aconsejable crear una partición de intercambio en cada disco para que puedan utilizarse para operaciones de lectura/escritura en paralelo. Cuanto más rápido se pueda intercambiar un disco, más rápido se ejecutará el sistema cuando se deba acceder a los datos del espacio de intercambio. Al elegir entre discos rotatorios y de estado sólido, es mejor para el rendimiento colocar el espacio de intercambio en el hardware de estado sólido.

Conviene señalar que se pueden usar _archivos_ de espacio de intercambio como alternativa a las _particiones_ de intercambio; esto es mas útil en sistemas con espacio de disco muy limitado.

#### ¿Qué es la Partición del Sistema EFI (ESP)?

Al instalar Gentoo en un sistema que usa UEFI para arrancar el sistema operativo (en lugar de BIOS), es esencial que se cree una Partición del Sistema EFI (ESP). Las instrucciones a continuación contienen las indicaciones necesarias para manejar correctamente esta operación. **No se requiere la partición del sistema EFI al arrancar en modo BIOS/Heredado.**

La ESP _debe_ ser una variante de FAT (En ocasiones se muestra como _vfat_ en los sistemas Linux). La [UEFI especificación](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) oficial cita que el firmware UEFI reconocerá sistemas de archivos FAT12, 16, o 32, aunque se recomienda FAT32 para la ESP. Después de la partición, formatee el ESP en consecuencia:

`root #` `mkfs.fat -F 32 /dev/sda1`

**Importante**

¡Si el ESP no está formateado con una variante FAT, el firmware UEFI del sistema no encontrará el cargador de arranque (o el núcleo de Linux) y lo más probable es que no pueda arrancar el sistema!

### ¿Qué es la partición de arranque BIOS?

Una _partición de arranque del BIOS_ es una partición muy pequeña (de 1 a 2 MB) en la que los cargadores de arranque como GRUB2 pueden colocar datos adicionales que no caben en el almacenamiento asignado. Solo es necesario cuando un disco está formateado con una etiqueta de disco GPT, pero el firmware del sistema se iniciará a través de GRUB2 en el modo de inicio BIOS/MBR DOS heredado. **_No es necesario_ al arrancar en modo EFI/UEFI, y _tampoco es necesario_ cuando se utiliza una etiqueta de disco MBR/Legacy DOS.** No se habilitará una _partición de arranque del BIOS_ en esta guía.

## Partición del disco con GPT para UEFI

Las siguientes partes explican cómo crear un ejemplo de diseño de partición para un único dispositivo de disco GPT que se ajustará a la especificación UEFI y la especificación de particiones detectables (DPS). DPS es una especificación proporcionada como parte de la especificación del grupo de la API del espacio de usuario (UAPI) de Linux y se recomienda, pero es completamente opcional. Las especificaciones se implementan utilizando la utilidad fdisk, que forma parte del paquete [sys-apps/util-linux](https://packages.gentoo.org/packages/sys-apps/util-linux).

La tabla proporciona valores predeterminados recomendados para una instalación trivial de Gentoo. Se pueden agregar particiones adicionales según las preferencias personales o los objetivos de diseño del sistema.

Ruta del dispositivo (sysfs)
punto de montaje
Sistema de archivos
UUID de DPS (Tipo-UUID)
Descripción
/dev/sda1/efivfat
c12a7328-f81f-11d2-ba4b-00a0c93ec93b
Detalles de la partición del sistema EFI (ESP).
/dev/sda2N/A. Intercambio no se monta en el sistema de archivos como un archivo de dispositivo.0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Detalles de la partición de intercambio.
/dev/sda3/xfs
44479540-f297-41b2-9af7-d131d5f0458a
Detalles de la partición raíz.

### Examinar el esquema de particionamiento actual

fdisk es una popular y potente herramienta que permite dividir el disco en particiones. Arranca fdisk sobre tu unidad de disco (en nuestro ejemplo usamos el dispositivo de disco /dev/sda):

`root #` `fdisk /dev/sda`

Use la tecla `p` para mostrar el esquema de particionamiento actual del disco:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

```

Este disco en particular se ha configurado para albergar dos sistemas de archivos Linux (cada uno con su correspondiente partición listada como "Linux") así como una partición de intercambio (listada como "Linux swap").

### Creando una nueva etiqueta de disco / eliminando todas las particiones

Al presionar la tecla `g`, se eliminarán instantáneamente todas las particiones de disco existentes y se creará una nueva etiqueta de disco GPT:

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

Alternativamente, para mantener una etiqueta de disco GPT existente (consulte el resultado de `p` arriba), considere eliminar las particiones existentes una por una del disco. Presione `d` para eliminar una partición. Por ejemplo, para eliminar una /dev/sda1 existente:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

La partición ha sido marcada para su borrado. Ya no aparecerá al mostrar la lista de particiones ( `p`), pero no será eliminada hasta que guarde los cambios realizados. Esto permite anular la operación si se ha cometido una equivocación - en este caso presione `q` inmediatamente y la tecla `Enter` a continuación y no se eliminarán las particiones.

Presione `p` de forma repetida para ver el listado de particiones y presione `d` junto con el número de la partición para borrarla. Acabará con la tabla de particiones vacía:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

```

Ahora que la tabla de particiones que está en memoria está vacía, estamos preparados para crear las particiones.

### Creando la partición del sistema EFI (ESP)

**Nota**

Es posible utilizar una ESP más pequeña, pero no se recomienda, especialmente porque puede compartirse con otros sistemas operativos.

Primero cree una pequeña partición del sistema EFI, que también se montará como /efi. Escriba `n` para crear una nueva partición, seguido de `1` para seleccionar la primera partición. Cuando se le solicite el primer sector, asegúrese de que comience en 2048 (que puede ser necesario para el cargador de arranque) y presione `Enter`. Cuando se le solicite el último sector, escriba +1G para crear una partición de 1 gibibyte de tamaño:

`Command (m for help):` `n`

```
Partition number (1-128, default 1): 1
First sector (2048-1953525134, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525134, default 1953523711): +1G

Created a new partition 1 of type 'Linux filesystem' and of size 1 GiB.
Partition #1 contains a vfat signature.

<div lang="en" dir="ltr" class="mw-content-ltr">
Do you want to remove the signature? [Y]es/[N]o: Y
The signature will be removed by a write command.

```

Marque la partición como una partición del sistema EFI:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 1
Changed type of partition 'Linux filesystem' to 'EFI System'.

```

### Crear la partición de intercambio

A continuación, para crear la partición de intercambio, presione `n` para crear una nueva partición, luego presione `2` para crear la segunda partición, /dev/sda2. Cuando se le solicite el primer sector, presione `Enter`. Cuando se le solicite el último sector, escriba +4G (o cualquier otro tamaño necesario para el espacio de intercambio) para crear una partición de 4 GiB de tamaño.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (2099200-1953525134, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525134, default 1953523711): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

Una vez que haya hecho esto, presione `t` para definir el tipo de partición, `2` para seleccionar la partición que acaba de crear y entonces "19" para fijar el tipo "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Crear la partición raíz

Finalmente, para crear la partición raíz, presione `n` para crear una nueva partición. Luego presione `3` para crear la tercera partición, /dev/sda3. Cuando se le solicite el primer sector, presione `Enter`. Cuando se le solicite el último sector, presione `Enter` para crear una partición que ocupe el resto del espacio restante en el disco.

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):

Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Nota**

No es necesario configurar el tipo de partición raíz en "Linux root (x86-64)" y el sistema funcionará normalmente si está configurado en el tipo "Linux filesystem". Este tipo de sistema de archivos sólo es necesario en los casos en los que se utiliza un gestor de arranque que lo admita (es decir, systemd-boot) y no se desea un archivo fstab.

Después de crear la partición raíz, presione `t` para configurar el tipo de partición, `3` para seleccionar la partición que acaba de crear y luego escriba _23_ para configurar el tipo de partición en "Linux Root (x86-64)".

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23

Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

Después de completar estos pasos, presionar `p` debería mostrar una tabla de particiones similar a la siguiente:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

```

### Almacenar la tabla de particiones

Presione `w` para guardar el diseño de la partición y salir de la utilidad fdisk:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Ahora que las particiones están disponibles, el siguiente paso de instalación es rellenarlas con sistemas de archivos.

## Particionado del disco con MBR para arranque BIOS / heredado

La siguiente tabla proporciona un diseño de partición recomendado para una instalación trivial de arranque de MBR DOS/BIOS heredado. Se pueden agregar particiones adicionales según las preferencias personales o los objetivos de diseño del sistema.

Ruta del dispositivo (sysfs)
punto de montaje
Sistema de archivos
UUID de DPS (PARTUUID)
Descripción
/dev/sda1/bootext4
N/A
Detalles de la partición de arranque MBR DOS/BIOS heredado.
/dev/sda2N/A. Intercambio no se monta en el sistema de archivos como un archivo de dispositivo.0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Detalles de la partición de intercambio.
/dev/sda3/xfs
44479540-f297-41b2-9af7-d131d5f0458a
Detalles de la partición raíz.

Cambie el diseño del particionado según sus preferencias personales.

### Ver el diseño de particionado actual

Lance fdisk sobre el disco (en nuestro ejemplo, usamos /dev/sda):

`root #` `fdisk /dev/sda`

Utilice la tecla `p` para mostrar la configuración de particionado actual del disco:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

Este disco en particular estaba configurado hasta ahora para albergar dos sistemas de archivos Linux (cada uno en su partición correspondiente listada como "Linux") así como una partición de intercambio (listada como "Linux swap"), usando una tabla GPT.

### Creando una nueva etiqueta de disco / eliminando todas las particiones

Al presionar `o` se eliminarán instantáneamente todas las particiones de disco existentes y se creará una nueva etiqueta de disco MBR (también llamada etiqueta de disco DOS):

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternativamente, para mantener una etiqueta de disco de DOS existente (consulte el resultado de `p` arriba), considere eliminar las particiones existentes en el disco una por una. Presione `d` para eliminar una partición. Por ejemplo, para eliminar una /dev/sda1 existente:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

La partición ahora está preparada para su eliminación. Ya no aparecerá al imprimir la lista de particiones ( `p`, pero no se borrará hasta que se hayan guardado los cambios. Esto permite a los usuarios cancelar la operación si se cometió un error, en ese caso, escriba `q` inmediatamente y presione `Enter` y la partición no se eliminará.

Presione repetidamente `p` para imprimir una lista de particiones y luego presione `d` y el número de la partición para eliminarla. Finalmente, la tabla de particiones estará vacía:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

El disco ahora está preparado para crear nuevas particiones.

### Creando la partición de arranque

Primero, cree una pequeña partición que se montará como /boot. Presione `n` para crear una nueva partición, seguido de `p` para una partición _primaria_ y `1` para seleccionar la primera partición primaria. Cuando se le solicite el primer sector, asegúrese de que comience desde 2048 (que puede ser necesario para el cargador de arranque) y presione `Entrar`. Cuando se le solicite el último sector, escriba +1G para crear una partición de 1 GB de tamaño:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G

Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Marque la partición como de arranque presionando la tecla `a` y presionando `Enter`:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Nota: si hay más de una partición disponible en el disco, entonces se deberá seleccionar la partición que se marcará como de arranque.

### Creando la partición de intercambio

A continuación, para crear la partición de intercambio, presione `n` para crear una nueva partición, luego `p`, luego escriba `2` para crear la segunda partición primaria, /dev/sda2. Cuando se le solicite el primer sector, presione `Enter`. Cuando se le solicite el último sector, escriba +4G (o cualquier otro tamaño necesario para el espacio de intercambio) para crear una partición de 4GB de tamaño.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

Una vez hecho lo anterior, presione `t` para establecer el tipo de partición, `2` para seleccionar la partición que acaba de crear y luego escriba _82_ para establecer el tipo de partición en "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Creando la partición raíz

Finalmente, para crear la partición raíz, presione `n` para crear una nueva partición. Luego presione `p` y `3` para crear la tercera partición primaria, /dev/sda3. Cuando se le solicite el primer sector, presione `Enter`. Cuando se le solicite el último sector, presione `Enter` para crear una partición que ocupe todo el espacio restante en el disco:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):

Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

Después de completar estos pasos, presionar `p` debería mostrar una tabla de particiones similar a esta:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### Guardar el diseño del particionado

Presione `w` para escribir el diseño de la partición y salir de fdisk:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Ahora es el momento de poner sistemas de archivos en las particiones.

## Crear los sistemas de archivos

**Advertencia**

Cuando utilice una unidad SSD o NVMe, es aconsejable comprobar si hay actualizaciones de firmware. Algunos SSD Intel en particular (600p y 6000p) requieren una actualización de firmware para [posible corrupción de datos](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) inducida por patrones de uso de E/S de XFS. El problema está en el nivel de firmware y no en ningún fallo del sistema de archivos XFS. La utilidad smartctl puede ayudar a verificar el modelo del dispositivo y la versión del firmware.

### Introducción

Ahora que ya se han creado las particiones, es el momento de crear en ellas un sistema de archivos. En la próxima sección se describen los distintos sistemas de archivos soportados en Linux. Los lectores que ya sepan los sistemas de archivos que pueden usar deben ir a [Aplicar un sistema de archivos a una partición](/wiki/Handbook:X86/Installation/Disks/es#Applying_a_filesystem_to_a_partition "Handbook:X86/Installation/Disks/es"). En caso contrario siga leyendo para conocer los sistemas de archivos disponibles...

### Sistemas de archivos

Linux ofrece soporta para varias docenas de sistemas de archivos, aunque muchos de ellos se utilizan para desplegar situaciones específicas. Solo algunos de los sistemas de archivos estables se pueden encontrar en la arquitectura x86. Se recomienda leer acerca de los sistemas de archivos y el estado de soporte en el que se encuentran antes de seleccionar uno demasiado experimental para particiones importantes. **XFS es el sistema de archivos recomendado para la mayoría de situaciones y plataformas.** Abajo se muestra una lista que no es exhaustiva

[XFS](/wiki/XFS "XFS")Sistema de archivos con registro de metadatos que viene con un sólido conjunto de funciones y está optimizado para la escalabilidad. Se ha actualizado continuamente para incluir características modernas. El único inconveniente es que las particiones XFS aún no se pueden reducir, aunque se está trabajando en ello. Es notable que XFS admite enlaces de referencia y copia en escritura (CoW), lo cual es particularmente útil en sistemas Gentoo debido a la cantidad de compilaciones que completan los usuarios. XFS es el sistema de archivos moderno y multiuso recomendado para todas las plataformas. Requiere una partición de al menos 300 MB.[ext4](/wiki/Ext4 "Ext4")Ext4 es un sistema de archivos confiable y multiuso para todas las plataformas, aunque carece de características modernas como reflinks.[VFAT](/wiki/VFAT "VFAT")También conocido como FAT32, es compatible con Linux pero no admite la configuración de permisos estándar de UNIX. Se utiliza principalmente para interoperabilidad/intercambio con otros sistemas operativos (Microsoft Windows o macOS de Apple), pero también es necesario para algunos firmware del gestor de arranque del sistema (como UEFI). Los usuarios de sistemas UEFI necesitarán una [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition") formateada con VFAT para poder arrancar.[btrfs](/wiki/Btrfs/es "Btrfs/es")Sistema de archivos de nueva generación. Proporciona funciones avanzadas como instantáneas, autorreparación mediante sumas de comprobación, compresión transparente, subvolúmenes y RAID integrado. No se garantiza que los núcleos anteriores a 5.4.y sean seguros de usar con btrfs en producción porque las correcciones para problemas graves solo están presentes en las versiones más recientes de las ramas del núcleo LTS. RAID 5/6 y grupos de cuotas no son seguros en todas las versiones de btrfs.[F2FS](/wiki/F2FS "F2FS")El sistema de archivos compatible con Flash fue creado originalmente por Samsung para su uso con memoria flash NAND. Es una opción decente al instalar Gentoo en tarjetas microSD, unidades USB u otros dispositivos de almacenamiento basados en flash.[NTFS](/wiki/NTFS "NTFS")Este sistema de archivos de "Nueva Tecnología" es el sistema de archivos insignia de Microsoft Windows desde Windows NT 3.1. De manera similar a VFAT, no almacena la configuración de permisos de UNIX ni los atributos extendidos necesarios para que BSD o Linux funcionen correctamente, por lo que no debe usarse como sistema de archivos raíz en la mayoría de los casos. Debe usarse "sólo" para interoperabilidad o intercambio de datos con sistemas Microsoft Windows (tenga en cuenta el énfasis en "sólo").[ZFS](/wiki/ZFS "ZFS")**Importante:** Los grupos de ZFS solo se pueden crear con las ISOs de admincd y LiveGUI. Para más información, consulte [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Sistema de archivos de nueva generación creado por Matthew Ahrens y Jeff Bonwick. Se diseñó en base a varias ideas clave: la administración del almacenamiento debe ser sencilla, la redundancia debe ser gestionada por el sistema de archivos, los sistemas de archivos nunca deben desconectarse para su reparación, las simulaciones automatizadas de los peores escenarios antes de enviar el código son importantes y la integridad de los datos es primordial.[bcachefs](/wiki/Bcachefs "Bcachefs")**Importante:** bcachefs aún está marcado como experimental en el núcleo, por lo que es importante realizar copias de seguridad periódicas de los datos importantes en un medio que no sea bcachefs.bcachefs es un sistema de archivos B-tree con todas las funciones basado en [bcache](/wiki/Bcache "Bcache"). Incluye funciones como Copia si Escritura (CoW), compresión y cifrado. bcachefs es comparable a Btrfs y ZFS. Una característica destacada es la compatibilidad nativa con almacenamiento por niveles, que permite usar una o más unidades de disco rápidas (como SSD basadas en flash o discos NVMe) como caché para una o más unidades de disco más lentas en un grupo, a la vez que gestiona de forma transparente los archivos activos y pasivos según la actividad.

Puede encontrar información más extensa sobre los sistemas de archivos en el [Filesystem Article](/wiki/Filesystem/es "Filesystem/es") mantenido por la comunidad.

### Creación de un sistema de archivos en una partición

**Nota**

Asegúrese de instalar el paquete de utilidades de espacio de usuario correspondiente para el sistema de archivos elegido antes de reiniciar. Habrá un recordatorio para hacerlo cerca del final del proceso de instalación.

Para crear un sistema de archivos en una partición o volumen, existen utilidades de espacio de usuario disponibles para todos los sistemas de archivos. Hacer clic en el nombre del sistema de archivos de la tabla de abajo para obtener información de cada sistema de archivos:

Filesystem
Comando de creación
¿Está en el entorno live?
Paquete
[XFS](/wiki/XFS "XFS")mkfs.xfs Sí
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4 "Ext4")mkfs.ext4 Sí
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat Sí
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/es "Btrfs/es")mkfs.btrfs Sí
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")mkfs.f2fs Sí
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Sí
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... No
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)[bcachefs](/wiki/Bcachefs "Bcachefs")mkfs.bcachefs Sí (Solo AMD64 a partir del 12-04-2025)
[sys-fs/bcachefs-tools](https://packages.gentoo.org/packages/sys-fs/bcachefs-tools)

\|}

**Importante**

El manual recomienda nuevas particiones como parte del proceso de instalación, pero es importante tener en cuenta que ejecutar cualquier comando mkfs borrará todos los datos contenidos en la partición. Cuando sea necesario, asegúrese de realizar una copia de seguridad adecuada de todos los datos que existan antes de crear un nuevo sistema de archivos.

Por ejemplo, para tener la partición raíz (/dev/sda3) como xfs como se usa en la estructura de partición de ejemplo, los siguientes comandos deberían ser usados:

`root #` `mkfs.xfs /dev/sda3`

#### Sistema de archivos de partición de arranque BIOS heredado

Los sistemas que arrancan a través de BIOS heredado con una etiqueta de disco MBR/DOS pueden usar cualquier formato de sistema de archivos admitido por el gestor de arranque.

Por ejemplo, para formatear con XFS:

`root #` `mkfs.xfs /dev/sda1`

#### Particiones ext4 pequeñas

Cuando se utiliza el sistema de archivos ext4 en una partición pequeña (menos de 8 GiB), el sistema de archivos debe crearse con las opciones adecuadas para reservar suficientes inodos. Esto se puede especificar usando la opción `-T small`:

`root #` `mkfs.ext4 -T small /dev/<device>`

Hacerlo cuadruplicará el número de inodos para un sistema de archivos determinado, ya que su "bytes por inodo" se reduce de uno cada 16 kB a uno cada 4 kB.

### Activar la partición de intercambio

mkswap es la orden utilizada para inicializar particiones de intercambio:

`root #` `mkswap /dev/sda2`

**Nota**

Las instalaciones que se iniciaron anteriormente, pero que no finalizaron el proceso de instalación, pueden reanudar la instalación desde este punto del manual. Utilice este enlace como enlace permanente: [Las instalaciones reanudadas comienzan aquí](/index.php?title=Handbook:X86/Instalaci%C3%B3n/Discos&action=edit&redlink=1 "Handbook:X86/Instalación/Discos (page does not exist)").

Para activar la partición, use swapon:

`root #` `swapon /dev/sda2`

Este paso de 'activación' sólo es necesario porque la partición de intercambio se creó recientemente dentro del entorno live. Una vez que se haya reiniciado el sistema, siempre que la partición de intercambio esté definida correctamente dentro de fstab u otro mecanismo de montaje, el espacio de intercambio se activará automáticamente.

## Montar la partición raíz

Es posible que a ciertos entornos live les falte el punto de montaje sugerido para la partición raíz de Gentoo (/mnt/gentoo), o puntos de montaje para particiones adicionales creadas en la sección de particionamiento:

`root #` `mkdir --parents /mnt/gentoo`

Continúe creando puntos de montaje adicionales necesarios para cualquier partición adicional (personalizada) creada durante los pasos anteriores utilizando el comando mkdir.

Una vez creados los puntos de montaje, es hora de hacer que las particiones sean accesibles mediante el comando mount.

Monte la partición raíz:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continúe montando particiones adicionales (personalizadas) según sea necesario usando el comando mount.

**Nota**

Si necesita que su /tmp/ resida en una partición separada, asegúrese de cambiar los permisos después de montarla:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Lo mismo debe ser aplicado a /var/tmp.

Más adelante en las instrucciones, se montará el sistema de archivos proc (una interfaz virtual con el núcleo), así como otros pseudosistemas de archivos del núcleo. Pero primero se debe extraer [el archivo de stage de Gentoo](/wiki/Handbook:X86/Installation/Stage/es "Handbook:X86/Installation/Stage/es").

[← Configurando la red](/wiki/Handbook:X86/Installation/Networking/es "Previous part") [Inicio](/wiki/Handbook:X86/es "Handbook:X86/es") [Instalar el archivo de Stage →](/wiki/Handbook:X86/Installation/Stage/es "Next part")