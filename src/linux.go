package src

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"log"
	"sort"
)

type Linux struct{}

func NewLinux() *Linux {
	return &Linux{}
}

// GetMemoryInfo 内存信息
func (l *Linux) GetMemoryInfo() (*Memory, error) {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting virtual memory info:", err)
		return nil, err
	}
	GBNumber := uint64(1024 * 1024 * 1024) //单位GB
	//总内存
	totalMemory := MixCompute("a/b", map[rune]float64{
		'a': float64(virtualMemory.Total),
		'b': float64(GBNumber),
	})
	//使用内存
	usedMemory := MixCompute("a/b", map[rune]float64{
		'a': float64(virtualMemory.Used),
		'b': float64(GBNumber),
	})
	//使用比例
	usedMemoryRate := float64(0)
	if virtualMemory.Used != 0 {
		usedMemoryRate = MixCompute("a/b*c", map[rune]float64{
			'a': float64(virtualMemory.Used),
			'b': float64(virtualMemory.Total),
			'c': 100,
		})
	}

	//可用内存
	freeMemory := MixCompute("a/b", map[rune]float64{
		'a': float64(virtualMemory.Available),
		'b': float64(GBNumber),
	})
	//可用内存比例
	freeMemoryRate := float64(0)
	if virtualMemory.Available != 0 {
		freeMemoryRate = MixCompute("a/b*c", map[rune]float64{
			'a': float64(virtualMemory.Available),
			'b': float64(virtualMemory.Total),
			'c': 100,
		})
	}
	//活跃内存
	activeMemory := MixCompute("a/b", map[rune]float64{
		'a': float64(virtualMemory.Active),
		'b': float64(GBNumber),
	})
	//活跃内存比例
	activeMemoryRate := float64(0)
	if virtualMemory.Active != 0 {
		activeMemoryRate = MixCompute("a/b*c", map[rune]float64{
			'a': float64(virtualMemory.Active),
			'b': float64(virtualMemory.Total),
			'c': 100,
		})
	}
	//非活跃内存
	inactiveMemory := MixCompute("a/b", map[rune]float64{
		'a': float64(virtualMemory.Inactive),
		'b': float64(GBNumber),
	})
	//非活跃内存比例
	inactiveMemoryRate := float64(0)
	if virtualMemory.Inactive != 0 {
		inactiveMemoryRate = MixCompute("a/b*c", map[rune]float64{
			'a': float64(virtualMemory.Inactive),
			'b': float64(virtualMemory.Total),
			'c': 100,
		})
	}
	//固定内存
	wiredMemory := MixCompute("a/b", map[rune]float64{
		'a': float64(virtualMemory.Wired),
		'b': float64(GBNumber),
	})
	//固定内存比例
	wiredMemoryRate := float64(0)
	if virtualMemory.Wired != 0 {
		wiredMemoryRate = MixCompute("a/b*c", map[rune]float64{
			'a': float64(virtualMemory.Wired),
			'b': float64(virtualMemory.Total),
			'c': 100,
		})
	}
	m := &Memory{
		TotalMemory:    totalMemory,
		UsedMemory:     usedMemory,
		UsedMemoryRate: usedMemoryRate,
		FreeMemory:     freeMemory,
		FreeMemoryRate: freeMemoryRate,
		Active:         activeMemory,
		ActiveRate:     activeMemoryRate,
		Inactive:       inactiveMemory,
		InactiveRate:   inactiveMemoryRate,
		Wired:          wiredMemory,
		WiredRate:      wiredMemoryRate,
	}
	return m, nil
}

// GetCpuInfo cpu信息
func (l *Linux) GetCpuInfo() (*Cpu, error) {
	// 获取 CPU 利用率信息
	cpuInfos, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Println("Error getting CPU info:", err)
		return nil, err
	}

	cpuList := []*CpuInfo{}
	// 输出每个 CPU 核心的利用率
	total := float64(0)
	for i, utilization := range cpuInfos {
		cupName := fmt.Sprintf("CPU Core %d", i)
		l := CpuInfo{
			CpuName: cupName,
			CpuRate: utilization,
		}
		total += utilization
		cpuList = append(cpuList, &l)
	}

	aVGCpuRate := MixCompute("a/b", map[rune]float64{
		'a': total,
		'b': float64(len(cpuInfos)),
	})

	m := &Cpu{
		AVGCpuRate: aVGCpuRate,
		CpuInfos:   cpuList,
	}
	return m, nil

}

