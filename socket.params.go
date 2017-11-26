package chrome

/*
SocketParams represents parameters for a WebSocket command
*/
type SocketParams struct {
	Expression            string `json:"expression"`                      // Expression to evaluate.
	ObjectGroup           string `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	IncludeCommandLineAPI bool   `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	Silent                bool   `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	ContextID             int    `json:"contextId,omitempty"`             // Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ReturnByValue         bool   `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview       bool   `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	UserGesture           bool   `json:"userGesture,omitempty"`           // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise          bool   `json:"awaitPromise,omitempty"`          // Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}
