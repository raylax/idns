package protocol

type Protocol struct {
	TransactionId int

	Flags

	QuestionRRs RRs

	AnswerRRs RRs

	AuthorityRRs RRs

	AdditionalRRs RRs

	Questions []Question

	Answers []Answer

	Authorities []Authority

	Additions []Addition
}

type RRs int

const (
	MaskBit1 = 0x0001
	MaskBit2 = 0x000f
)
