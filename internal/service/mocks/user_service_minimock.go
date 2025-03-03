// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/oganes5796/instagram-clon/internal/service.UserService -o user_service_minimock.go -n UserServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/oganes5796/instagram-clon/internal/domain"
)

// UserServiceMock implements mm_service.UserService
type UserServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGenerateToken          func(ctx context.Context, username string, password string) (s1 string, err error)
	funcGenerateTokenOrigin    string
	inspectFuncGenerateToken   func(ctx context.Context, username string, password string)
	afterGenerateTokenCounter  uint64
	beforeGenerateTokenCounter uint64
	GenerateTokenMock          mUserServiceMockGenerateToken

	funcRegister          func(ctx context.Context, user *domain.UserInfo) (i1 int64, err error)
	funcRegisterOrigin    string
	inspectFuncRegister   func(ctx context.Context, user *domain.UserInfo)
	afterRegisterCounter  uint64
	beforeRegisterCounter uint64
	RegisterMock          mUserServiceMockRegister
}

// NewUserServiceMock returns a mock for mm_service.UserService
func NewUserServiceMock(t minimock.Tester) *UserServiceMock {
	m := &UserServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GenerateTokenMock = mUserServiceMockGenerateToken{mock: m}
	m.GenerateTokenMock.callArgs = []*UserServiceMockGenerateTokenParams{}

	m.RegisterMock = mUserServiceMockRegister{mock: m}
	m.RegisterMock.callArgs = []*UserServiceMockRegisterParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserServiceMockGenerateToken struct {
	optional           bool
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockGenerateTokenExpectation
	expectations       []*UserServiceMockGenerateTokenExpectation

	callArgs []*UserServiceMockGenerateTokenParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// UserServiceMockGenerateTokenExpectation specifies expectation struct of the UserService.GenerateToken
type UserServiceMockGenerateTokenExpectation struct {
	mock               *UserServiceMock
	params             *UserServiceMockGenerateTokenParams
	paramPtrs          *UserServiceMockGenerateTokenParamPtrs
	expectationOrigins UserServiceMockGenerateTokenExpectationOrigins
	results            *UserServiceMockGenerateTokenResults
	returnOrigin       string
	Counter            uint64
}

// UserServiceMockGenerateTokenParams contains parameters of the UserService.GenerateToken
type UserServiceMockGenerateTokenParams struct {
	ctx      context.Context
	username string
	password string
}

// UserServiceMockGenerateTokenParamPtrs contains pointers to parameters of the UserService.GenerateToken
type UserServiceMockGenerateTokenParamPtrs struct {
	ctx      *context.Context
	username *string
	password *string
}

// UserServiceMockGenerateTokenResults contains results of the UserService.GenerateToken
type UserServiceMockGenerateTokenResults struct {
	s1  string
	err error
}

// UserServiceMockGenerateTokenOrigins contains origins of expectations of the UserService.GenerateToken
type UserServiceMockGenerateTokenExpectationOrigins struct {
	origin         string
	originCtx      string
	originUsername string
	originPassword string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGenerateToken *mUserServiceMockGenerateToken) Optional() *mUserServiceMockGenerateToken {
	mmGenerateToken.optional = true
	return mmGenerateToken
}

// Expect sets up expected params for UserService.GenerateToken
func (mmGenerateToken *mUserServiceMockGenerateToken) Expect(ctx context.Context, username string, password string) *mUserServiceMockGenerateToken {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &UserServiceMockGenerateTokenExpectation{}
	}

	if mmGenerateToken.defaultExpectation.paramPtrs != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by ExpectParams functions")
	}

	mmGenerateToken.defaultExpectation.params = &UserServiceMockGenerateTokenParams{ctx, username, password}
	mmGenerateToken.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGenerateToken.expectations {
		if minimock.Equal(e.params, mmGenerateToken.defaultExpectation.params) {
			mmGenerateToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGenerateToken.defaultExpectation.params)
		}
	}

	return mmGenerateToken
}

