package protocol

type Protocol struct {

	TransactionId int16

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

type RRs int16

const (
	MASK_BIT_1	= 0x0001
	MASK_BIT_4	= 0x000f
)