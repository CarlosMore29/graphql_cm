package graphql_cm_test

import (
	"fmt"

	"github.com/CarlosMore29/graphql_cm"

	"testing"
)

func TestGeneral(t *testing.T) {

	var obj *graphql_cm.APIQL = graphql_cm.InitAPIQL()
	fmt.Println(obj)
	// if false {
	// 	t.Error("Error al marshal:  %s")
	// }
}
