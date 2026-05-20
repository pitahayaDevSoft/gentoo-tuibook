# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Kernel/de "Handbuch:MIPS/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Kernel "Handbook:MIPS/Installation/Kernel (100% translated)")
- español
- [français](/wiki/Handbook:MIPS/Installation/Kernel/fr "Handbook:MIPS/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Kernel/it "Handbook:MIPS/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Kernel/pl "Handbook:MIPS/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Kernel/pt-br "Manual:MIPS/Instalação/Kernel (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Kernel/ru "Handbook:MIPS/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Kernel/ta "கையேடு:MIPS/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Kernel/zh-cn "手册：MIPS/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Kernel/ja "ハンドブック:MIPS/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Kernel/ko "Handbook:MIPS/Installation/Kernel/ko (100% translated)")

[Manual MIPS](/wiki/Handbook:MIPS/es "Handbook:MIPS/es")[Instalación](/wiki/Handbook:MIPS/Full/Installation/es "Handbook:MIPS/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:MIPS/Installation/About/es "Handbook:MIPS/Installation/About/es")[Elegir los medios](/wiki/Handbook:MIPS/Installation/Media/es "Handbook:MIPS/Installation/Media/es")[Configurar la red](/wiki/Handbook:MIPS/Installation/Networking/es "Handbook:MIPS/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:MIPS/Installation/Disks/es "Handbook:MIPS/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:MIPS/Installation/Stage/es "Handbook:MIPS/Installation/Stage/es")[Instalar el sistema base](/wiki/Handbook:MIPS/Installation/Base/es "Handbook:MIPS/Installation/Base/es")Configurar el núcleo[Configurar el sistema](/wiki/Handbook:MIPS/Installation/System/es "Handbook:MIPS/Installation/System/es")[Instalar las herramientas](/wiki/Handbook:MIPS/Installation/Tools/es "Handbook:MIPS/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:MIPS/Installation/Bootloader/es "Handbook:MIPS/Installation/Bootloader/es")[Terminar](/wiki/Handbook:MIPS/Installation/Finalizing/es "Handbook:MIPS/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:MIPS/Full/Working/es "Handbook:MIPS/Full/Working/es")[Introducción a Portage](/wiki/Handbook:MIPS/Working/Portage/es "Handbook:MIPS/Working/Portage/es")[Ajustes USE](/wiki/Handbook:MIPS/Working/USE/es "Handbook:MIPS/Working/USE/es")[Características de Portage](/wiki/Handbook:MIPS/Working/Features/es "Handbook:MIPS/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:MIPS/Working/Initscripts/es "Handbook:MIPS/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:MIPS/Working/EnvVar/es "Handbook:MIPS/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:MIPS/Full/Portage/es "Handbook:MIPS/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:MIPS/Portage/Files/es "Handbook:MIPS/Portage/Files/es")[Variables](/wiki/Handbook:MIPS/Portage/Variables/es "Handbook:MIPS/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:MIPS/Portage/Branches/es "Handbook:MIPS/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:MIPS/Portage/Tools/es "Handbook:MIPS/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:MIPS/Portage/CustomTree/es "Handbook:MIPS/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:MIPS/Portage/Advanced/es "Handbook:MIPS/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:MIPS/Full/Networking/es "Handbook:MIPS/Full/Networking/es")[Comenzar](/wiki/Handbook:MIPS/Networking/Introduction/es "Handbook:MIPS/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:MIPS/Networking/Advanced/es "Handbook:MIPS/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:MIPS/Networking/Modular/es "Handbook:MIPS/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:MIPS/Networking/Wireless/es "Handbook:MIPS/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:MIPS/Networking/Extending/es "Handbook:MIPS/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:MIPS/Networking/Dynamic/es "Handbook:MIPS/Networking/Dynamic/es")

## Contents

- [1Opcional: Instalar firmware y/o microcódigo](#Opcional:_Instalar_firmware_y.2Fo_microc.C3.B3digo)
  - [1.1Firmware](#Firmware)
    - [1.1.1Sugerido: Linux Firmware](#Sugerido:_Linux_Firmware)
      - [1.1.1.1Carga de Firmware](#Carga_de_Firmware)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Cargador de arranque](#Cargador_de_arranque)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2Esquema tradicional, otros cargadores de arranque (por ejemplo, (e)lilo, syslinux, etc.)](#Esquema_tradicional.2C_otros_cargadores_de_arranque_.28por_ejemplo.2C_.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
- [3Configuración y compilación del núcleo](#Configuraci.C3.B3n_y_compilaci.C3.B3n_del_n.C3.BAcleo)
  - [3.1Alternativa: Configuración manual](#Alternativa:_Configuraci.C3.B3n_manual)
  - [3.2Instalar las fuentes del núcleo](#Instalar_las_fuentes_del_n.C3.BAcleo)
    - [3.2.1Proceso modprobed-db](#Proceso_modprobed-db)
    - [3.2.2Proceso manual](#Proceso_manual)
      - [3.2.2.1Habilitar las opciones requeridas](#Habilitar_las_opciones_requeridas)
    - [3.2.3Opcional: módulos del núcleo firmados](#Opcional:_m.C3.B3dulos_del_n.C3.BAcleo_firmados)
  - [3.3Preparar la configuración](#Preparar_la_configuraci.C3.B3n)
  - [3.4Personalizar la configuración](#Personalizar_la_configuraci.C3.B3n)
  - [3.5Compilar e instalar](#Compilar_e_instalar)

## Opcional: Instalar firmware y/o microcódigo

### Firmware

#### Sugerido: Linux Firmware

En muchos sistemas, se requiere firmware no FOSS para el funcionamiento de ciertos dispositivos. El paquete [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) contiene firmware para muchos dispositivos, pero no para todos.

**Consejo**

La mayoría de las tarjetas inalámbricas y GPU,s. requieren firmware para funcionar.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Nota**

La instalación de determinados paquetes de firmware suele requerir la aceptación de las licencias de firmware asociadas. Si es necesario, visite la [sección de manejo de licencias](/wiki/Handbook:MIPS/Working/Portage/es#Licenses "Handbook:MIPS/Working/Portage/es") del Manual para obtener ayuda sobre cómo aceptar licencias.

##### Carga de Firmware

Los archivos de firmware suelen cargarse junto con el módulo del núcleo correspondiente. Esto significa que el firmware debe integrarse en el núcleo mediante _CONFIG\_EXTRA\_FIRMWARE_ si el módulo del núcleo está configurado como _Y_ en lugar de _M_. En la mayoría de los casos, integrar un módulo que requiere firmware puede complicar o interrumpir la carga.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") puede usarse para automatizar la instalación del núcleo, la generación de [initramfs](/wiki/Initramfs "Initramfs"), la generación de [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") y/o la configuración del gestor de arranque, entre otras cosas. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implementa dos métodos para lograr esto: el installkernel tradicional, originado en Debian, y el kernel-install de [systemd](/wiki/Systemd/es "Systemd/es"). La elección de uno u otro depende, entre otros factores, del gestor de arranque del sistema. Por defecto, el kernel-install de systemd se usa en los perfiles de systemd, mientras que el installkernel tradicional es el predeterminado para otros perfiles.

### Cargador de arranque

Ahora es el momento de pensar qué cargador de arranque desea el usuario para el sistema. Si no está seguro, siga la subsección 'Diseño tradicional' a continuación.

**Importante**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Los usuarios de GRUB pueden usar kernel-install de systemd o el installkernel tradicional de Debian. El indicador USE [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") alterna entre estas implementaciones. Para ejecutar automáticamente grub-mkconfig al instalar el núcleo, active el indicador [USE](/wiki/USE "USE") [grub](https://packages.gentoo.org/useflags/grub) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)").

**Nota**

GRUB requires kernels to be installed to /boot.

ARCHIVO **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

}}

**Nota**

systemd-boot requires kernels to be installed to /efi.

**Nota**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Nota**

EFIstub requires kernels to be installed to /efi.

#### Esquema tradicional, otros cargadores de arranque (por ejemplo, (e)lilo, syslinux, etc.)

El esquema tradicional con /boot (p. ej., (e)LILO, syslinux, etc.) se usa por defecto si los indicadores USE [grub](https://packages.gentoo.org/useflags/grub) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") y [uki](https://packages.gentoo.org/useflags/uki) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") no están habilitados. No se requiere ninguna acción adicional.

### Initramfs

Un **i** nitial **ram**-based **f** ile **s** ystem, o [initramfs](/wiki/Initramfs "Initramfs"), puede ser necesario para que un sistema arranque. Diversos casos pueden requerirlo, pero los más comunes incluyen:

- Núcleos donde los controladores de almacenamiento/sistema de archivos son módulos.
- Diseños con /usr/ o /var/ en particiones separadas.
- Sistemas de archivos raíz cifrados.

**Consejo**

Los [núcleos de distribución](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") están diseñados para usarse con un initramfs, ya que muchos controladores de almacenamiento y sistemas de archivos se construyen como módulos.

Además de montar el sistema de archivos raíz, un initramfs también puede realizar otras tareas como:

- Ejecución de **f** ile **s** ystem **c** onsistency chec **k** fsck, una herramienta para comprobar y reparar la consistencia de un sistema de archivos en caso de apagado incorrecto del sistema.
- Proporcionar un entorno de recuperación en caso de fallos tras el arranque.

[Installkernel](/wiki/Installkernel "Installkernel") puede generar automáticamente un initramfs al instalar el núcleo si el indicador USE [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") o [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") está habilitado:

ARCHIVO **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub/es "EFI stub/es") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fes "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

ARCHIVO **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

## Configuración y compilación del núcleo

**Consejo**

Usar dist-kernel en el primer arranque puede ser una buena idea, ya que proporciona un método muy sencillo para descartar problemas del sistema y de configuración del núcleo. Tener siempre un núcleo funcional al que recurrir puede acelerar la depuración y aliviar la ansiedad cuando al actualizar el sistema ya no arranca.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Importante**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Clasificadas de menor a mayor complicación:

[La manera completamente manual](/wiki/Handbook:MIPS/Installation/Kernel/es#Alternative:_Manual_configuration "Handbook:MIPS/Installation/Kernel/es")Las nuevas fuentes del núcleo se instalan a través del administrador de paquetes del sistema. El núcleo se configura, compila e instala manualmente utilizando eselect kernel y una serie de comandos make. Las futuras actualizaciones del núcleo repiten el proceso manual de configuración, compilación e instalación de los archivos del núcleo. Este es el proceso más complejo, pero ofrece el máximo control sobre el proceso de actualización del núcleo.[Estrategia híbrida: Genkernel](/wiki/Handbook:MIPS/Installation/Kernel/es#Alternative:_Genkernel "Handbook:MIPS/Installation/Kernel/es")Aquí usamos el término híbrido, pero tenga en cuenta que las fuentes de dist-kernel y del manual incluyen métodos para lograr el mismo objetivo. Las nuevas fuentes del núcleo se instalan a través del administrador de paquetes del sistema. Los administradores del sistema pueden usar la herramienta genkernel de Gentoo para configurar, compilar, e instalar automáticamente el núcleo Linux, sus módulos asociados y (opcionalmente, pero _**no**_ habilitado por defecto) un archivo initramfs archivo. Es posible proporcionar un archivo de configuración del núcleo personalizado si es necesaria la personalización. La configuración, compilación e instalación futuras del núcleo requieren la participación del administrador del sistema en la forma de ejecutar eselect kernel, genkernel y potencialmente otros comandos para cada actualización. Esta opción solo debe considerarse para usuarios que saben que necesitan genkernel

El eje alrededor del cual se construyen todas las distribuciones es el núcleo Linux. Es la capa entre los programas del usuario y el hardware del sistema. Aunque el manual proporciona a sus usuarios varias fuentes posibles del núcleo, hay disponible una lista más completa y con descripciones más detalladas en la [Página de descripción general del núcleo](/wiki/Kernel/Overview "Kernel/Overview").

**Consejo**

Tareas de instalación del núcleo como copiar la imagen del núcleo a /boot o la [Partición del Sistema EFI](/index.php?title=Partici%C3%B3n_del_Sistema_EFI&action=edit&redlink=1 "Partición del Sistema EFI (page does not exist)"), generar un [initramfs](/wiki/Initramfs "Initramfs") y/o [Imagen unificada del núcleo](/index.php?title=Imagen_unificada_del_n%C3%BAcleo&action=edit&redlink=1 "Imagen unificada del núcleo (page does not exist)"), que actualiza la configuración del gestor de arranque, se puede automatizar con [installkernel](/wiki/Installkernel "Installkernel"). Es posible que los usuarios deseen configurar e instalar [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) antes de continuar. Consulte la [Sección de instalación del núcleo a continuación](/index.php?title=Handbook:MIPS/Instalaci%C3%B3n/Kernel&action=edit&redlink=1 "Handbook:MIPS/Instalación/Kernel (page does not exist)") para obtener más información.

### Alternativa: Configuración manual

### Instalar las fuentes del núcleo

Al instalar y compilar el núcleo para sistemas basados en mips, Gentoo recomienda el paquete [sys-kernel/mips-sources](https://packages.gentoo.org/packages/sys-kernel/mips-sources).

Elija una fuente del núcleo adecuada e instálela usando emerge:

`root #` `emerge --ask sys-kernel/mips-sources`

Esto instalará las fuentes del núcleo Linux en /usr/src/ usando la versión específica del kernel en el nombre de la ruta. No creará un enlace simbólico de forma automática a no ser que el indicador USE [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") esté habilitado en el paquete de fuentes del núcleo elegido.

Es una convencion que se mantenga el enlace simbólico /usr/src/linux, de modo que se refiera a las fuentes que correspondan con el núcleo que se está ejecutando actualmente. Sin embargo, este enlace simbólico no se creará por defecto. Una manera fácil de crear el enlace simbólico es utilizar el módulo kernel de eselect.

Para obtener más información sobre la finalidad del enlace simbólico y cómo administrarlo, consulte [Kernel/Upgrade](/wiki/Kernel/Upgrade/es "Kernel/Upgrade/es").

Primero, enumere todos los núcleos instalados:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.1-gentoo

```

Para crear un enlace simbólico llamado linux, use:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.19.1-gentoo

```

Configurar manualmente un núcleo suele considerarse uno de los procedimientos más difíciles que debe realizar un administrador de sistemas. Y nada es menos cierto: después de configurar algunos núcleos, ¡nadie recuerda lo difícil que fue! Hay dos maneras en que un usuario de Gentoo puede administrar un sistema de núcleo manual, las cuales se listan a continuación:

**Importante**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### Proceso modprobed-db

Una forma muy sencilla de administrar el núcleo es instalar primero [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) y usar [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) para recopilar información sobre las necesidades del sistema. modprobed-db es una herramienta que monitoriza el sistema mediante crontab para agregar todos los módulos de todos los dispositivos a lo largo de su vida útil y garantizar que todo lo que el usuario necesita sea compatible. Por ejemplo, si se agrega un mando de Xbox después de la instalación, modprobed-db agregará los módulos que se compilarán la próxima vez que se recompile el núcleo. Puede encontrar más información sobre este tema en el artículo [Modprobed-db](/wiki/Modprobed-db "Modprobed-db").

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:MIPS/Installation/Kernel#Distribution_kernels.2Fes "Handbook:MIPS/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.
}}

#### Proceso manual

Este método permite al usuario tener control total sobre la construcción de su núcleo con la mínima ayuda de herramientas externas que desee. Algunos podrían considerarlo una complicación sin más.

Sin embargo, con esta elección una cosa sí es cierta: es vital conocer el sistema para configurar manualmente un núcleo. La mayor cantidad de información se puede obtener instalando [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) que contiene la orden lspci:

`root #` `emerge --ask sys-apps/pciutils`

**Nota**

Dentro de la jaula chroot, se pueden ignorar con tranquilidad las advertencias sobre pcilib (como _pcilib: cannot open /sys/bus/pci/devices_) que pudiera mostrar lspci.

Otra fuente de información sobre nuestro sistema consiste en ejecutar lsmod para ver los módulos del nucleo que ha usado el CD de instalación y tener así buenas indicaciones sobre qué habilitar.

Ahora vaya al directorio de origen del núcleo.

`root #` `cd /usr/src/linux
`

**Consejo**

Para ver la lista completa con los argumentos de make disponibles para el núcleo, ejecute `make help`.

El núcleo cuenta con un método para autodetectar los módulos que se utilizan actualmente en el CD de instalación, lo que ofrece un excelente punto de partida para que el usuario configure los suyos propios. Esto se puede activar con:

`root #` `make localmodconfig`

Ahora es el momento de configurar usando nconfig:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:MIPS/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fes "Handbook:MIPS/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:MIPS/Installation/Kernel#Compiling_and_installing.2Fes "Handbook:MIPS/Installation/Kernel")

###### Habilitar las opciones requeridas

#### Opcional: módulos del núcleo firmados

Para firmar automáticamente los módulos del núcleo, habilite CONFIG\_MODULE\_SIG\_ALL:

KERNEL **Firmar módulos del núcleo CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Opcionalmente, cambie el algoritmo hash si lo desea.

Para exigir que todos los módulos estén firmados con una firma válida, habilite CONFIG\_MODULE\_SIG\_FORCE también:

KERNEL **Forzar el firmado de los módulos del núcleo CONFIG\_MODULE\_SIG\_FORCE**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Para usar una clave personalizada, especifique la ubicación de esta clave en CONFIG\_MODULE\_SIG\_KEY. Si no se especifica, el sistema de compilación del núcleo generará una clave. Se recomienda generar uno manualmente. Esto se puede hacer con:

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

OpenSSL hará algunas preguntas sobre el usuario para el que se genera la clave; se recomienda completar estas preguntas lo más detalladamente posible.

Guarde la clave en un lugar seguro; como mínimo, solo el usuario root debe poder leerla. Verifique esto con:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

Si esto genera algo distinto a lo anterior, corrija los permisos con:

`root #` `chown root:root kernel_key.pem
`

`root #` `chmod 400 kernel_key.pem
`

KERNEL **Especificar la clave para el firmado CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

Para firmar también módulos del núcleo externos instalados por otros paquetes a través de `linux-mod-r1.eclass`, habilite el indicador USE [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") globalmente:

ARCHIVO **`/etc/portage/make.conf`** **Habilitar el firmado de módulos**

```
USE="modules-sign"

# Opcional, cuando se usen claves personalizadas.
MODULES_SIGN_KEY="/ruta/a/kernel_key.pem"
MODULES_SIGN_CERT="/ruta/a/kernel_key.pem" # Solo es necesario si MODULES_SIGN_KEY no contiene también el certificado
MODULES_SIGN_HASH="sha512" # sha512 es el predeterminado

```

**Nota**

MODULES\_SIGN\_KEY y MODULES\_SIGN\_CERT pueden apuntar a archivos diferentes. Para este ejemplo, el archivo pem generado por OpenSSL incluye tanto la clave como el certificado adjunto y, por lo tanto, ambas variables se configuran con el mismo valor.

### Preparar la configuración

**Importante**

En las máquinas Origin 200/2000, Indigo2 Impact (R10000), Octane/Octane2 y O2, se necesita un núcleo de 64 bits para arrancar el sistema. En estas máquinas, hacer emerge de [sys-devel/kgcc64](https://packages.gentoo.org/packages/sys-devel/kgcc64) para crear un compilador cruzado con el que construir núcleos de 64 bits.

En muchos de los sistemas soportados se dispone de .configs ocultos entre los fuentes del núcleo. No todos los sistemas tienen estas configuraciones distribuidas así. Aquéllos que los tienen se pueden configurar mediante las órdenes descritas en la tabla de abajo.

Sistema
Orden de configuración
Servidores Cobalt
make cobalt\_defconfigIndy, Indigo2 (R4k), Challenge S
make ip22\_defconfigOrigin 200/2000
make ip27\_defconfigIndigo2 Impact (R10k)
make ip28\_defconfigO2
make ip32\_defconfig

Todas las imágenes de instalación de Gentoo ofrecen una opción de configuración del núcleo que es parte de la propia imagen y accesible a través de /proc/config.gz. Esto se puede utilizar en muchos casos. Es mejor si se consulta la misma que la del núcleo que está corriendo en ese momento. Para extraerla, simplemente lanzarla a través de zcat como se muestra abajo.

`root #` `zcat /proc/config.gz > .config`

**Importante**

Esta configuración del núcleo está preparada para una imagen de arranque por red. Esto es, espera encontrar una imagen de sistema de archivos raíz en algún lugar cercano como un directorio para initramfs o un dispositivo de bucle para initrd. Cuando se lance make menuconfig, no se debe olvidar entrar en la configuración general (General Setup) y deshabilitar las opciones para initramfs.

### Personalizar la configuración

Una vez se encuentra una configuración, descargarla en el directorio de fuentes del núcleo y renombrarla a .config. Una vez hecho esto, lanzar make oldconfig para poner todo al día conforme a las instrucciones de arriba y personalizar la configuración antes de compilar.

`root #` `cd /usr/src/linux
`

`root #` `cp /ruta/a/config-ejemplo .config
`

`root #` `make oldconfig`

Por ahora, sólo aprete la tecla `ENTER` (o `Return`) en cada pregunta para aceptar las opciones predeterminadas...

`root #` `make menuconfig`

**Importante**

En la sección de hacking del núcleo existe una opción titulada "Are You Using A Cross Compiler?". Esto le indica a los ficheros Makefile del núcleo que antepongan "mips-linux-" (o mipsel-linux ... etc) a las órdenes gcc y as cuando se construya el núcleo. Esta opción se debería deshabilitar incluso si se está realizando una compilación cruzada. En lugar de esto, si se necesita realizar una compilación cruzada, especificar el prefijo utilizando la variable CROSS\_COMPILE tal y como se muestra en la siguiente sección.

**Importante**

Existe un problema conocido con JFS y ALSA en sistemas Octane systems en los que ALSA no funciona. Dada la naturaleza experimental de JFS en MIPS, se recomienda no utilizar JFS en este momento.

### Compilar e instalar

Ahora que el núcleo está configurado, es hora de compilarlo e instalarlo. Salga de la configuración e inicie el proceso de compilación:

**Nota**

En las máquinas de 64 bits, especificar CROSS\_COMPILE=mips64-unknown-linux-gnu- (o mips64el-... si se trata de un sistema little-endian) para utilizar el compilador de 64 bits.

Para compilar de forma nativa:

`root #` `make vmlinux modules modules_install`

Para compilar de forma cruzada en la máquina destino, ajustar mips64-unknown-linux-gnu- de la forma adecuada:

`root #` `make vmlinux modules modules_install CROSS_COMPILE=mips64-unknown-linux-gnu-`

Cuando se compile en otra máquina como una x86, utilizar las siguientes órdenes para compilar el núcleo e instalar los módulos en un directorio específica que se va a transferir a la máquina destino.

`root #` `make vmlinux modules CROSS_COMPILE=mips64-unknown-linux-gnu-
`

`root #` `make modules_install INSTALL_MOD_PATH=/algúnlugar`

**Importante**

Cuando se compile un núcleo de 64 bits para el Indy, Indigo2 (R4k), Challenge S u O2, utilizar el objetivo vmlinux.32 en lugar de vmlinux, de lo contrario, la máquina no arrancará. Esto se hace para saltarse la PROM que no reconoce el formato ELF64.

`root #` `make vmlinux.32`

**Nota**

Es posible habilitar compilaciones en paralelo usando `make -jX`, siendo X el número de compilaciones paralelas que se permite que lance el proceso. Es similar a las instrucciones anteriores sobre /etc/portage/make.conf con la variable `MAKEOPTS`.

Lo de arriba creará vmlinux.32, que es el núcleo final.

Cuando se ha terminado de compilar el núcleo, copiar la imagen a /boot/.

**Nota**

En los servidores Cobalt, el cargador de arranque espera ver una imagen de núcleo comprimida. Recuerde hacer gzip -9 al fichero una vez está en /boot/.

`root #` `cp vmlinux /boot/kernel-6.19.1-gentoo`

Para lo servidores Cobalt, comprimir la imagen del núcleo:

`root #` `gzip -9v /boot/kernel-6.19.1-gentoo`

Continúe la instalación con [Configuring the system](/wiki/Handbook:MIPS/Installation/System/es "Handbook:MIPS/Installation/System/es").

[← Instalar el sistema base](/wiki/Handbook:MIPS/Installation/Base/es "Previous part") [Inicio](/wiki/Handbook:MIPS/es "Handbook:MIPS/es") [Configurar el sistema →](/wiki/Handbook:MIPS/Installation/System/es "Next part")