package doer

//go:generate $GOPATH/bin/mockgen -destination=../mocks/mock_doer.go -package=mocks unit-test-demo/doer Doer

type Doer interface {
	DoSomething(int, string) error
}
