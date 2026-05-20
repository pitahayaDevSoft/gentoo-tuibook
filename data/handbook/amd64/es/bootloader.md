# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- español
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Bootloader/hu "Handbook:AMD64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Bootloader/pl "Handbook:AMD64/Installation/Bootloader/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64 "Handbook:AMD64")[Installation](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[About the installation](/wiki/Handbook:AMD64/Installation/About/es "Handbook:AMD64/Installation/About/es")[Choosing the media](/wiki/Handbook:AMD64/Installation/Media/es "Handbook:AMD64/Installation/Media/es")[Configuring the network](/wiki/Handbook:AMD64/Installation/Networking/es "Handbook:AMD64/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:AMD64/Installation/Disks/es "Handbook:AMD64/Installation/Disks/es")[The stage file](/wiki/Handbook:AMD64/Installation/Stage/es "Handbook:AMD64/Installation/Stage/es")[Installing base system](/wiki/Handbook:AMD64/Installation/Base/es "Handbook:AMD64/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:AMD64/Installation/Kernel/es "Handbook:AMD64/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:AMD64/Installation/System/es "Handbook:AMD64/Installation/System/es")[Installing tools](/wiki/Handbook:AMD64/Installation/Tools/es "Handbook:AMD64/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:AMD64/Installation/Finalizing/es "Handbook:AMD64/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage introduction](/wiki/Handbook:AMD64/Working/Portage/es "Handbook:AMD64/Working/Portage/es")[USE flags](/wiki/Handbook:AMD64/Working/USE/es "Handbook:AMD64/Working/USE/es")[Portage features](/wiki/Handbook:AMD64/Working/Features/es "Handbook:AMD64/Working/Features/es")[Initscript system](/wiki/Handbook:AMD64/Working/Initscripts/es "Handbook:AMD64/Working/Initscripts/es")[Environment variables](/wiki/Handbook:AMD64/Working/EnvVar/es "Handbook:AMD64/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[Files and directories](/wiki/Handbook:AMD64/Portage/Files/es "Handbook:AMD64/Portage/Files/es")[Variables](/wiki/Handbook:AMD64/Portage/Variables/es "Handbook:AMD64/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:AMD64/Portage/Branches/es "Handbook:AMD64/Portage/Branches/es")[Additional tools](/wiki/Handbook:AMD64/Portage/Tools/es "Handbook:AMD64/Portage/Tools/es")[Custom package repository](/wiki/Handbook:AMD64/Portage/CustomTree/es "Handbook:AMD64/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:AMD64/Portage/Advanced/es "Handbook:AMD64/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[Getting started](/wiki/Handbook:AMD64/Networking/Introduction/es "Handbook:AMD64/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:AMD64/Networking/Advanced/es "Handbook:AMD64/Networking/Advanced/es")[Modular networking](/wiki/Handbook:AMD64/Networking/Modular/es "Handbook:AMD64/Networking/Modular/es")[Wireless](/wiki/Handbook:AMD64/Networking/Wireless/es "Handbook:AMD64/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:AMD64/Networking/Extending/es "Handbook:AMD64/Networking/Extending/es")[Dynamic management](/wiki/Handbook:AMD64/Networking/Dynamic/es "Handbook:AMD64/Networking/Dynamic/es")

## Contents

- [1Seleccionar un gestor de arranque](#Seleccionar_un_gestor_de_arranque)
- [2Predeterminado: GRUB](#Predeterminado:_GRUB)
  - [2.1Emerge](#Emerge)
  - [2.2Instalación](#Instalaci.C3.B3n)
    - [2.2.1Sistemas DOS/BIOS obsoleto](#Sistemas_DOS.2FBIOS_obsoleto)
    - [2.2.2Sistemas UEFI](#Sistemas_UEFI)
      - [2.2.2.1Opcional: Arranque Seguro](#Opcional:_Arranque_Seguro)
      - [2.2.2.2Depurar GRUB](#Depurar_GRUB)
  - [2.3Configuración](#Configuraci.C3.B3n)
- [3Alternativa 1: systemd-boot](#Alternativa_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Instalación](#Instalaci.C3.B3n_2)
  - [3.3Opcional: Arranque Seguro](#Opcional:_Arranque_Seguro_2)
- [4Alternativa 2: Módulo EFI](#Alternativa_2:_M.C3.B3dulo_EFI)
  - [4.1Imagen del Núcleo Unificada](#Imagen_del_N.C3.BAcleo_Unificada)
- [5Otras Alternativas](#Otras_Alternativas)
- [6Reiniciar el sistema](#Reiniciar_el_sistema)

## Seleccionar un gestor de arranque

Una vez se haya configurado el núcleo Linux configurado, instalado las herramientas del sistema y editado los ficheros de configuración, es el momento de instalar la última pieza importante de una instalación Linux: cargador de arranque.

El cargador de arranque es el responsable de arrancar el núcleo Linux en el momento del inicio. Sin él, el sistema no sabría cómo proceder cuando se pulsa el botón de encendido.

Para **amd64**, documentamos cómo configurar [GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/es#Default:_GRUB "Handbook:AMD64/Blocks/Bootloader/es") para sistemas basados ​​en DOS/BIOS, y [GRUB](#Default:_GRUB), [systemd-boot](#Alternative_1:_systemd-boot) o [EFI Stub](#Alternative_2:_efibootmgr) para sistemas UEFI.

En esta sección del Manual se realizó una delimitación entre realizar un _emerge_ del paquete del gestor de arranque e _instalar_ un gestor de arranque en un disco del sistema. Aquí el término _emerge_ será utilizado para pedirle a [Portage](/wiki/Portage "Portage") que instale el paquete para que esté disponible en el sistema. El término _install_ significará la copia de los ficheros del gestor de arranque o la modificación física de las secciones del disco apropiadas para dejar al gestor de arranque _activado y listo para operar_ en el próximo reinicio.

## Predeterminado: GRUB

De manera predeterminada, la mayoría de los sistemas Gentoo utilizan [GRUB](/wiki/GRUB/es "GRUB/es") (disponible en el paquete [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)), que es el sucesor directo de [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy"). Sin ninguna configuración adicional, GRUB soporta antiguos sistemas BIOS ("pc"). Con una pequeña configuración, necesaria antes de construir, GRUB puede soportar más de media docena de plataformas adicionales. Para obtener más información, consulte la sección [de prerequisitos](/wiki/GRUB/es#Requisitos_previos "GRUB/es") del artículo sobre [GRUB](/wiki/GRUB/es "GRUB/es").

### Emerge

Cuando se utilice un antiguo sistema con BIOS que soporte solo tablas de particiones MBR, no son necesarias configuraciones extra para realizar un emerge de GRUB:

`root #` `emerge --ask --verbose sys-boot/grub`

Una nota para los usuarios de UEFI: Al lanzar la orden de arriba, se mostrarán los valores habilitados de GRUB\_PLATFORMS antes de hacer emerge. Cuando se utilicen sistemas con UEFI, los usuarios necesitarán asegurarse de que `GRUB_PLATFORMS="efi-64"` está habilitado (ya que es el caso por defecto). Si ese no es el caso para la configuración, se necesitará añadir `GRUB_PLATFORMS="efi-64"` al fichero /etc/portage/make.conf _antes_ de hacer emerge de GRUB de modo que el paquete se construya con funcionalidad EFI:

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

Si se realizó emerge de GRUB sin habilitar previamente `GRUB_PLATFORMS="efi-64"`, se puede añadir esta línea la línea (tal y como se muestra arriba) a make.conf y se pueden recalcular las dependencias para el [conjunto de paquetes world](/wiki/World_set_(Portage) "World set (Portage)") pasando las opciones `--update --newuse` a emerge:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

El software GRUB ahora se ha instalado en el _sistema_, pero aún no se ha instalado como un _gestor de arranque_ secundario.

### Instalación

A continuación, instalar los archivos necesarios de GRUB en el directorio /boot/grub/ mediante la orden grub-install. Se supone que el primer disco (desde el que se inicia el sistema) es /dev/sda alguna de las órdenes siguientes lo hará:

#### Sistemas DOS/BIOS obsoleto

Para sistemas DOS/BIOS obsoleto:

`root #` `grub-install /dev/sda`

#### Sistemas UEFI

**Importante**

Asegúrese de que la partición EFI de sistema está montada _antes_ de lanzar grub-install. Es posible que grub-install instale el fichero GRUB para EFI (grubx64.efi) en el directorio incorrecto **sin** ofrecer _ningún_ tipo de indicación de que se ha utilizado el directorio incorrecto.

Para sistemas UEFI:

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

Tras una instalación exitosa, el resultado debe coincidir con el resultado del comando anterior. Si la salida no coincide exactamente, entonces siga con [Depurar GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/es#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/es"), en caso contrario siga con el [paso Configurar](/wiki/Handbook:AMD64/Blocks/Bootloader/es#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/es").

##### Opcional: Arranque Seguro

Para arrancar correctamente con el arranque seguro habilitado, el certificado de firma debe ser aceptado por el firmware [UEFI](/wiki/UEFI "UEFI") o debe usarse [shim](/wiki/Shim "Shim") como precargador. Shim está prefirmado con el certificado de Microsoft de terceros, aceptado por defecto por la mayoría de las placas base UEFI.

La configuración del firmware UEFI para aceptar claves personalizadas depende del proveedor del firmware, lo cual queda fuera del alcance de este manual. A continuación, se muestra cómo configurar shim. Se asume que el usuario ya ha seguido las instrucciones de las secciones anteriores para generar una clave de firma y para configurar Portage para usarla. De no ser así, vuelva primero a la sección [Instalación del núcleo](/wiki/Handbook:AMD64/Installation/Kernel/es "Handbook:AMD64/Installation/Kernel/es").

El paquete [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) instala un ejecutable EFI independiente, precompilado y firmado, siempre que el indicador USE [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") esté activado. Instale los paquetes necesarios y copie el grub independiente, Shim y MokManager en el mismo directorio de la partición del sistema EFI. Por ejemplo:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

A continuación, el registro de la clave de firma en la MOKlist de shim, requiere claves en formato DER, mientras que sbsign y el sistema de compilación del núcleo esperan claves en formato PEM. En secciones anteriores del manual se mostró un ejemplo para generar dicha clave de firma PEM; esta clave ahora debe convertirse al formato DER:

`root #` `openssl x509 -in /rura/a/kernel_key.pem -inform PEM -out /ruta/a/kernel_key.der -outform DER`

**Nota**

La ruta utilizada aquí debe ser la ruta al archivo pem que contiene el certificado que pertenece a la clave generada. En este ejemplo, tanto la clave como el certificado están en el mismo archivo pem.

Ahora el certificado convertido se puede importar a la MOKlist de Shims; este comando le pedirá que establezca una contraseña para la solicitud de importación:

`root #` `mokutil --import /ruta/a/kernel_key.der`

**Nota**

Cuando el núcleo en ejecución ya confía en el certificado que se está importando, se mostrará el mensaje "Already in kernel trusted keyring.". En tal caso, vuelva a ejecutar el comando anterior con el argumento --ignore-keyring.

A continuación, registre Shim con el firmware UEFI. En el siguiente comando, `boot-disk` y `boot-partition-id` deben reemplazarse por el identificador de disco y partición de la partición del sistema EFI:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Tenga en cuenta que esta versión independiente de grub, precompilada y firmada, lee el archivo grub.cfg desde una ubicación diferente a la habitual. En lugar de la ruta predeterminada /boot/grub/grub.cfg, el archivo de configuración debe estar en el mismo directorio que el ejecutable EFI de grub, por ejemplo, /efi/EFI/Gentoo/grub.cfg. Al usar [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) para instalar el núcleo y actualizar la configuración de grub, se puede usar la variable de entorno GRUB\_CFG para sobrescribir la ubicación habitual del archivo de configuración de grub.

Por ejemplo:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

O, vía [installkernel](/wiki/Installkernel "Installkernel"):

ARCHIVO **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Nota**

El proceso de importación no finalizará hasta que se reinicie el sistema. Tras completar todos los pasos del manual, reinicie el sistema y Shim se cargará. Encontrará la solicitud de importación registrada por mokutil. La aplicación MokManager se iniciará y solicitará la contraseña establecida al crear la solicitud de importación. Siga las instrucciones en pantalla para completar la importación del certificado. A continuación, reinicie el sistema en el menú UEFI y active la configuración de Arranque Seguro.

##### Depurar GRUB

Al depurar GRUB, hay un par de correcciones rápidas que pueden funcionar en una instalación de arranque sin tener que reiniciar en un nuevo entorno de imagen live.

En el caso de que se muestre "EFI variables are not supported on this system" en algún lugar de la pantalla, es probable que la imagen live no se haya iniciado en modo EFI y se encuentre actualmente en modo de inicio BIOS obsoleto. La solución es probar la [opción removible de GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/es#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/es") que se menciona a continuación. Esto sobrescribirá el archivo EFI ejecutable ubicado en /EFI/BOOT/BOOTX64.EFI. Al reiniciar en modo EFI, el firmware de la placa base puede ejecutar esta entrada de inicio predeterminada y ejecutar GRUB.

Si grub-install devuelve un error que dice "Could not prepare Boot variable: Read-only file system" y el entorno live se inició correctamente en modo UEFI, entonces debería ser posible volver a montar el sistema especial efivar como lectura-escritura y luego vuelva a ejecutar el [comando antes mencionado grub-install](/wiki/Handbook:AMD64/Blocks/Bootloader/es#GRUB_Install_EFI_systems_command "Handbook:AMD64/Blocks/Bootloader/es"):

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

Esto se debe a que ciertos entornos Gentoo no oficiales no montan el sistema de archivos especial EFI de forma predeterminada. Si el comando anterior no se ejecuta, reinicie usando un entorno de imagen live oficial de Gentoo en modo EFI.

Algunos fabricantes de placas base con implementaciones UEFI deficientes parecen admitir _sólo_ la ubicación del directorio /EFI/BOOT para el archivo .EFI en la partición del sistema EFI (ESP). El instalador de GRUB puede crear el archivo .EFI en esta ubicación automáticamente agregando la opción `--removable` al comando de instalación. Asegúrese de que el ESP haya sido montado antes de ejecutar el siguiente comando; suponiendo que está montado en /efi (como se definió anteriormente), ejecute:

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

Esto crea el directorio 'predeterminado' definido por la especificación UEFI y luego crea un archivo con el nombre predeterminado: BOOTX64.EFI.

### Configuración

A continuación, generar la configuración de GRUB basada en la configuración de usuario especificado en el archivo /etc/default/grub y en los guiones /etc/grub.d. En la mayoría de los casos, no se necesita ninguna configuración por parte de los usuarios ya que GRUB detectará automáticamente el núcleo que debe iniciar (la versión más alta disponible en /boot/) y cuál es el sistema de ficheros raíz. También es posible añadir argumentos del núcleo en /etc/default/grub mediante la variable GRUB\_CMDLINE\_LINUX.

Para generar la configuración final de GRUB, lance la orden grub-mkconfig:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

La salida de la orden debe mostrar que se ha encontrado al menos una imagen de Linux, como las que son necesarias para arrancar el sistema. Si se utiliza un initramfs o se utilizó genkernel para construir el núcleo, la imagen correcta initrd debería poder detectarse de esta forma. Si no es el caso, hay que ir a /boot/ y comprobar los contenidos mediante la orden ls. Si los archivos no están, se debe regresar a la configuración del núcleo y a las instrucciones de instalación.

**Consejo**

La utilidad os-prober en conjunto con GRUB para detectar otros sistemas operativos en los discos conectados. Windows 7, 8.1, 10 y otras distribuciones de Linux son detectables. Para esto escenarios de doble booteo debes hacer un emerge del paquete [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) y luego volver a ejecutar grub-mkconfig (como se muestra arriba). Si se encuentran problemas en la detección, asegúrate de leer el artículo de [GRUB](/wiki/GRUB "GRUB") entero _antes_ de pedir soporte a la comunidad de Gentoo.

## Alternativa 1: systemd-boot

Otra opción es [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), que funciona tanto en equipos OpenRC como systemd. Es un cargador en cadena ligero y funciona bien con el arranque seguro.

### Emerge

Para instalar systemd-boot, active el indicador USE [boot](https://packages.gentoo.org/useflags/boot) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") y reinstale [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (para sistemas systemd) o [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (para sistemas OpenRC):

ARCHIVO **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

ó

`root #` `emerge --ask sys-apps/systemd-utils`

### Instalación

Ahora, instale el cargador de arranque systemd-boot en la [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"):

`root #` `bootctl install`

**Importante**

Asegúrese de que la partición del sistema EFI se haya montado _antes_ de ejecutar bootctl install.

Cuando utilice este gestor de arranque, antes de reiniciar, verifique que exista una nueva entrada de arranque usando:

`root #` `bootctl list`

**Advertencia**

Las líneas de comando para cargar núcleos para las nuevas entradas de systemd-boot se leen desde /etc/kernel/cmdline o /usr/lib/kernel/cmdline. Si ninguno de estos archivos está presente, se reutiliza la línea de comandos del núcleo actualmente en ejecución (/proc/cmdline). Por lo tanto, en instalaciones nuevas, podría ocurrir que la línea de comandos del núcleo del Live CD se use accidentalmente para iniciar el nuevo núcleo. La línea de comando para cargar núcleos para las entradas registradas se pueden verificar con:

`root #` `bootctl list`

. Si ésto no muestra la línea de comando deseada, cree /etc/kernel/cmdline con la línea de comando para cargar el núcleo correcta y reinstale el núcleo.

Si no existe ninguna entrada nueva, asegúrese de que el paquete [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) se haya instalado con los indicadores USE [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") y [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") activados y vuelva a ejecutar la instalación del núcleo.

Para núcleos de distribución:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

Para un núcleo configurado y compilado de forma manual:

`root #` `make install`

**Importante**

Al instalar núcleos para systemd-boot, no se agrega ningún argumento de línea de comandos del núcleo root= de forma predeterminada. En los sistemas systemd que utilizan initramfs, los usuarios pueden confiar en [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fes "Systemd") para encontrar automáticamente la partición raíz en el arranque. De lo contrario, los usuarios deben especificar manualmente la ubicación de la partición raíz configurando root= en /etc/kernel/cmdline así como cualquier otro argumento de la línea de comandos del núcleo que deba ser usado. Y luego reinstalar el núcleo como se describe arriba.

### Opcional: Arranque Seguro

Cuando el indicador USE [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") está activado, Portage firmará automáticamente el ejecutable EFI systemd-boot. Además, bootctl install instalará automáticamente la versión firmada.

Para iniciar correctamente con el arranque seguro habilitado, el certificado utilizado debe ser aceptado por el firmware [UEFI](/wiki/UEFI "UEFI") o se debe utilizar [shim](/wiki/Shim "Shim") como precargador. Shim está prefirmado con el certificado Microsoft de terceros, aceptado de forma predeterminada por la mayoría de las placas base UEFI.

La configuración del firmware UEFI para aceptar claves personalizadas depende del proveedor del firmware, lo cual queda fuera del alcance de este manual. A continuación, se muestra cómo configurar shim. Se asume que el usuario ya ha seguido las instrucciones de las secciones anteriores para generar una clave de firma y configurado Portage para usarla. De no ser así, vuelva a la sección [Instalación del kernel](/index.php?title=Handbook:AMD64/Instalaci%C3%B3n/Kernel/es&action=edit&redlink=1 "Handbook:AMD64/Instalación/Kernel/es (page does not exist)").

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

La MOKlist de Shims MOK requiere claves en formato DER, mientras que sbsign y el sistema de compilación del núcleo esperan claves en formato PEM. En las secciones anteriores del manual se mostró un ejemplo para generar dicha clave de firma PEM; esta clave ahora debe convertirse al formato DER:

`root #` `openssl x509 -in /ruta/a/kernel_key.pem -inform PEM -out /ruta/a/kernel_key.der -outform DER`

**Nota**

La ruta utilizada aquí debe ser la ruta al archivo pem que contiene el certificado que pertenece a la clave generada. En este ejemplo, tanto la clave como el certificado están en el mismo archivo pem.

Luego, el certificado convertido se puede importar a la MOKlist de Shim:

`root #` `mokutil --import /ruta/a/kernel_key.der`

**Nota**

Cuando el núcleo en ejecución ya confía en el certificado que se está importando, se mostrará el mensaje "Already in kernel trusted keyring.". En tal caso, vuelva a ejecutar el comando anterior añadiendo el argumento --ignore-keyring.

Y finalmente registramos Shim con el firmware UEFI. En el siguiente comando, `boot-disk` y `boot-partition-id` deben reemplazarse con el disco y el identificador de partición de la partición del sistema EFI:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Nota**

El proceso de importación no finalizará hasta que se reinicie el sistema. Tras completar todos los pasos del manual, reinicie el sistema y Shim se cargará. Encontrará la solicitud de importación registrada por mokutil. La aplicación MokManager se iniciará y solicitará la contraseña establecida al crear la solicitud de importación. Siga las instrucciones en pantalla para completar la importación del certificado. A continuación, reinicie el sistema en el menú UEFI y active la configuración de Arranque Seguro.

## Alternativa 2: Módulo EFI

Los sistemas informáticos con firmware basado en UEFI técnicamente no necesitan cargadores de arranque secundarios (como por ejemplo, GRUB) para arrancar los núcleos. Los cargadores de arranque secundarios existen para _extender_ la funcionalidad del firmware UEFI durante el proceso de arranque. Usar GRUB (consulte la sección anterior) suele ser más fácil y seguro porque ofrece una manera más flexible para modificar rápidamente los parámetros del kernel en el momento del arranque.

Los administradores de sistemas que deseen adoptar un enfoque minimalista, aunque más rígido, para iniciar el sistema pueden evitar los cargadores de arranque secundarios e iniciar el núcleo Linux como un [EFI stub](/wiki/EFI_stub/es "EFI stub/es").

La aplicación [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) es una herramienta que se utiliza para interactuar con el firmware UEFI - gestor de arranque principal del sistema. Normalmente, se usa para agregar o eliminar entradas de inicio a la lista de entradas de inicio del firmware. También puede actualizar la configuración del firmware para que los núcleos Linux que se agregaron previamente como entradas de arranque puedan ejecutarse con opciones adicionales. Estas interacciones se realizan a través de estructuras de datos especiales llamadas variables EFI (de ahí la necesidad de soporte en ele núcleo para las variables EFI).

Asegúrese de revisar el artículo [EFI stub](/wiki/EFI_stub/es "EFI stub/es") del núcleo _antes_ de continuar. El núcleo debe tener opciones específicas habilitadas para que el firmware UEFI pueda iniciarlo directamente. Puede que sea necesario recompilar el núcleo para poder incorporar este soporte.

También es una buena idea consultar el artículo [efibootmgr](/wiki/Efibootmgr "Efibootmgr") para obtener información adicional.

**Nota**

Para reiterar, efibootmgr _no_ es un requisito para iniciar un sistema UEFI; simplemente es necesario agregar una entrada para un EFI stub para el núcleo en el firmware UEFI. Cuando se construye apropiadamente con soporte EFI stub, el propio núcleo Linux se puede iniciar _directamente_. Se pueden incorporar opciones adicionales para la línea de comandos del núcleo Linux (hay una opción de configuración del núcleo llamada CONFIG\_CMDLINE. De manera similar, el soporte para initramfs también se puede 'integrar' en el núcleo.

Para arrancar el núcleo directamente desde el firmware, la imagen del núcleo debe estar presente en la partición del sistema EFI. Esto se puede lograr activando el indicador USE [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/es (page does not exist)](/index.php?title=USE_flag/es&action=edit&redlink=1 "USE flag/es (page does not exist)") en [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel). Dado que no se garantiza que el arranque EFI Stub funcione en todos los sistemas UEFI, el indicador USE está desactivado en la rama estable; debe aceptarse el paquete de la rama de pruebas para que installkernel utilice esta función.

ARCHIVO **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

ARCHIVO **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel efistub

```

Después reinstale [installkernel](/wiki/Installkernel "Installkernel"), cree el directorio /efi/EFI/Gentoo y reinstale el núcleo:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

Para núcleos de la distrubución:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

Para núcleos gestionados manualmente:

`root #` `make install`

Como alternativa, cuando no se utiliza [installkernel](/wiki/Installkernel "Installkernel"), el núcleo se puede copiar manualmente a la partición del sistema EFI, nombrándolo bzImage.efi:

`root #` `mkdir -p /efi/EFI/Gentoo
`

`root #` `cp /boot/vmlinuz-* /efi/EFI/Gentoo/bzImage.efi
`

Instale el paquete efibootmgr:

`root #` `emerge --ask sys-boot/efibootmgr`

Cree una entrada de inicio llamada "Gentoo" para el EFI stub para el núcleo recién compilado dentro del firmware UEFI:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi"`

**Nota**

El uso de una barra invertida ( `\`) como separador de ruta de directorio es obligatorio cuando se utilizan definiciones UEFI.

Si se utiliza un sistema de archivos de inicio en RAM (initramfs), cópielo también a la partición del sistema EFI y luego agréguele la opción de arranque adecuada:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi" --unicode "initrd=\EFI\Gentoo\initramfs.img"`

**Consejo**

El firmware puede analizar opciones adicionales de línea de comando del núcleo al especificarlas junto con la opción initrd=... como se muestra arriba.

Una vez realizados estos cambios, cuando el sistema se reinicie, aparecerá una entrada en el menú de inicio llamada "gentoo".

### Imagen del Núcleo Unificada

Si [installkernel](/wiki/Installkernel "Installkernel") se configuró para crear e instalar imágenes del núcleo unificadas. La imagen del núcleo unificada ya debería estar instalada en el directorio EFI/Linux en la partición del sistema EFI; si este no es el caso, asegúrese de que el directorio exista y luego ejecute la instalación del núcleo nuevamente como se describe anteriormente en el manual.

Para agregar una entrada de arranque directo para la imagen unificada del núcleo instalada:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Otras Alternativas

Para otras opciones que no están cubiertas en el Manual, consulte la lista completa de [cargadores de arranque](/wiki/Bootloader/es "Bootloader/es") disponibles.

## Reiniciar el sistema

Salga del entorno chroot y desmonte todas las particiones que continúen montadas. Luego escriba la orden mágica da inicio a la auténtica prueba final: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

¡No olvide retirar la imagen live, de lo contrario puede ser arrancada de nuevo en lugar del sistema Gentoo recién instalado!

Una vez reiniciado en el nuevo entorno Gentoo, es aconsejable finalizar con [Finalizando la instalación de Gentoo](/wiki/Handbook:AMD64/Installation/Finalizing/es "Handbook:AMD64/Installation/Finalizing/es").

[← Instalar las herramientas](/wiki/Handbook:AMD64/Installation/Tools/es "Previous part") [Inicio](/wiki/Handbook:AMD64/es "Handbook:AMD64/es") [Finalizar →](/wiki/Handbook:AMD64/Installation/Finalizing/es "Next part")