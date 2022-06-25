//BU ÖZELLİK SADECE LİNUX İÇİN GEÇERLİDİR
package main
import (
     "os"
     "os/exec"
   . "fmt"
    
)
 func main(){
var işlemler [4]string
işlemler[0] = "1-Isı değerlerini görüntüle\n"
işlemler[1] = "2-çalışan işlemleri görüntüle\n"
işlemler[2] = "3-disk kullanımını listele\n" 
işlemler[3] = "4-kullanılan linux dağıtımı hakkında bilgi\n"

	Println(işlemler)
	Println("yapmak istediğiniz işlemin numarasını giriniz")
	
	var işlem int 
	Scanln(&işlem)

	switch işlem{
	case 1:
		cmd := exec.Command("sensors")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case 2:
		cmd := exec.Command("ps","-ax")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case 3:
		cmd := exec.Command("df","-h")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case 4:
		cmd := exec.Command("uname","-a")
		cmd.Stdout = os.Stdout
		cmd.Run()
		
	}
}
