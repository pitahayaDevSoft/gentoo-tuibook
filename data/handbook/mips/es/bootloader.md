# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- español
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS "Handbook:MIPS")[Installation](/wiki/Handbook:MIPS/Full/Installation "Handbook:MIPS/Full/Installation")[About the installation](/wiki/Handbook:MIPS/Installation/About/es "Handbook:MIPS/Installation/About/es")[Choosing the media](/wiki/Handbook:MIPS/Installation/Media/es "Handbook:MIPS/Installation/Media/es")[Configuring the network](/wiki/Handbook:MIPS/Installation/Networking/es "Handbook:MIPS/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:MIPS/Installation/Disks/es "Handbook:MIPS/Installation/Disks/es")[The stage file](/wiki/Handbook:MIPS/Installation/Stage/es "Handbook:MIPS/Installation/Stage/es")[Installing base system](/wiki/Handbook:MIPS/Installation/Base/es "Handbook:MIPS/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:MIPS/Installation/Kernel/es "Handbook:MIPS/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:MIPS/Installation/System/es "Handbook:MIPS/Installation/System/es")[Installing tools](/wiki/Handbook:MIPS/Installation/Tools/es "Handbook:MIPS/Installation/Tools/es")Configuring the bootloader[Finalizing](/wiki/Handbook:MIPS/Installation/Finalizing/es "Handbook:MIPS/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:MIPS/Full/Working "Handbook:MIPS/Full/Working")[Portage introduction](/wiki/Handbook:MIPS/Working/Portage/es "Handbook:MIPS/Working/Portage/es")[USE flags](/wiki/Handbook:MIPS/Working/USE/es "Handbook:MIPS/Working/USE/es")[Portage features](/wiki/Handbook:MIPS/Working/Features/es "Handbook:MIPS/Working/Features/es")[Initscript system](/wiki/Handbook:MIPS/Working/Initscripts/es "Handbook:MIPS/Working/Initscripts/es")[Environment variables](/wiki/Handbook:MIPS/Working/EnvVar/es "Handbook:MIPS/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:MIPS/Full/Portage "Handbook:MIPS/Full/Portage")[Files and directories](/wiki/Handbook:MIPS/Portage/Files/es "Handbook:MIPS/Portage/Files/es")[Variables](/wiki/Handbook:MIPS/Portage/Variables/es "Handbook:MIPS/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:MIPS/Portage/Branches/es "Handbook:MIPS/Portage/Branches/es")[Additional tools](/wiki/Handbook:MIPS/Portage/Tools/es "Handbook:MIPS/Portage/Tools/es")[Custom package repository](/wiki/Handbook:MIPS/Portage/CustomTree/es "Handbook:MIPS/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:MIPS/Portage/Advanced/es "Handbook:MIPS/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:MIPS/Full/Networking "Handbook:MIPS/Full/Networking")[Getting started](/wiki/Handbook:MIPS/Networking/Introduction/es "Handbook:MIPS/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:MIPS/Networking/Advanced/es "Handbook:MIPS/Networking/Advanced/es")[Modular networking](/wiki/Handbook:MIPS/Networking/Modular/es "Handbook:MIPS/Networking/Modular/es")[Wireless](/wiki/Handbook:MIPS/Networking/Wireless/es "Handbook:MIPS/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:MIPS/Networking/Extending/es "Handbook:MIPS/Networking/Extending/es")[Dynamic management](/wiki/Handbook:MIPS/Networking/Dynamic/es "Handbook:MIPS/Networking/Dynamic/es")

[Handbook:MIPS/Blocks/Bootloader/es](/index.php?title=Handbook:MIPS/Blocks/Bootloader/es&action=edit&redlink=1 "Handbook:MIPS/Blocks/Bootloader/es (page does not exist)")

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

Una vez reiniciado en el nuevo entorno Gentoo, es aconsejable finalizar con [Finalizando la instalación de Gentoo](/wiki/Handbook:MIPS/Installation/Finalizing/es "Handbook:MIPS/Installation/Finalizing/es").

[← Instalar las herramientas](/wiki/Handbook:MIPS/Installation/Tools/es "Previous part") [Inicio](/wiki/Handbook:MIPS/es "Handbook:MIPS/es") [Finalizar →](/wiki/Handbook:MIPS/Installation/Finalizing/es "Next part")