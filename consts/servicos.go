package consts

import (
	"encoding/json"
)

type TipoServico string
type Language struct {
	Codigo    TipoServico
	Descricao string
}

const (
	SERVICE_PAC_41068                                      TipoServico = "41068"
	SERVICE_PAC_04510                                      TipoServico = "04510"
	SERVICE_PAC_GRANDES_FORMATOS                           TipoServico = "41300"
	SERVICE_PAC_CONTRATO_GRANDES_FORMATOS                  TipoServico = "04693"
	SERVICE_PAC_CONTRATO_UO                                TipoServico = "04812"
	SERVICE_SEDEX_41556                                    TipoServico = "41556"
	SERVICE_SEDEX_12                                       TipoServico = "40169"
	SERVICE_SEDEX_10                                       TipoServico = "40215"
	SERVICE_SEDEX_12_CONTRATO_AGENCIA_03140                TipoServico = "03140"
	SERVICE_SEDEX_10_CONTRATO_AGENCIA_03158                TipoServico = "03158"
	SERVICE_SEDEX_10_PACOTE                                TipoServico = "40886"
	SERVICE_SEDEX_HOJE_40290                               TipoServico = "40290"
	SERVICE_SEDEX_HOJE_40878                               TipoServico = "40878"
	SERVICE_SEDEX_A_VISTA                                  TipoServico = "04014"
	SERVICE_SEDEX_VAREJO_A_COBRAR                          TipoServico = "40045"
	SERVICE_SEDEX_AGRUPADO                                 TipoServico = "41009"
	SERVICE_SEDEX_REVERSO                                  TipoServico = "40380"
	SERVICE_SEDEX_PAGAMENTO_NA_ENTREGA                     TipoServico = "04189"
	SERVICE_SEDEX_CONTRATO_UO                              TipoServico = "04316"
	SERVICE_PAC_PAGAMENTO_NA_ENTREGA                       TipoServico = "04685"
	SERVICE_CARTA_COMERCIAL_A_FATURAR                      TipoServico = "10065"
	SERVICE_CARTA_REGISTRADA                               TipoServico = "10014"
	SERVICE_SEDEX_CONTRATO_AGENCIA                         TipoServico = "04162"
	SERVICE_PAC_CONTRATO_AGENCIA                           TipoServico = "04669"
	SERVICE_SEDEX_REVERSO_CONTRATO_AGENCIA                 TipoServico = "04170"
	SERVICE_PAC_REVERSO_CONTRATO_AGENCIA                   TipoServico = "04677"
	SERVICE_CARTA_COM_A_FATURAR_SELO_E_SE                  TipoServico = "12556"
	SERVICE_CARTA_COMERCIAL_REGISTRADA_CTR_EP_MAQ_FRAN     TipoServico = "10707"
	SERVICE_MINI_ENVIOS_04227                              TipoServico = "04227"
	SERVICE_MINI_ENVIOS_04235                              TipoServico = "04235"
	SERVICE_MINI_ENVIOS_04391                              TipoServico = "04391"
	SERVICE_SEDEX_CONTRATO_GRANDES_FORMATOS_LM             TipoServico = "04146"
	SERVICE_SEDEX_CONTRATO_AGENCIA_LM                      TipoServico = "04154"
	SERVICE_SEDEX_REVERSO_LM                               TipoServico = "04243"
	SERVICE_SEDEX_CONTRATO_UO_LM                           TipoServico = "04278"
	SERVICE_PAC_CONTRATO_GRANDES_FORMATOS_LM               TipoServico = "04883"
	SERVICE_PAC_CONTRATO_AGENCIA_LM                        TipoServico = "04367"
	SERVICE_PAC_REVERSO_LM                                 TipoServico = "04375"
	SERVICE_PAC_CONTRATO_UO_LM                             TipoServico = "04332"
	SERVICE_SEDEX_CONTRATO_AGENCIA_PAGAMENTO_NA_ENTREGA_LM TipoServico = "04151"
	SERVICE_PAC_CONTRATO_AGENCIA_PAGAMENTO_NA_ENTREGA_LM   TipoServico = "04308"
	SERVICE_SEDEX_CONTRATO_AGENCIA_TA                      TipoServico = "04553"
	SERVICE_PAC_CONTRATO_AGENCIA_TA                        TipoServico = "04596"
	SERVICE_PAC_CONTRATO_AGENCIA_03298                     TipoServico = "03298"
	SERVICE_SEDEX_CONTRATO_AGENCIA_03220                   TipoServico = "03220"
	// NOVOS CODIGOS DE SERVICO DOS CORREIOS BRONZE (1o sem 2020)
	SERVICE_PAC_CONTRATO_AGENCIA_03085     TipoServico = "03085"
	SERVICE_SEDEX_CONTRATO_AGENCIA_03050   TipoServico = "03050"
	SERVICE_CARTA_REGISTRADA_AGENCIA_80250 TipoServico = "80250"
)