// ExpectCtxParam1 sets up expected param ctx for UserService.GenerateToken
func (mmGenerateToken *mUserServiceMockGenerateToken) ExpectCtxParam1(ctx context.Context) *mUserServiceMockGenerateToken {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &UserServiceMockGenerateTokenExpectation{}
	}

	if mmGenerateToken.defaultExpectation.params != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Expect")
	}

	if mmGenerateToken.defaultExpectation.paramPtrs == nil {
		mmGenerateToken.defaultExpectation.paramPtrs = &UserServiceMockGenerateTokenParamPtrs{}
	}
	mmGenerateToken.defaultExpectation.paramPtrs.ctx = &ctx
	mmGenerateToken.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGenerateToken
}

// ExpectUsernameParam2 sets up expected param username for UserService.GenerateToken
func (mmGenerateToken *mUserServiceMockGenerateToken) ExpectUsernameParam2(username string) *mUserServiceMockGenerateToken {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &UserServiceMockGenerateTokenExpectation{}
	}

	if mmGenerateToken.defaultExpectation.params != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Expect")
	}

	if mmGenerateToken.defaultExpectation.paramPtrs == nil {
		mmGenerateToken.defaultExpectation.paramPtrs = &UserServiceMockGenerateTokenParamPtrs{}
	}
	mmGenerateToken.defaultExpectation.paramPtrs.username = &username
	mmGenerateToken.defaultExpectation.expectationOrigins.originUsername = minimock.CallerInfo(1)

	return mmGenerateToken
}

// ExpectPasswordParam3 sets up expected param password for UserService.GenerateToken
func (mmGenerateToken *mUserServiceMockGenerateToken) ExpectPasswordParam3(password string) *mUserServiceMockGenerateToken {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &UserServiceMockGenerateTokenExpectation{}
	}

	if mmGenerateToken.defaultExpectation.params != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Expect")
	}

	if mmGenerateToken.defaultExpectation.paramPtrs == nil {
		mmGenerateToken.defaultExpectation.paramPtrs = &UserServiceMockGenerateTokenParamPtrs{}
	}
	mmGenerateToken.defaultExpectation.paramPtrs.password = &password
	mmGenerateToken.defaultExpectation.expectationOrigins.originPassword = minimock.CallerInfo(1)

	return mmGenerateToken
}

// Inspect accepts an inspector function that has same arguments as the UserService.GenerateToken
func (mmGenerateToken *mUserServiceMockGenerateToken) Inspect(f func(ctx context.Context, username string, password string)) *mUserServiceMockGenerateToken {
	if mmGenerateToken.mock.inspectFuncGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("Inspect function is already set for UserServiceMock.GenerateToken")
	}

	mmGenerateToken.mock.inspectFuncGenerateToken = f

	return mmGenerateToken
}

// Return sets up results that will be returned by UserService.GenerateToken
func (mmGenerateToken *mUserServiceMockGenerateToken) Return(s1 string, err error) *UserServiceMock {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &UserServiceMockGenerateTokenExpectation{mock: mmGenerateToken.mock}
	}
	mmGenerateToken.defaultExpectation.results = &UserServiceMockGenerateTokenResults{s1, err}
	mmGenerateToken.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGenerateToken.mock
}

// Set uses given function f to mock the UserService.GenerateToken method
func (mmGenerateToken *mUserServiceMockGenerateToken) Set(f func(ctx context.Context, username string, password string) (s1 string, err error)) *UserServiceMock {
	if mmGenerateToken.defaultExpectation != nil {
		mmGenerateToken.mock.t.Fatalf("Default expectation is already set for the UserService.GenerateToken method")
	}

	if len(mmGenerateToken.expectations) > 0 {
		mmGenerateToken.mock.t.Fatalf("Some expectations are already set for the UserService.GenerateToken method")
	}

	mmGenerateToken.mock.funcGenerateToken = f
	mmGenerateToken.mock.funcGenerateTokenOrigin = minimock.CallerInfo(1)
	return mmGenerateToken.mock
}

// When sets expectation for the UserService.GenerateToken which will trigger the result defined by the following
// Then helper
func (mmGenerateToken *mUserServiceMockGenerateToken) When(ctx context.Context, username string, password string) *UserServiceMockGenerateTokenExpectation {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("UserServiceMock.GenerateToken mock is already set by Set")
	}

	expectation := &UserServiceMockGenerateTokenExpectation{
		mock:               mmGenerateToken.mock,
		params:             &UserServiceMockGenerateTokenParams{ctx, username, password},
		expectationOrigins: UserServiceMockGenerateTokenExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGenerateToken.expectations = append(mmGenerateToken.expectations, expectation)
	return expectation
}

