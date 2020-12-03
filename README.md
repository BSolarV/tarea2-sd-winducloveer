# Tarea 1

## WinduCloveer
> Carlos Jara Almendra - 201773036-5  
> Bastián Solar Vargas - 201773003-k

## Maquinas Viruales:

El usuario de las máquinas es: sd

* Máquina 1: Logistica  
	> ip/hostname: 10.10.28.63  
	> contraseña: RQykXsIOZSDOuzd
 
* Máquina 2: Clientes
	> ip/hostname: 10.10.28.64  
	> contraseña: FcPkvnGbEWEAlie
 
* Máquina 3: Camiones   
	> ip/hostname: 10.10.28.65  
	> contraseña: uNzXQlZUsGbKgND
 
* Máquina 4: Finanzas
	> ip/hostname: 10.10.28.66  
	> contraseña: xysmRmDVuHkoWLk

# Consideraciones:
## Generales  
* Las conexiones se definen en base a las maquinas virtuales asignadas, de forma que los ejecutables funcionarán efectivamente en su máquina correspondiente.  
	> Logistica -> Maquina 1: 10.10.28.63  
	> Clientes -> Maquina 2: 10.10.28.64  
	> Camiones -> Maquina 3: 10.10.28.65  
	> Finanzas -> Maquina 4: 10.10.28.66  
* Los sistemas requieren un criterio de cierre que se definiria a criterio del grupo para cada cual.
* Se muestra en pantalla información que puede resultar útil para el usuario.
* Se supone todas las entradas a los sistemas serán validas (sin *typos* o dentro de valores realistas y no negativos)

## logistica:
* El sistema debe terminar su ejecución luego de un tiempo definido de inactividad ingresado al iniciar la ejecución.
	> Tiempo de inactividad es considerado como el *gap* de tiempo entre acciones críticas (recibir orden de cliente; Enviar paquete a camión; Recibir paquete de camión)
	> Tiempo máximo de inactividad débe ser de al menos 7 veces el tiempo de viaje de los camiones. (ej: si el tiempo de viajes de camiones es de 60 segundos, el mínimo tiempo máximo de inactividad para logística sería 7 minutos)
* El codigo de seguimiento no requiere un formato definido, por lo cual se utilizará el ID del Paquete en el sistema. En caso de ser Retail el codigo de seguimiento será 0 pues no peden ser seguidos.
* Solamente se puede hacer seguimiento de paquetes recibidos en la instancia actual del sistema y no de paquetes escritos en el archivo registro por instancias previas.
* Se ignora el campo ID que envien los clientes, pues los clientes no deben definir el id, es labor del sistema.
* El archivo de registro se hara en un arhivo llamado *registroLogistica.csv*, el cual no llevará encabezado pues se considera sabido que el orden de columnas será:
	> timestamp | ID Paquete | Tipo | Producto(Descripcion) | Valor | Origen | Destino | CodigoDeSeguimiento


## Clientes:
* El sistema pregunta al inicia que tipo de cliente simulará (cliente Retail o cliente Pyme).
* El sistema debe terminar su ejecucion cuando termine de enviar todos los paquetes que tenga en su csv respectivo (csv para *Retail* y *Pymes*)
* Se pueden tener multiples clientes enviando paquetes, pero solo se leerá el archivo correspondiente al tipo de cliente seleccionado al inicio.
* Tanto el codigo de seguimiento como el estado del paquete se mostrarán en pantalla debido a que no se requiere un archivo de registros para la instancia del sistema.
* Los csv con paquetes a entregar estarán en la carpeta *files* bajos los nombres *retail.csv* para paquetes de retail y *pymes.csv* para paquetes de pymes.
* Los csv tendran formatos (esquema de columnas) identicos a los de ejemplo.
* Para correcta lectura de archivos se deben tener dentro de una carpeta *files* el directorio desde el cual se ejecute.
	> * Si se ejecuta desde la carpeta raiz del sistema (\~/Tarea1) se utilizarán los archivos de ejemplo que se encuentran en "\~/Tarea1/files".   
	> * Si se ejecuta desde la carpeta "\~/Tarea1/bin" se requerirán archivos en "\~/Tarea1/bin/files".    
	> * **Se recomienda seguir las instrucciones de ejecucion y situarse en la carpeta raiz del proyecto (\~/Tarea1) para la ejecución.**
* Para las acciones del cliente se trabajarán mediante probabilidades definidas al inicio de la ejecución.

## Camiones:
* El sistema debe terminar su ejecución cuando reconozca que ***Logistica*** terminó sus procesos.
* Para considerar si reintentar un paquete de retail se siguio la siguiente idea:  
	* 1 intento -> costo 0  
	* 2 intentos -> costo 10   
		* Prioritarios
			* *0.8\*valor + 0.2\*0.3\*valor* = ganancia estimada   
			* Condicion de 2 intentos: 
				> *0.8\*valor + 0.2\*0.3\*valor* > 10  
		* Normales
			* *0.8\*valor* = ganancia estimada   
			* Condicion de 2 intentos:   
				> *0.8\*valor > 10*
	* 3 intentos -> costo 20  
		* Prioritarios
			* *0.8\*valor + 0.2\*0.3\*valor* = ganancia estimada
			* Condicion de 3 intentos: 
				> 0.8\*valor + 0.2\*0.3\*valor > 20 
		* Normales
			* 0.8*valor = ganancia estimada
			* Condicion de 3 intentos: 
				> 0.8*valor > 20 

## Finanzas:
* La maquina tendrá montado correctamente un servidor de RabbitMQ, con usuario *WinduCloveer* y clave *secret* con permisos en *"/"* y administrador.
* El sistema debe terminar su ejecución luego de un tiempo definido de inactividad definido al iniciar la ejecución.
* El registro se mostrará cada una cantidad definida de segundos que se solicitará al inicio de la ejecución.
* El archivo de registro se hara en un arhivo llamado *registroFinanzas.csv*, el cual no llevará encabezado pues se considera sabido que el orden de columnas será:
	> Id Paquete | Descripcion | Tipo | Intentos | Estado | ValorOriginal | Ganancia/Costo 

# Ejecución

## Registros:
* Para limpiar los registros se incluye el comando *clearRegisters* en el Makefile.
	> make clearRegisters

## Logística:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Logistica* que creará un binario del mismo nombre en el directorio *bin*.
	> make BuildLogistica
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Logistica

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