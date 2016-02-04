package jsonstruct

type Status struct {
  Cluster struct {
    ClusterName string
    LocalNode struct {
      Id string
      HostName string
      Master string
      NumberOfNodesSeen float64
    }
  }
  Os struct {
    Name string
    Version string
    Cores float64
    LodeAverage float64
  }
  Jvm struct {
    Name string
    Vendor string
    Version string
    StartTime float64
    UpTime float64
  }
  Memory struct {
    Heap map[string]float64
    NonHeap map[string]float64
  }
  Gc struct {
    CollectionTime float64
    CollectionCount float64
  }
  Index struct {
    Status string
    ActiveShards float64
    ActivePrimaryShards float64
    ActiveReplicas float64
    UnassignedShards float64
    RelocatingShards float64
    InitializingShards float64
    Documents string
    PrimaryShardsStoreSize string
    TotalStoreSize string
  }
  Product struct {
    Name string
    Version string
    Edition string
  }
}
