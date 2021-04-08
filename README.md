# Operación Fuego de Quasar

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.
El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

![](https://matthcep.s3.amazonaws.com/Screen+Shot+2021-01-24+at+6.41.27+PM.png)

# Ejecutar aplicación local

Para ejecutar la aplicación en local, seguir los siguientes pasos:

Requerimientos:

- Tener instalado <abbr title="Golang">Go</abbr>

Clonar el repositorio:

`$ git clone https://github.com/folivarez/Quasar.git`

`$ cd Quasar`

Una vez dentro de la carpeta del proyecto ejecutar los siguientes:

`$ go run app.go`

Para correr los test:

`$ cd test`

`$ go test -v`

## Implementación de la solución

Para el desarrollo de la siguiente prueba, se utilizó las siguientes tecnologías:

- Golang vgo1.14
- Google Cloud Platform (App Engine)

El concepto usado para la solución del problema se llama trilateración, la trilateración usa las localizaciones conocidas de dos o más puntos de referencia, y la distancia medida entre el sujeto (Nave) y cada punto de referencia (Satelites). Para determinar de forma única y precisa la localización relativa de un punto en un plano bidimensional usando solo trilateración, se necesitan generalmente al menos 3 puntos de referencia. Para este caso, solo se tienen en cuenta 3 puntos, de acuerdo al requerimiento.

![](https://i.imgur.com/JMrsnDr.png)

En el diagrama se muestra la intercepción de todos los puntos en un plano bidemencional.

## Endpoints

**Topsecret**
Obtener la posicion y el mensaje completo recibido por los 3 satelites
`POST -> /topsecret`
Ejemplo Payload:

    {
    	"satellites": [
    		{
    			"name": "kenobi",
    			"distance": 989.945,
    			"message": ["este","","","mensaje",""]
    		},
    		{
    			"name": "skywalker",
    			"distance": 608.276,
    			"message": ["","es","","","secreto"]
    		},
    		{
    			"name": "sato",
    			"distance": 500,
    			"message": ["este","","un","",""]
    		}
    	]
    }

Ejemplo respuesta 200:

{
    "message": "este es un mensaje secreto",
    "position": {
        "x": -312.1124559649421,
        "y": 154.5355105513955
    }
}

Example failed 400 response

    {
      "error": "No se puede calcular la posicion con la informacion recibida"
    }

<br>
<br>

## Endpoint Productivo

[https://quasar-310020.uc.r.appspot.com](https://quasar-310020.uc.r.appspot.com)
