package test

import (
	"context"
	"testing"

	"github.com/samuelves/correiosgo/consts"
	"github.com/samuelves/correiosgo/servicos/consulta"
	"github.com/stretchr/testify/assert"
)

func TestSimpleRequest(t *testing.T) {
	r := consulta.NewFreteRequest("01243000", "65299970").SetServicos(consts.SERVICE_PAC_04510)
	resp, err := consulta.CalcularFrete(context.Background(), r)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
