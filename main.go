package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
CONTAINER
=========
- application install - binary to HD
- Run - process creates to the RAM

- app1 needs go1.16 installed in the system/computer
- app2 needs go1.25 installed in the system/computer
- but normally we cannot install same app/go different version in the computer

- In Linux - Namespace
- we can solve this problem using Namespace

- When OS runs - full things are within ONE namespace
- Now when we install software - eveythings are within that ONE namepace
- Namespace is a separate world

SOLUTION
--------
- We can create separate NAMESPACE - separate world
- We cannot get other software from another namespace
- NAMESPACE ISOLATE "WHAT" THE PROCESS WILL SEE
	- process can things like files, network, PIDs....

- CGROUP / Control Group
- ISOLATE "HOW MUCH" A PROCESS CAN SEE

- now we can crate 2 namespce and install go1.16 and go1.25
- now we can install app1 on namespace A and install app2 on namespace B
- when app1 will run, Kernal will check which namespace it's coming from

- namespace A will get centain level Kernal, HD , RAM, CPU...access
- This namepace is separate world
- BUT ALL NAMESPACE ARE SHARING THE SAME KERNAL, RAM, HD...

- There can be more proecsses in this namespace world
- This world does not know about ours and think everything is it's own
- This separate world is a CONTAINER!!!!!!!

=> Using Namespace and Cgroups, isolate the computer hardware - a separate world is being created
=> This separate world is a CONTAINER

=> Image: Screenshots of running compter.
=> We can create lots of container using that image



Youtube: https://www.youtube.com/watch?v=Ch_1jkj5uTY&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=11
*/
