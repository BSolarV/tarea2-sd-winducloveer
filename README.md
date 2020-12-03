# Tarea 1

## WinduCloveer
> Carlos Jara Almendra - 201773036-5  
> Bastián Solar Vargas - 201773003-k

## Maquinas Viruales:

El usuario de las máquinas es: sd

* Máquina 1: datanode  
	> ip/hostname: 10.10.28.63  
	> contraseña: RQykXsIOZSDOuzd
 
* Máquina 2: datanode
s	> ip/hostname: 10.10.28.64  
	> contraseña: FcPkvnGbEWEAlie
 
* Máquina 3: datanode   
	> ip/hostname: 10.10.28.65  
	> contraseña: uNzXQlZUsGbKgND
 
* Máquina 4: namenode
	> ip/hostname: 10.10.28.66  
	> contraseña: xysmRmDVuHkoWLk

* Todas son client

# Consideraciones:

# Ejecución

## Registros:
* Para limpiar los registros se incluye el comando *clearRegisters* en el Makefile.
	> make clearRegisters

## namenode:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Logistica* que creará un binario del mismo nombre en el directorio *bin*.
	> make buildNamenode
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/namenode

## Clientes:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Clientes* que creará un binario del mismo nombre en el directorio *bin*.
	> make BuildClientes
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Clientes

## Camiones:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Camiones* que creará un binario del mismo nombre en el directorio *bin*.
	> make BuildCamiones
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Camiones 

## Finanzas:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Finanzas* que creará un binario del mismo nombre en el directorio *bin*.
	> make BuildFinanzas
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:
	> ./bin/Finanzas