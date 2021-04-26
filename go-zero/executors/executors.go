/**
 * @link  https://zeromicro.github.io/go-zero/executors.html
 * 在 go-zero 中，executors 充当任务池，做多任务缓冲，使用做批量处理的任务。
 * 如：clickhouse 大批量 insert，sql batch insert。同时也可以在 go-queue 看到 executors 【在 queue 里面使用的是 ChunkExecutor ，限定任务提交字节大小】。
 * 所以当你存在以下需求，都可以使用这个组件：
 *  - 批量提交任务
 *  - 缓冲一部分任务，惰性提交
 *  - 延迟任务提交
 */
package executors

import (
	"github.com/tal-tech/go-zero/core/executors"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"time"
)

type DailyTask struct {
	insertExecutor *executors.BulkExecutor
	mysqlConn      sqlx.SqlConn
}

func (dt DailyTask) Init() {
	dt.insertExecutor = executors.NewBulkExecutor(
		func(tasks []interface{}) {

		},
		executors.WithBulkTasks(10240),
		executors.WithBulkInterval(time.Second*3),
	)
}

func (dt DailyTask) insertNewData(ch chan interface{}) {
	//for items := range ch {
	//	if v, ok := items; !ok {
	//		continue
	//	}
	//}
	_ = dt.insertExecutor.Add([]interface{}{})
	dt.insertExecutor.Flush()
	dt.insertExecutor.Wait()

}
