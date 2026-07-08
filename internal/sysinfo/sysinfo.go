package sysinfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDetails() {
	fmt.Println("--- System Information ---")

	// 1. OS
	osCmd := exec.Command("uname", "-sr")
	osOut, _ := osCmd.Output()
	fmt.Printf("OS:          %s", string(osOut))

	// 2. Motherboard
	vendorCmd := exec.Command("cat", "/sys/class/dmi/id/board_vendor")
	vendorOut, _ := vendorCmd.Output()
	vendorStr := strings.TrimSpace(string(vendorOut))

	nameCmd := exec.Command("cat", "/sys/class/dmi/id/board_name")
	nameOut, _ := nameCmd.Output()
	nameStr := strings.TrimSpace(string(nameOut))

	mbStr := vendorStr + " " + nameStr
	if strings.TrimSpace(mbStr) == "" {
		mbStr = "Unknown Motherboard"
	}
	fmt.Printf("Motherboard: %s\n", mbStr)

	// 3. CPU
	cpuCmd := exec.Command("sh", "-c", "lscpu | grep 'Model name:'")
	cpuOut, _ := cpuCmd.Output()
	cpuStr := strings.TrimSpace(string(cpuOut))
	if cpuStr == "" {
		altCPUCmd := exec.Command("sh", "-c", "grep -m 1 'model name' /proc/cpuinfo")
		altCPUOut, _ := altCPUCmd.Output()
		cpuStr = strings.TrimSpace(string(altCPUOut))
	}
	cpuStr = strings.TrimPrefix(cpuStr, "Model name:")
	cpuStr = strings.TrimPrefix(cpuStr, "model name	:")
	fmt.Printf("CPU:         %s\n", strings.TrimSpace(cpuStr))

	// 4. RAM
	ramCmd := exec.Command("sh", "-c", "grep 'MemTotal' /proc/meminfo")
	ramOut, _ := ramCmd.Output()
	ramStr := strings.TrimSpace(string(ramOut))
	ramStr = strings.TrimPrefix(ramStr, "MemTotal:")
	fmt.Printf("RAM (Total): %s\n", strings.TrimSpace(ramStr))

	// 5. GPU
	gpuCmd := exec.Command("sh", "-c", "lspci | grep -E \"VGA|3D\" | sed 's/.*: //'")
	gpuOut, _ := gpuCmd.Output()
	gpuStr := strings.TrimSpace(string(gpuOut))
	fmt.Printf("GPU:         %s\n", gpuStr)

	// 6. Disk Space
	diskCmd := exec.Command("sh", "-c", "df -h / | tail -n 1 | awk '{print $4 \" free (used \" $5 \")\"}'")
	diskOut, _ := diskCmd.Output()
	fmt.Printf("Disk (/):    %s", string(diskOut))

	// 7. Local IP
	ipCmd := exec.Command("sh", "-c", "hostname -I | awk '{print $1}'")
	ipOut, _ := ipCmd.Output()
	fmt.Printf("Local IP:    %s", string(ipOut))
}