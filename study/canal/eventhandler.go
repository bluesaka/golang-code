package canal

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
)

/**
refer https://github.com/go-mysql-org/go-mysql-elasticsearch/blob/master/river/sync.go
 */
type syncEventHandler struct{}

func (h *syncEventHandler) OnRotate(rotateEvent *replication.RotateEvent) error {
	return nil
}

// OnTableChanged is called when the table is created, altered, renamed or dropped.
// You need to clear the associated data like cache with the table.
// It will be called before OnDDL.
func (h *syncEventHandler) OnTableChanged(schema string, table string) error {
	fmt.Println("OnTableChanged:", schema, table)
	return nil
}

func (h *syncEventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	fmt.Println("OnDDL:", nextPos, queryEvent)
	return nil
}

func (h *syncEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Println("OnRow:", e)
	return nil
}

func (h *syncEventHandler) OnXID(nextPos mysql.Position) error {
	fmt.Println("OnXID:", nextPos)
	return nil
}

func (h *syncEventHandler) OnGTID(gtid mysql.GTIDSet) error {
	fmt.Println("OnGTID:", gtid)
	return nil
}

// OnPosSynced Use your own way to sync position. When force is true, sync position immediately.
func (h *syncEventHandler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	fmt.Println("OnPosSynced:", pos, set, force)
	return nil
}

func (h *syncEventHandler) String() string {
	return "my handler"
}
