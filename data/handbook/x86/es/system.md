# System

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/System/de "Handbuch:X86/Installation/System (100% translated)")
- [English](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System (100% translated)")
- español
- [français](/wiki/Handbook:X86/Installation/System/fr "Handbook:X86/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/System/it "Handbook:X86/Installation/System (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/System/hu "Handbook:X86/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/System/pl "Handbook:X86/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/System/pt-br "Manual:X86/Instalação/Sistema (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/System/cs "Handbook:X86/Installation/System/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/System/ru "Handbook:X86/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/System/ta "கையேடு:X86/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/System/zh-cn "手册：X86/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/System/ja "ハンドブック:X86/インストール/システムの設定 (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/System/ko "Handbook:X86/Installation/System/ko (100% translated)")

[Manual X86](/wiki/Handbook:X86/es "Handbook:X86/es")[Instalación](/wiki/Handbook:X86/Full/Installation/es "Handbook:X86/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:X86/Installation/About/es "Handbook:X86/Installation/About/es")[Elegir los medios](/wiki/Handbook:X86/Installation/Media/es "Handbook:X86/Installation/Media/es")[Configurar la red](/wiki/Handbook:X86/Installation/Networking/es "Handbook:X86/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:X86/Installation/Disks/es "Handbook:X86/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:X86/Installation/Stage/es "Handbook:X86/Installation/Stage/es")[Instalar el sistema base](/wiki/Handbook:X86/Installation/Base/es "Handbook:X86/Installation/Base/es")[Configurar el núcleo](/wiki/Handbook:X86/Installation/Kernel/es "Handbook:X86/Installation/Kernel/es")Configurar el sistema[Instalar las herramientas](/wiki/Handbook:X86/Installation/Tools/es "Handbook:X86/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:X86/Installation/Bootloader/es "Handbook:X86/Installation/Bootloader/es")[Terminar](/wiki/Handbook:X86/Installation/Finalizing/es "Handbook:X86/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:X86/Full/Working/es "Handbook:X86/Full/Working/es")[Introducción a Portage](/wiki/Handbook:X86/Working/Portage/es "Handbook:X86/Working/Portage/es")[Ajustes USE](/wiki/Handbook:X86/Working/USE/es "Handbook:X86/Working/USE/es")[Características de Portage](/wiki/Handbook:X86/Working/Features/es "Handbook:X86/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:X86/Working/Initscripts/es "Handbook:X86/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:X86/Working/EnvVar/es "Handbook:X86/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:X86/Full/Portage/es "Handbook:X86/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:X86/Portage/Files/es "Handbook:X86/Portage/Files/es")[Variables](/wiki/Handbook:X86/Portage/Variables/es "Handbook:X86/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:X86/Portage/Branches/es "Handbook:X86/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:X86/Portage/Tools/es "Handbook:X86/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:X86/Portage/CustomTree/es "Handbook:X86/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:X86/Portage/Advanced/es "Handbook:X86/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:X86/Full/Networking/es "Handbook:X86/Full/Networking/es")[Comenzar](/wiki/Handbook:X86/Networking/Introduction/es "Handbook:X86/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:X86/Networking/Advanced/es "Handbook:X86/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:X86/Networking/Modular/es "Handbook:X86/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:X86/Networking/Wireless/es "Handbook:X86/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:X86/Networking/Extending/es "Handbook:X86/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:X86/Networking/Dynamic/es "Handbook:X86/Networking/Dynamic/es")

## Contents

- [1Información del sistema de ficheros](#Informaci.C3.B3n_del_sistema_de_ficheros)
  - [1.1Etiquetas e Identificadores únicos (UUIDs) del sistema de archivo](#Etiquetas_e_Identificadores_.C3.BAnicos_.28UUIDs.29_del_sistema_de_archivo)
  - [1.2Etiquetas de particiones y UUIDs](#Etiquetas_de_particiones_y_UUIDs)
  - [1.3Acerca de fstab](#Acerca_de_fstab)
  - [1.4Crear el archivo fstab](#Crear_el_archivo_fstab)
    - [1.4.1Sistemas DOS/BIOS obsoleto](#Sistemas_DOS.2FBIOS_obsoleto)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2Información de la red](#Informaci.C3.B3n_de_la_red)
  - [2.1Hostname](#Hostname)
    - [2.1.1Establecer hostname (OpenRC o systemd)](#Establecer_hostname_.28OpenRC_o_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Red](#Red)
    - [2.2.1DHCP mediante dhcpd (con cualquier sistema de inicio)](#DHCP_mediante_dhcpd_.28con_cualquier_sistema_de_inicio.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.2.1Configurar la red](#Configurar_la_red)
      - [2.2.2.2Inicio automático de red en el arranque](#Inicio_autom.C3.A1tico_de_red_en_el_arranque)
  - [2.3El archivo hosts](#El_archivo_hosts)
  - [2.4Opcional: Hacer que funcione PCMCIA](#Opcional:_Hacer_que_funcione_PCMCIA)
- [3Información del sistema](#Informaci.C3.B3n_del_sistema)
  - [3.1Contraseña del usuario root](#Contrase.C3.B1a_del_usuario_root)
  - [3.2Configuración de inicio y arranque](#Configuraci.C3.B3n_de_inicio_y_arranque)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Información del sistema de ficheros

#### Etiquetas e Identificadores únicos (UUIDs) del sistema de archivo

Tanto MBR (BIOS) como GPT incluyen soporte para etiquetas del _sistema de archivo_ y para UUIDs del _sistema de archivo_. Estos atributos pueden estar definidos en /etc/fstab como alternativas a usar por el comando mount cuando intente encontrar y montar los dispositivos de bloques. Las etiquetas del sistema de archivo y los UUIDs son identificados por el prefijo LABEL y UUID y pueden ser visualizados con el comando blkid.

`root #` `blkid`

**Advertencia**

Si se destruye el sistema de ficheros dentro de una partición, entonces los valores de la etiqueta del sistema de ficheros y del UUID también serán alterados o eliminados.

Para mayor especificidad, se recomienda a los lectores que utilizan tablas de particiones de estilo MBR que utilicen UUID en lugar de etiquetas para especificar volúmenes montables en /etc/fstab.

**Importante**

Los UUIDs de los sistemas de ficheros en un volumen LVM y sus instantáneas LVM snapshots son idénticos, por lo tanto se debe evitar el uso de UUIDs para montar volúmenes LVM.

#### Etiquetas de particiones y UUIDs

Los sistemas con soporte para etiquetas de disco GPT ofrecen opciones 'robustas' adicionales para definir particiones en /etc/fstab. Las etiquetas de partición y los UUID de partición se pueden usar para identificar las particiones individuales del dispositivo de bloque, independientemente del sistema de archivos que se haya elegido para la partición en sí. Las etiquetas de partición y los UUID se identifican mediante los prefijos PARTLABEL y/o PARTUUID y se pueden ver claramente en la terminal ejecutando el comando blkid.

La salida para un sistema EFI **amd64** que utiliza los UUID de especificación de partición detectable puede ser como la siguiente:

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

Si bien no siempre es cierto para las etiquetas de partición, el uso de un UUID para identificar una partición en fstab proporciona una garantía de que el gestor de arranque no se confundirá al buscar un determinado volumen, incluso si se cambia el _sistema de archivos_ o reescrito posteriormente. Usar los archivos de dispositivos de bloque predeterminados más antiguos (/dev/sd\*N) para definir las particiones en fstab tiene su riesgo para los sistemas a los que se agregan o eliminan regularmente dispositivos de bloque SATA.

El nombrado de los dispositivos de bloque depende de una variedad de factores, entre ellos cómo y en qué orden se conectan los discos al sistema. Se podrían incluso mostrar en un orden diferente dependiendo de qué dispositivos detecta el núcleo en primer lugar en los momentos iniciales del proceso de arranque. Dicho esto, a menos que el administrador del sistema juegue constantemente con el orden de los discos, usar los nombres predeterminados de los dispositivos es un método simple y sencillo.

### Acerca de fstab

En Linux, todas las particiones utilizadas por el sistema se deben listar en [/etc/fstab](/wiki//etc/fstab/es "/etc/fstab/es"). Este archivo contiene los puntos de montaje de esas particiones (dónde se encuentran en la estructura del sistema de archivos), cómo se deben montar y con qué opciones especiales (de forma automática o no, si los usuarios pueden montar o no, etc.)

### Crear el archivo fstab

**Nota**

Si el sistema de inicio que se utiliza es systemd, los UUID de la partición cumplen con la Especificación de Partición Detectable como se indica en [Preparación de los discos](/wiki/Handbook:X86/Installation/Disks/es "Handbook:X86/Installation/Disks/es"), y el sistema utiliza UEFI, entonces se puede omitir la creación de un fstab, ya que systemd monta automáticamente las particiones que siguen la especificación.

El archivo en/etc/fstab utiliza un sintaxis similar a una una tabla. Cada línea consta de seis campos, separados por espacios en blanco (espacios, tabuladores o una mezcla de ambos). Cada campo tiene su propio significado:

1. El primer campo muestra el dispositivo de bloques o sistema de archivo remoto que debe ser montado. Varios tipos de identificadores de dispositivo están disponibles para nodos de dispositivo de bloques, incluyendo rutas al archivo especial de dispositivo, etiquetas e identificadores únicos (UUIDs) del sistema de archivo y etiquetas e identificadores únicos (UUIDs) de particiones.
2. El segundo campo muestra el punto de montaje en el que la partición se debe montar.
3. El tercer campo muestra el tipo de sistema de ficheros usado por la partición.
4. El cuarto campo muestra las opciones de montaje usadas por mount cuando se quiere montar la partición. Como cada sistema de ficheros tiene sus propias opciones de montaje, se recomienda a los administradores de sistemas leer la página del manual de mount (man mount) para un listado completo. Las opciones de montaje múltiples se deben separar por comas.
5. El quinto campo lo utiliza dump para determinar si la partición debe ser volcada o no. Esto generalmente se puede dejar a `0` (cero).
6. El sexto campo lo utiliza fsck para determinar el orden en que los sistemas de ficheros se deben revisar en caso de que el sistema no se apagara correctamente. Para el sistema de ficheros raíz se debe definir a `1` mientras que para el resto debería ser `2` (o 0 si no se necesita comprobación del sistema de archivo).

**Importante**

El archivo predeterminado /etc/fstab proporcionado en los archivos stage de Gentoo _no_ es un archivo fstab válido, sino una plantilla que se puede usar para ingresar valores reales.

`root #` `nano /etc/fstab`

#### Sistemas DOS/BIOS obsoleto

Echemos un vistazo a cómo anotar las opciones para la partición /boot. Este es sólo un ejemplo y debe modificarse de acuerdo con las decisiones de partición tomadas anteriormente en la instalación.
En el ejemplo de particionamiento para **x86**, /boot normalmente es la partición /dev/sda1, con xfs como sistema de archivos recomendado. Es necesario comprobarlo durante el arranque, por lo que anotaríamos:

ARCHIVO **`/etc/fstab`** **Ejemplo de línea en /etc/fstab para arranque DOS/BIOS Obsoleto**

```
# Ajuste las diferencias de formato y/o particiones adicionales creadas en el paso "Preparación de los discos".
/dev/sda1   /boot     ext4    defaults        0 2

```

Algunos administradores de sistemas quieren que la partición /boot no se monte automáticamente para mejorar la seguridad de su sistema. Esas personas deberían sustituir `defaults` por `noauto`. Esto significa que esos usuarios deberán montar manualmente esta partición cada vez que quieran usarla.

Añadir las reglas que coinciden con el esquema de esquema de particionamiento decido anteriormente y añadir las reglas para dispositivos tales como lector(es) de CD-ROM, y por supuesto, si se utilizan otras particiones o unidades, añadirlos también.

Abajo se muestra un ejemplo más elaborado de un fichero /etc/fstab:

ARCHIVO **`/etc/fstab`** **Un ejemplo completo de /etc/fstab para un sistema DOS/BIOS obsoleto**

```
# Ajuste las diferencias de formato y/o particiones adicionales creadas en el paso "Preparación de los discos".
/dev/sda1   /boot        ext4    defaults    0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

#### DPS UEFI PARTUUID

Below is an example of an /etc/fstab file for a disk formatted with a GPT disklabel and Discoverable Partition Specification (DPS) UUIDs set for UEFI firmware:

ARCHIVO **`/etc/fstab`** **GPT disklabel DPS PARTUUID fstab example**

```
# Adjust any formatting difference and additional partitions created from the "Preparing the disks" step.
# This example shows a GPT disklabel with Discoverable Partition Specification (DSP) UUID set:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b   /efi        vfat    umask=0077,tz=UTC            0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=44479540-f297-41b2-9af7-d131d5f0458a   /           xfs    defaults,noatime              0 1

```

Cuando se utiliza `auto` en el tercer campo, hace que la orden mount averigüe el sistema de ficheros. Esto se recomienda para los medios extraíbles ya que se pueden crear con uno o más de un sistema de ficheros. La opción `user` en el cuarto campo permite que los usuarios que no sean root puedan montar el CD.

Para mejorar el rendimiento, la mayoría de los usuarios querrían agregar la opción de montaje `noatime`, lo que da como resultado un sistema más rápido ya que los tiempos de acceso no se registran (de todos modos, generalmente no son necesarios). Esto también se recomienda para sistemas con unidades de estado sólido (SSD). Es posible que los usuarios deseen considerar `lazytime` en su lugar.

**Consejo**

Debido a la degradación del rendimiento, no se recomienda definir la opción de montaje `discard` en /etc/fstab. Por lo general, es mejor programar los descartes de bloques periódicamente mediante un programador de tareas como cron o un temporizador (systemd). Consulte [Ejecuciones periódicas de fstrim](/wiki/SSD#Periodic_fstrim_jobs "SSD") para obtener más información.

Compruebe el fichero /etc/fstab y salga para continuar.

## Información de la red

Es importante tener en cuenta que las siguientes secciones se proporcionan para ayudar al lector a configurar rápidamente su sistema para formar parte de una red de área local.

Para los sistemas que ejecutan OpenRC, una referencia más detallada para la configuración de red está disponible en la sección [configuración de red avanzada](/wiki/Handbook:X86/Networking/Introduction/es "Handbook:X86/Networking/Introduction/es"), que se trata casi al final del manual. Es posible que los sistemas con necesidades de red más específicas deban revisarse primero allí y luego regresar aquí para continuar con el resto de la instalación.

Para una configuración de red systemd más específica, consulte la [parte de redes](/wiki/Systemd/es#Network "Systemd/es") del artículo [systemd](/wiki/Systemd/es "Systemd/es").

### Hostname

Una de las elecciones que debe hacer el administrador del sistema es dar nombre a su PC. Esto parece bastante fácil, pero muchos usuarios tienen dificultades para encontrar el nombre apropiado para el host (hostname). Para poder avanzar, sepa que la decisión no es definitiva, se puede cambiar después. En los ejemplos a continuación, se usa el nombre de host _tux_.

#### Establecer hostname (OpenRC o systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

Para configurar el hostname del sistema para un sistema que actualmente está "ejecutando" systemd, se puede utilizar la utilidad hostnamectl. Sin embargo, durante el proceso de instalación, se debe utilizar el comando [systemd-firstboot](/wiki/Handbook:X86/Installation/System/es#Init_and_boot_configuration_systemd "Handbook:X86/Installation/System/es") (consulte más adelante en el manual).

Para establecer el nombre de host como "tux", se ejecutaría:

`root #` `hostnamectl hostname tux`

Vea la ayuda ejecutando hostnamectl --help o man 1 hostnamectl.

### Red

Hay _muchas_ opciones disponibles para configurar interfaces de red. Esta sección cubre solo algunos métodos. Elija el que parezca más adecuado según la configuración necesaria.

#### DHCP mediante dhcpd (con cualquier sistema de inicio)

En la mayoría de redes LAN opera un servidor DHCP. Si es este el caso, entonces se recomienda usar el programa dhcpd para obtener una dirección IP.

Para instalarlo:

`root #` `emerge --ask net-misc/dhcpcd`

Para habilitarlo y luego iniciar el servicio en sistemas OpenRC:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

Para habilitar el servicio en sistemas systemd:

`root #` `systemctl enable dhcpcd`

Con estos pasos completados, la próxima vez que arranque el sistema, dhcpcd debería obtener una dirección IP del servidor DHCP. Consulte el artículo [Dhcpcd](/wiki/Dhcpcd/es "Dhcpcd/es") para obtener más detalles.

#### netifrc (OpenRC)

**Consejo**

Esta es una forma particular de configurar la red usando [Netifrc](/wiki/Netifrc "Netifrc") en OpenRC. Existen otros métodos para configuraciones más simples como [Dhcpcd](/wiki/Dhcpcd/es "Dhcpcd/es").

##### Configurar la red

Durante la instalación de Gentoo Linux, se configuró la red. Sin embargo, eso fue para el entorno vivo de instalación y no para el entorno instalado. Ahora se realiza la configuración de la red para el sistema de Gentoo Linux que se está instalando.

**Nota**

Se puede obtener información más detallada sobre redes, incluyendo temas más avanzados como bonding, bridging, 802.1 Q VLANs o conexiones de red inalámbrica en la sección de [configuración avanzada de la red](/wiki/Handbook:X86/Networking/Introduction/es "Handbook:X86/Networking/Introduction/es").

Toda la información de red se recopila en /etc/conf.d/net. Utiliza una sencilla - pero no tan intuitiva - sintaxis. ¡No tema! Todo se explica a continuación. Hay disponible un ejemplo completamente comentado que abarca muchas configuraciones diferentes en /usr/share/doc/netifrc-\*/net.example.bz2.

En primer lugar se debe instalar [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc):

`root #` `emerge --ask --noreplace net-misc/netifrc`

Por defecto se usa DHCP. Para que funcione, se debe instalar un cliente DHCP. Esto se describe más adelante cuando se describa la instalación de las herramientas del sistema necesarias.

Si la conexión de red se debe configurar con opciones específicas DHCP o porque no se utiliza DHCP en absoluto, entonces abra /etc/conf.d/net:

`root #` `nano /etc/conf.d/net`

Defina tanto config\_eth0 como routes\_eth0 para introducir información de la dirección IP y del enrutamiento:

**Nota**

Esto asume que el interfaz de red se llama eth0. Esto, sin embargo, depende mucho del sistema. Se recomienda asumir que el interfaz se llama igual que cuando se nombra el interfaz arrancando desde los medios de instalación en caso de que éstos sean lo suficientemente recientes. Se puede encontrar más información en la sección [Nombrado de las interfaces de red](/wiki/Handbook:X86/Networking/Advanced/es#Network_interface_naming "Handbook:X86/Networking/Advanced/es").

ARCHIVO **`/etc/conf.d/net`** **Definición de IP estática**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

Para utilizar DHCP, se debe definir config\_eth0:

ARCHIVO **`/etc/conf.d/net`** **Definición DHCP**

```
config_eth0="dhcp"

```

Por favor, lea /usr/share/doc/netifrc-\*/net.example.bz2 para obtener una lista de opciones de configuración adicionales. Asegúrese de leer también la página del manual de DHCP si necesita definir determinadas opciones.

Si el sistema tiene varias interfaces de red, entonces repita los pasos anteriores para config\_eth1, config\_eth2, etc.

Ahora guarde la configuración y salga para continuar.

##### Inicio automático de red en el arranque

Para que los interfaces de red se activen en el arranque, se necesita añadirlos al nivel de ejecución por defecto (default).

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Si el sistema dispone de varios interfaces de red, entonces se necesita crear los archivos net.\* necesarios tal y como se hizo con net.eth0.

Si después de arrancar el sistema se descubre que el nombre de la interfaz de red (que actualmente está documentado como `eth0`) está equivocado, entonces tendremos que seguir los siguientes pasos para corregirlo:

1. Actualizar el archivo /etc/conf.d/net indicando el nombre correcto de la interfaz (como `enp3s0` o `enp5s0` en lugar de `eth0`).
2. Crear un nuevo enlace simbólico (como /etc/init.d/net.enp3s0).
3. Eliminar el enlace simbólico antiguo (rm /etc/init.d/net.eth0).
4. Añadir el nuevo enlace al nivel de ejecución por defecto (default).
5. Eliminar el enalce anterior con rc-update del net.eth0 default.

### El archivo hosts

Un importante próximo paso puede ser informar a este nuevo sistema sobre otros hosts en su entorno de red. Los nombres de host de red se pueden definir en el archivo /etc/hosts. Agregar nombres de host aquí permitirá la resolución de nombres de host a direcciones IP para hosts que no sean resueltos por el servidor de nombres.

`root #` `nano /etc/hosts`

ARCHIVO **`/etc/hosts`** **Rellenar la información de red**

```
# Esto define el presente sistema y debe estar configurado
127.0.0.1     tux.reddecasa tux localhost
::1           tux.reddecasa tux localhost

# Definiciones opcionales de sistemas adicionales en la red
192.168.0.5   juana.reddecasa juana
192.168.0.6   benito.reddecasa benito

```

Guarde y salga del editor para continuar.

### Opcional: Hacer que funcione PCMCIA

Los sistemas **x86** que necesitan compatibilidad con PCMCIA ahora deben instalar el paquete [sys-apps/pcmciautils](https://packages.gentoo.org/packages/sys-apps/pcmciautils). Tenga en cuenta que el soporte para esta antigua tecnología de 16 bits está siendo [eliminado del núcleo principal de Linux](https://lore.kernel.org/all/Y07d7rMvd5++85BJ@owl.dominikbrodowski.net/) [.org/pub/scm/linux/kernel/git/gregkh/char-misc.git/commit/?h=char-misc-next&id=9b12f050c76f090cc6d0aebe0ef76fed79ec3f15 comenzando con dispositivos 'char'](https://git.kernel) en v6.4.0.

`root #` `emerge --ask sys-apps/pcmciautils`

## Información del sistema

### Contraseña del usuario root

Establezca la contraseña del usuario root con la orden passwd.

`root #` `passwd`

Se creará más adelante una cuenta de usuario normal para las operaciones diarias.

### Configuración de inicio y arranque

#### OpenRC

Cuando se usa OpenRC con Gentoo, se utiliza /etc/rc.conf para configurar los servicios, el arranque y parada de un sistema. Abra etc/rc.conf y disfrute de todos los comentarios presentes en el archivo. Revise la configuración y cambie lo que sea necesario.

`root #` `nano /etc/rc.conf`

A continuación, abra /etc/conf.d/keymaps para gestionar la configuración del teclado. Edítelo para configurar y seleccionar el teclado correcto.

`root #` `nano /etc/conf.d/keymaps`

Tenga un cuidado especial la variable keymap. Si el mapa de teclado incorrecto está activado, entonces se producirán resultados extraños cuando tecleemos.

Para terminar, edite /etc/conf.d/hwclock para definir las opciones del reloj. Edítelo conforme a las preferencias personales.

`root #` `nano /etc/conf.d/hwclock`

Si el reloj hardware no está utilizando UTC, entonces es necesario definir `clock="local"` en el archivo, de lo contrario, el sistema podría mostrar un comportamiento de desfase en el reloj.

#### systemd

En primer lugar, se recomienda ejecutar systemd-machine-id-setup y luego systemd-firstboot, que preparará varios componentes del sistema para configurarlos correctamente para el primer arranque en el nuevo entorno de systemd. Al pasar las siguientes opciones, se le pedirá al usuario que establezca una configuración regional, zona horaria, nombre de host, contraseña de root y valores de shell de root. También asignará una ID de máquina aleatoria a la instalación:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

A continuación los usuarios deben ejecutar systemctl para restablecer todos los archivos de unidad instalados a los valores de política preestablecidos:

**Nota**

It is possible that a failure warning will print out after running the following command. This is normal, as Gentoo's LiveCD uses OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

Es posible ejecutar todos los cambios preestablecidos, pero esto puede restablecer cualquier servicio que ya se haya configurado durante el proceso:

`root #` `systemctl preset-all`

Estos dos pasos ayudarán a garantizar una transición fluida desde el entorno live hasta el primer arranque de la instalación.

[← Configurar el núcleo](/wiki/Handbook:X86/Installation/Kernel/es "Previous part") [Inicio](/wiki/Handbook:X86/es "Handbook:X86/es") [Instalar las herramientas →](/wiki/Handbook:X86/Installation/Tools/es "Next part")