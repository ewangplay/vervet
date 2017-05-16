package vervet

type Config interface {
    GetUrlPattern() (string, error)
    GetLogIdLiteral() (string, error)
    GetErrorCodeLiteral() (string, error)
    GetErrorMessageLiteral() (string, error)
    GetTimeCostLiteral() () (string, error)
    GetRequestUrlLiteral() (string, error)
}
