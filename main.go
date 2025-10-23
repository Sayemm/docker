package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
VIRTUAL MACHINE?
---------------
- 6xVCPU - 6 Core
- 6 Process can run parallely
- install Hypervisor - HP process will create
	-> can create Virtual Machine (Virtual CPU, RAM, HD, NIC...)
	-> can virtualize physical hardware
	-> Then map between physical (host) and virtual
	-> got virtual computer
		- so we can install an OS (guest OS) in here

- 6 Core, RAM: 10GB, HD: 100GB
- VM1: 4 Core, RAM: 4GB, HD: 50GB
- VM3: 4 Core, RAM: 4GB, HD: 50GB
- VM3: 4 Core, RAM: 4GB, HD: 50GB
-> 12 Core!! Physical 6 core will be divided
-> 16 GB - will be divided -  Balloing technique
=> One VM does not know about another VM

Youtube: https://www.youtube.com/watch?v=5YNAGk6zSag&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=10
*/
