package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
Linux Basic Commands
--------------------
- sayem@Mac ~ %
	- user name is sayem who logged in to the machine Mac
	- ~ -> home directory
	- % -> general user (% also meaning z shell, $ means bash shell, # means super user)

- Linux File System
	- / -> root directory
		- /bin -> binaries (basic commands - ls, cat..)
		- /boot-> kernel & bootloader
		- /dev -> devices (like software needs for wifi adapter)
		- /etc -> configurations
		- /home-> various user home
		- /lib -> libraries
		- /media> mounts (external hardisks)
		- /proc-> kernal & process information
		- /sbin-> system binaries
		- /tmp -> temporary file
		- /usr -> user programs (all applications)
		- /var -> logs, cache
		........................

- pwd (present working directory)
- ls
	- ls -l  (long detail view)
	- ls -a  (hidden files)
	- ls -la (hiddend and long list)

- cd
	- cd / -> go to root path
	- cd ~ -> home directory
	- cd - -> last use immediate directory
	- cd .. -> parent directory

- mkdir app               -> create app folder
- mkdir -p sayem/sayed    -> folder within folder
- touch filename (a.txt)     -> create a file

- cp file1.txt file2.txt     -> create file2.txt and copy the content from file1.txt
- cp -r sayed/ munni/        -> create another folder munni and recursively copy all the contents from sayed
- cp file2.txt ../sayed/file3.txt -> copy file2.txt to another folder and create file3.txt

=> relative path => ../sayed/file3.txt
=> absolute path => /app/sayem/sayem => start from the root

- echo hello world >> a.txt  -> append content to the file
- echo hello world > a.txt.  -> remove old content and just add "hello world"
- cat filenane (a.txt)       -> see the contents of the file
- head -n 5 a.txt            -> see first 5 lines from the txt
- tail -n 5 a.txt            -> see last 5 lines from the txt
- rm filename (a.txt)        -> remove a file not directory
- rm -r foldername           -> remove the folder recursively (delete one by one in that folder then remote the main folder)
- rm -rf foldername          -> forcefully remove the folder

- mv file3.txt file4.txt     -> rename file3 to file4





Youtube: https://www.youtube.com/watch?v=6WL7VOX55kg&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=19
*/
