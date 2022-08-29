package repositories_interface

//go:generate mockgen -destination=../../test/mock/imail_mock_repository.go -package=mock GoWeb/repository/interface IMailRep
type IMailRep interface {
	Send(string) bool
}
