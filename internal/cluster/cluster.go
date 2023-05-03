package cluster

type ClusterClient struct {
	Logger *ClusterLogger
}

func NewClient() (*ClusterClient, error) {
	// Creating CC example
	clusterClient := &ClusterClient{}
	// Initializing logger
	clusterClient.Logger = initClusterLogger()
	return clusterClient, nil
}
