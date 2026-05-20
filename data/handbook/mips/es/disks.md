# Disks

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Disks/de "Handbuch:MIPS/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks (100% translated)")
- español
- [français](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Disks/pt-br "Manual:MIPS/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Disks/ru "Handbook:MIPS/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Disks/ta "கையேடு:MIPS/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Disks/zh-cn "手册：MIPS/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Disks/ja "ハンドブック:MIPS/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Disks/ko "Handbook:MIPS/Installation/Disks/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS "Handbook:MIPS")[Installation](/wiki/Handbook:MIPS/Full/Installation "Handbook:MIPS/Full/Installation")[About the installation](/wiki/Handbook:MIPS/Installation/About/es "Handbook:MIPS/Installation/About/es")[Choosing the media](/wiki/Handbook:MIPS/Installation/Media/es "Handbook:MIPS/Installation/Media/es")[Configuring the network](/wiki/Handbook:MIPS/Installation/Networking/es "Handbook:MIPS/Installation/Networking/es")Preparing the disks[The stage file](/wiki/Handbook:MIPS/Installation/Stage/es "Handbook:MIPS/Installation/Stage/es")[Installing base system](/wiki/Handbook:MIPS/Installation/Base/es "Handbook:MIPS/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:MIPS/Installation/Kernel/es "Handbook:MIPS/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:MIPS/Installation/System/es "Handbook:MIPS/Installation/System/es")[Installing tools](/wiki/Handbook:MIPS/Installation/Tools/es "Handbook:MIPS/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:MIPS/Installation/Bootloader/es "Handbook:MIPS/Installation/Bootloader/es")[Finalizing](/wiki/Handbook:MIPS/Installation/Finalizing/es "Handbook:MIPS/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:MIPS/Full/Working "Handbook:MIPS/Full/Working")[Portage introduction](/wiki/Handbook:MIPS/Working/Portage/es "Handbook:MIPS/Working/Portage/es")[USE flags](/wiki/Handbook:MIPS/Working/USE/es "Handbook:MIPS/Working/USE/es")[Portage features](/wiki/Handbook:MIPS/Working/Features/es "Handbook:MIPS/Working/Features/es")[Initscript system](/wiki/Handbook:MIPS/Working/Initscripts/es "Handbook:MIPS/Working/Initscripts/es")[Environment variables](/wiki/Handbook:MIPS/Working/EnvVar/es "Handbook:MIPS/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:MIPS/Full/Portage "Handbook:MIPS/Full/Portage")[Files and directories](/wiki/Handbook:MIPS/Portage/Files/es "Handbook:MIPS/Portage/Files/es")[Variables](/wiki/Handbook:MIPS/Portage/Variables/es "Handbook:MIPS/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:MIPS/Portage/Branches/es "Handbook:MIPS/Portage/Branches/es")[Additional tools](/wiki/Handbook:MIPS/Portage/Tools/es "Handbook:MIPS/Portage/Tools/es")[Custom package repository](/wiki/Handbook:MIPS/Portage/CustomTree/es "Handbook:MIPS/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:MIPS/Portage/Advanced/es "Handbook:MIPS/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:MIPS/Full/Networking "Handbook:MIPS/Full/Networking")[Getting started](/wiki/Handbook:MIPS/Networking/Introduction/es "Handbook:MIPS/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:MIPS/Networking/Advanced/es "Handbook:MIPS/Networking/Advanced/es")[Modular networking](/wiki/Handbook:MIPS/Networking/Modular/es "Handbook:MIPS/Networking/Modular/es")[Wireless](/wiki/Handbook:MIPS/Networking/Wireless/es "Handbook:MIPS/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:MIPS/Networking/Extending/es "Handbook:MIPS/Networking/Extending/es")[Dynamic management](/wiki/Handbook:MIPS/Networking/Dynamic/es "Handbook:MIPS/Networking/Dynamic/es")

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

[Handbook:MIPS/Blocks/Disks/es](/index.php?title=Handbook:MIPS/Blocks/Disks/es&action=edit&redlink=1 "Handbook:MIPS/Blocks/Disks/es (page does not exist)")

## Crear los sistemas de archivos

**Advertencia**