// Then sets up UserService.GenerateToken return parameters for the expectation previously defined by the When method
func (e *UserServiceMockGenerateTokenExpectation) Then(s1 string, err error) *UserServiceMock {
	e.results = &UserServiceMockGenerateTokenResults{s1, err}
	return e.mock
}

// Times sets number of times UserService.GenerateToken should be invoked
func (mmGenerateToken *mUserServiceMockGenerateToken) Times(n uint64) *mUserServiceMockGenerateToken {
	if n == 0 {
		mmGenerateToken.mock.t.Fatalf("Times of UserServiceMock.GenerateToken mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGenerateToken.expectedInvocations, n)
	mmGenerateToken.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGenerateToken
}

func (mmGenerateToken *mUserServiceMockGenerateToken) invocationsDone() bool {
	if len(mmGenerateToken.expectations) == 0 && mmGenerateToken.defaultExpectation == nil && mmGenerateToken.mock.funcGenerateToken == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGenerateToken.mock.afterGenerateTokenCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGenerateToken.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GenerateToken implements mm_service.UserService
func (mmGenerateToken *UserServiceMock) GenerateToken(ctx context.Context, username string, password string) (s1 string, err error) {
	mm_atomic.AddUint64(&mmGenerateToken.beforeGenerateTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmGenerateToken.afterGenerateTokenCounter, 1)

	mmGenerateToken.t.Helper()

	if mmGenerateToken.inspectFuncGenerateToken != nil {
		mmGenerateToken.inspectFuncGenerateToken(ctx, username, password)
	}

	mm_params := UserServiceMockGenerateTokenParams{ctx, username, password}

	// Record call args
	mmGenerateToken.GenerateTokenMock.mutex.Lock()
	mmGenerateToken.GenerateTokenMock.callArgs = append(mmGenerateToken.GenerateTokenMock.callArgs, &mm_params)
	mmGenerateToken.GenerateTokenMock.mutex.Unlock()

	for _, e := range mmGenerateToken.GenerateTokenMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmGenerateToken.GenerateTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGenerateToken.GenerateTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmGenerateToken.GenerateTokenMock.defaultExpectation.params
		mm_want_ptrs := mmGenerateToken.GenerateTokenMock.defaultExpectation.paramPtrs

		mm_got := UserServiceMockGenerateTokenParams{ctx, username, password}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGenerateToken.t.Errorf("UserServiceMock.GenerateToken got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGenerateToken.GenerateTokenMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.username != nil && !minimock.Equal(*mm_want_ptrs.username, mm_got.username) {
				mmGenerateToken.t.Errorf("UserServiceMock.GenerateToken got unexpected parameter username, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGenerateToken.GenerateTokenMock.defaultExpectation.expectationOrigins.originUsername, *mm_want_ptrs.username, mm_got.username, minimock.Diff(*mm_want_ptrs.username, mm_got.username))
			}

			if mm_want_ptrs.password != nil && !minimock.Equal(*mm_want_ptrs.password, mm_got.password) {
				mmGenerateToken.t.Errorf("UserServiceMock.GenerateToken got unexpected parameter password, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGenerateToken.GenerateTokenMock.defaultExpectation.expectationOrigins.originPassword, *mm_want_ptrs.password, mm_got.password, minimock.Diff(*mm_want_ptrs.password, mm_got.password))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGenerateToken.t.Errorf("UserServiceMock.GenerateToken got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGenerateToken.GenerateTokenMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGenerateToken.GenerateTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmGenerateToken.t.Fatal("No results are set for the UserServiceMock.GenerateToken")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmGenerateToken.funcGenerateToken != nil {
		return mmGenerateToken.funcGenerateToken(ctx, username, password)
	}
	mmGenerateToken.t.Fatalf("Unexpected call to UserServiceMock.GenerateToken. %v %v %v", ctx, username, password)
	return
}

// GenerateTokenAfterCounter returns a count of finished UserServiceMock.GenerateToken invocations
func (mmGenerateToken *UserServiceMock) GenerateTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGenerateToken.afterGenerateTokenCounter)
}

// GenerateTokenBeforeCounter returns a count of UserServiceMock.GenerateToken invocations
func (mmGenerateToken *UserServiceMock) GenerateTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGenerateToken.beforeGenerateTokenCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.GenerateToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGenerateToken *mUserServiceMockGenerateToken) Calls() []*UserServiceMockGenerateTokenParams {
	mmGenerateToken.mutex.RLock()

	argCopy := make([]*UserServiceMockGenerateTokenParams, len(mmGenerateToken.callArgs))
	copy(argCopy, mmGenerateToken.callArgs)

	mmGenerateToken.mutex.RUnlock()

	return argCopy
}

// MinimockGenerateTokenDone returns true if the count of the GenerateToken invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockGenerateTokenDone() bool {
	if m.GenerateTokenMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GenerateTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GenerateTokenMock.invocationsDone()
}

// MinimockGenerateTokenInspect logs each unmet expectation
func (m *UserServiceMock) MinimockGenerateTokenInspect() {
	for _, e := range m.GenerateTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.GenerateToken at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGenerateTokenCounter := mm_atomic.LoadUint64(&m.afterGenerateTokenCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GenerateTokenMock.defaultExpectation != nil && afterGenerateTokenCounter < 1 {
		if m.GenerateTokenMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to UserServiceMock.GenerateToken at\n%s", m.GenerateTokenMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to UserServiceMock.GenerateToken at\n%s with params: %#v", m.GenerateTokenMock.defaultExpectation.expectationOrigins.origin, *m.GenerateTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGenerateToken != nil && afterGenerateTokenCounter < 1 {
		m.t.Errorf("Expected call to UserServiceMock.GenerateToken at\n%s", m.funcGenerateTokenOrigin)
	}

	if !m.GenerateTokenMock.invocationsDone() && afterGenerateTokenCounter > 0 {
		m.t.Errorf("Expected %d calls to UserServiceMock.GenerateToken at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GenerateTokenMock.expectedInvocations), m.GenerateTokenMock.expectedInvocationsOrigin, afterGenerateTokenCounter)
	}
}

type mUserServiceMockRegister struct {
	optional           bool
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockRegisterExpectation
	expectations       []*UserServiceMockRegisterExpectation

	callArgs []*UserServiceMockRegisterParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// UserServiceMockRegisterExpectation specifies expectation struct of the UserService.Register
type UserServiceMockRegisterExpectation struct {
	mock               *UserServiceMock
	params             *UserServiceMockRegisterParams
	paramPtrs          *UserServiceMockRegisterParamPtrs
	expectationOrigins UserServiceMockRegisterExpectationOrigins
	results            *UserServiceMockRegisterResults
	returnOrigin       string
	Counter            uint64
}

// UserServiceMockRegisterParams contains parameters of the UserService.Register
type UserServiceMockRegisterParams struct {
	ctx  context.Context
	user *domain.UserInfo
}

// UserServiceMockRegisterParamPtrs contains pointers to parameters of the UserService.Register
type UserServiceMockRegisterParamPtrs struct {
	ctx  *context.Context
	user **domain.UserInfo
}

// UserServiceMockRegisterResults contains results of the UserService.Register
type UserServiceMockRegisterResults struct {
	i1  int64
	err error
}

// UserServiceMockRegisterOrigins contains origins of expectations of the UserService.Register
type UserServiceMockRegisterExpectationOrigins struct {
	origin     string
	originCtx  string
	originUser string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmRegister *mUserServiceMockRegister) Optional() *mUserServiceMockRegister {
	mmRegister.optional = true
	return mmRegister
}

// Expect sets up expected params for UserService.Register
func (mmRegister *mUserServiceMockRegister) Expect(ctx context.Context, user *domain.UserInfo) *mUserServiceMockRegister {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserServiceMockRegisterExpectation{}
	}

	if mmRegister.defaultExpectation.paramPtrs != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by ExpectParams functions")
	}

	mmRegister.defaultExpectation.params = &UserServiceMockRegisterParams{ctx, user}
	mmRegister.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmRegister.expectations {
		if minimock.Equal(e.params, mmRegister.defaultExpectation.params) {
			mmRegister.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegister.defaultExpectation.params)
		}
	}

	return mmRegister
}

// ExpectCtxParam1 sets up expected param ctx for UserService.Register
func (mmRegister *mUserServiceMockRegister) ExpectCtxParam1(ctx context.Context) *mUserServiceMockRegister {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserServiceMockRegisterExpectation{}
	}

	if mmRegister.defaultExpectation.params != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Expect")
	}

	if mmRegister.defaultExpectation.paramPtrs == nil {
		mmRegister.defaultExpectation.paramPtrs = &UserServiceMockRegisterParamPtrs{}
	}
	mmRegister.defaultExpectation.paramPtrs.ctx = &ctx
	mmRegister.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmRegister
}

// ExpectUserParam2 sets up expected param user for UserService.Register
func (mmRegister *mUserServiceMockRegister) ExpectUserParam2(user *domain.UserInfo) *mUserServiceMockRegister {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserServiceMockRegisterExpectation{}
	}

	if mmRegister.defaultExpectation.params != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Expect")
	}

	if mmRegister.defaultExpectation.paramPtrs == nil {
		mmRegister.defaultExpectation.paramPtrs = &UserServiceMockRegisterParamPtrs{}
	}
	mmRegister.defaultExpectation.paramPtrs.user = &user
	mmRegister.defaultExpectation.expectationOrigins.originUser = minimock.CallerInfo(1)

	return mmRegister
}

