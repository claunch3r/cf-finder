package main
import (
    "fmt"
    "os"
    "net"
    "bufio"
    "flag"
    "strings"
)

func txtToList(filename string) ([]string, error){

    file, err := os.Open(filename)
    if err != nil {
	return nil, err
    }
    defer file.Close()
    
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	lines = append(lines, scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
    	return nil, err
    }
    
    return lines, nil
}

func checkIP(cfNetworks []string, targetIPs []string) (err error) {
    for _, targetIP := range targetIPs {
        ip := net.ParseIP(targetIP)
        if ip == nil {
            fmt.Println("\033[33m"+"[ERR] [An incorrect IP address was provided]"+"\033[0m", targetIP)
            continue
        }
        
        isCloudflare := false
        
        for _, cfNetwork := range cfNetworks {
            _, cfSubnet, _ := net.ParseCIDR(cfNetwork)
            if cfSubnet.Contains(ip) {
                fmt.Println("\033[31m"+"[!] [Cloudflare]"+"\033[0m", targetIP)
                isCloudflare = true
                break
            }
        }
        
        if !isCloudflare {
            fmt.Println("\033[32m"+"[+] [Not Cloudflare]"+"\033[0m", targetIP)
        }
    }
    
    return nil
}

func main(){
    fmt.Println(`
        ____      _____           __         
  _____/ __/     / __(_)___  ____/ /__  _____
 / ___/ /_______/ /_/ / __ \/ __  / _ \/ ___/
/ /__/ __/_____/ __/ / / / / /_/ /  __/ /    
\___/_/       /_/ /_/_/ /_/\__,_/\___/_/

                  created by @Claunch3r     
    `)
    
    var file string
    var target string
    flag.StringVar(&target, "target", "Unknown", "target ip")
    flag.StringVar(&file, "file", "Unknown", "target ip list")
    flag.Parse()
    
    cfNetworks := []string{"173.245.48.0/20","103.21.244.0/22","103.22.200.0/22","103.31.4.0/22",
    "141.101.64.0/18","108.162.192.0/18","190.93.240.0/20","188.114.96.0/20",
    "197.234.240.0/22","198.41.128.0/17","162.158.0.0/15","104.16.0.0/13",
    "104.24.0.0/14","172.64.0.0/13","131.0.72.0/22",
    "2400:cb00::/32","2606:4700::/32","2803:f800::/32","2405:b500::/32",
    "2405:8100::/32","2a06:98c0::/29","2c0f:f248::/32"}
    
    if file != "Unknown" {
    
    	targetIPs, err := txtToList(file)
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	checkIP(cfNetworks, targetIPs)
    	
    } else if target != "Unknown" {
    
    	targetIP := strings.Split(target, ",")
    	checkIP(cfNetworks, targetIP)
    	
    } else {
    	fmt.Println("\033[33m"+"[ERR] [The IP addresses were not provided]"+"\033[0m")
    }
}




