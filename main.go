package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
DOCKER ECOSYSTEM
================
- docker -> it's a platform or ecosystem - not a tools/application
- ecosysmtem - chain / multiple things are working together

- docker cli
- docker engine
- docker desktop (Linux VM)
- docker images (screenshot of container)
- docker Hub (website/database where images are being stored)
- docker Compose

* docker run hello-world (image)
- docker cli/client receives this command
- sends this to docker engine
	- dockerd -> containerd (cache?) -> docker Hub
	- then containerd tells runc to convert hello-world image to a container
	- runc communicates with Kernel with create a separate namespace


Youtube: https://www.youtube.com/watch?v=lDklV7p3h8U&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=14
*/
