package parser

type Result string

const (
	RESULT_SUCCESS    Result = "SUCCESS"
	RESULT_COMMENT    Result = "COMMENT"
	RESULT_UNEXPECTED Result = "UNEXPECTED"
	RESULT_TERMINATOR Result = "TERMINATOR"
	RESULT_EOF        Result = "EOF"
)

// switch res {
// case RESULT_SUCCESS:
// case RESULT_COMMENT:
// case RESULT_UNEXPECTED:
// case RESULT_TERMINATOR:
// case RESULT_EOF:
// default:
// }
