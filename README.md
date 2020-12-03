# Tarea 1

## WinduCloveer
> Carlos Jara Almendra - 201773036-5  
> Bastián Solar Vargas - 201773003-k

## Maquinas Virtuales:

El usuario de las máquinas es: sd

* Máquina 1: datanode  
	> ip/hostname: 10.10.28.63  
	> contraseña: RQykXsIOZSDOuzd
 
* Máquina 2: datanode
	> ip/hostname: 10.10.28.64  
	> contraseña: FcPkvnGbEWEAlie
 
* Máquina 3: datanode   
	> ip/hostname: 10.10.28.65  
	> contraseña: uNzXQlZUsGbKgND
 
* Máquina 4: namenode
	> ip/hostname: 10.10.28.66  
	> contraseña: xysmRmDVuHkoWLk

* Todas son client

# Consideraciones:

## Generales 
* No se harán inputs incorrectos
* Se dejarán los binarios compilados para evitar problemas con la versión de Go instalada en las máquinas virtuales ya que ésta es de una versión menor.
* En el Método Centralizado cuando llegue una propuesta el NameNode la atenderá por orden de llegada


## Datanodes
* Al iniciar se preguntará si desea debbuguear, solo recibe como respuesta *yes* o *no* sin comillas
* La opción de debbugear mostrará unas marcas de tiempo que nos ayudará a determinar cuanto demora cada algoritmo (Centralizado o Distribuido)
* También pregunta el número de datanode, un datanode no puede tener un número igual a otro. En ese caso el programa terminará su ejecución
* Se dejaron prints en consola para dilucidar si el programa funciona o no

## NameNode
* Al iniciar se preguntará si desea debbuguear, solo recibe como respuesta *yes* o *no* sin comillas
* La opción de debbugear mostrará unas marcas de tiempo que nos ayudará a determinar cuanto demora cada algoritmo (Centralizado o Distribuido)
* Se dejaron prints en consola para dilucidar si el programa funciona o no
* Se escribirá un archivo Log.txt. Si éste no existe será creado por el nameNode
* Cada vez que se ejecute el NameNode, si Eexiste Log.txt intentará cargar en memoria los libros que estaban escritos en el Log.txt.

## Cliente
* Debe elegir cuál de los 2 archivos ejecutar. o el clientCentralized (Ejecuta el sistema Centralizado) o el clientDistributed (Ejecuta el sistema Distribuido)
* Al iniciar se preguntará si desea debbuguear, solo recibe como respuesta *yes* o *no* sin comillas
* La opción de debbugear mostrará unas marcas de tiempo que nos ayudará a determinar cuanto demora cada algoritmo (Centralizado o Distribuido)
* El cliente podrá Subir, Descargar o Visualizar libros y también salir del sistema. Esto lo hará ingresando los números dados en la consola para hacer cada acción.
* Si elige subir, se le preguntará el nombre que le quiere dar al archivo a guardar en el Log
* Luego de eso le pedira que ingrese el nombre del archivo. En las máquinas habrán 4 archivos .pdf
* Se debe subir como: *nombreArchivo.pdf* 
* Los 4 nombres son:
	>Dracula-Stoker_Bram.pdf
	>Frankenstein-Mary_Shelley.pdf
	>Orgullo_y_prejuicio-Jane_Austen.pdf
	>Peter_Pan-J._M._Barrie.pdf

# Ejecución

## Registros:
* Para limpiar los registros se incluye el comando *clearRegisters* en el Makefile.
	> make clearRegisters

## Namenode:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *namenode* que creará un binario del mismo nombre en el directorio *bin*.
	> make buildNamenode
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea2) ejecutar:  
	> ./bin/namenode
* Se dejarán los binarios compilados para evitar problemas con la versión de Go instalada en las máquinas virtuales ya que ésta es de una versión menor.

## Datanode:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *datanode* que creará un binario del mismo nombre en el directorio *bin*.
	> make buildDatanode
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea2) ejecutar:  
	> ./bin/datanode
* Se dejarán los binarios compilados para evitar problemas con la versión de Go instalada en las máquinas virtuales ya que ésta es de una versión menor.

## Clientes:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *clientCentralized* o *clientDistributed* que creará un binario del mismo nombre correspondiente en el directorio *bin*.
	> make buildClientCentralized
	> make buildClientDistributed
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea2) ejecutar: 
	> ./bin/clientCentralized
	> ./bin/clientDistributed
* Se dejarán los binarios compilados para evitar problemas con la versión de Go instalada en las máquinas virtuales ya que ésta es de una versión menor.
