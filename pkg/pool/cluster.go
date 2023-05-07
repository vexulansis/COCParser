package pool

type Cluster struct {
	Pool         *Pool
	Generator    *Generator
	ErrorHandler *ErrorHandler
}
type ClusterConfig struct {
	PoolSize         int
	GeneratorSize    int
	ErrorHandlerSize int
}

func NewCluster(name string, config ClusterConfig) {

}
