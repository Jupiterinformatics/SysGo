//BU ÖZELLİK SADECE LİNUX İÇİN GEÇERLİDİR
package main
import (
     "os"
     "os/exec"
   . "fmt"
    
)
 func main(){
var processs [5]string
processs[0] = "1-show heat values\n"
processs[1] = "2-show work processs\n"
processs[2] = "3-list disk usage\n" 
processs[3] = "4-take information about distro\n"
processs[4] = "5-list ram usage\n"

	Println(processs)
	Println("procces no :")
	
	var process int 
	Scanln(&process)

	switch process{
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
	case 5:
		cmd := exec.Command("free","-h")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
