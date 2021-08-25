package gomock

type Matcher interface {
	Matches(x interface{}) bool
	String() string
}