Cuando utilice una unidad SSD o NVMe, es aconsejable comprobar si hay actualizaciones de firmware. Algunos SSD Intel en particular (600p y 6000p) requieren una actualización de firmware para [posible corrupción de datos](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) inducida por patrones de uso de E/S de XFS. El problema está en el nivel de firmware y no en ningún fallo del sistema de archivos XFS. La utilidad smartctl puede ayudar a verificar el modelo del dispositivo y la versión del firmware.

### Introducción

Ahora que ya se han creado las particiones, es el momento de crear en ellas un sistema de archivos. En la próxima sección se describen los distintos sistemas de archivos soportados en Linux. Los lectores que ya sepan los sistemas de archivos que pueden usar deben ir a [Aplicar un sistema de archivos a una partición](/wiki/Handbook:MIPS/Installation/Disks/es#Applying_a_filesystem_to_a_partition "Handbook:MIPS/Installation/Disks/es"). En caso contrario siga leyendo para conocer los sistemas de archivos disponibles...

### Sistemas de archivos

Linux ofrece soporta para varias docenas de sistemas de archivos, aunque muchos de ellos se utilizan para desplegar situaciones específicas. Solo algunos de los sistemas de archivos estables se pueden encontrar en la arquitectura mips. Se recomienda leer acerca de los sistemas de archivos y el estado de soporte en el que se encuentran antes de seleccionar uno demasiado experimental para particiones importantes. **XFS es el sistema de archivos recomendado para la mayoría de situaciones y plataformas.** Abajo se muestra una lista que no es exhaustiva

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

Por ejemplo, para tener la partición raíz (/dev/sda5) como xfs como se usa en la estructura de partición de ejemplo, los siguientes comandos deberían ser usados:

`root #` `mkfs.xfs /dev/sda5`

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

`root #` `mkswap /dev/sda10`

**Nota**

Las instalaciones que se iniciaron anteriormente, pero que no finalizaron el proceso de instalación, pueden reanudar la instalación desde este punto del manual. Utilice este enlace como enlace permanente: [Las instalaciones reanudadas comienzan aquí](/index.php?title=Handbook:MIPS/Instalaci%C3%B3n/Discos&action=edit&redlink=1 "Handbook:MIPS/Instalación/Discos (page does not exist)").

Para activar la partición, use swapon:

`root #` `swapon /dev/sda10`

Este paso de 'activación' sólo es necesario porque la partición de intercambio se creó recientemente dentro del entorno live. Una vez que se haya reiniciado el sistema, siempre que la partición de intercambio esté definida correctamente dentro de fstab u otro mecanismo de montaje, el espacio de intercambio se activará automáticamente.

## Montar la partición raíz

Es posible que a ciertos entornos live les falte el punto de montaje sugerido para la partición raíz de Gentoo (/mnt/gentoo), o puntos de montaje para particiones adicionales creadas en la sección de particionamiento:

`root #` `mkdir --parents /mnt/gentoo`

Continúe creando puntos de montaje adicionales necesarios para cualquier partición adicional (personalizada) creada durante los pasos anteriores utilizando el comando mkdir.

Una vez creados los puntos de montaje, es hora de hacer que las particiones sean accesibles mediante el comando mount.

Monte la partición raíz:

`root #` `mount /dev/sda5 /mnt/gentoo`

Continúe montando particiones adicionales (personalizadas) según sea necesario usando el comando mount.

**Nota**

Si necesita que su /tmp/ resida en una partición separada, asegúrese de cambiar los permisos después de montarla:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Lo mismo debe ser aplicado a /var/tmp.

Más adelante en las instrucciones, se montará el sistema de archivos proc (una interfaz virtual con el núcleo), así como otros pseudosistemas de archivos del núcleo. Pero primero se debe extraer [el archivo de stage de Gentoo](/wiki/Handbook:MIPS/Installation/Stage/es "Handbook:MIPS/Installation/Stage/es").

[← Configurando la red](/wiki/Handbook:MIPS/Installation/Networking/es "Previous part") [Inicio](/wiki/Handbook:MIPS/es "Handbook:MIPS/es") [Instalar el archivo de Stage →](/wiki/Handbook:MIPS/Installation/Stage/es "Next part")