# Media

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Media/de "Handbook:MIPS/Installation/Media/de (57% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Media "Handbook:MIPS/Installation/Media (100% translated)")
- español
- [français](/wiki/Handbook:MIPS/Installation/Media/fr "Handbook:MIPS/Installation/Media/fr (57% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Media/it "Handbook:MIPS/Installation/Media/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Media/hu "Handbook:MIPS/Installation/Media/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Media/pl "Handbook:MIPS/Installation/Media/pl (14% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Media/pt-br "Handbook:MIPS/Installation/Media/pt-br (57% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Media/ru "Handbook:MIPS/Installation/Media/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Media/ta "Handbook:MIPS/Installation/Media/ta (57% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Media/zh-cn "Handbook:MIPS/Installation/Media/zh-cn (57% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Media/ja "Handbook:MIPS/Installation/Media/ja (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Media/ko "Handbook:MIPS/Installation/Media/ko (57% translated)")

[Manual MIPS](/wiki/Handbook:MIPS/es "Handbook:MIPS/es")[Instalación](/wiki/Handbook:MIPS/Full/Installation/es "Handbook:MIPS/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:MIPS/Installation/About/es "Handbook:MIPS/Installation/About/es")Elegir los medios[Configurar la red](/wiki/Handbook:MIPS/Installation/Networking/es "Handbook:MIPS/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:MIPS/Installation/Disks/es "Handbook:MIPS/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:MIPS/Installation/Stage/es "Handbook:MIPS/Installation/Stage/es")[Instalar el sistema base](/wiki/Handbook:MIPS/Installation/Base/es "Handbook:MIPS/Installation/Base/es")[Configurar el núcleo](/wiki/Handbook:MIPS/Installation/Kernel/es "Handbook:MIPS/Installation/Kernel/es")[Configurar el sistema](/wiki/Handbook:MIPS/Installation/System/es "Handbook:MIPS/Installation/System/es")[Instalar las herramientas](/wiki/Handbook:MIPS/Installation/Tools/es "Handbook:MIPS/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:MIPS/Installation/Bootloader/es "Handbook:MIPS/Installation/Bootloader/es")[Terminar](/wiki/Handbook:MIPS/Installation/Finalizing/es "Handbook:MIPS/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:MIPS/Full/Working/es "Handbook:MIPS/Full/Working/es")[Introducción a Portage](/wiki/Handbook:MIPS/Working/Portage/es "Handbook:MIPS/Working/Portage/es")[Ajustes USE](/wiki/Handbook:MIPS/Working/USE/es "Handbook:MIPS/Working/USE/es")[Características de Portage](/wiki/Handbook:MIPS/Working/Features/es "Handbook:MIPS/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:MIPS/Working/Initscripts/es "Handbook:MIPS/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:MIPS/Working/EnvVar/es "Handbook:MIPS/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:MIPS/Full/Portage/es "Handbook:MIPS/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:MIPS/Portage/Files/es "Handbook:MIPS/Portage/Files/es")[Variables](/wiki/Handbook:MIPS/Portage/Variables/es "Handbook:MIPS/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:MIPS/Portage/Branches/es "Handbook:MIPS/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:MIPS/Portage/Tools/es "Handbook:MIPS/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:MIPS/Portage/CustomTree/es "Handbook:MIPS/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:MIPS/Portage/Advanced/es "Handbook:MIPS/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:MIPS/Full/Networking/es "Handbook:MIPS/Full/Networking/es")[Comenzar](/wiki/Handbook:MIPS/Networking/Introduction/es "Handbook:MIPS/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:MIPS/Networking/Advanced/es "Handbook:MIPS/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:MIPS/Networking/Modular/es "Handbook:MIPS/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:MIPS/Networking/Wireless/es "Handbook:MIPS/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:MIPS/Networking/Extending/es "Handbook:MIPS/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:MIPS/Networking/Dynamic/es "Handbook:MIPS/Networking/Dynamic/es")

## Contents

- [1Requerimientos hardware](#Requerimientos_hardware)
- [2Notas de instalación](#Notas_de_instalaci.C3.B3n)
- [3Vistazo rápido al arranque desde la red](#Vistazo_r.C3.A1pido_al_arranque_desde_la_red)
  - [3.1Congigurar TFTP y DHCP](#Congigurar_TFTP_y_DHCP)
- [4Arranque en red de estaciones SGI](#Arranque_en_red_de_estaciones_SGI)
  - [4.1Descargar una imagen de arranque en red](#Descargar_una_imagen_de_arranque_en_red)
  - [4.2Configuración DHCP para un cliente SGI](#Configuraci.C3.B3n_DHCP_para_un_cliente_SGI)
  - [4.3Opciones del núcleo](#Opciones_del_n.C3.BAcleo)
  - [4.4Arrancar los demonios](#Arrancar_los_demonios)
  - [4.5Arranque en red de la estación SGI](#Arranque_en_red_de_la_estaci.C3.B3n_SGI)
  - [4.6Solución de problemas](#Soluci.C3.B3n_de_problemas)
- [5Arranque en red en estaciones Cobalt](#Arranque_en_red_en_estaciones_Cobalt)
  - [5.1Vista general al procedimiento de arranque en red](#Vista_general_al_procedimiento_de_arranque_en_red)
  - [5.2Descargar una imagen de arranque en red Cobalt](#Descargar_una_imagen_de_arranque_en_red_Cobalt)
  - [5.3Configuración del servidor NFS](#Configuraci.C3.B3n_del_servidor_NFS)
  - [5.4Configuración DHCP para una máquina Cobalt](#Configuraci.C3.B3n_DHCP_para_una_m.C3.A1quina_Cobalt)
  - [5.5Arrancar los demonios](#Arrancar_los_demonios_2)
  - [5.6Arranque en red de la máquina Cobalt](#Arranque_en_red_de_la_m.C3.A1quina_Cobalt)
  - [5.7Solución de problemas](#Soluci.C3.B3n_de_problemas_2)
- [6Usar un CD de instalación](#Usar_un_CD_de_instalaci.C3.B3n)

## Requerimientos hardware

CPU (Big Endian port)
Clases de CPU-MIPS3, MIPS4, MIPS5 o MIPS64
CPU (Little Endian port)
Clases de CPU-MIPS4, MIPS5 o MIPS64
Memoria
128 MB
Espacio en disco
3.0 GB (excluyendo espacio swap)
Espacio swap
Al menos 256 MB

Para más información, lea los [Requerimientos hardware de MIPS](/wiki/MIPS/Hardware_Requirements "MIPS/Hardware Requirements")

## Notas de instalación

En muchas arquitecturas, el procesador ha evolucionado a través de varias generaciones, cada una se construye con los fundamentos de la anterior. La arquitectura MIPS no es ninguna excepción. Hay varias generaciones de CPUs que se cubren en la arquitectura MIPS. Para elegir el archivo comprimido con la imagen de arranque desde red y CFLAGS correctamente es necesario conocer de qué familia de CPU dispone el sistema. Estas familias están descritas como las arquitecturas de conjuntos de instrucciones (Instruction Set Architecture o ISA).

MIPS ISA
32/64 bits
CPUs cubiertas
MIPS 1
32 bits
R2000, R3000
MIPS 2
32 bits
R6000
MIPS 3
64 bits
R4000, R4400, R4600, R4700
MIPS 4
64 bits
R5000, RM5000, RM7000 R8000, R9000, R10000, R12000, R14000, R16000
MIPS 5
4 bits
Todavía ninguno
MIPS32
32 bits
Series AMD Alchemy, 4kc, 4km, y muchas otras... Hay algunas revisiones en el ISA MIPS32.
MIPS64
64 bits
Broadcom SiByte SB1, 5kc ... etc... Hay algunas revisiones en el ISA MIPS64.

**Nota**

El nivel ISA MIPS5 lo diseñó Silicon Graphics en 1994, sin embargo nunca se llegó a utilizar como una CPU en el mundo real. Vive como parte del ISA MIPS64.

**Nota**

Las ISAs MIPS32 y MIPS64 son una causa común de confusiones. El nivel ISA MIPS64 es realmente un superconjunto del ISA MIPS5 ISA, de modo que incluye todas las instrucciones hasta MIPS5 e ISAs anteriores. MIPS32 es el subconjunto de 32 bits de MIPS64. Existe debido a que la mayoría de las aplicaciones requieren únicamente procesado de 32 bits.

Igualmente, una cuestión importante es familiarizarse con el concepto de extremidad (endianness). La extremidad se refiere al forma en que la CPU lee las palabras de la memoria principal. Una palabra se puede leer como extremo mayor (big endian) en el que el byte más significativo se lee primero o extremo menor (little endian) en el que el byte menos significativo se lee primer. Las máquinas Intel x86 son normalmente de extremo menor, en cambio las máquinas Apple y Sparc son de extremo mayor. En MIPS se utilizan ambos. Para diferenciarlos, se añade _el_ a la arquitectura para denotar extremo menor (little endian).

Arquitectura
32 o 64 bits
Extremidad (Endianess)
Máquinas cubiertas
mips
32 bits
Big Endian
Silicon Graphics
mipsel
32 bits
Little Endian
Servidores Cobalt
mips64
64 bits
Big Endian
Silicon Graphics
mips64el
64 bits
Little Endian
Cobalt Servers

Para el lector que desee saber más sobre ISAs, se recomienda visitar los siguientes sitios web:

- [Sitio Web Linux/MIPS: MIPS ISA](http://www.linux-mips.org/wiki/Instruction_Set_Architecture)
- [Sitio Web Linux/MIPS: Extremidad (Endianness)](http://www.linux-mips.org/wiki/Endianess)
- [Sitio Web Linux/MIPS: Procesadores](http://www.linux-mips.org/wiki/Processors)
- [Wikipedia: Conjunto de instrucciones](https://es.wikipedia.org/wiki/Conjunto_de_instrucciones)

## Vistazo rápido al arranque desde la red

En esta sección cubriermos las necesidades para arrancar de forma correcta una estación de trabajo Silicon Graphics o una appliance de un servidor Cobalt Server desde la red. Esto es solo una guía rápida, no está pensada para incidir en todas las partes. Para obtener más información se recomienda leer el artículo sobre [nodos sin disco](/wiki/Diskless_nodes/es "Diskless nodes/es").

Dependiendo de la máquina, se necesita cierto hardware para arrancar desde la red e instalar Linux.

- En general:
  - DHCP/BOAMD Alchemy series, 4kc, 4km, many others... Hay pocas revisiones en el servidor ISA.OTP MIPS32 (se recomienda ISC DHCPd)
  - Paciencia (y mucha)
- Para las estaciones de trabajo Silicon Graphics:
  - Servidor TFTP server (se recomienda tftp-hpa)
  - Cuando se necesita utilizar la consola serie:
    - MiniDIN8 --> Cable serie RS-232 (únicamente necesario para sistemas IP22 y IP28 systems)
    - Cable de módem nulo (Null-modem)
    - Terminal compatible VT100 o ANSI que pueda trabajar a 9600 baudios
- Para los servidores Cobalt (NO el Qube original):
  - Servidor NFS
  - Cable de módem nulo (Null-modem)
  - Terminal compatible VT100 o ANSI que pueda trabajar a 115200 baudios

**Nota**

Las máquinas SGI utilizand un conector MiniDIN 8 para l , and most electronics stores should stock the plugs requiredos puertos serie. Aparentemente los cables de los módemsy Apple funcionan igual de bien que los cables serie pero debido a que las máquinas Apple están equipadas con USB y módems internos éstos son más difíciles de encontrar. Se pude obtener un diagrama del cableado desde la [Linux/MIPS Wiki](http://www.linux-mips.org/wiki/Serial_Cable) y en muchas tiendas de electrónica se deberían poder encontrar.

**Nota**

Para el terminal, esto podría ser un auténtico terminal VT100/ANSI o podría ser un ordenador PC corriendo un software de emulación de terminales (como HyperTerminal, Minicom, seyon, Telex, xc, screen, ... el que se prefiera). No importa la plataforma que corra la máquina siempre y cunado disponga de un terminal RS-232 disponible y el software apropiado.

**Nota**

Esta guía NO cubre el Qube original. El appliance servidor Qube original no dispone de puerto serie en su configuración por defecto y por tanto no es posible instalar Gentoo en ellas sin la ayuda de un destornillador y una máquina subrogada para realizar la instalación.

### Congigurar TFTP y DHCP

Tal y como se ha mencionado anteriormente, esto no es una guía completa, se trata de una configuración mínima que hará que las cosas funcionen. Utilícela cuando se realice una configuración desde cero o utilice las sugerencias para modificar una configuración existente para dar soporte al arranque por red.

Merece la pena comentar que no es necesario que los servidores utilizados estén corriendo Gentoo Linux, podrían correr FreeBSD o cualquier plataforma tipo Unix. Sin embargo, en esta guía se asume que se va a utilizar Gentoo Linux. Se se quiere, también se puede correr TFTP/NFS en una máquina distinta a la que corre el servidor DHCP.

**Advertencia**

El equipo MIPS de Gentoo no puede ayudar en el caso en el que se utilicen otros sistemas operativos como servidores de arranque en red.

Primer paso: Configurar DHCP. Para que el demonio DHCP ISC responda a las peticiones BOOTP (tal y como requiere la BOOTROM de SGI y Cobalt), en primer lugar se debe habilitar BOOTP dinámico en el rango de direcciones que se usa. A continuación configure una entrada para cada cliente con punteros a la imagen de inicio.

`root #` `emerge --ask net-misc/dhcp`

Una vez instalado, crear el fichero /etc/dhcp/dhcpd.conf. A continuación se muestra una configuración mínima para empezar a trabajar.

ARCHIVO **`/etc/dhcp/dhcpd.conf`** **Archivo dhcpd.conf básico**

```
# Indicar a dhcpd que deshabilite DNS dinámico.
# dhcpd no arrancará si no se ha definido esto.
ddns-update-style none;

# Crear una subred:
subnet 192.168.10.0 netmask 255.255.255.0 {
  # Conjunto de direcciones para nuestros clientes. ¡No olvidar 'dynamic-bootp'!
  pool {
    range dynamic-bootp 192.168.10.1 192.168.10.254;
  }

  # Servidores DNS y pasarela por defecto. Sustituir por los valores apropiados  option domain-name-servers 203.1.72.96, 202.47.56.17;
  option routers 192.168.10.1;

  # Indicar al servidor DHCP su autoridad para esta subred.
  authoritative;

  # Permitir la utilización de BOOTP en esta subred.
  allow bootp;
}

```

Con esta configuración se pueden añadir cualquier número de clientes dentro de la cláusula de la subred. Más adelante en este guía se cubre esto.

Siguiente paso: Configurar el servidor TFTP. Se recomienda utilizar tftp-hpa ya que es el único demonio TFTP del que se tiene constancia que funciona correctamente. Proceder a su instalación tal y como se muestra abajo:

`root #` `emerge --ask net-ftp/tftp-hpa`

Esto creará /tftproot para almacenar las imágenes del arranque en red. Mover esto a otro lugar si es necesario. En esta guía se asume que se almacena en la localización por defecto.

## Arranque en red de estaciones SGI

### Descargar una imagen de arranque en red

Dependiendo del sistema en el que se realiza la instalación, hay varias imágenes disponibles para la descarga. Están etiquetadas conforme al tipo de sistema y procesador para los que se han compilado. Los tipo de máquina son los siguientes:

Código del nombre
Máquinas
IP22
Indy, \*Indigo 2, Challenge S
IP26
\*Indigo 2 Power
IP27
Origin 200, Origin 2000
IP28
\*Indigo 2 Impact
IP30
Octane
IP32
O2

**Nota**

_Indigo 2_. Es un error común mezclar el IRIS Indigo (IP12 w/ R3000 CPU o IP20 con una CPU R4000, en ninguno de ellos se corre Linux), el Indigo 2 (IP22, que corre Linux correctamente), el Indigo 2 Power basado en R8000 (que no corre Linux) y el Indigo 2 Impact basado en R10000 (IP28, el cual es altamente experimental). Por favor, recordar que son máquinas diferentes.

También se debe recordar que el nombre de fichero r4k hace referencia a los procesadores de la serie R4000, r5k los de la serie R5000, rm5k los de la RM5200 y r10k para los de la R10000. Las imágenes están disponibles en los [servidores réplica de Gentoo](https://www.gentoo.org/downloads/mirrors/).

### Configuración DHCP para un cliente SGI

Una vez descargado el fichero descomprimirlo en el directorio /tftproot/. (Utilizar bzip2 -d para descomprimiro). A continuación editar el fichero /etc/dhcp/dhcpd.conf y añadir la entrada apropiada para el cliente SGI.

ARCHIVO **`/etc/dhcp/dhcpd.conf`** **Extracto para estación de trabajo SGI**

```
subnet xxx.xxx.xxx.xxx netmask xxx.xxx.xxx.xxx {
  # ... aquí se incluye el contenido usual ...

  # Estación de trabajo SGI... cambiar 'sgi' por el nombre de su máquina SGI.
  host sgi {

    # Dirección MAC de la máquina SGI. Normalmente esta información está escrita
    # en la parte trasera o en la base de la máquina
    ha machinerdware ethernet 08:00:69:08:db:77;

    # Servidor TFTP desde el que se van a realizar las descargas (por defecto el mismo que el servidor DHCP)
    next-server 192.168.10.1;

    # Dirección IP que se asignará a la máquina SGI
    fixed-address 192.168.10.3;

    # Nombre del fichero en la PROM para descargar e iniciar
    filename "/gentoo-r4k.img";
  }
}

```

### Opciones del núcleo

Ya casi hemos terminado, sin embargo todavía se deben realizar un par de ajustes. Lanzar una consola con privilegios de root.

Deshabilitar "Path Maximum Transfer Unit" pues de lo contrario la PROM SGI no encontrará el núcleo:

`root #` `echo 1 > /proc/sys/net/ipv4/ip_no_pmtu_disc`

Ajustar el rango de puertos utilizables por la PROM de SGI:

`root #` `echo "2048 32767" > /proc/sys/net/ipv4/ip_local_port_range`

`root #` `echo "2048 32767" > /proc/sys/net/ipv4/ip_local_port_range`

Esto debería ser suficiente para que el servidor Linux se entienda con la PROM de SGI.

### Arrancar los demonios

Llegados a este punto, arrancar los demonios.

`root #` `/etc/init.d/dhcp start
`

`root #` `/etc/init.d/in.tftpd start
`

Si todo ha ido bien en el último paso entonces todo está preparado para poner en marcha la estación de trabajo y seguir con la guía. Si el servidor DHCP no está funcionando por cualquier razón, se puede intentar lanzar dhcpd en la línea de órdenes y comprobar qué puede estar ocurriendo. Si todo está en su sitio, se debería lanzar sin problemas y pasar a segundo plano, de lo contrario se mostrará 'exiting.' justo debajo del problema que ha encontrado.

Una forma fácil de verificar si el demonio tftp está corriendo es teclear la siguiente orden y confirmar la salida:

`root #` `netstat -al | grep ^udp`

```
udp        0      0 *:bootpc                *:*
udp        0      0 *:631                   *:*
udp        0      0 *:xdmcp                 *:*
udp        0      0 *:tftp                  *:* <-- (look for this line)

```

### Arranque en red de la estación SGI

Bien, ya está todo configurado, el servidor DHCP está corriendo y también el TFTP. Llegó el momento de arrancar la máquina SGI. Encender la unidad y cuando se muestre "Running power-on diagnostics" en la pantalla, bien hacer clic en "Stop For Maintenance" o pulsar la tecla `Escape`. Se mostrará un menú similar al siguiente.

`Running power-on diagnostics`

```
System Maintenance Menu

1) Start System
2) Install System Software
3) Run Diagnostics
4) Recover System
5) Enter Command Monitor
Option?

```

Teclear `5` para entrar en el monitor de órdenes. Desde este monitor, lanzar el proceso BootP:

`>>` `bootp(): root=/dev/ram0`

A partir de este punto, la máquina debería comenzar a descargar la imagen y aproximadamente veinte segundos más tarde, comenzar el inicio de Linux. Si todo va bien, se iniciará un intérprete de ordenes busybox (ash) tal y como se muestra abajo y por tanto se puede continuar con la instalación de Gentoo Linux.

CÓDIGO **Cuando las cosas van bien...**

```
init started:  BusyBox v1.00-pre10 (2004.04.27-02:55+0000) multi-call binary

Gentoo Linux; http://www.gentoo.org/
 Copyright 2001-2004 Gentoo Technologies, Inc.; Distributed under the GPL

 Gentoo/MIPS Netboot for Silicon Graphics Machines
 Build Date: April 26th, 2004

 * To configure networking, do the following:

 * For Static IP:
 * /bin/net-setup <IP Address> <Gateway Address> [telnet]

 * For Dynamic IP:
 * /bin/net-setup dhcp [telnet]

 * If you would like a telnetd daemon loaded as well, pass "telnet"
 * As the final argument to /bin/net-setup.

Please press Enter to activate this console.

```

### Solución de problemas

Si la máquina se niega a descargar su imagen, se puede deber a dos cosas:

1. No se han seguido correctamente las instrucciones, o
2. Necesita algo de persuasión (¡Suelte ese martillo!)

A continuación se muestra una lista de cosas a comprobar:

- dhcpd está entregando una dirección IP a la máquina SGI. Deberían aparecer algunos mensajes mostrando peticones BOOTP en los registros del sistema. La orden tcpdump puede ser de utilidad.
- Se han definido correctamente los permisos en la carpeta tftp (nomalmente /tftproot/ debería ser legible por todo el mundo)
- Comprobar los registros del sistema para ver qué está reportando el servidor tftp (quizá algunos errores)

Si se ha comprobado todo en el servidor y se siguen produciendo demoras indefinidas o errores en la máquian SGI, se puede teclear lo siguiente en la consola.

`>>` `resetenv
`

`>>` `unsetenv netaddr
`

`>>` `unsetenv dlserver
`

`>>` `init
`

`>>` `bootp(): root=/dev/ram0`

## Arranque en red en estaciones Cobalt

### Vista general al procedimiento de arranque en red

Al contrario que las máquinas SGI, los servidores Cobalt utilizan NFS para transferir sus núcleos para el arranque. Arrancar la máquina manteniendo pulsadas las teclas izquierda y derecha de movimiento del cursor mientras se enciende la máquina. La máquina entonces intentará obtener una dirección IP a través de BOOTP, montar el directorio /nfsroot/ desde el servidor mediante NFS, a continuación intentar descargar e iniciar el fichero vmlinux\_raq-2800.gz (dependiendo del modelo) que se asume es un fichero binario ELF estándar.

### Descargar una imagen de arranque en red Cobalt

En [http://distfiles.gentoo.org/experimental/mips/historical/netboot/cobalt/](http://distfiles.gentoo.org/experimental/mips/historical/netboot/cobalt/) están disponibles las imágenes de arranque necesarias para poner en funcionamiento un servidor Cobalt. Los ficheros deben tener un nombre nfsroot-NÚCLEO-COLO-FECHA-cobalt.tar. Seleccionar el más reciente y desempaquetarlo en / como se muestra abajo:

`root #` `tar -C / -xvf nfsroot-2.6.13.4-1.19-20051122-cobalt.tar`

### Configuración del servidor NFS

Debido a que la máquina utiliza NFS para descargar su imagen, es necesario exportar /nfsroot/ en el servidor. Instalar el paquete [net-fs/nfs-utils](https://packages.gentoo.org/packages/net-fs/nfs-utils):

`root #` `emerge --ask net-fs/nfs-utils`

Una vez hecho esto, escribir lo siguiente en el fichero /etc/exports.

ARCHIVO **`/etc/exports`** **Exportar el directorio /nfsroot**

```
/nfsroot      *(ro,sync)

```

Una vez hecho esto, arrancar el servidor NFS:

`root #` `/etc/init.d/nfs start`

Si el servidor NFS ya estaba corriendo, indicarle que eche un vistazo de nuevo a sus exportaciones utilizando exportfs.

`root #` `exportfs -av`

### Configuración DHCP para una máquina Cobalt

Ahora se realiza la parte DHCP de todo esto, lo cual es relativamente sencillo. Añadir lo siguiente al fichero /etc/dhcp/dhcpd.conf.

ARCHIVO **`/etc/dhcp/dhcpd.conf`** **Extracto del servidor Cobalt**

```
subnet xxx.xxx.xxx.xxx netmask xxx.xxx.xxx.xxx {
  # ... contenido habitual ...

  # Configuración para el servidor Cobalt
  # Definir aquí el nombre del equipo:
  host qube {
    # Ruta al directorio nfsroot.
    # Esto se utilizar principalmente cuando se utilizar la opción de arranque TFTP en CoLo
    # Esto no debería cambiarse
    option root-path "/nfsroot";

    # Dirección MAC del servidor Cobalt
    hardware ethernet 00:10:e0:00:86:3d;

    # Servidor desde el que se descarga la imagen
    next-server 192.168.10.1;

    # Dirección IP del servidor Cobalt
    fixed-address 192.168.10.2;

    # Localización del fichero default.colo relativa a /nfsroot
    # Esto no debería cambiarse
    filename "default.colo";
  }
}

```

### Arrancar los demonios

Ahora se deben arrancar los demonios. Escribir lo siguiente:

`root #` `/etc/init.d/dhcp start
`

`root #` `/etc/init.d/nfs start`

Si todo fue bien en el último paso, debería estar todo preparado para encender la máquina y continuar con esta guía. En caso de que el servidor DHCP no esté funcionando correctamente por la razón que sea, se puede intentar lanzar dhcpd en la línea de órdenes y ver qué nos cuenta. Debería simplemente lanzarse en segundo plano, de lo contrario mostrará 'exiting.' (saliendo) justo debajo del error que se ha producido.

### Arranque en red de la máquina Cobalt

Ahora es el momento de poner en marcha la máquina Cobalt. Conectarse a un cable de módem nulo y configurar el terminal serie para que utilice 115200 baudios, 8 bits, sin paridad, un bit de parada y emulación VT100. Una vez se haya hecho esto, mantener pulsadas las teclas izquierda y derecha de movimiento del cursor mientras se enciende la máquina.

En el panel se debería mostrar "Net Booting", y visualizar alguna actividad de red seguida de CoLo. En el panel, bajar hasta la opción "Network (NFS)" y pulsar la tecla `Enter`. Observar que la máquina comienza su arranque en la consola serie.

`...`

```
elf: 80080000 <-- 00001000 6586368t + 192624t
elf: entry 80328040
net: interface down
CPU revision is: 000028a0
FPU revision is: 000028a0
Primary instruction cache 32kB, physically tagged, 2-way, linesize 32 bytes.
Primary data cache 32kB 2-way, linesize 32 bytes.
Linux version 2.4.26-mipscvs-20040415 (root@khazad-dum) (gcc version 3.3.3...)
Determined physical RAM map:
 memory: 08000000 @ 00000000 (usable)
Initial ramdisk at: 0x80392000 (3366912 bytes)
On node 0 totalpages: 32768
zone(0): 32768 pages.
zone(1): 0 pages.
zone(2): 0 pages.
Kernel command line: console=ttyS0,115200 root=/dev/ram0
Calibrating delay loop... 249.85 BogoMIPS
Memory: 122512k/131072k available (2708k kernel code, 8560k reserved, 3424k dat)

```

Se mostrará un intérprete ash de busybox tal y como se muestra abajo desde el cual se puede continuar con la instalación de Gentoo Linux.

CÓDIGO **Cuando las cosas van bien...**

```
VFS: Mounted root (ext2 filesystem) readonly.
Freeing unused kernel memory: 280k freed
init started:  BusyBox v1.00-pre10 (2004.04.27-02:55+0000) multi-call binary

Gentoo Linux; http://www.gentoo.org/
 Copyright 2001-2004 Gentoo Technologies, Inc.; Distributed under the GPL

 Gentoo/MIPS Netboot for Cobalt Microserver Machines
 Build Date: April 26th, 2004

 * To configure networking, do the following:

 * For Static IP:
 * /bin/net-setup <IP Address> <Gateway Address> [telnet]

 * For Dynamic IP:
 * /bin/net-setup dhcp [telnet]

 * If you would like a telnetd daemon loaded as well, pass "telnet"
 * As the final argument to /bin/net-setup.

Please press Enter to activate this console.

```

### Solución de problemas

Si la máquina se niega a descargar su imagen, se puede deber a dos cosas:

1. No se han seguido correctamente las instrucciones, o
2. Necesita algo de persuasión (¡Suelte ese martillo!)

A continuación se muestra una lista de cosas a comprobar:

- dhcpd le está entregando una dirección IP a la máquina Cobalt. Observar los mensajes acerca de peticiones BOOTP en los registros del sistema. tcpdump puede también ser de utilidad.
- Los permisos se han definido correctamtente en la carpeta /nfsroot/ (debería ser legible por todo el mundo).
- Asegurarse que el servidor NFS está corriendo y exportando el directorio /nfsroot/. Comprobar esto mediante exportfs -v en el servidor.

## Usar un CD de instalación

En máquinas Silicon Graphics, es posible arrancar desde un CD para instalar sistemas operativos (así es como se instala IRIX por ejemplo). Recientemente, imágenes para tales CD de arranque para instalar Gentoo han sido posibles. Estos CD,s. están diseñados para funcionar de la misma manera.

De momento el Live CD de Gentoo/MIPS sólo funcionará en estaciones de trabajo SGI Indy, Indigo y O2 equipadas con CPU,s. series R4000 y R5000, sin embargo otras plataformas pueden ser posibles en un futuro.

Las imágenes del Live CD se pueden encontrar bajo el directorio experimental/mips/livecd/ en los servidores espejo de Gentoo.

**Advertencia**

Estos CD,s. son altamente experimentales en este momento. Podrían no funcionar. Por favor informe tanto de exitos como de fracasos en Bugzilla, [en este hilo de los foros](https://forums.gentoo.org/viewtopic.php?t=242518) o en el canal IRC #gentoo-mips.

[← Acerca de la instalación](/wiki/Handbook:MIPS/Installation/About/es "Previous part") [Inicio](/wiki/Handbook:MIPS/es "Handbook:MIPS/es") [Configurando la red →](/wiki/Handbook:MIPS/Installation/Networking/es "Next part")