// Inspect accepts an inspector function that has same arguments as the UserService.Register
func (mmRegister *mUserServiceMockRegister) Inspect(f func(ctx context.Context, user *domain.UserInfo)) *mUserServiceMockRegister {
	if mmRegister.mock.inspectFuncRegister != nil {
		mmRegister.mock.t.Fatalf("Inspect function is already set for UserServiceMock.Register")
	}

	mmRegister.mock.inspectFuncRegister = f

	return mmRegister
}

// Return sets up results that will be returned by UserService.Register
func (mmRegister *mUserServiceMockRegister) Return(i1 int64, err error) *UserServiceMock {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserServiceMockRegisterExpectation{mock: mmRegister.mock}
	}
	mmRegister.defaultExpectation.results = &UserServiceMockRegisterResults{i1, err}
	mmRegister.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmRegister.mock
}

// Set uses given function f to mock the UserService.Register method
func (mmRegister *mUserServiceMockRegister) Set(f func(ctx context.Context, user *domain.UserInfo) (i1 int64, err error)) *UserServiceMock {
	if mmRegister.defaultExpectation != nil {
		mmRegister.mock.t.Fatalf("Default expectation is already set for the UserService.Register method")
	}

	if len(mmRegister.expectations) > 0 {
		mmRegister.mock.t.Fatalf("Some expectations are already set for the UserService.Register method")
	}

	mmRegister.mock.funcRegister = f
	mmRegister.mock.funcRegisterOrigin = minimock.CallerInfo(1)
	return mmRegister.mock
}

