package baseorbitdb

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("orbitdb.baseorbitdb")
}
