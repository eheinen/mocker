// ------------------MOCKER--------------------
// DO NOT EDIT THIS FILE, CODE AUTO-GENERATED
// ------------------MOCKER--------------------

package mocks

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockExecutePerson
// No argument and return expected
func MockExecutePerson(mocking *mock.Mock) *mock.Call {
	return mocking.On("ExecutePerson")
}

// AssertExecutePerson
// No argument and return expected
func AssertExecutePerson(t *testing.T, mocking *mock.Mock, times int) {
	mocking.AssertNumberOfCalls(t, "ExecutePerson", times)
}

// MockFindPerson
// *person.Person,  error
func MockFindPerson(mocking *mock.Mock, expectedReturns ...interface{}) *mock.Call {
	return mocking.On("FindPerson").Return(expectedReturns...)
}

// AssertFindPerson
// No argument and return expected
func AssertFindPerson(t *testing.T, mocking *mock.Mock, times int, expectedReturns ...interface{}) {
	if times != 0 {
		mocking.AssertCalled(t, "FindPerson", expectedReturns...)
	}

	mocking.AssertNumberOfCalls(t, "FindPerson", times)
}

// MockPostPerson
// (person person.Person, ctx context.Context)
// error
func MockPostPerson(mocking *mock.Mock, expectedReturns ...interface{}) *mock.Call {
	return mocking.On("PostPerson",
		mock.AnythingOfType("person.Person"),
		mock.Anything,
	).Return(expectedReturns...)
}

// AssertPostPerson
// (person person.Person, ctx context.Context)
func AssertPostPerson(t *testing.T, mocking *mock.Mock, times int, expectedReturns ...interface{}) {
	if times != 0 {
		mocking.AssertCalled(t, "PostPerson", expectedReturns...)
	}

	mocking.AssertNumberOfCalls(t, "PostPerson", times)
}

// MockUpdatePerson
// (person person.Person, ctx context.Context)
// error
func MockUpdatePerson(mocking *mock.Mock, expectedReturns ...interface{}) *mock.Call {
	return mocking.On("UpdatePerson",
		mock.AnythingOfType("person.Person"),
		mock.Anything,
	).Return(expectedReturns...)
}

// AssertUpdatePerson
// (person person.Person, ctx context.Context)
func AssertUpdatePerson(t *testing.T, mocking *mock.Mock, times int, expectedReturns ...interface{}) {
	if times != 0 {
		mocking.AssertCalled(t, "UpdatePerson", expectedReturns...)
	}

	mocking.AssertNumberOfCalls(t, "UpdatePerson", times)
}