// When sets expectation for the UserService.Register which will trigger the result defined by the following
// Then helper
func (mmRegister *mUserServiceMockRegister) When(ctx context.Context, user *domain.UserInfo) *UserServiceMockRegisterExpectation {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserServiceMock.Register mock is already set by Set")
	}

	expectation := &UserServiceMockRegisterExpectation{
		mock:               mmRegister.mock,
		params:             &UserServiceMockRegisterParams{ctx, user},
		expectationOrigins: UserServiceMockRegisterExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmRegister.expectations = append(mmRegister.expectations, expectation)
	return expectation
}

// Then sets up UserService.Register return parameters for the expectation previously defined by the When method
func (e *UserServiceMockRegisterExpectation) Then(i1 int64, err error) *UserServiceMock {
	e.results = &UserServiceMockRegisterResults{i1, err}
	return e.mock
}

// Times sets number of times UserService.Register should be invoked
func (mmRegister *mUserServiceMockRegister) Times(n uint64) *mUserServiceMockRegister {
	if n == 0 {
		mmRegister.mock.t.Fatalf("Times of UserServiceMock.Register mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmRegister.expectedInvocations, n)
	mmRegister.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmRegister
}

func (mmRegister *mUserServiceMockRegister) invocationsDone() bool {
	if len(mmRegister.expectations) == 0 && mmRegister.defaultExpectation == nil && mmRegister.mock.funcRegister == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmRegister.mock.afterRegisterCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmRegister.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Register implements mm_service.UserService
func (mmRegister *UserServiceMock) Register(ctx context.Context, user *domain.UserInfo) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmRegister.beforeRegisterCounter, 1)
	defer mm_atomic.AddUint64(&mmRegister.afterRegisterCounter, 1)

	mmRegister.t.Helper()

	if mmRegister.inspectFuncRegister != nil {
		mmRegister.inspectFuncRegister(ctx, user)
	}

	mm_params := UserServiceMockRegisterParams{ctx, user}

	// Record call args
	mmRegister.RegisterMock.mutex.Lock()
	mmRegister.RegisterMock.callArgs = append(mmRegister.RegisterMock.callArgs, &mm_params)
	mmRegister.RegisterMock.mutex.Unlock()

	for _, e := range mmRegister.RegisterMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmRegister.RegisterMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRegister.RegisterMock.defaultExpectation.Counter, 1)
		mm_want := mmRegister.RegisterMock.defaultExpectation.params
		mm_want_ptrs := mmRegister.RegisterMock.defaultExpectation.paramPtrs

		mm_got := UserServiceMockRegisterParams{ctx, user}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmRegister.t.Errorf("UserServiceMock.Register got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmRegister.RegisterMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmRegister.t.Errorf("UserServiceMock.Register got unexpected parameter user, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmRegister.RegisterMock.defaultExpectation.expectationOrigins.originUser, *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegister.t.Errorf("UserServiceMock.Register got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmRegister.RegisterMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegister.RegisterMock.defaultExpectation.results
		if mm_results == nil {
			mmRegister.t.Fatal("No results are set for the UserServiceMock.Register")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmRegister.funcRegister != nil {
		return mmRegister.funcRegister(ctx, user)
	}
	mmRegister.t.Fatalf("Unexpected call to UserServiceMock.Register. %v %v", ctx, user)
	return
}

// RegisterAfterCounter returns a count of finished UserServiceMock.Register invocations
func (mmRegister *UserServiceMock) RegisterAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.afterRegisterCounter)
}

// RegisterBeforeCounter returns a count of UserServiceMock.Register invocations
func (mmRegister *UserServiceMock) RegisterBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.beforeRegisterCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.Register.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegister *mUserServiceMockRegister) Calls() []*UserServiceMockRegisterParams {
	mmRegister.mutex.RLock()

	argCopy := make([]*UserServiceMockRegisterParams, len(mmRegister.callArgs))
	copy(argCopy, mmRegister.callArgs)

	mmRegister.mutex.RUnlock()

	return argCopy
}

