package tsvadaptor

import (
	"github.com/ProtoML/ProtoML/formatadaptor/delimiteradaptor"
)
 
func New() *delimiteradaptor.DelimiterAdaptor {
	return delimiteradaptor.New('\t')
}