func GetServiceDescription(codigo TipoServico) string {

	servicesJson := `[
		{
		"codigo": "41068",
		"descricao": "PAC"
		},
		{
			"codigo": "04510",
			"descricao": "PAC"
		},
		{
			"codigo": "41300",
			"descricao": "PAC"
		},
		{
			"codigo": "04693",
			"descricao": "PAC"
		},
		{
			"codigo": "04812",
			"descricao": "PAC"
		},
		{
			"codigo": "41556",
			"descricao": "SEDEX"
		},
		{
			"codigo": "40169",
			"descricao": "SEDEX 12"
		},
		{
		"codigo":"40215",
		"descricao": "SEDEX 10"
		},
		{
		"codigo":"03140",
		"descricao": "SEDEX 12"
		},
		{
		"codigo":"03158",
		"descricao": "SEDEX 10"
		},
		{
		"codigo":"40886",
		"descricao": "SEDEX 10 PACOTE"
		},
		{
		"codigo":"40290",
		"descricao": "SEDEX HOJE"
		},
		{
			"codigo":"40878",
			"descricao": "SEDEX HOJE"
		},
		{
		"codigo":"04014",
		"descricao": "SEDEX A VISTA"
		},
		{
		"codigo":"40045",
		"descricao": "SEDEX VAREJO A COBRAR"
		},
		{
			"codigo":"41009",
			"descricao": "SEDEX AGRUPADO"
		},
		{
		"codigo":"40380",
		"descricao": "SEDEX REVERSO"
		},
		{
		"codigo":"04189",
		"descricao": "SEDEX PAGAMENTO NA ENTREGA"
		},
		{
		"codigo":"04316",
		"descricao": "SEDEX UO"
		},
		{
		"codigo":"04685",
		"descricao": "PAC PAGAMENTO NA ENTREGA"
		},
		{
			"codigo":"10065",
			"descricao": "CARTA COMERCIAL A FATURAR"
		},
		{
		"codigo":"10014",
		"descricao": "CARTA REGISTRADA"
		},
		{
			"codigo":"04162",
			"descricao": "SEDEX"
		},
		{
		"codigo":"04669",
		"descricao": "PAC"
		},
		{
		"codigo":"04170",
		"descricao": "SEDEX REVERSO"
		},
		{
		"codigo":"04677",
		"descricao": "PAC REVERSO"
		},
		{
		"codigo":"12556",
		"descricao": "CARTA COM A FATURAR SELO E SE"
		},
		{
		"codigo":"10707",
		"descricao": "CARTA COMERCIAL REGISTRADA CTR EP MAQ FRAN"
		},
		{
		"codigo":"04227",
		"descricao": "MINI ENVIOS"
		},
		{
		"codigo":"04235",
		"descricao": "MINI ENVIOS"
		},
		{
		"codigo":"04391",
		"descricao": "MINI ENVIOS"
		},
		{
		"codigo":"04146",
		"descricao": "SEDEX GRANDES FORMATOS LM"
		},
		{
		"codigo":"04154",
		"descricao": "SEDEX LM"
		},
		{
		"codigo":"04243",
		"descricao": "SEDEX REVERSO LM"
		},
		{
		"codigo":"04278",
		"descricao": "SEDEX UO LM"
		},
		{
		"codigo":"04883",
			"descricao": "PAC GRANDES FORMATOS LM"
		},
		{
		"codigo":"04367",
			"descricao": "PAC LM"
		},
		{
		"codigo":"04375",
			"descricao": "PAC REVERSO LM"
		},
		{
		"codigo":"04332",
			"descricao": "PAC UO LM"
		},
		{
		"codigo":"04151",
		"descricao": "SEDEX PAGAMENTO NA ENTREGA LM"
		},
		{
		"codigo":"04308",
		"descricao": "PAC PAGAMENTO NA ENTREGA LM"
		},
		{
		"codigo":"04553",
		"descricao": "SEDEX TA"
		},
		{
		"codigo":"04596",
		"descricao": "PAC TA"
		},
		{
		"codigo":"03298",
		"descricao": "PAC"
		},
		{
		"codigo":"03220",
		"descricao": "SEDEX"
		},
		{
		"codigo":"03085",
		"descricao": "PAC"
		},
		{
		"codigo":"03050",
		"descricao": "SEDEX"
		},
		{
		"codigo":"80250",
		"descricao": "CARTA REGISTRADA"
		}
]`
	bytes := []byte(servicesJson)
	var services []Language
	err := json.Unmarshal(bytes, &services)
	if err != nil {
		panic(err)
	}
	var descricao string
	for l := range services {
		if services[l].Codigo == codigo {
			descricao = services[l].Descricao
		}
	}
	return descricao
}
