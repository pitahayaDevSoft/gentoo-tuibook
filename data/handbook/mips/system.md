# System

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/System/de "Handbuch:MIPS/Installation/System (100% translated)")
- English
- [español](/wiki/Handbook:MIPS/Installation/System/es "Manual de Gentoo: MIPS/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/System/fr "Handbook:MIPS/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/System/it "Handbook:MIPS/Installation/System/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/System/hu "Handbook:MIPS/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/System/pl "Handbook:MIPS/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/System/pt-br "Manual:MIPS/Instalação/Sistema (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/System/ru "Handbook:MIPS/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/System/ta "கையேடு:MIPS/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/System/zh-cn "手册：MIPS/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/System/ja "ハンドブック:MIPS/インストール/システム (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/System/ko "Handbook:MIPS/Installation/System/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS "Handbook:MIPS")[Installation](/wiki/Handbook:MIPS/Full/Installation "Handbook:MIPS/Full/Installation")[About the installation](/wiki/Handbook:MIPS/Installation/About "Handbook:MIPS/Installation/About")[Choosing the media](/wiki/Handbook:MIPS/Installation/Media "Handbook:MIPS/Installation/Media")[Configuring the network](/wiki/Handbook:MIPS/Installation/Networking "Handbook:MIPS/Installation/Networking")[Preparing the disks](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks")[The stage file](/wiki/Handbook:MIPS/Installation/Stage "Handbook:MIPS/Installation/Stage")[Installing base system](/wiki/Handbook:MIPS/Installation/Base "Handbook:MIPS/Installation/Base")[Configuring the kernel](/wiki/Handbook:MIPS/Installation/Kernel "Handbook:MIPS/Installation/Kernel")Configuring the system[Installing tools](/wiki/Handbook:MIPS/Installation/Tools "Handbook:MIPS/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader")[Finalizing](/wiki/Handbook:MIPS/Installation/Finalizing "Handbook:MIPS/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:MIPS/Full/Working "Handbook:MIPS/Full/Working")[Portage introduction](/wiki/Handbook:MIPS/Working/Portage "Handbook:MIPS/Working/Portage")[USE flags](/wiki/Handbook:MIPS/Working/USE "Handbook:MIPS/Working/USE")[Portage features](/wiki/Handbook:MIPS/Working/Features "Handbook:MIPS/Working/Features")[Initscript system](/wiki/Handbook:MIPS/Working/Initscripts "Handbook:MIPS/Working/Initscripts")[Environment variables](/wiki/Handbook:MIPS/Working/EnvVar "Handbook:MIPS/Working/EnvVar")[Working with Portage](/wiki/Handbook:MIPS/Full/Portage "Handbook:MIPS/Full/Portage")[Files and directories](/wiki/Handbook:MIPS/Portage/Files "Handbook:MIPS/Portage/Files")[Variables](/wiki/Handbook:MIPS/Portage/Variables "Handbook:MIPS/Portage/Variables")[Mixing software branches](/wiki/Handbook:MIPS/Portage/Branches "Handbook:MIPS/Portage/Branches")[Additional tools](/wiki/Handbook:MIPS/Portage/Tools "Handbook:MIPS/Portage/Tools")[Custom package repository](/wiki/Handbook:MIPS/Portage/CustomTree "Handbook:MIPS/Portage/CustomTree")[Advanced features](/wiki/Handbook:MIPS/Portage/Advanced "Handbook:MIPS/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:MIPS/Full/Networking "Handbook:MIPS/Full/Networking")[Getting started](/wiki/Handbook:MIPS/Networking/Introduction "Handbook:MIPS/Networking/Introduction")[Advanced configuration](/wiki/Handbook:MIPS/Networking/Advanced "Handbook:MIPS/Networking/Advanced")[Modular networking](/wiki/Handbook:MIPS/Networking/Modular "Handbook:MIPS/Networking/Modular")[Wireless](/wiki/Handbook:MIPS/Networking/Wireless "Handbook:MIPS/Networking/Wireless")[Adding functionality](/wiki/Handbook:MIPS/Networking/Extending "Handbook:MIPS/Networking/Extending")[Dynamic management](/wiki/Handbook:MIPS/Networking/Dynamic "Handbook:MIPS/Networking/Dynamic")

## Contents

- [1Filesystem information](#Filesystem_information)
  - [1.1Filesystem labels and UUIDs](#Filesystem_labels_and_UUIDs)
  - [1.2Partition labels and UUIDs](#Partition_labels_and_UUIDs)
  - [1.3About fstab](#About_fstab)
  - [1.4Creating the fstab file](#Creating_the_fstab_file)
    - [1.4.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2Networking information](#Networking_information)
  - [2.1Hostname](#Hostname)
    - [2.1.1Set the hostname (OpenRC or systemd)](#Set_the_hostname_.28OpenRC_or_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Network](#Network)
    - [2.2.1DHCP via dhcpcd (any init system)](#DHCP_via_dhcpcd_.28any_init_system.29)
    - [2.2.2NetworkManger (any init system)](#NetworkManger_.28any_init_system.29)
    - [2.2.3netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.3.1Configuring the network](#Configuring_the_network)
      - [2.2.3.2Automatically start networking at boot](#Automatically_start_networking_at_boot)
    - [2.2.4systemd-networkd (systemd)](#systemd-networkd_.28systemd.29)
    - [2.2.5Optional: Networking tools](#Optional:_Networking_tools)
      - [2.2.5.1Installing a PPPoE client](#Installing_a_PPPoE_client)
      - [2.2.5.2Install wireless networking tools](#Install_wireless_networking_tools)
  - [2.3The hosts file](#The_hosts_file)
- [3System information](#System_information)
  - [3.1Root password](#Root_password)
  - [3.2Init and boot configuration](#Init_and_boot_configuration)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Filesystem information\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-1 "Edit section: ")\]

### Filesystem labels and UUIDs\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-2 "Edit section: ")\]

Both MBR (BIOS) and GPT include support for _filesystem_ labels and _filesystem_ UUIDs. These attributes can be defined in /etc/fstab as alternatives for the mount command to use when attempting to find and mount block devices. Filesystem labels and UUIDs are identified by the LABEL and UUID prefix and can be viewed with the blkid command:

`root #` `blkid`

**Warning**

If the filesystem inside a partition is wiped, then the filesystem label and the UUID values will be subsequently altered or removed.

For uniqueness, readers who are using MBR-style partition tables are advised to use UUIDs rather than labels to specify mountable volumes in /etc/fstab.

**Important**

UUIDs of the filesystem on a LVM volume and its LVM snapshots are identical, therefore using UUIDs to mount LVM volumes should be avoided.

### Partition labels and UUIDs\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-3 "Edit section: ")\]

Systems with GPT disklabel support offer additional 'robust' options to define partitions in /etc/fstab. Partition labels and partition UUIDs can be used to identify the block device's individual partition(s), regardless of what filesystem has been chosen for the partition itself. Partition labels and UUIDs are identified by the PARTLABEL and/or PARTUUID prefixes and can be viewed nicely in the terminal by running the blkid command.

Output for an **amd64** EFI system using the Discoverable Partition Specification UUIDs may like the following:

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

While not always true for partition labels, using a UUID to identify a partition in fstab provides a guarantee that the bootloader will not be confused when looking for a certain volume, even if the _filesystem_ is changed or re-written in the future. Using the older default block device files (/dev/sd\*N) for defining the partitions in fstab is risky for systems that have SATA block devices regularly added or removed.

The naming for block device files depends on a number of factors, including how and in what order the disks are attached to the system. They also could show up in a different order depending on which of the devices are detected by the kernel first during the early boot process. With this being stated, unless the system administrator intends to constantly fiddle with the disk ordering, using default block device files is a simple and straightforward approach.

### About fstab\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-4 "Edit section: ")\]

Under Linux, all partitions used by the system must be listed in [/etc/fstab](/wiki//etc/fstab "/etc/fstab"). This file contains the mount points of those partitions (where they are seen in the file system structure), how they should be mounted and with what special options (automatically or not, whether users can mount them or not, etc.)

### Creating the fstab file\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-5 "Edit section: ")\]

**Note**

If the init system being used is systemd, the partition UUIDs conform to the Discoverable Partition Specification as given in [Preparing the disks](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks"), and the system uses UEFI, then creating an fstab can be skipped, since systemd auto-mounts partitions that follow the spec.

The /etc/fstab file uses a table-like syntax. Every line consists of six fields, separated by whitespace (space(s), tabs, or a mixture of the two). Each field has its own meaning:

1. The first field shows the block special device or remote filesystem to be mounted. Several kinds of device identifiers are available for block special device nodes, including paths to device files, filesystem labels and UUIDs, and partition labels and UUIDs.
2. The second field shows the mount point at which the partition should be mounted.
3. The third field shows the type of filesystem used by the partition.
4. The fourth field shows the mount options used by mount when it wants to mount the partition. As every filesystem has its own mount options, so system admins are encouraged to read the mount man page (man mount) for a full listing. Multiple mount options are comma-separated.
5. The fifth field is used by dump to determine if the partition needs to be dumped or not. This can generally be left as `0` (zero).
6. The sixth field is used by fsck to determine the order in which filesystems should be checked if the system wasn't shut down properly. The root filesystem should have `1` while the rest should have `2` (or `0` if a filesystem check is not necessary).

**Important**

The default /etc/fstab file provided in Gentoo stage files is _not_ a valid fstab file but instead a template that can be used to enter in relevant values.

`root #` `nano /etc/fstab`

#### DOS/Legacy BIOS systems\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-6 "Edit section: ")\]

Let us take a look at how to write down the options for the /boot partition. This is just an example, and should be modified according to the partitioning decisions made earlier in the installation.
In the **mips** partitioning example, /boot is usually the /dev/sda1 partition, with xfs recommended for the filesystem. It needs to be checked during boot, so we would write down:

FILE **`/etc/fstab`** **An example DOS/Legacy BIOS boot line for /etc/fstab**

```
# Adjust for any formatting differences and/or additional partitions created from the "Preparing the disks" step
/dev/sda1   /boot     ext2    defaults        0 2

```

Some system administrators want the /boot partition to not be mounted automatically to improve their system's security. Those people should substitute the `defaults` with `noauto`. This does mean that those users will need to manually mount this partition every time they want to use it.

Add the rules that match the previously decided partitioning scheme and append rules for devices such as CD-ROM drive(s), and of course, if other partitions or drives are used, for those too.

Below is a more elaborate example of an /etc/fstab file:

FILE **`/etc/fstab`** **A full /etc/fstab example for a DOS/Legacy BIOS system**

```
# Adjust for any formatting differences and/or additional partitions created from the "Preparing the disks" step
/dev/sda1   /boot        ext2    defaults    0 2
/dev/sda10   none         swap    sw                   0 0
/dev/sda5   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

#### DPS UEFI PARTUUID\[ [edit](/index.php?title=Handbook:MIPS/Blocks/Fstab&action=edit&section=T-1 "Edit section: ")\]

Below is an example of an /etc/fstab file for a disk formatted with a GPT disklabel and Discoverable Partition Specification (DPS) UUIDs set for UEFI firmware:

FILE **`/etc/fstab`** **GPT disklabel DPS PARTUUID fstab example**

```
# Adjust any formatting difference and additional partitions created from the "Preparing the disks" step.
# This example shows a GPT disklabel with Discoverable Partition Specification (DSP) UUID set:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b                                  0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=   /           xfs    defaults,noatime              0 1

```

When `auto` is used in the third field, it makes the mount command guess what the filesystem would be. This is recommended for removable media as they can be created with one of many filesystems. The `user` option in the fourth field makes it possible for non-root users to mount the CD.

To improve performance, most users would want to add the `noatime` mount option, which results in a faster system since access times are not registered (those are not needed generally anyway). This is also recommended for systems with solid state drives (SSDs). Users may wish to consider `lazytime` instead.

**Tip**

Due to degradation in performance, defining the `discard` mount option in /etc/fstab is not recommended. It is generally better to schedule block discards on a periodic basis using a job scheduler such as cron or a timer (systemd). See [Periodic fstrim jobs](/wiki/SSD#Periodic_fstrim_jobs "SSD") for more information.

Double-check the /etc/fstab file, then save and quit to continue.

## Networking information\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-8 "Edit section: ")\]

It is important to note the following sections are provided to help the reader quickly setup their system to partake in a local area network.

For systems running OpenRC, a more detailed reference for network setup is available in the [advanced network configuration](/wiki/Handbook:MIPS/Networking/Introduction "Handbook:MIPS/Networking/Introduction") section, which is covered near the end of the handbook. Systems with more specific network needs may need to skip ahead, then return here to continue with the rest of the installation.

For more specific systemd network setup, please review see the [networking portion](/wiki/Systemd#Network "Systemd") of the [systemd](/wiki/Systemd "Systemd") article.

### Hostname\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-9 "Edit section: ")\]

One of the choices the system administrator has to make is name their PC. This seems to be quite easy, but lots of users are having difficulties finding the appropriate name for the hostname. To speed things up, know that the decision is not final - it can be changed afterwards. In the examples below, the hostname _tux_ is used.

#### Set the hostname (OpenRC or systemd)\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-10 "Edit section: ")\]

`root #` `echo tux > /etc/hostname`

#### systemd\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-11 "Edit section: ")\]

To set the system hostname for a system currently _running_ systemd, the hostnamectl utility may be used. During the installation process however, [systemd-firstboot](/wiki/Handbook:MIPS/Installation/System#Init_and_boot_configuration_systemd "Handbook:MIPS/Installation/System") command must be used instead (see later on in handbook).

For setting the hostname to "tux", one would run:

`root #` `hostnamectl hostname tux`

View help by running hostnamectl --help or man 1 hostnamectl.

### Network\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-12 "Edit section: ")\]

The Gentoo Linux installation process involves configuring networking in _two parts_. The configuration done in earlier steps enables networking for the _live environment_. Now it is time to apply the network configuration that will be used in the _installed environment_.

More detailed information about networking, including advanced topics like bonding, bridging, 802.1Q VLANs or wireless networking is covered in the [advanced network configuration](/wiki/Handbook:MIPS/Networking/Introduction "Handbook:MIPS/Networking/Introduction") section.

**Important**

Users should avoid using multiple network management daemons (eg. Network Manager, netifrc, systemd-networkd).

**Note**

This section is not exhaustive and covers a few common methods.

#### DHCP via dhcpcd (any init system)\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-13 "Edit section: ")\]

**Tip**

Ethernet users with a home router generally prefer this option.

DHCP provides a reliable way for systems to obtain an IP address automatically. Most LANs operate a DHCP server, usually integrated with a router which provides internet access.

Systems that do not use WiFi can use dhcpcd directly to configure the network automatically. Alternatively, DHCP can be used by the other methods described in the next sections, which either provide their own DHCP client implementation, or use [net-misc/dhcpcd](https://packages.gentoo.org/packages/net-misc/dhcpcd).

To install:

`root #` `emerge --ask net-misc/dhcpcd`

To enable the service on OpenRC systems:

`root #` `rc-update add dhcpcd default
`

To enable the service on systemd systems:

`root #` `systemctl enable dhcpcd`

The next time the system boots, dhcpcd should obtain an IP address from the DHCP server. See the [Dhcpcd](/wiki/Dhcpcd "Dhcpcd") article for more details.

#### NetworkManger (any init system)\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-14 "Edit section: ")\]

**Tip**

WiFi users generally prefer this option.

[NetworkManager](/wiki/NetworkManager "NetworkManager") provides a simple method to manage networks using Ethernet, WiFi, PPPoE or celluar via CLI, TUI or GUI.

First add the [networkmanager](https://packages.gentoo.org/useflags/networkmanager) [USE flag](/wiki/USE_flag "USE flag") USE flag to /etc/portage/make.conf:

FILE **`/etc/portage/make.conf`**

```
...
USE="${USE} networkmanager"
...

```

Non desktop profiles will also need to add:

FILE **`/etc/portage/package.use/networkmanager`**

```
net-wireless/wpa_supplicant dbus

```

Next, emerge the package:

`root #` `emerge --ask net-misc/networkmanager`

Finally, enable the service to start at boot time:

OpenRC:

`root #` `rc-update add NetworkManager default`

systemd:

`root #` `systemctl enable NetworkManager`

For more information see the [NetworkManager](/wiki/NetworkManager "NetworkManager") article.

#### netifrc (OpenRC)\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-15 "Edit section: ")\]

##### Configuring the network\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-16 "Edit section: ")\]

All networking information is gathered in /etc/conf.d/net. It uses a straightforward - yet perhaps not intuitive - syntax. Do not fear! Everything is explained below. A fully commented example that covers many different configurations is available in /usr/share/doc/netifrc-\*/net.example.bz2.

First install [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc):

`root #` `emerge --ask --noreplace net-misc/netifrc`

DHCP is used by default. For DHCP to work, a DHCP client needs to be installed.

If the network connection needs to be configured because of specific DHCP options or because DHCP is not used at all, then open /etc/conf.d/net:

`root #` `nano /etc/conf.d/net`

Set both config\_eth0 and routes\_eth0 to enter IP address information and routing information:

**Note**

This assumes that the network interface will be called eth0. This is, however, very system dependent. It is recommended to assume that the interface is named the same as the interface name when booted from the installation media if the installation media is sufficiently recent. More information can be found in the [Network interface naming](/wiki/Handbook:MIPS/Networking/Advanced#Network_interface_naming "Handbook:MIPS/Networking/Advanced") section.

FILE **`/etc/conf.d/net`** **Static IP definition**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

To use DHCP, define config\_eth0:

FILE **`/etc/conf.d/net`** **DHCP definition**

```
config_eth0="dhcp"

```

Please read /usr/share/doc/netifrc-\*/net.example.bz2 for a list of additional configuration options. Be sure to also read up on the DHCP client man page if specific DHCP options need to be set.

If the system has several network interfaces, then repeat the above steps for config\_eth1, config\_eth2, etc.

Now save the configuration and exit to continue.

##### Automatically start networking at boot\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-17 "Edit section: ")\]

To have the network interfaces activated at boot, they need to be added to the default runlevel.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

If the system has several network interfaces, then the appropriate net.\* files need to be created just like we did with net.eth0.

If, after booting the system, it is discovered the network interface name (which is currently documented as `eth0`) was wrong, then execute the following steps to rectify:

1. Update the /etc/conf.d/net file with the correct interface name (like `enp3s0` or `enp5s0`, instead of `eth0`).
2. Create new symbolic link (like /etc/init.d/net.enp3s0).
3. Remove the old symbolic link (rm /etc/init.d/net.eth0).
4. Add the new one to the default runlevel.
5. Remove the old one using rc-update del net.eth0 default.

#### systemd-networkd (systemd)\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-18 "Edit section: ")\]

[systemd](/wiki/Systemd "Systemd") has it's own network management like OpenRC has netifrc.

To configure systemd-networkd, create a \*.network file under the /etc/systemd/network directory. Files are written in INI style syntax, and usually given a name starting with two digits which determines the order they are loaded. It is recommended to use numbers lower than 70 for machine-specific configuration, to ensure it is not overridden by systemd-provided files. See the [systemd.network(5)](https://man.archlinux.org/man/systemd.network.5.en) [Special:MyLanguage/man page](/wiki/Special:MyLanguage/man_page "Special:MyLanguage/man page") man page for reference.

A \*.network file contains a `[Match]` section, which specifies which interfaces the file will be applied to, and further sections, especially `[Network]` which describe the configuration to applied. As configuration files are loaded in alphabetical order, the first file matching an interface is used. As such, configuration files with less specific `[Match]` criteria would usually be given names with higher numbers than those with more specific criteria.

Below are some examples:

Wired DHCP connection:

FILE **`/etc/systemd/network/50-dhcp.network`** **DHCP**

```
[Match]
Name=en*

[Network]
DHCP=yes

```

Alternatively, systemd-networkd provides some useful configurations in /usr/share/systemd/network, which can be enabled by symlinking from /etc/systemd/network. A configuration similar to the above is 89-ethernet.network.example.

`root #` `ln -s /usr/share/systemd/network/89-ethernet.network.example /etc/systemd/network/89-ethernet.network`

Wired connection with static IP address (uses [CIDR](https://en.wikipedia.org/wiki/CIDR "wikipedia:CIDR") notation for subnetting):

FILE **`/etc/systemd/network/50-static.network`** **Static**

```
[Match]
Name=enp1s0

[Network]
Address=192.168.1.10/24
Gateway=192.168.1.1
DNS=192.168.1.1

```

WiFi can also be configured with the help of [Wpa supplicant](/wiki/Wpa_supplicant "Wpa supplicant") or [Iwd](/wiki/Iwd "Iwd"):

FILE **`/etc/systemd/network/50-wireless.network`** **WiFi**

```
[Match]
Name=wlan0

[Network]
DHCP=yes
IgnoreCarrierLoss=3s

```

To enable the service to run at next boot up:

`root #` `systemctl enable systemd-networkd.service`

#### Optional: Networking tools\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-19 "Edit section: ")\]

**Important**

After rebooting into the new system, **wireless networks will be unavailable** (SSIDs unlisted) if they use the **insecure and obsolete** _WEP (WEP)_ or _TKIP (WPA or WPA+WPA2)_ protocols, even if they might be available from the boot-media system. There is an [insecure workaround](/wiki/Warning_about_insecure_obsolete_network_hardware#Insecure_workaround "Warning about insecure obsolete network hardware") to enable connecting to these obsolete networks which can be used before reboot, though users should be aware of the security implications.

If extra tools are required when using netifrc or systemd-networkd to configure networking on the first reboot, it is necessary to install them now while networking is available in the live environment.

##### Installing a PPPoE client\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-20 "Edit section: ")\]

If PPP is used to connect to the internet, install the [net-dialup/ppp](https://packages.gentoo.org/packages/net-dialup/ppp) package:

`root #` `emerge --ask net-dialup/ppp`

##### Install wireless networking tools\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-21 "Edit section: ")\]

If the system will be connecting to wireless networks, install the [net-wireless/iw](https://packages.gentoo.org/packages/net-wireless/iw) package for Open or WEP networks and/or the [net-wireless/wpa\_supplicant](https://packages.gentoo.org/packages/net-wireless/wpa_supplicant) package for WPA or WPA2 networks. iw is also a useful basic diagnostic tool for scanning wireless networks.

`root #` `emerge --ask net-wireless/iw net-wireless/wpa_supplicant`

### The hosts file\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-22 "Edit section: ")\]

An important next step may be to inform this new system about other hosts in its network environment. Network host names can be defined in the /etc/hosts file. Adding host names here will enable host name to IP addresses resolution for hosts that are not resolved by the nameserver.

`root #` `nano /etc/hosts`

FILE **`/etc/hosts`** **Filling in the networking information**

```
# This defines the current system and must be set
127.0.0.1     tux.homenetwork tux localhost
::1           tux.homenetwork tux localhost

# Optional definition of extra systems on the network
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

Save and exit the editor to continue.

## System information\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-24 "Edit section: ")\]

### Root password\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-25 "Edit section: ")\]

Set the root password using the passwd command.

`root #` `passwd`

Later an additional regular user account will be created for daily operations.

### Init and boot configuration\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-26 "Edit section: ")\]

#### OpenRC\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-27 "Edit section: ")\]

When using OpenRC with Gentoo, it uses /etc/rc.conf to configure the services, startup, and shutdown of a system. Open up /etc/rc.conf and enjoy all the comments in the file. Review the settings and change where needed.

`root #` `nano /etc/rc.conf`

Next, open /etc/conf.d/keymaps to handle keyboard configuration. Edit it to configure and select the right keyboard.

`root #` `nano /etc/conf.d/keymaps`

Take special care with the keymap variable. If the wrong keymap is selected, then weird results will come up when typing on the keyboard.

Finally, edit /etc/conf.d/hwclock to set the clock options. Edit it according to personal preference.

`root #` `nano /etc/conf.d/hwclock`

If the hardware clock is not using UTC, then it is necessary to set `clock="local"` in the file. Otherwise the system might show clock skew behavior.

Systems setup from a desktop stage3 should run dbus-uuidgen at this point:

`root #` `dbus-uuidgen --ensure=/etc/machine-id`

This isn't required on non desktop stage3 created systems as [sys-apps/dbus](https://packages.gentoo.org/packages/sys-apps/dbus) will automatically generate /etc/machine-id when the package is first installed.

#### systemd\[ [edit](/index.php?title=Handbook:Parts/Installation/System&action=edit&section=T-28 "Edit section: ")\]

First, it is recommended to run systemd-machine-id-setup and then systemd-firstboot which will prepare various components of the system are set correctly for the first boot into the new systemd environment. The passing the following options will include a prompt for the user to set a locale, timezone, hostname, root password, and root shell values. It will also assign a random machine ID to the installation:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Next users should run systemctl to reset all installed unit files to the preset policy values:

**Note**

It is possible that a failure warning will print out after running the following command. This is normal, as Gentoo's LiveCD uses OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

It's possible to run the full preset changes but this may reset any services which were already configured during the process:

`root #` `systemctl preset-all`

These two steps will help ensure a smooth transition from the live environment to the installation's first boot.

[← Configuring the kernel](/wiki/Handbook:MIPS/Installation/Kernel "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Installing tools →](/wiki/Handbook:MIPS/Installation/Tools "Next part")