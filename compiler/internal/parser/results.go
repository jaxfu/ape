package parser

type Result string

const (
	RESULT_SUCCESS    Result = "RESULT_SUCCESS"
	RESULT_TERMINATOR Result = "RESULT_TERMINATOR"
	RESULT_UNEXPECTED Result = "RESULT_UNEXPECTED"
	RESULT_EOF        Result = "RESULT_EOF"
)

// switch res {
// case RESULT_SUCCESS:
// case RESULT_TERMINATOR:
// case RESULT_UNEXPECTED:
// case RESULT_EOF:
// default:
// }
