package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
Docker Hands On
---------------
- docker images
- docker ps                      - only running container
- docker ps -a                   - all container
- docker rmi -f ubuntu:24.04     - remove the image forcefully
- docker pull ubuntu:24.04       - pull from docker hub and cache it to local computer
- docker run -it ubuntu:24.04 bash
	- docker run -it ubuntu:24.04 sh (will create another container from the same image)


- when we run a container from an image it gets a random name, we can give that name
	-> docker run --name my_ubuntu -it ubuntu:24.04
	- helps us to identify the right container as we might create multiple container from the same image
	- we can use the container id as well to identify the container
	-> docker exec -it my_ubuntu bash (using the name -> running the same container from another terminal)

- (-it) (-i -t will also work)
	-i flag means standard input (open for input)
	-t flag means pseudo TTY(terminal)
	- to run it on interactive mode
	- docker run --name my_ubuntu -i ubuntu:24.04 (this will also work but not with terminal)

- docker run --name my_ubuntu1 ubuntu:24.04
	- without -it, will be on Exited Status
	- (docker ps ) won't be able to find, (docker ps -a) will find that container

- docker run ubuntu:24.04
	- docker run   -> main command
	- ubuntu:24.04 -> which container
	- bash         -> when the container is completely created from image ubuntu:24.04
	               -> the program I want to run





Youtube: https://www.youtube.com/watch?v=JQbaqSFzTvs&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=21
*/
