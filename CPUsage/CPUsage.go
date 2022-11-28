package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Percentage of CPU usage in 1 second intervals:") // n-> 1,2,3....n
	var firstIdle, lastIdle uint64

	for i := 0; i < 11; i++ {
		file, err := os.Open("/proc/stat")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		//cat /proc/stat => cpu  1061998 2252 267626 14677568 11996 0 14516 0 0 0
		firstLine := scanner.Text()[5:] // delete cpu[][]
		file.Close()
		//1061998 2252 267626 14677568 11996 0 14516 0 0 0 #After deletion
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		split := strings.Fields(firstLine)
		idleTime, _ := strconv.ParseUint(split[3], 10, 64) // ("14677568",10(Decimal),64(uint64)
		totalTime := uint64(0)
		for _, s := range split {
			u, _ := strconv.ParseUint(s, 10, 64)
			totalTime += u // 1061998+2252+267626...+0 -> sum of split
		}

		if i > 0 {
			alfaIdleTime := idleTime - firstIdle
			alfaTotalTime := totalTime - lastIdle
			//fmt.Println(idleTime, totalTime, alfaIdleTime, alfaTotalTime)
			Idle_Cpu_Percent := float64(alfaIdleTime) / float64(alfaTotalTime) // Free CPU %
			cpuUsagePerSec := (1.0 - Idle_Cpu_Percent) * 100.0
			if cpuUsagePerSec >= 10 {
				fmt.Printf("after %d second: CPUsage: %.1f \n", i, cpuUsagePerSec)
				fmt.Println("WARNING --- HIGH CPU USAGE!!!---CPU usage exceeded 10% ")
			} else {
				fmt.Printf("after %d second: CPUsage: %.1f \n", i, cpuUsagePerSec)
			}
		}
		firstIdle = idleTime
		lastIdle = totalTime

		time.Sleep(time.Second * 1)
	}
}
