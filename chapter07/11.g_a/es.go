package main

// batch

//var shareCli = esClientWithBuffer{}

type esClientWithBuffer struct {
	batchBuffer [][]interface{}
	messageChan chan interface{}
	shortBuffer []interface{} //满了之后扔到batchBuffer
}

func (cli *esClientWithBuffer) push() {
	// todo 队列操作
}

func (cli *esClientWithBuffer) prepareBatch() {
	for msg := range cli.messageChan {
		if len(cli.shortBuffer) == batchSize {
			cli.batchBuffer = append(cli.batchBuffer, cli.shortBuffer)
			cli.shortBuffer = []interface{}{}
		}
		cli.shortBuffer = append(cli.shortBuffer, msg)
	}
}

func pushToElasticSearchService(data interface{}) {
	//	todo ...

}

var batchSize int = 20
