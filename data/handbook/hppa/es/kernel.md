# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbuch:HPPA/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Kernel "Handbook:HPPA/Installation/Kernel (100% translated)")
- español
- [français](/wiki/Handbook:HPPA/Installation/Kernel/fr "Handbook:HPPA/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Kernel/it "Handbook:HPPA/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Kernel/hu "Handbook:HPPA/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Kernel/pl "Handbook:HPPA/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Kernel/pt-br "Manual:HPPA/Instalação/Kernel (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Kernel/ru "Handbook:HPPA/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Kernel/ta "கையேடு:HPPA/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Kernel/zh-cn "手册：HPPA/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Kernel/ja "ハンドブック:HPPA/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Kernel/ko "Handbook:HPPA/Installation/Kernel/ko (100% translated)")

[Manual HPPA](/wiki/Handbook:HPPA/es "Handbook:HPPA/es")[Instalación](/wiki/Handbook:HPPA/Full/Installation/es "Handbook:HPPA/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:HPPA/Installation/About/es "Handbook:HPPA/Installation/About/es")[Elegir los medios](/wiki/Handbook:HPPA/Installation/Media/es "Handbook:HPPA/Installation/Media/es")[Configurar la red](/wiki/Handbook:HPPA/Installation/Networking/es "Handbook:HPPA/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:HPPA/Installation/Disks/es "Handbook:HPPA/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:HPPA/Installation/Stage/es "Handbook:HPPA/Installation/Stage/es")[Instalar el sistema base](/wiki/Handbook:HPPA/Installation/Base/es "Handbook:HPPA/Installation/Base/es")Configurar el núcleo[Configurar el sistema](/wiki/Handbook:HPPA/Installation/System/es "Handbook:HPPA/Installation/System/es")[Instalar las herramientas](/wiki/Handbook:HPPA/Installation/Tools/es "Handbook:HPPA/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:HPPA/Installation/Bootloader/es "Handbook:HPPA/Installation/Bootloader/es")[Terminar](/wiki/Handbook:HPPA/Installation/Finalizing/es "Handbook:HPPA/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:HPPA/Full/Working/es "Handbook:HPPA/Full/Working/es")[Introducción a Portage](/wiki/Handbook:HPPA/Working/Portage/es "Handbook:HPPA/Working/Portage/es")[Ajustes USE](/wiki/Handbook:HPPA/Working/USE/es "Handbook:HPPA/Working/USE/es")[Características de Portage](/wiki/Handbook:HPPA/Working/Features/es "Handbook:HPPA/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:HPPA/Working/Initscripts/es "Handbook:HPPA/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:HPPA/Working/EnvVar/es "Handbook:HPPA/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:HPPA/Full/Portage/es "Handbook:HPPA/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:HPPA/Portage/Files/es "Handbook:HPPA/Portage/Files/es")[Variables](/wiki/Handbook:HPPA/Portage/Variables/es "Handbook:HPPA/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:HPPA/Portage/Branches/es "Handbook:HPPA/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:HPPA/Portage/Tools/es "Handbook:HPPA/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:HPPA/Portage/CustomTree/es "Handbook:HPPA/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:HPPA/Portage/Advanced/es "Handbook:HPPA/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:HPPA/Full/Networking/es "Handbook:HPPA/Full/Networking/es")[Comenzar](/wiki/Handbook:HPPA/Networking/Introduction/es "Handbook:HPPA/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:HPPA/Networking/Advanced/es "Handbook:HPPA/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:HPPA/Networking/Modular/es "Handbook:HPPA/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:HPPA/Networking/Wireless/es "Handbook:HPPA/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:HPPA/Networking/Extending/es "Handbook:HPPA/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:HPPA/Networking/Dynamic/es "Handbook:HPPA/Networking/Dynamic/es")

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
  - [3.1Distribution kernels](#Distribution_kernels)
    - [3.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [3.1.2Installing a distribution kernel](#Installing_a_distribution_kernel)
    - [3.1.3Upgrading and cleaning up](#Upgrading_and_cleaning_up)
    - [3.1.4Post-install/upgrade tasks](#Post-install.2Fupgrade_tasks)
      - [3.1.4.1Manually rebuilding the initramfs or Unified Kernel Image](#Manually_rebuilding_the_initramfs_or_Unified_Kernel_Image)
  - [3.2Alternativa: Configuración manual](#Alternativa:_Configuraci.C3.B3n_manual)
  - [3.3Instalar las fuentes del núcleo](#Instalar_las_fuentes_del_n.C3.BAcleo)
    - [3.3.1Proceso modprobed-db](#Proceso_modprobed-db)
    - [3.3.2Proceso manual](#Proceso_manual)
      - [3.3.2.1Habilitar las opciones requeridas](#Habilitar_las_opciones_requeridas)
    - [3.3.3Opcional: módulos del núcleo firmados](#Opcional:_m.C3.B3dulos_del_n.C3.BAcleo_firmados)
  - [3.4Compilación e instalación](#Compilaci.C3.B3n_e_instalaci.C3.B3n)

## Opcional: Instalar firmware y/o microcódigo

### Firmware

#### Sugerido: Linux Firmware

En muchos sistemas, se requiere firmware no FOSS para el funcionamiento de ciertos dispositivos. El paquete [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) contiene firmware para muchos dispositivos, pero no para todos.

**Consejo**

La mayoría de las tarjetas inalámbricas y GPU,s. requieren firmware para funcionar.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Nota**

La instalación de determinados paquetes de firmware suele requerir la aceptación de las licencias de firmware asociadas. Si es necesario, visite la [sección de manejo de licencias](/wiki/Handbook:HPPA/Working/Portage/es#Licenses "Handbook:HPPA/Working/Portage/es") del Manual para obtener ayuda sobre cómo aceptar licencias.

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

[Propuesta de automatización total: núcleos de la distribución](/wiki/Handbook:HPPA/Installation/Kernel/es#Distribution_kernels "Handbook:HPPA/Installation/Kernel/es")Un [Núcleo de la Distribución](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") se usa para configurar, compilar e instalar automáticamente el núcleo Linux, sus módulos asociados y (opcionalmente, pero habilitado por defecto) un archivo initramfs. Las actualizaciones futuras del núcleo están completamente automatizadas, ya que se manejan a través del administrador de paquetes, como cualquier otro paquete del sistema. Es posible [proporcionar un archivo de configuración del núcleo personalizado](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") si es necesaria la personalización. Este es el proceso menos complicado y es perfecto para los nuevos usuarios de Gentoo debido a que funciona de forma inmediata y ofrece una participación mínima por parte del administrador del sistema.[La manera completamente manual](/wiki/Handbook:HPPA/Installation/Kernel/es#Alternative:_Manual_configuration "Handbook:HPPA/Installation/Kernel/es")Las nuevas fuentes del núcleo se instalan a través del administrador de paquetes del sistema. El núcleo se configura, compila e instala manualmente utilizando eselect kernel y una serie de comandos make. Las futuras actualizaciones del núcleo repiten el proceso manual de configuración, compilación e instalación de los archivos del núcleo. Este es el proceso más complejo, pero ofrece el máximo control sobre el proceso de actualización del núcleo.[Estrategia híbrida: Genkernel](/wiki/Handbook:HPPA/Installation/Kernel/es#Alternative:_Genkernel "Handbook:HPPA/Installation/Kernel/es")Aquí usamos el término híbrido, pero tenga en cuenta que las fuentes de dist-kernel y del manual incluyen métodos para lograr el mismo objetivo. Las nuevas fuentes del núcleo se instalan a través del administrador de paquetes del sistema. Los administradores del sistema pueden usar la herramienta genkernel de Gentoo para configurar, compilar, e instalar automáticamente el núcleo Linux, sus módulos asociados y (opcionalmente, pero _**no**_ habilitado por defecto) un archivo initramfs archivo. Es posible proporcionar un archivo de configuración del núcleo personalizado si es necesaria la personalización. La configuración, compilación e instalación futuras del núcleo requieren la participación del administrador del sistema en la forma de ejecutar eselect kernel, genkernel y potencialmente otros comandos para cada actualización. Esta opción solo debe considerarse para usuarios que saben que necesitan genkernel

El eje alrededor del cual se construyen todas las distribuciones es el núcleo Linux. Es la capa entre los programas del usuario y el hardware del sistema. Aunque el manual proporciona a sus usuarios varias fuentes posibles del núcleo, hay disponible una lista más completa y con descripciones más detalladas en la [Página de descripción general del núcleo](/wiki/Kernel/Overview "Kernel/Overview").

**Consejo**

Tareas de instalación del núcleo como copiar la imagen del núcleo a /boot o la [Partición del Sistema EFI](/index.php?title=Partici%C3%B3n_del_Sistema_EFI&action=edit&redlink=1 "Partición del Sistema EFI (page does not exist)"), generar un [initramfs](/wiki/Initramfs "Initramfs") y/o [Imagen unificada del núcleo](/index.php?title=Imagen_unificada_del_n%C3%BAcleo&action=edit&redlink=1 "Imagen unificada del núcleo (page does not exist)"), que actualiza la configuración del gestor de arranque, se puede automatizar con [installkernel](/wiki/Installkernel "Installkernel"). Es posible que los usuarios deseen configurar e instalar [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) antes de continuar. Consulte la [Sección de instalación del núcleo a continuación](/index.php?title=Handbook:HPPA/Instalaci%C3%B3n/Kernel&action=edit&redlink=1 "Handbook:HPPA/Instalación/Kernel (page does not exist)") para obtener más información.

### Distribution kernels

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ are ebuilds that cover the complete process of unpacking, configuring, compiling, and installing the kernel. The primary advantage of this method is that the kernels are updated to new versions by the package manager as part of @world upgrade. This requires no more involvement than running an emerge command. Distribution kernels default to a configuration supporting the majority of hardware, however two mechanisms are offered for customization: savedconfig and config snippets. See the project page for [more details on configuration.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Optional: Signed kernel modules

The kernel modules in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) are already signed. To sign the modules of kernels built from source enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/es "/etc/portage/make.conf/es"):

ARCHIVO **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"

# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.

```

If MODULES\_SIGN\_KEY is not specified the kernel build system will generate a key, it will be stored in /usr/src/linux-x.y.z/certs. It is recommended to manually generate a key to ensure that it will be the same for each kernel release. A key may be generated with:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**Nota**

The MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may be different files. For this example the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

OpenSSL will ask some questions about the user generating the key, it is recommended to fill in these questions as detailed as possible.

Store the key in a safe location, at the very least the key should be readable only by the root user. Verify this with:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

If this outputs anything other then the above, correct the permissions with:

`root #` `chown root:root kernel_key.pem`

`root #` `chmod 400 kernel_key.pem`

#### Installing a distribution kernel

To build a kernel with Gentoo patches from source, type:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

System administrators who want to avoid compiling the kernel sources locally can instead use precompiled kernel images:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Importante**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:HPPA/Installation/Kernel#Initramfs.2Fes "Handbook:HPPA/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:HPPA/Installation/Kernel#Initramfs.2Fes "Handbook:HPPA/Installation/Kernel").

#### Upgrading and cleaning up

Once the kernel is installed, the package manager will automatically update it to newer versions. The previous versions will be kept until the package manager is requested to clean up stale packages. To reclaim disk space, stale packages can be trimmed by periodically running emerge with the `--depclean` option:

`root #` `emerge --depclean`

Alternatively, to specifically clean up old kernel versions:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Consejo**

By design, emerge only removes the kernel build directory. It does not actually remove the kernel modules, nor the installed kernel image. To completely clean-up old kernels, the [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) tool may be used.

#### Post-install/upgrade tasks

An upgrade of a distribution kernel is capable of triggering an automatic rebuild for external kernel modules installed by other packages (for example: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) or [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). This automated behaviour is enabled by enabling the [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") USE flag. When required, this same flag will also trigger re-generation of the [initramfs](/wiki/Initramfs "Initramfs").

It is highly recommended to enable this flag globally via /etc/portage/make.conf when using a distribution kernel:

ARCHIVO **`/etc/portage/make.conf`** **Enabling USE=dist-kernel**

```
USE="dist-kernel"

```

##### Manually rebuilding the initramfs or Unified Kernel Image

If required, manually trigger such rebuilds by, after a kernel upgrade, executing:

`root #` `emerge --ask @module-rebuild`

If any kernel modules (e.g. ZFS) are needed at early boot, rebuild the initramfs afterward via:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:HPPA/Installation/System/es "Handbook:HPPA/Installation/System/es").

### Alternativa: Configuración manual

### Instalar las fuentes del núcleo

Al instalar y compilar el núcleo para sistemas basados en hppa, Gentoo recomienda el paquete [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Elija una fuente del núcleo adecuada e instálela usando emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

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

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:HPPA/Installation/Kernel#Distribution_kernels.2Fes "Handbook:HPPA/Installation/Kernel").

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:HPPA/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fes "Handbook:HPPA/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:HPPA/Installation/Kernel#Compiling_and_installing.2Fes "Handbook:HPPA/Installation/Kernel")

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

### Compilación e instalación

**Importante**

Para compilar un núcleo de 64 bits, primero emerge kgcc64. Sin embargo, no se recomienda ejecutar un núcleo de 64 bits. Solo ejecute un núcleo de 64 bits si el sistema tiene más de 4GB de RAM o si el servidor lo requiere, es decir, en el A500.

Ahora que el núcleo está configurado, es hora de compilarlo e instalarlo. Salga de la configuración e inicie el proceso de compilación:

`root #` `make && make modules_install`

If building a 64-bit kernel, do this instead (it's necessary even for native builds, see [here](https://www.spinics.net/lists/kernel/msg3987729.html)):

`root #` `CROSS_COMPILE=hppa64-unknown-linux-gnu- make && CROSS_COMPILE=hppa64-unknown-linux-gnu- make modules_install`

**Nota**

Es posible habilitar compilaciones paralelas usando `make -jX` siendo X el número de tareas paralelas que el proceso de compilación puede lanzar. Esto es similar a las instrucciones sobre /etc/portage/make.conf anteriores, con la variable `MAKEOPTS`.

Cuando el núcleo haya terminado de compilarse, copie la imagen del kernel en /boot/. Use el nombre que sea apropiado para la elección del kernel y recuérdelo, ya que lo necesitará más adelante cuando configure el gestor de arranque. Recuerde reemplazar kernel-6.19.1-gentoo con el nombre y la versión del núcleo instalado.

`root #` `cp vmlinux /boot/kernel-6.19.1-gentoo`

Continúe la instalación con [Configuring the system](/wiki/Handbook:HPPA/Installation/System/es "Handbook:HPPA/Installation/System/es").

[← Instalar el sistema base](/wiki/Handbook:HPPA/Installation/Base/es "Previous part") [Inicio](/wiki/Handbook:HPPA/es "Handbook:HPPA/es") [Configurar el sistema →](/wiki/Handbook:HPPA/Installation/System/es "Next part")