package vervet

type Config interface {
    GetUrlPattern() (string, error)
}
