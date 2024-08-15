package src

type (
	LinuxResp struct {
		MemoryInfo  *Memory        `json:"memoryInfo"`
		CpuInfo     *Cpu           `json:"cpuInfo"`
		DiskInfo    *Disk          `json:"diskInfo"`
		NetInfo     *ConnIO        `json:"netInfo"`
		ProcessInfo []*ProcessInfo `json:"processInfo"`
	}

	Memory struct {
		TotalMemory    float64 `json:"totalMemory"`    //总内存--单位G
		UsedMemory     float64 `json:"usedMemory"`     //使用内存--单位G
		UsedMemoryRate float64 `json:"usedMemoryRate"` //使用内存比例
		FreeMemory     float64 `json:"freeMemory"`     //空闲内存--单位G
		FreeMemoryRate float64 `json:"freeMemoryRate"` //空闲内存比例
		Active         float64 `json:"active"`         //活跃内存--单位G --正在被程序使用的内存页面。这些页面可能被频繁访问,不太容易被换出到磁盘。
		ActiveRate     float64 `json:"activeRate"`     //活跃内存比例
		Inactive       float64 `json:"inactive"`       //非活跃内存 --单位G --最近未被频繁访问的页面,可能被换出到磁盘以释放内存空间。
		InactiveRate   float64 `json:"inactiveRate"`   //非活跃内存比例
		Wired          float64 `json:"wired"`          //固定内存 --单位G --由内核使用的内存,如swap分区,内存映射文件,共享内存,和共享库。
		WiredRate      float64 `json:"wiredRate"`      //固定内存比例
	}

	CpuInfo struct {
		CpuName string  `json:"cpuName"` //核心名字
		CpuRate float64 `json:"cpuRate"` //核心利用率
	}

	Cpu struct {
		AVGCpuRate float64    `json:"AVGCpuRate"` //cpu平均利用率
		CpuInfos   []*CpuInfo `json:"cpuInfos"`
	}

	//磁盘分区IO
	IOCountersStat struct {
		ReadCount        uint64 `json:"readCount"`        //从磁盘读取的总次数
		MergedReadCount  uint64 `json:"mergedReadCount"`  //合并的读操作次数。操作系统可能会将多个小的读操作合并成一个更大的读操作来提高效率。
		WriteCount       uint64 `json:"writeCount"`       //写入磁盘的总次数
		MergedWriteCount uint64 `json:"mergedWriteCount"` //合并的写操作次数
		ReadBytes        uint64 `json:"readBytes"`        //从磁盘读取的总字节数
		WriteBytes       uint64 `json:"writeBytes"`       //写入磁盘的总字节数
		ReadTime         uint64 `json:"readTime"`         //读操作总耗时(以毫秒为单位)
		WriteTime        uint64 `json:"writeTime"`        //写操作总耗时(以毫秒为单位)
		IopsInProgress   uint64 `json:"iopsInProgress"`   //当前正在进行的 I/O 操作数
		IoTime           uint64 `json:"ioTime"`           //磁盘 I/O 总耗时(以毫秒为单位)
		WeightedIO       uint64 `json:"weightedIO"`       //加权 I/O 操作时间。这是一个综合指标,它考虑了 I/O 操作的大小和等待时间
		Name             string `json:"name"`             //磁盘设备名称
		SerialNumber     string `json:"serialNumber"`     //磁盘序列号
		Label            string `json:"label"`            //磁盘标签
	}
	SKCountersStat struct {
		Filesystem string  `json:"filesystem"` //文件系统
		MountEd    string  `json:"mountEd"`    //挂载点
		Total      float64 `json:"total"`      //总大小
		Used       float64 `json:"used"`       //已使用大小
		UsedRate   float64 `json:"usedRate"`   //已使用比例
		Free       float64 `json:"free"`       //剩余大小
		FreeRate   float64 `json:"freeRate"`   //剩余比例
	}

	//磁盘
	Disk struct {
		DiskSK []*SKCountersStat `json:"diskSk"` //磁盘信息
		DiskIO []*IOCountersStat `json:"diskIO"` //磁盘分区IO信息
	}
	Addr struct {
		IP   string `json:"ip"`
		Port uint32 `json:"port"`
	}

	//网络连接
	ConnectionStat struct {
		Fd     uint32  `json:"fd"`         //连接对应的文件描述符
		Family uint32  `json:"family"`     //表示地址族
		Type   uint32  `json:"type"`       //表示连接的类型，1TCP 流式套接字 2UDP 数据报套接字 3原始套接字 4可靠数据报套接字 5有序数据报套接字
		Laddr  Addr    `json:"localaddr"`  //本地地址，主机和端口
		Raddr  Addr    `json:"remoteaddr"` //远程地址,主机和端口
		Status string  `json:"status"`     //连接的状态，ESTABLISHED、LISTEN、CLOSE_WAIT 等
		Uids   []int32 `json:"uids"`       //与此连接相关的用户 ID 列表
		Pid    int32   `json:"pid"`        //此连接相关的进程 ID
	}
	//网络传输
	ConnIOCountersStat struct {
		Name        string `json:"name"`        // 网络接口名称
		BytesSent   uint64 `json:"bytesSent"`   // 发送的字节数
		BytesRecv   uint64 `json:"bytesRecv"`   // 接收的字节数
		PacketsSent uint64 `json:"packetsSent"` // 发送的数据包数量
		PacketsRecv uint64 `json:"packetsRecv"` // 接收的数据包数量
		Errin       uint64 `json:"errin"`       // 接收过程中发生的错误总数
		Errout      uint64 `json:"errout"`      // 发送过程中发生的错误总数
		Dropin      uint64 `json:"dropin"`      // 丢弃的入站数据包数量
		Dropout     uint64 `json:"dropout"`     // 丢弃的出站数据包数量(在 OSX 和 BSD 上始终为 0)
		Fifoin      uint64 `json:"fifoin"`      // 接收过程中 FIFO 缓冲区发生错误的次数
		Fifoout     uint64 `json:"fifoout"`     // 发送过程中 FIFO 缓冲区发生错误的次数
	}
	//网络
	ConnIO struct {
		ConnIOCountersStat []*ConnIOCountersStat `json:"connIOCountersStat"` //网络传输
		ConnectionStat     []*ConnectionStat     `json:"connectionStat"`     //网络连接
	}
	//进程信息
	ProcessInfo struct {
		Pid      int32   `json:"pid"`
		MemUsage float64 `json:"memUsage"` //内存使用率
		CpuUsage float64 `json:"cpuUsage"` //cpu使用率
		Command  string  `json:"command"`  //进程启动信息
	}
)
