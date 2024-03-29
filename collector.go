// A Stats collector for StatHat

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"time"
	"strconv"
)

func main() {
	client := http.Client{}

	for {

		si := syscall.Sysinfo_t{}
		syscall.Sysinfo(&si)

		l := strconv.FormatFloat(float64(si.Loads[0])    /  65536,'f', 2 ,    )
		fmt.Println(l, si.Loads[0]/65536.0, si.Loads[0])


		time.Sleep(1 * time.Second)
		continue

		resp, err := client.PostForm("https://api.stathat.com/ez", url.Values{
			"stat":  {"loadavg"},
			"ezkey": {"YLJRun7adtSFKR2u"},
			"value": {l},
		})
		if err != nil {
			// Problems...
			fmt.Println("Error sending stat... : %v", err)
			os.Exit(1)
		}
		fmt.Println(resp)
		resp.Body.Close()
		time.Sleep(1 * time.Second)
	}
}
