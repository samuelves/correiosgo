package consts

type TipoServico string

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

/*
const servicesJson json := `[
	{
		"codigo": "41068",
		"descricao": "PAC",
	},
	{
		"codigo": "04510",
		"descricao": "PAC",
	},
	{
		"codigo": "41300",
		"descricao": "PAC",
	},
	{
		"codigo": "04693",
		"descricao": "PAC",
	},
	{
		"codigo": "04812",
		"descricao": "PAC",
	},
	{
		"codigo": "41556",
		"descricao": "SEDEX",
	},
	{
		"codigo":40169,
		"descricao": "SEDEX 12",
	},
	{
		"codigo":40215,
		"descricao": "SEDEX_10",
		},
	{
		"codigo":03140,
		"descricao": "SEDEX_12",
		},
	{
		"codigo":03158,
		"descricao": "SERVICE_SEDEX_10_CONTRATO_AGENCIA_03158",
		},
	{
		"codigo":40886,
		"descricao": "SERVICE_SEDEX_10_PACOTE",
		},
	{
		"codigo":40290,
		"descricao": "SERVICE_SEDEX_HOJE_40290",
		},
	{
		"codigo":40878,
		"descricao": "SERVICE_SEDEX_HOJE_40878",
		},
	{
		"codigo":04014,
		"descricao": "SERVICE_SEDEX_A_VISTA",
		},
	{
		"codigo":40045,
		"descricao": "SERVICE_SEDEX_VAREJO_A_COBRAR",
		},
	{
		"codigo":41009,
		"descricao": "SERVICE_SEDEX_AGRUPADO",
	},
	{
		"codigo":40380,
		"descricao": "SERVICE_SEDEX_REVERSO",
		},
	{
		"codigo":04189,
		"descricao": "SERVICE_SEDEX_PAGAMENTO_NA_ENTREGA",
		},
	{
		"codigo":04316,
		"descricao": "SERVICE_SEDEX_CONTRATO_UO",
		},
	{
		"codigo":04685,
		"descricao": "SERVICE_PAC_PAGAMENTO_NA_ENTREGA",
	},
	{
		"codigo":10065,
		"descricao": "SERVICE_CARTA_COMERCIAL_A_FATURAR",
	},
	{
		"codigo":10014,
		"descricao": "SERVICE_CARTA_REGISTRADA",
		},
	{
		"codigo":04162,
		"descricao": "SERVICE_SEDEX_CONTRATO_AGENCIA",
	},
	{
		"codigo":04669,
		"descricao": "SERVICE_PAC_CONTRATO_AGENCIA",
		},
	{
		"codigo":04170,
		"descricao": "SERVICE_SEDEX_REVERSO_CONTRATO_AGENCIA",
		},
	{
		"codigo":04677,
		"descricao": "SERVICE_PAC_REVERSO_CONTRATO_AGENCIA",
		},
	{
		"codigo":12556,
		"descricao": "SERVICE_CARTA_COM_A_FATURAR_SELO_E_SE",
		},
	{
		"codigo":10707,
		"descricao": "SERVICE_CARTA_COMERCIAL_REGISTRADA_CTR_EP_MAQ_FRAN",
		},
	{
		"codigo":04227,
		"descricao": "SERVICE_MINI_ENVIOS_04227",
		},
	{
		"codigo":04235,
		"descricao": "SERVICE_MINI_ENVIOS_04235",
		},
	{
		"codigo":04391,
		"descricao": "SERVICE_MINI_ENVIOS_04391",
		},
	{
		"codigo":04146,
		"descricao": "SERVICE_SEDEX_CONTRATO_GRANDES_FORMATOS_LM",
		},
	{
		"codigo":04154,
		"descricao": "SERVICE_SEDEX_CONTRATO_AGENCIA_LM",
		},
	{
		"codigo":04243,
		"descricao": "SERVICE_SEDEX_REVERSO_LM",
	},
	{
		"codigo":04278,
		"descricao": "SERVICE_SEDEX_CONTRATO_UO_LM",
		},
	{
		"codigo":04883,
			"descricao": "SERVICE_PAC_CONTRATO_GRANDES_FORMATOS_LM",
		},
	{
		"codigo":04367,
			"descricao": "SERVICE_PAC_CONTRATO_AGENCIA_LM",
		},
	{
		"codigo":04375,
			"descricao": "SERVICE_PAC_REVERSO_LM",
		},
	{
		"codigo":04332,
			"descricao": "SERVICE_PAC_CONTRATO_UO_LM",
		},
	{
		"codigo":04151,
			"descricao": "SERVICE_SEDEX_CONTRATO_AGENCIA_PAGAMENTO_NA_ENTREGA_LM",
		},
	{
		"codigo":04308,
	"descricao": "SERVICE_PAC_CONTRATO_AGENCIA_PAGAMENTO_NA_ENTREGA_LM",
	},
	{
		"codigo":04553,
	"descricao": "SERVICE_SEDEX_CONTRATO_AGENCIA_TA",
	},
	{
		"codigo":04596,
	"descricao": "SERVICE_PAC_CONTRATO_AGENCIA_TA",
	},
	{
	"codigo":03298,
	"descricao": "SERVICE_PAC_CONTRATO_AGENCIA_03298",
	},
	{
	"codigo":03220,
	"descricao": "SERVICE_SEDEX_CONTRATO_AGENCIA_03220",
	},
	{
	"codigo":03085,
	"descricao": "SERVICE_PAC_CONTRATO_AGENCIA_03085",
	},
	{
	"codigo":03050,
	"descricao": "SERVICE_SEDEX_CONTRATO_AGENCIA_03050",
	},
	{
	"codigo":80250,
	"descricao": "SERVICE_CARTA_REGISTRADA_AGENCIA_80250",
	}
]`

*/