// GetDiskInfo 磁盘信息
func (l *Linux) GetDiskInfo() (*Disk, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	GBNumber := uint64(1024 * 1024 * 1024) //单位GB
	skList := make([]*SKCountersStat, 0)
	for _, part := range partitions {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			continue
		}

		total := MixCompute("a/b", map[rune]float64{
			'a': float64(usage.Total),
			'b': float64(GBNumber),
		})
		used := MixCompute("a/b", map[rune]float64{
			'a': float64(usage.Used),
			'b': float64(GBNumber),
		})
		free := MixCompute("a/b", map[rune]float64{
			'a': float64(usage.Free),
			'b': float64(GBNumber),
		})
		freeRate := float64(0)
		if usage.Total != 0 {
			freeRate = MixCompute("(a/b)*c", map[rune]float64{
				'a': float64(usage.Free),
				'b': float64(usage.Total),
				'c': 100,
			})
		}
		d := SKCountersStat{
			Filesystem: part.Device,
			MountEd:    part.Mountpoint,
			Total:      total,
			Used:       used,
			UsedRate:   usage.UsedPercent,
			Free:       free,
			FreeRate:   freeRate,
		}
		skList = append(skList, &d)
	}

	//显示磁盘分区IO信息
	diskIOs, err := disk.IOCounters()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	diskIOList := make([]*IOCountersStat, 0)
	for _, v := range diskIOs {
		skIo := IOCountersStat{
			ReadCount:        v.ReadCount,
			MergedReadCount:  v.MergedReadCount,
			WriteCount:       v.WriteCount,
			MergedWriteCount: v.MergedWriteCount,
			ReadBytes:        v.ReadBytes,
			WriteBytes:       v.WriteBytes,
			ReadTime:         v.ReadTime,
			WriteTime:        v.WriteTime,
			IopsInProgress:   v.IopsInProgress,
			IoTime:           v.IoTime,
			WeightedIO:       v.WeightedIO,
			Name:             v.Name,
			SerialNumber:     v.SerialNumber,
			Label:            v.Label,
		}
		diskIOList = append(diskIOList, &skIo)
	}
	dList := &Disk{
		DiskSK: skList,
		DiskIO: diskIOList,
	}

	return dList, nil
}

// GetNetInfo 网络信息
func (l *Linux) GetNetInfo() (*ConnIO, error) {
	counters, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	ConnIOCountersStatList := make([]*ConnIOCountersStat, 0)
	copier.Copy(&ConnIOCountersStatList, counters)
	sort.Slice(ConnIOCountersStatList, func(i, j int) bool {
		return ConnIOCountersStatList[i].BytesSent > ConnIOCountersStatList[j].BytesSent
	})

	infoList, err := net.Connections("all") //可填入tcp、udp、tcp4、udp4等等
	if err != nil {
		return nil, err
	}
	ConnectionStatList := make([]*ConnectionStat, 0)
	copier.Copy(&ConnectionStatList, infoList)
	resp := &ConnIO{
		ConnIOCountersStat: ConnIOCountersStatList,
		ConnectionStat:     ConnectionStatList,
	}
	return resp, err
}

// GetProcessInfo 进程信息
func (l *Linux) GetProcessInfo() ([]*ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		fmt.Println("Error getting processes:", err)
		return nil, err
	}
	ProcessInfoList := make([]*ProcessInfo, 0)
	// 遍历进程列表
	for _, proc := range procs {
		// 获取进程 ID
		pid := proc.Pid
		cmdline, err := proc.Cmdline()
		if err != nil {
			continue
		}

		// 获取 CPU 使用率
		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			continue
		}

		// 获取内存使用量
		memUsage := float64(0)
		memInfo, err := proc.MemoryInfo()
		if err != nil {
			continue
		} else {
			// 计算内存使用率
			totalMem, err := mem.VirtualMemory()
			if err != nil {
				continue
			} else {
				memUsage = float64(memInfo.RSS) / float64(totalMem.Total) * 100
			}
		}
		p := ProcessInfo{
			Pid:      pid,
			MemUsage: memUsage,
			CpuUsage: cpuPercent,
			Command:  cmdline,
		}
		ProcessInfoList = append(ProcessInfoList, &p)
	}
	//排序
	sort.Slice(ProcessInfoList, func(i, j int) bool {
		return ProcessInfoList[i].MemUsage > ProcessInfoList[j].MemUsage
	})
	return ProcessInfoList, nil
}
