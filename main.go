package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
LINUX PACKAGES
--------------
- complete OS already has a package manager which can install/uninstall packages (vlc, git) those are within a repository
	- high level - apt
		=> to install a package it might need something beforehand in the system
		=> apt will find those dependencies first and install those with dpkg first before install the main package
		=> can upgrade (new version)

	- low level - dpkg
		=> this installs packages in the end => apt talks with dpkg to install the packages
		=> if we directly try to install a package using dpkg it might fail, as it won't tell us which things it needs beforehand to install that package

- ubuntu package manager  - (apt, apt-get = high, dpkg = low)
- redhat, centos, fedora - (yum, dnf = high, rpm)
- arch - packman
- openSuse - zypper
=> Each package manager has its own repository system
- repository has an URL
- package has a format (.deb for ubuntu, .rpm for yum)

- apt install git
	=> apt will search git in the repository (where everything is stored for that package manager)
	=> apt has the URL and then apt will check whether that URL has git or not

- package (git, vlc) includes
	=> program (binary)
	=> metadata
	=> dependencies
	=> install scripts (pre-install, post-install, pre-remove, post-remove)
	=> config files






================> WHY DO WE NEED PACKAGE MANAGER IN LINUX

Linux applications depend on system-provided shared libraries.
Apps like VLC, Chrome, and Firefox do not include their own audio, video, network, or GUI libraries—they rely on the same common libraries installed in the Linux system.
Because of this, Linux needs a package manager (like apt) to automatically install and manage all these required shared libraries and keep them updated.

Windows applications usually bring their own libraries inside their .exe or .msi installers.
Each app bundles its own audio, video, GUI, and network DLLs, so they don’t rely on system-wide libraries.
Since Windows apps include their dependencies themselves, a package manager is not required for normal software installation.


Youtube: https://www.youtube.com/watch?v=W7uuHbEv528&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=18
*/
