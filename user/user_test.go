package user

import (
	"errors"
	"testing"

	"unit-test-demo/match"
	"unit-test-demo/mocks"

	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDoer := mocks.NewMockDoer(mockController)
	testUser := User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	// mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	// Expect Do to be called once with any value and "Hello GoMock" as parameters, and return nil from the mocked call.
	// mockDoer.EXPECT().DoSomething(gomock.Any(), "Hello GoMock")

	// Using our own matchers
	// Expect Do to be called once with 123 and any string as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, match.OfType("string")).Return(nil).Times(1)

	// GoMock provides a way to assert that one call must happen after another call, the .After method
	// callFirst := mockDoer.EXPECT().DoSomething(1, "first this")
	// callA := mockDoer.EXPECT().DoSomething(2, "then this").After(callFirst)
	// callB := mockDoer.EXPECT().DoSomething(2, "or this").After(callFirst)

	// GoMock also provides a convenience function gomock.InOrder to specify that the calls must be performed in the exact order given. This is less flexible than using .After directly, but can make your tests more readable for longer sequences of calls
	gomock.InOrder(
		mockDoer.EXPECT().DoSomething(1, "first this"),
		mockDoer.EXPECT().DoSomething(2, "then this"),
		mockDoer.EXPECT().DoSomething(3, "then this"),
		mockDoer.EXPECT().DoSomething(4, "finally this"),
	)

	err := testUser.Use()
	if err != nil {
		panic(err)
	}
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockController)
	testUser := User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)

	err := testUser.Use()
	if err != nil {
		t.Fail()
	}
}
