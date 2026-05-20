# Disks

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Disks/de "Handbuch:SPARC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Disks "Handbook:SPARC/Installation/Disks (100% translated)")
- español
- [français](/wiki/Handbook:SPARC/Installation/Disks/fr "Handbook:SPARC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Disks/hu "Handbook:SPARC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Disks/pt-br "Handbook:SPARC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Disks/ru "Handbook:SPARC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Disks/ta "கையேடு:SPARC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Disks/zh-cn "手册：SPARC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Disks/ja "ハンドブック:SPARC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Disks/ko "Handbook:SPARC/Installation/Disks/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[About the installation](/wiki/Handbook:SPARC/Installation/About/es "Handbook:SPARC/Installation/About/es")[Choosing the media](/wiki/Handbook:SPARC/Installation/Media/es "Handbook:SPARC/Installation/Media/es")[Configuring the network](/wiki/Handbook:SPARC/Installation/Networking/es "Handbook:SPARC/Installation/Networking/es")Preparing the disks[The stage file](/wiki/Handbook:SPARC/Installation/Stage/es "Handbook:SPARC/Installation/Stage/es")[Installing base system](/wiki/Handbook:SPARC/Installation/Base/es "Handbook:SPARC/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:SPARC/Installation/Kernel/es "Handbook:SPARC/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:SPARC/Installation/System/es "Handbook:SPARC/Installation/System/es")[Installing tools](/wiki/Handbook:SPARC/Installation/Tools/es "Handbook:SPARC/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:SPARC/Installation/Bootloader/es "Handbook:SPARC/Installation/Bootloader/es")[Finalizing](/wiki/Handbook:SPARC/Installation/Finalizing/es "Handbook:SPARC/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage introduction](/wiki/Handbook:SPARC/Working/Portage/es "Handbook:SPARC/Working/Portage/es")[USE flags](/wiki/Handbook:SPARC/Working/USE/es "Handbook:SPARC/Working/USE/es")[Portage features](/wiki/Handbook:SPARC/Working/Features/es "Handbook:SPARC/Working/Features/es")[Initscript system](/wiki/Handbook:SPARC/Working/Initscripts/es "Handbook:SPARC/Working/Initscripts/es")[Environment variables](/wiki/Handbook:SPARC/Working/EnvVar/es "Handbook:SPARC/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[Files and directories](/wiki/Handbook:SPARC/Portage/Files/es "Handbook:SPARC/Portage/Files/es")[Variables](/wiki/Handbook:SPARC/Portage/Variables/es "Handbook:SPARC/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:SPARC/Portage/Branches/es "Handbook:SPARC/Portage/Branches/es")[Additional tools](/wiki/Handbook:SPARC/Portage/Tools/es "Handbook:SPARC/Portage/Tools/es")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree/es "Handbook:SPARC/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:SPARC/Portage/Advanced/es "Handbook:SPARC/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[Getting started](/wiki/Handbook:SPARC/Networking/Introduction/es "Handbook:SPARC/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:SPARC/Networking/Advanced/es "Handbook:SPARC/Networking/Advanced/es")[Modular networking](/wiki/Handbook:SPARC/Networking/Modular/es "Handbook:SPARC/Networking/Modular/es")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless/es "Handbook:SPARC/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:SPARC/Networking/Extending/es "Handbook:SPARC/Networking/Extending/es")[Dynamic management](/wiki/Handbook:SPARC/Networking/Dynamic/es "Handbook:SPARC/Networking/Dynamic/es")

## Contents

- [1Introducción a los dispositivos de bloque](#Introducci.C3.B3n_a_los_dispositivos_de_bloque)
  - [1.1Dispositivos de bloque](#Dispositivos_de_bloque)
- [2Crear los sistemas de archivos](#Crear_los_sistemas_de_archivos)
  - [2.1Introducción](#Introducci.C3.B3n)
  - [2.2Sistemas de archivos](#Sistemas_de_archivos)
  - [2.3Creación de un sistema de archivos en una partición](#Creaci.C3.B3n_de_un_sistema_de_archivos_en_una_partici.C3.B3n)
    - [2.3.1Sistema de archivos de partición de arranque BIOS heredado](#Sistema_de_archivos_de_partici.C3.B3n_de_arranque_BIOS_heredado)
    - [2.3.2Particiones ext4 pequeñas](#Particiones_ext4_peque.C3.B1as)
  - [2.4Activar la partición de intercambio](#Activar_la_partici.C3.B3n_de_intercambio)
- [3Montar la partición raíz](#Montar_la_partici.C3.B3n_ra.C3.ADz)

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

[Handbook:SPARC/Blocks/Disks/es](/index.php?title=Handbook:SPARC/Blocks/Disks/es&action=edit&redlink=1 "Handbook:SPARC/Blocks/Disks/es (page does not exist)")

## Crear los sistemas de archivos

**Advertencia**

Cuando utilice una unidad SSD o NVMe, es aconsejable comprobar si hay actualizaciones de firmware. Algunos SSD Intel en particular (600p y 6000p) requieren una actualización de firmware para [posible corrupción de datos](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) inducida por patrones de uso de E/S de XFS. El problema está en el nivel de firmware y no en ningún fallo del sistema de archivos XFS. La utilidad smartctl puede ayudar a verificar el modelo del dispositivo y la versión del firmware.

### Introducción

Ahora que ya se han creado las particiones, es el momento de crear en ellas un sistema de archivos. En la próxima sección se describen los distintos sistemas de archivos soportados en Linux. Los lectores que ya sepan los sistemas de archivos que pueden usar deben ir a [Aplicar un sistema de archivos a una partición](/wiki/Handbook:SPARC/Installation/Disks/es#Applying_a_filesystem_to_a_partition "Handbook:SPARC/Installation/Disks/es"). En caso contrario siga leyendo para conocer los sistemas de archivos disponibles...

### Sistemas de archivos

Linux ofrece soporta para varias docenas de sistemas de archivos, aunque muchos de ellos se utilizan para desplegar situaciones específicas. Solo algunos de los sistemas de archivos estables se pueden encontrar en la arquitectura sparc. Se recomienda leer acerca de los sistemas de archivos y el estado de soporte en el que se encuentran antes de seleccionar uno demasiado experimental para particiones importantes. **XFS es el sistema de archivos recomendado para la mayoría de situaciones y plataformas.** Abajo se muestra una lista que no es exhaustiva

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

Por ejemplo, para tener la partición raíz (/dev/sda1) como xfs como se usa en la estructura de partición de ejemplo, los siguientes comandos deberían ser usados:

`root #` `mkfs.xfs /dev/sda1`

#### Sistema de archivos de partición de arranque BIOS heredado

Los sistemas que arrancan a través de BIOS heredado con una etiqueta de disco MBR/DOS pueden usar cualquier formato de sistema de archivos admitido por el gestor de arranque.

Por ejemplo, para formatear con XFS:

`root #` `mkfs.xfs `

#### Particiones ext4 pequeñas

Cuando se utiliza el sistema de archivos ext4 en una partición pequeña (menos de 8 GiB), el sistema de archivos debe crearse con las opciones adecuadas para reservar suficientes inodos. Esto se puede especificar usando la opción `-T small`:

`root #` `mkfs.ext4 -T small /dev/<device>`

Hacerlo cuadruplicará el número de inodos para un sistema de archivos determinado, ya que su "bytes por inodo" se reduce de uno cada 16 kB a uno cada 4 kB.

### Activar la partición de intercambio

mkswap es la orden utilizada para inicializar particiones de intercambio:

`root #` `mkswap /dev/sda2`

**Nota**

Las instalaciones que se iniciaron anteriormente, pero que no finalizaron el proceso de instalación, pueden reanudar la instalación desde este punto del manual. Utilice este enlace como enlace permanente: [Las instalaciones reanudadas comienzan aquí](/index.php?title=Handbook:SPARC/Instalaci%C3%B3n/Discos&action=edit&redlink=1 "Handbook:SPARC/Instalación/Discos (page does not exist)").

Para activar la partición, use swapon:

`root #` `swapon /dev/sda2`

Este paso de 'activación' sólo es necesario porque la partición de intercambio se creó recientemente dentro del entorno live. Una vez que se haya reiniciado el sistema, siempre que la partición de intercambio esté definida correctamente dentro de fstab u otro mecanismo de montaje, el espacio de intercambio se activará automáticamente.

## Montar la partición raíz

Es posible que a ciertos entornos live les falte el punto de montaje sugerido para la partición raíz de Gentoo (/mnt/gentoo), o puntos de montaje para particiones adicionales creadas en la sección de particionamiento:

`root #` `mkdir --parents /mnt/gentoo`

Continúe creando puntos de montaje adicionales necesarios para cualquier partición adicional (personalizada) creada durante los pasos anteriores utilizando el comando mkdir.

Una vez creados los puntos de montaje, es hora de hacer que las particiones sean accesibles mediante el comando mount.

Monte la partición raíz:

`root #` `mount /dev/sda1 /mnt/gentoo`

Continúe montando particiones adicionales (personalizadas) según sea necesario usando el comando mount.

**Nota**

Si necesita que su /tmp/ resida en una partición separada, asegúrese de cambiar los permisos después de montarla:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Lo mismo debe ser aplicado a /var/tmp.

Más adelante en las instrucciones, se montará el sistema de archivos proc (una interfaz virtual con el núcleo), así como otros pseudosistemas de archivos del núcleo. Pero primero se debe extraer [el archivo de stage de Gentoo](/wiki/Handbook:SPARC/Installation/Stage/es "Handbook:SPARC/Installation/Stage/es").

[← Configurando la red](/wiki/Handbook:SPARC/Installation/Networking/es "Previous part") [Inicio](/wiki/Handbook:SPARC/es "Handbook:SPARC/es") [Instalar el archivo de Stage →](/wiki/Handbook:SPARC/Installation/Stage/es "Next part")