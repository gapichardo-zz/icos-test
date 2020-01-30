# icos-test
<H1>Repositorio para el artículo "Aplicación de almacenamiento de archivos en la nube con alarma de modificación"</H1>
<H2>Resumen:</H2>
<p>El objetivo principal de este artículo es mostrar la versatilidad, facilidad de uso y potencial que tienen las soluciones arquitectadas e implementadas en la nube, en este caso, el ejemplo desarrollado se implementó utilizando servicios de la IBM Cloud.</p>

<p>La solución presentada en este artículo es una aplicación Web desarrollada en Node.js, utilizando Express como framework de backend. Esta aplicación se ejecuta en contenedores Docker desplegada en IBM Kubernetes Service, se despliega una forma HTML básica para seleccionar un archivo que, utilizando las API’s de ICOS, se realiza la carga del archivo. La actualización del componente de almacenamiento (bucket) dispara un “trigger” que a su vez ejecuta una “acción” desarrollada en GoLang y alojada en IBM Functions, la cual envía un SMS alertando la carga del archivo por medio del servicio de Twilio.</p>

<H2>Resumen de Arquitectura:</H2>
<images></images>
