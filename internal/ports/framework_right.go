package ports

type DbPort interface {
    CloseDbConnection()
    AddToHistory ( answer int32, operatino string) error
}