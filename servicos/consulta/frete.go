package consulta

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/samuelves/correiosgo/consts"
	"github.com/samuelves/correiosgo/utils"
	"github.com/shopspring/decimal"
)

var FreteEndpoint = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx"

type RequestMode int

type FreteRequestStruct struct {
	CepOrigem        string
	CepDestino       string
	PesoKg           decimal.Decimal
	ComprimentoCm    decimal.Decimal
	AlturaCm         decimal.Decimal
	LarguraCm        decimal.Decimal
	Servicos         []consts.TipoServico
	ValorDeclarado   decimal.Decimal
	AvisoRecebimento bool
	CdEmpresa        string
	DsSenha          string
	Mode             RequestMode
}

type ServicoResponseStruct struct {
	Tipo                  consts.TipoServico
	Preco                 decimal.Decimal
	PrazoEntregaDias      int
	PrecoSemAdicionais    decimal.Decimal
	PrecoMaoPropria       decimal.Decimal
	PrecoAvisoRecebimento decimal.Decimal
	PrecoValorDeclarado   decimal.Decimal
	EntregaDomiciliar     bool
	EntregaSabado         bool
	Erro                  *ServicoResponseErrorStruct
	ErroMsg               string
}

type ServicoResponseErrorStruct struct {
	Codigo consts.TipoErro
}

type FreteResponseStruct struct {
	Servicos map[consts.TipoServico]ServicoResponseStruct
}

type servicoRespStruct struct {
	Codigo                string
	Valor                 string
	PrazoEntrega          int
	ValorSemAdicionais    string
	ValorMaoPropria       string
	ValorAvisoRecebimento string
	ValorValorDeclarado   string
	EntregaDomiciliar     string
	EntregaSabado         string
	Erro                  int
	MsgErro               string
}

const (
	RequestModeAuto RequestMode = iota
	RequestModeSingle
	RequestModeCombined
)

func (r *FreteRequestStruct) SetServicos(srvs ...consts.TipoServico) *FreteRequestStruct {
	r.Servicos = make([]consts.TipoServico, 0)
	r.Servicos = append(r.Servicos, srvs...)
	return r
}

func (r *FreteRequestStruct) SetContrato(cdEmpresa string) *FreteRequestStruct {
	r.CdEmpresa = cdEmpresa
	return r
}

func (r *FreteRequestStruct) SetSenha(dsSenha string) *FreteRequestStruct {
	r.DsSenha = dsSenha
	return r
}

func (r *FreteRequestStruct) SetPeso(peso decimal.Decimal) *FreteRequestStruct {
	r.PesoKg = peso
	return r
}

func (r *FreteRequestStruct) SetDimensoes(comprimentoCm decimal.Decimal, alturaCm decimal.Decimal, larguraCm decimal.Decimal) *FreteRequestStruct {
	r.ComprimentoCm = comprimentoCm
	r.AlturaCm = alturaCm
	r.LarguraCm = larguraCm
	return r
}

func (r *FreteRequestStruct) AppendServico(srv consts.TipoServico) *FreteRequestStruct {
	for _, v := range r.Servicos {
		if v == srv {
			return r
		}
	}
	r.Servicos = append(r.Servicos, srv)
	return r
}

func NewFreteRequest(cepOrigem, cepDestino string) *FreteRequestStruct {
	return &FreteRequestStruct{
		CepOrigem:      cepOrigem,
		CepDestino:     cepDestino,
		PesoKg:         decimal.NewFromFloat(0.5),
		ComprimentoCm:  decimal.NewFromFloat(16.0),
		LarguraCm:      decimal.NewFromFloat(11.0),
		AlturaCm:       decimal.NewFromFloat(5.0),
		Servicos:       []consts.TipoServico{consts.SERVICE_PAC_41068, consts.SERVICE_SEDEX_41556},
		ValorDeclarado: decimal.NewFromFloat(0.0),
	}
}

