package types

type Participante struct {
	Alinhamento   string `json:"alinhamento"`
	Espirito      int    `json:"espirito"`
	Familia       string `json:"familia"`
	FotoExtension string `json:"foto_extension"`
	ID            int    `json:"id"`
	Nome          string `json:"nome"`
	Potencia      int    `json:"potencia"`
	Presenca      int    `json:"presenca"`
	Profissao     string `json:"profissao"`
	Responsa      int    `json:"responsa"`
	Simpatia      int    `json:"simpatia"`
	Xp            int    `json:"xp"`
}

type Elenco struct {
	Desc              string         `json:"desc"`
	Etiqueta          string         `json:"etiqueta"`
	FotoExtension     string         `json:"foto_extension"`
	Nome              string         `json:"nome"`
	ParticipanteIndex int            `json:"participante_index"`
	Participantes     []Participante `json:"participantes"`
}
