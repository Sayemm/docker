package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
VM vs CONTAINER
===============
- VM does not share Kernel
- VM has it's own OS ie. Kernel

- Container usees the Kernel of host OS

how docker is related?
----------------------
- most important part of docker is docker engine
- docker enginer can create a Container using computers Kernel
- separate namespace,.. for each container

- for windows/mac - docker desktop installs a mini/lightweight VM and mini basic linux in that VM


Youtube: https://www.youtube.com/watch?v=HWt2IXk59aw&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=12
*/