func CalcularFrete(ctx context.Context, req *FreteRequestStruct) (*FreteResponseStruct, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}
	// desde 2019, os Correios não aceitam consultas múltiplas caso não seja
	// informado o código da empresa + senha
	if len(req.Servicos) > 1 &&
		((req.Mode == RequestModeAuto && req.CdEmpresa == "") || (req.Mode == RequestModeSingle)) {
		reqs := make([]*FreteRequestStruct, len(req.Servicos))
		for k := range req.Servicos {
			clone := &FreteRequestStruct{
				CepOrigem:        utils.FixCEP(req.CepOrigem),
				CepDestino:       utils.FixCEP(req.CepDestino),
				PesoKg:           req.PesoKg,
				ComprimentoCm:    req.ComprimentoCm,
				AlturaCm:         req.AlturaCm,
				LarguraCm:        req.LarguraCm,
				Servicos:         []consts.TipoServico{req.Servicos[k]},
				ValorDeclarado:   req.ValorDeclarado,
				AvisoRecebimento: req.AvisoRecebimento,
				CdEmpresa:        req.CdEmpresa,
				DsSenha:          req.DsSenha,
			}
			reqs[k] = clone
		}
		r00 := &FreteResponseStruct{
			Servicos: make(map[consts.TipoServico]ServicoResponseStruct),
		}
		for i, v := range reqs {
			rsp, err := CalcularFrete(ctx, v)
			if err != nil && len(reqs) == i+1 {
				return r00, err
			} else if err != nil {
				continue
			}
			for k2, v2 := range rsp.Servicos {
				r00.Servicos[k2] = v2
			}
		}
		return r00, nil
	}
	v := url.Values{}
	v.Set("sCepOrigem", strings.Trim(req.CepOrigem, "-"))
	v.Set("sCepDestino", strings.Trim(req.CepDestino, "-"))
	v.Set("nVlPeso", req.PesoKg.String())
	v.Set("nCdFormato", "1")
	v.Set("nVlComprimento", req.ComprimentoCm.String())
	v.Set("nVlAltura", req.AlturaCm.String())
	v.Set("nVlLargura", req.LarguraCm.String())
	v.Set("StrRetorno", "xml")
	svcs := make([]string, len(req.Servicos))
	for k, v := range req.Servicos {
		svcs[k] = string(v)
	}
	v.Set("nCdServico", strings.Join(svcs, ","))
	v.Set("nVlValorDeclarado", req.ValorDeclarado.String())
	if req.AvisoRecebimento {
		v.Set("sCdAvisoRecebimento", "S")
	}
	if req.CdEmpresa != "" {
		v.Set("nCdEmpresa", req.CdEmpresa)
		v.Set("sDsSenha", req.DsSenha)
	}

	rq0, _ := http.NewRequest(http.MethodGet, FreteEndpoint+"?"+v.Encode(), nil)
	rq0 = rq0.WithContext(ctx)

	cresp, err := http.DefaultClient.Do(rq0)
	if err != nil {
		return nil, err
	}
	defer cresp.Body.Close()

	rrbuf := new(bytes.Buffer)
	io.Copy(rrbuf, cresp.Body)
	p := xml.NewDecoder(rrbuf)
	p.CharsetReader = utils.CharsetReader

	vlov := struct {
		XMLName string              `xml:"Servicos"`
		Values  []servicoRespStruct `xml:"cServico"`
	}{}

	err = p.Decode(&vlov)
	if err != nil {
		fmt.Println("CORREIOS: " + rrbuf.String())
		return nil, err
	}
	//
	output := &FreteResponseStruct{
		Servicos: make(map[consts.TipoServico]ServicoResponseStruct),
	}
	//
	for _, v := range vlov.Values {
		v2 := ServicoResponseStruct{}
		v2.Tipo = consts.TipoServico(v.Codigo)
		v2.Preco, _ = decimal.NewFromString(utils.FixDecimals(v.Valor))
		v2.PrazoEntregaDias = v.PrazoEntrega
		v2.PrecoSemAdicionais, _ = decimal.NewFromString(utils.FixDecimals(v.ValorSemAdicionais))
		v2.PrecoMaoPropria, _ = decimal.NewFromString(utils.FixDecimals(v.ValorMaoPropria))
		v2.PrecoAvisoRecebimento, _ = decimal.NewFromString(utils.FixDecimals(v.ValorAvisoRecebimento))
		v2.PrecoValorDeclarado, _ = decimal.NewFromString(utils.FixDecimals(v.ValorValorDeclarado))
		v2.EntregaDomiciliar = (v.EntregaDomiciliar == "S")
		v2.EntregaSabado = (v.EntregaSabado == "S")
		if v.Erro != 0 {
			er9 := &ServicoResponseErrorStruct{
				Codigo: consts.TipoErro(v.Erro),
			}
			v2.Erro = er9
			v2.ErroMsg = v.MsgErro
		}
		output.Servicos[v2.Tipo] = v2
	}
	return output, nil
}
