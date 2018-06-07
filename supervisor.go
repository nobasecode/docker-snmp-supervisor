package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
)

// author = 'Mohamed Laabid'


//get a list of docker containers
func inspect() map[string][]string{
    
    containers := make(map[string][]string)
        
    out, err := exec.Command("sh","-c","docker inspect --format='{{.Name}}|{{.State.Running}}|{{.Config.Image}}|{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -qa)").Output()  
          if err != nil {log.Fatal(err)}
    
    s := strings.Split(string(out), "\n")
    

    for i := range s[:len(s)-1] {

        infos := strings.Split(s[i], "|")

        containers[infos[0][1:3]] = append(containers[infos[0][1:3]], infos[1])
        containers[infos[0][1:3]] = append(containers[infos[0][1:3]], infos[2])
        if infos[3] == "" {infos[3]= "none"}
        containers[infos[0][1:3]] = append(containers[infos[0][1:3]], infos[3])

    }
    
    return containers

}

func container_info(name string) []string{

    containers := inspect()

    return containers[name]

}

func snmpget(user string,password string,ipadrr string,oid string) string {

        out, err := exec.Command("snmpget", "-u", user, "-l", "authPriv", "-a", "MD5", "-x", "DES", "-A", password, "-X", password, ipadrr, oid).Output()
                if err != nil {
                       // log.Fatal(err)
                        return "dead"
                } else {
        return string(out)                    
        }
}

func main(){


    fmt.Print("Container name: ")
    var input string
    fmt.Scanln(&input)

    fmt.Println("Running time : "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.2.1.1.3.0"))
    fmt.Println("Total Processes : "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.2.1.25.1.6.0"))
    fmt.Println("Total Swap Size: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.4.3.0"))
    fmt.Println("Available Swap Space: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.4.4.0"))
    fmt.Println("Total RAM used: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.4.6.0"))
    fmt.Println("Total RAM Free: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.4.11.0"))
    fmt.Println("Total size of the disk/partion (kBytes): "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.9.1.6.1"))
    fmt.Println("Available space on the disk: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.9.1.7.1"))
    fmt.Println("Used space on the disk: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.9.1.8.1"))
    fmt.Println("Percentage of space used on disk: "+snmpget("bootstrap","azertyui",container_info(input)[2],"1.3.6.1.4.1.2021.9.1.9.1"))

}   