// MinimockRegisterDone returns true if the count of the Register invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockRegisterDone() bool {
	if m.RegisterMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.RegisterMock.invocationsDone()
}

// MinimockRegisterInspect logs each unmet expectation
func (m *UserServiceMock) MinimockRegisterInspect() {
	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.Register at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterRegisterCounter := mm_atomic.LoadUint64(&m.afterRegisterCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterMock.defaultExpectation != nil && afterRegisterCounter < 1 {
		if m.RegisterMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to UserServiceMock.Register at\n%s", m.RegisterMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to UserServiceMock.Register at\n%s with params: %#v", m.RegisterMock.defaultExpectation.expectationOrigins.origin, *m.RegisterMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegister != nil && afterRegisterCounter < 1 {
		m.t.Errorf("Expected call to UserServiceMock.Register at\n%s", m.funcRegisterOrigin)
	}

	if !m.RegisterMock.invocationsDone() && afterRegisterCounter > 0 {
		m.t.Errorf("Expected %d calls to UserServiceMock.Register at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.RegisterMock.expectedInvocations), m.RegisterMock.expectedInvocationsOrigin, afterRegisterCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGenerateTokenInspect()

			m.MinimockRegisterInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGenerateTokenDone() &&
		m.MinimockRegisterDone()
}
