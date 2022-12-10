package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Flags() (float64, int) {
	var DefaultCpuThreshold float64 = 20
	var DefaultInterval int = 1

	CpuThresholdPtr := flag.Float64("CpuThresold", DefaultCpuThreshold, "a float")
	IntervalPtr := flag.Int("Interval", DefaultInterval, "an int")
	flag.Parse()

	var MaxCPUsage float64
	if *CpuThresholdPtr > 0 {
		MaxCPUsage = (*CpuThresholdPtr)
	} else {
		MaxCPUsage = DefaultCpuThreshold //CPU usage threshold cannot be less than 0.
	}
	if MaxCPUsage == (DefaultCpuThreshold) {
		fmt.Printf("Default CPU threshold value selected: %.1f\n", DefaultCpuThreshold) //Default threshold Value
	} else {
		fmt.Printf("Determined max cpu threshold value: %.2f\n", *CpuThresholdPtr)
	}

	var ValueInterval int
	if *IntervalPtr > 0 {
		ValueInterval = (*IntervalPtr)
	} else {
		ValueInterval = DefaultInterval //Interval cannot be less than 0.
	}
	if ValueInterval == DefaultInterval {
		fmt.Printf("Default Interval Value Selected: %d\n ", DefaultInterval) //Default Interval
	} else {
		fmt.Printf("Determined Interval value: %d\n", *IntervalPtr)
	}

	return MaxCPUsage, ValueInterval
}

func openProcStatFile() string {
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
	return firstLine
}

func main() {
	MaxCPUsage, ValueInterval := Flags()

	fmt.Printf("\nPercentage of CPU usage in %d second intervals:\n", ValueInterval) // n-> 1,2,3....n
	var firstIdle, lastIdle uint64

	for i := 0; i < 10; i++ {
		split := strings.Fields(openProcStatFile())
		idleTime, _ := strconv.ParseUint(split[3], 10, 64) // ("14677568",10(Decimal),64(uint64)
		totalTime := uint64(0)

		for _, s := range split {
			u, _ := strconv.ParseUint(s, 10, 64)
			totalTime += u // 1061998+2252+267626...+0 -> sum of split
		}
		if i > 0 {
			alfaIdleTime := idleTime - firstIdle
			alfaTotalTime := totalTime - lastIdle
			Idle_Cpu_Percent := float64(alfaIdleTime) / float64(alfaTotalTime) // Free CPU %
			cpuUsagePerSec := (1.0 - Idle_Cpu_Percent) * 100.0
			if cpuUsagePerSec >= float64(MaxCPUsage) {
				fmt.Printf("after %d second: MaxCPUsage: %.1f \n", i*ValueInterval, cpuUsagePerSec)
				fmt.Printf("WARNING --- HIGH CPU USAGE!!!---CPU usage exceeded %.0f \n", MaxCPUsage)
			} else {
				fmt.Printf("after %d second: MaxCPUsage: %.1f \n", i*ValueInterval, cpuUsagePerSec)
			}
		}
		firstIdle = idleTime
		lastIdle = totalTime
		time.Sleep(time.Second * 1)
	}
}
