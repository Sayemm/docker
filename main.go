package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

/*
Managing User, Group & Permissions
----------------------------------

USER
====
- 2 types
	- root user (can do anything)
	- regular user (limited permission)
- we normally works as a regular/normal user

- useradd -m sayem   -> new user create (-m to create folder in home directory)
- id sayem           -> identification of sayem (uid=1001(sayem) gid=1001(sayem) groups=1001(sayem))
- cat /etc/passwd    -> see all users
- cd ~               -> cd to home directory of current user (# means root, root@71fc0329d044:~#)
- su - sayem         -> switch user sayem
- root@71fc0329d044:~# passwd sayem -> change pass from root

- userdel -r muaz    -> user delete (-r for deleting the folder)

- usermod -L sayem   -> lock that user so that nobody can log in
- usermod -U sayem   -> lock that user

GROUP
=====
- Group is created automatically when we create an user

- groupadd admin    -> create group with name admin
- groupdel admin    -> delete group
- cat /etc/group    -> see all groups (from root)
- groups            -> from user to find out the group where this user is included

- usermod -aG admin sayem -> sayem is being append group - admin
	- id sayem      => uid=1001(sayem) gid=1001(sayem) groups=1001(sayem),1002(admin)
- gpasswd -d sayem admin  -> remove sayem from group admin

- 2 types group
	- Primary Group
	- Secondary/Supplementary Group

- each user is connected to both group
	-> primary group of sayem => gid=1001(sayem)
	-> secondary => 1002(admin)


PERMISSION
==========

- 1 file (a.txt) and 1 folder (files), inside of this directory

ls -l
------
total 4
-rw-rw-r-- 1 sayem sayem    0 Dec 16 02:34 a.txt
drwxrwxr-x 2 sayem sayem 4096 Dec 16 02:34 files

-  -> means fle
d  -> means directory

rw-rw-r--
	- user sayem permission (rw-) (can read, write but cannot execute)
	- group sayem permission(rw-) (can read, write but cannot execute)
		- like another user on the same group can read, write but cannot execute
	- other user permission (r--) (can only read, cannot write / execute)
		- like other user from differnt group can only read
rwxrwxr-x
	- user sayem permission (rwx) (can read, write, execute)
	- group sayem permission(rwx) (can read, write, execute)
	- other user permission (r-x) (can only read, execute but cannot write)
rwx
	- read, write, execure
	- 3x3 (user, group, others)

- give execute permission to user sayem for a.txt
	- chmod u+x a.txt
	- $ ls -l
		total 4
		-rwxrw-r-- 1 sayem sayem    0 Dec 16 02:34 a.txt
		drwxrwxr-x 2 sayem sayem 4096 Dec 16 02:34 files

- remove write permission for user sayem for a.txt
	- chmod u-w a.txt
	- $ ls -l
		total 4
		-r-xrw-r-- 1 sayem sayem    0 Dec 16 02:34 a.txt
		drwxrwxr-x 2 sayem sayem 4096 Dec 16 02:34 files

- remote group write permissoin
	- chmod g-w a.txt
	- $ ls -l
		total 4
		-r-xr--r-- 1 sayem sayem    0 Dec 16 02:34 a.txt
		drwxrwxr-x 2 sayem sayem 4096 Dec 16 02:34 files

- changed ownership of that file from sayem to another one (root user can do that)
	- chown muaz a.txt

- for other - chmod o+w a.txt

1 => just a file no subdirectory
2 => subdiretory within folder files (empty then how 2?)
	- ls -a (show all as well as HIDDEN)
		.  ..

0    - 0 bytes as a.txt is empty right now
4096 - without content, folder needed 4096 bytes of memory

Youtube: https://www.youtube.com/watch?v=6wqYiB3Cobk&list=PLpCqPSEm2Xe8dVi8cCLM9jmRp-FtEIGil&index=20
*/
