package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
GNU Utilities
=============
- GNU
	- GNU is not unix
	- Started by - Free Software Foundation 9FSF) - Richard Stallman - 1983
	- Unix was an old OS - PAID/Proprietary
	- Richard wanted to make it free
	- Created Various FREE components for OS
	- So that everyone can create an OS for free
	GNU Components
		- GNU compiler collection (gcc - compile c)
		- GNU C Library (glibc - printf/scanf/fopen.. is not part of c -> needs glibc)
		- GNU Utils / GNU Coreutils
		- GNU Bash - Program (different types of bash program - altogether we call it SHELL)
		- GNU Debugger

- GNU Utils / GNU Coreutils (Commands)
	- Kernel manages hardware (CPU, RAM, HDD)
	- But as a user how can I use that?
	- User needs to manage (input-output) as well.
	- GNU Coreutils solves this problem
		- GNU coreutils runs on user space
		- GNU Coreutils helps to communicate with Kernals
		- GNU Coreutils has all the commnads
		- We can use those to communicate with the kernel/hardware
	- ls, cd, cat, grep, cp, rm
	- WE NEED TO WRITE THIS COMMAND SOMEWHERE - TERMINAL

- Desktop Environment (GNOME, KDE) gives that TERMINAL
	- desktop environment has Different things
		- Terminal
		- Window manager (UI)
		- File Explorer
		- Taskbar
	- What see - Graphical things
	- Ubuntu desktop environment name is GNOME

- Terminal
	- program that runs on user space
	- Shell: sh (bourne shell), ksh (Korn shell), bash (bourne again shell), zsh (Z shell),
		- INTERPRATE COMMAND
		- Bash takes the command that comes from the terminal
		- interprate command, remove something...
		- communicates with GNU Utils about that command
		- then gives that command to kernel

		- SHELL (Bridge between Terminal and Kernel)


Youtube: https://www.youtube.com/watch?v=mWdRPr9V-pw&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=16
*/
