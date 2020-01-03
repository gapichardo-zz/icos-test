/**
 *
 * main() Este código se ejecutará cuando sea llamado por el trigger asociado
 *
 * @param Las acciones en Cloud Functions actions aceptan parámetros, las cuales pueden ser objetos JSON.
 *
 * @return En este caso, el valor de regreso de laz función no se utiliza.
 *
 */
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

/*
  Las constantes declaradas aquí son para uso de Twilio y se obtienen del Portal de Twilio

*/
const (
	accountSid  = "<SID-DE-LA-CUENTA>"
	authToken   = "<TOKEN-DE-AUTORIZACION>"
	urlStr      = "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	twilioPhone = "<TELEFONO-ASIGNADO-POR-TWILIO>"
)

/*
   Al ser llamada la acción se ejecutará la llamada a la función Main,
   Los parametros que recibe se mapean en una estructura MAp de Go para
   ocuparlos en el envío del mensaje SMS.

*/

func Main(params map[string]interface{}) map[string]interface{} {

	var1 := "Nombre del Archivo:" + fmt.Sprintf("%v", params["key"])
	var2 := "Operación: " + fmt.Sprintf("%v", params["operation"])
	var3 := " en el Bucket: " + fmt.Sprintf("%v", params["bucket"])

	// Los numeros de los telefonos celulares que deben estar validados en el portal de Twilio
	// Pueden ser obtenidos de una base de datos en un ambiente productivo

	celulares := []string{"<TELEFONO-DE-ENVIO1>", "<TELEFONO-DE-ENVIO2>"}
	v := url.Values{}
	cliente := &http.Client{}
	fmt.Println(urlStr)

	for item, tel := range celulares {
		v.Set("To", tel)
		v.Set("From", twilioPhone)
		v.Set("Body", "Cambios "+var3+" -"+var1+" -"+var2)
		rb := *strings.NewReader(v.Encode())

		req, error := http.NewRequest("POST", urlStr, &rb)
		if error != nil {
			fmt.Println("Error:", error)
		}
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, _ := cliente.Do(req)
		fmt.Printf("Staus de envio: %d, con estatus %s\n", item+1, resp.Status)
	}

	msg := make(map[string]interface{})
	msg["body"] = "Hello " + "sms enviado" + "!"
	// opcionalmente se puede hacer log a la stdout (o stderr)
	fmt.Println("hello Go action")
	return msg
}
