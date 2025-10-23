package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
What is Kernal?
---------------
- 3xThread - 2xCore => 6 Core
- Core (Processing Unit (ALU, CU), Register Set)

- OS installs to HD
- Run the computer - loads OS code to RAM
- Now OS takes access to everything (RAM, HD, CPU...Hardware..)

- Core part/Brain of OS is Kernal
- Kernal Space: Where OS runs
- User Space: Where other applications run
- CPU has 2 mode as well based on what it's running right now (kernel mode, user mode)

- When process that is running of user space, needs anything like file acces...
	-> process cannot access it directly
	-> CPU currently on user modem as it's runnig a process other that Kernal
	-> CPU on user modem cannnot access any file/hardwre directly
	-> Process will system call to Kernal
	-> this created privacy/security like one process cannot access other process's file

- Kernal
	- Process management (process/thread scheduling..)
	- Memomy management (allocate/free memory)
	- Device management (NIC connection, hardware driver..)
	- File system management
	- system call
	- etc....

Youtube: https://www.youtube.com/watch?v=qhQ_4-o2kJQ&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=9
*/
