# Sistema de Anti-Ransomware Distribuido
## _Projeto desenvolvido por, Tomás Ferreira N20190881, João Rolo N20190861, Martim Costa N50039055_

_Documentation for project developed in the Distributed Systems Class_

In this documentation, it will be explained how you can replicate the system developed for the project, including commands for installing the necessary packages aswell commands for configuring the files necessary to run the project.

For this project we will use the following set-up,
### Machines
| VM | Purpose | Necessary Packages | OS | Ram | DiskSpace| 
| :---: | :---:| :---:| :----:|:---:|:---:|
| **UE01, UE02** | Run the secure server service, this machines will be the core of the system, since they will be the ones using the client server API, saving security copies, File Hashes, and re-uploading to the client server| Golang-1.16>= | Ubuntu 20.04 | It's prefered to use 2gb of ram during the instalation, but later u can downgrade it to 1gb of ram, since its enough to run the necessary application | Minimum 10gb |
| **UE03** | Run the client server service, which will contain the client files, that are going to be saved in the secure server, aswell as their hash | Golang-1.16>= | Ubuntu 20.04 | It's prefered to use 2gb of ram during the instalation, but later u can downgrade it to 1gb of ram, since its enough to run the necessary application | Minimum 10gb |
| **UE04** | Run the Nginx service, using the Nginx Load-Balancer service | Nginx1.18>= | Ubuntu 20.04 | It's prefered to use 2gb of ram during the instalation, but later u can downgrade it to 1gb of ram, since its enough to run the necessary application | Minimum 10gb |

After installing a the Ubuntu ISO file, you can create a VM in the virtualization software of your choice, with the settings referred above. You only need to do this once, since you can create linked clones from the first Vm created.
<img src="Clone1.png" width ="auto" height="400"/>