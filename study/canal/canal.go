package canal

import (
	"context"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"os"
)

/**

Binlog output looks like:
=== RotateEvent ===
Date: 1970-01-01 08:00:00
Log position: 0
Event size: 43
Position: 4
Next log name: mysql-bin.000001

=== FormatDescriptionEvent ===
Date: 2020-11-24 14:09:41
Log position: 123
Event size: 119
Version: 4
Server version: 5.7.30-log
Checksum algorithm: 1

=== QueryEvent ===
Date: 2020-11-24 14:36:00
Log position: 404
Event size: 185
Slave proxy ID: 3
Execution time: 0
Error code: 0
Schema: test-db
Query: CREATE USER 'canal'@'%' IDENTIFIED WITH 'mysql_native_password' AS '*E3619321C1A937C46A0D8BD1DAC39F93B27D4458'

=== AnonymousGTIDEvent ===
Date: 2020-11-24 14:36:00
Log position: 469
Event size: 65
Commit flag: 1
GTID_NEXT: 00000000-0000-0000-0000-000000000000:0
LAST_COMMITTED: 1
SEQUENCE_NUMBER: 2
Immediate commmit timestamp: 0 (<n/a>)
Orignal commmit timestamp: 0 (<n/a>)
Transaction length: 0
Immediate server version: 0
Orignal server version: 0

*/
func ReadBinlog() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "123456",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	// Start sync with specified binlog file and position
	streamer, err := syncer.StartSync(mysql.Position{
		Name: "",
		Pos:  1,
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	i := 0
	for {
		event, err := streamer.GetEvent(ctx)
		if err != nil {
			panic(err)
		}
		event.Dump(os.Stdout)
		i++
		if i >= 10 {
			break
		}
	}
}

func Sync() {
	cfg := canal.NewDefaultConfig()
	cfg.Password = "123456"
	// only care table canal_test_user and canal_test_teacher in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"canal_test_user", "canal_test_teacher"}

	client, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	// register event handler
	client.SetEventHandler(&syncEventHandler{})

	if err := client.Run(); err != nil {
		panic(err)
	}
}
