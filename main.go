package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
DOCKER ENGINE
=============
1. Create Container
2. delete Container
3. stop Container
4. clean Container
5. build image

parts / components of docker engine
-----------------------------------
1. dockerd - docker daemon
2. container runtime
	- containerd - high level container runtime
	- runc - low level container runtime

daemon - type of gurdian (perform good deeds-silently-invisibly)
	1. background process (when we don't delete/not stuck like go server(terminal session dhore rakhe))
	2. Listerns for requests
	3. Serve services to other process
	- starts when computer starts

container runtime
	- which creates/delete/stop/clear container
	- works with container

	containerd
	----------
	- manage container lifecycle (start, end)
	- manage container storage
	- manage image
	- provide API so that dockerd can comuunicate with containerd

	runc
	----
	- create - namespace, cgroup create
	- delete - - namespace, cgroup create delete
	- stop, clean container

when we tell docker enginer to create a container
------
- dockerd accepts the request
- order containerd to take the request
- containderd use runc to create container
- runc requests kernel to create namespace, cgroup etc

install 2 things
----------------
- docker engine
- docker CLI/Client

Youtube: https://www.youtube.com/watch?v=BjwoCFwooJU&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=13
*/
