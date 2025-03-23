package inst

import (
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"

	"github.com/openark/orchestrator/go/db"
)

func WriteClusterUserMessage(message ClusterUserMessage) error {
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`
			insert into
					cluster_user_messages (cluster_name, level, message, created_at)
				values
					(?, ?, ?, NOW())
			`,
			message.ClusterName, message.Level, message.Message)
		return log.Errore(err)
	}
	return ExecDBWriteFunc(writeFunc)
}

func ReadClusterUserMessage(clusterName string) (result []ClusterUserMessage, err error) {
	query := `
		select
			message_id, level, message, created_at 
		from
			cluster_user_messages
		where
			cluster_name = ?
		order by
			created_at DESC
		`
	err = db.QueryOrchestrator(query, sqlutils.Args(clusterName), func(m sqlutils.RowMap) error {
		message := ClusterUserMessage{
			ID:          m.GetInt64("message_id"),
			ClusterName: clusterName,
			Level:       m.GetString("level"),
			Message:     m.GetString("message"),
			CreatedAt:   m.GetTime("created_at"),
		}
		result = append(result, message)
		return nil
	})

	return result, log.Errore(err)
}

func AckClusterUserMessage(messageId int64) error {
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`delete from cluster_user_messages where message_id = ?`, messageId)
		return log.Errore(err)
	}
	return ExecDBWriteFunc(writeFunc)
}
