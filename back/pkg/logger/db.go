package logger

import "extractor/back/pkg/db"

const (
	GET      = 1
	ADD      = 2
	MODIFY   = 3
	REMOVE   = 4
	GENERATE = 5

	DATA       = 0
	REPORT     = 1
	PARAMETER  = 2
	DEPARTMENT = 3
)

func DB(methodID int32, typeID int32, tupleID int32, status bool) {
	db.InvokeExec(`INSERT INTO
					logs (
						[method_id],
						[type_id],
						[tuple_id],
						[status]
					)
					values (
						@p1, @p2, @p3, @p4
					)`,
		methodID, typeID, tupleID, status)
}
