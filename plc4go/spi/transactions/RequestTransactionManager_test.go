/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package transactions

import (
	"container/list"
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/apache/plc4x/plc4go/spi/pool"
	"github.com/apache/plc4x/plc4go/spi/testutils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNewRequestTransactionManager(t *testing.T) {
	type args struct {
		numberOfConcurrentRequests       int
		requestTransactionManagerOptions []options.WithOption
	}
	tests := []struct {
		name  string
		args  args
		setup func(t *testing.T, args *args)
		want  RequestTransactionManager
	}{
		{
			name: "just create one",
			want: &requestTransactionManager{
				workLog:  *list.New(),
				executor: sharedExecutorInstance,
			},
		},
		{
			name: "just create one with option",
			args: args{
				numberOfConcurrentRequests: 2,
				requestTransactionManagerOptions: []options.WithOption{
					WithCustomExecutor(sharedExecutorInstance),
				},
			},
			want: &requestTransactionManager{
				numberOfConcurrentRequests: 2,
				workLog:                    *list.New(),
				executor:                   sharedExecutorInstance,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, &tt.args)
			}
			if got := NewRequestTransactionManager(tt.args.numberOfConcurrentRequests, tt.args.requestTransactionManagerOptions...); !assert.Equal(t, tt.want, got) {
				t.Errorf("NewRequestTransactionManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCustomExecutor(t *testing.T) {
	type args struct {
		executor pool.Executor
	}
	tests := []struct {
		name string
		args args
		want options.WithOption
	}{
		{
			name: "with a option",
			args: args{
				executor: sharedExecutorInstance,
			},
			want: WithCustomExecutor(sharedExecutorInstance),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customExecutor := WithCustomExecutor(tt.args.executor)
			assert.NotNil(t, customExecutor)
		})
	}
}

func Test_requestTransactionManager_SetNumberOfConcurrentRequests(t *testing.T) {
	type fields struct {
		runningRequests            []*requestTransaction
		numberOfConcurrentRequests int
		transactionId              int32
		workLog                    list.List
		executor                   pool.Executor
	}
	type args struct {
		numberOfConcurrentRequests int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set a number",
		},
		{
			name: "set a number on running requests",
			fields: fields{
				runningRequests: []*requestTransaction{
					{}, // empty one is sufficient
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &requestTransactionManager{
				runningRequests:            tt.fields.runningRequests,
				numberOfConcurrentRequests: tt.fields.numberOfConcurrentRequests,
				transactionId:              tt.fields.transactionId,
				workLog:                    tt.fields.workLog,
				executor:                   tt.fields.executor,
			}
			r.SetNumberOfConcurrentRequests(tt.args.numberOfConcurrentRequests)
		})
	}
}

func Test_requestTransactionManager_StartTransaction(t *testing.T) {
	type fields struct {
		runningRequests                     []*requestTransaction
		numberOfConcurrentRequests          int
		transactionId                       int32
		workLog                             list.List
		executor                            pool.Executor
		shutdown                            bool
		traceTransactionManagerTransactions bool
		log                                 zerolog.Logger
	}
	tests := []struct {
		name       string
		fields     fields
		setup      func(t *testing.T, fields *fields)
		wantAssert func(t *testing.T, requestTransaction RequestTransaction) bool
	}{
		{
			name: "start one",
			setup: func(t *testing.T, fields *fields) {
				fields.log = testutils.ProduceTestingLogger(t)
			},
			wantAssert: func(t *testing.T, requestTransaction RequestTransaction) bool {
				assert.False(t, requestTransaction.IsCompleted())
				return true
			},
		},
		{
			name: "start one in shutdown",
			fields: fields{
				shutdown: true,
			},
			setup: func(t *testing.T, fields *fields) {
				fields.log = testutils.ProduceTestingLogger(t)
			},
			wantAssert: func(t *testing.T, requestTransaction RequestTransaction) bool {
				assert.True(t, requestTransaction.IsCompleted())
				assert.Error(t, requestTransaction.AwaitCompletion(context.Background()))
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, &tt.fields)
			}
			r := &requestTransactionManager{
				runningRequests:                     tt.fields.runningRequests,
				numberOfConcurrentRequests:          tt.fields.numberOfConcurrentRequests,
				transactionId:                       tt.fields.transactionId,
				workLog:                             tt.fields.workLog,
				executor:                            tt.fields.executor,
				shutdown:                            tt.fields.shutdown,
				traceTransactionManagerTransactions: tt.fields.traceTransactionManagerTransactions,
				log:                                 tt.fields.log,
			}
			if got := r.StartTransaction(); !assert.True(t, tt.wantAssert(t, got)) {
				t.Errorf("StartTransaction() = %v", got)
			}
		})
	}
}

func Test_requestTransactionManager_endRequest(t *testing.T) {
	type fields struct {
		runningRequests            []*requestTransaction
		numberOfConcurrentRequests int
		transactionId              int32
		workLog                    list.List
		executor                   pool.Executor
	}
	type args struct {
		transaction *requestTransaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "end request with unknown transaction",
			args: args{
				transaction: &requestTransaction{},
			},
			wantErr: true,
		},
		{
			name: "end request",
			args: args{
				transaction: &requestTransaction{},
			},
			fields: fields{
				runningRequests: []*requestTransaction{
					{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &requestTransactionManager{
				runningRequests:            tt.fields.runningRequests,
				numberOfConcurrentRequests: tt.fields.numberOfConcurrentRequests,
				transactionId:              tt.fields.transactionId,
				workLog:                    tt.fields.workLog,
				executor:                   tt.fields.executor,
			}
			if err := r.endRequest(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("endRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestTransactionManager_failRequest(t *testing.T) {
	type fields struct {
		runningRequests            []*requestTransaction
		numberOfConcurrentRequests int
		transactionId              int32
		workLog                    list.List
		executor                   pool.Executor
		log                        zerolog.Logger
	}
	type args struct {
		transaction *requestTransaction
		err         error
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockSetup func(t *testing.T, fields *fields, args *args)
		wantErr   bool
	}{
		{
			name: "fail a request",
			args: args{
				transaction: &requestTransaction{},
			},
			mockSetup: func(t *testing.T, fields *fields, args *args) {
				fields.log = testutils.ProduceTestingLogger(t)

				completionFuture := NewMockCompletionFuture(t)
				expect := completionFuture.EXPECT()
				expect.Cancel(true, nil).Return()
				args.transaction.completionFuture = completionFuture
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup(t, &tt.fields, &tt.args)
			}
			r := &requestTransactionManager{
				runningRequests:            tt.fields.runningRequests,
				numberOfConcurrentRequests: tt.fields.numberOfConcurrentRequests,
				transactionId:              tt.fields.transactionId,
				workLog:                    tt.fields.workLog,
				executor:                   tt.fields.executor,
				log:                        tt.fields.log,
			}
			if err := r.failRequest(tt.args.transaction, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("failRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestTransactionManager_getNumberOfActiveRequests(t *testing.T) {
	type fields struct {
		runningRequests            []*requestTransaction
		numberOfConcurrentRequests int
		transactionId              int32
		workLog                    list.List
		executor                   pool.Executor
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "get em",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &requestTransactionManager{
				runningRequests:            tt.fields.runningRequests,
				numberOfConcurrentRequests: tt.fields.numberOfConcurrentRequests,
				transactionId:              tt.fields.transactionId,
				workLog:                    tt.fields.workLog,
				executor:                   tt.fields.executor,
			}
			if got := r.getNumberOfActiveRequests(); got != tt.want {
				t.Errorf("getNumberOfActiveRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestTransactionManager_processWorklog(t *testing.T) {
	type fields struct {
		runningRequests            []*requestTransaction
		numberOfConcurrentRequests int
		transactionId              int32
		workLog                    list.List
		executor                   pool.Executor
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "process nothing",
		},
		{
			name: "process one",
			fields: fields{
				numberOfConcurrentRequests: 100,
				workLog: func() list.List {
					l := list.New()
					l.PushBack(&requestTransaction{})
					return *l
				}(),
				executor: sharedExecutorInstance,
			},
		},
		{
			name: "process two",
			fields: fields{
				numberOfConcurrentRequests: 100,
				workLog: func() list.List {
					l := list.New()
					l.PushBack(&requestTransaction{
						transactionId:    1,
						completionFuture: NewMockCompletionFuture(t),
					})
					l.PushBack(&requestTransaction{
						transactionId:    2,
						completionFuture: NewMockCompletionFuture(t),
					})
					return *l
				}(),
				executor: sharedExecutorInstance,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &requestTransactionManager{
				runningRequests:            tt.fields.runningRequests,
				numberOfConcurrentRequests: tt.fields.numberOfConcurrentRequests,
				transactionId:              tt.fields.transactionId,
				workLog:                    tt.fields.workLog,
				executor:                   tt.fields.executor,
			}
			r.processWorklog()
		})
	}
}

func Test_requestTransactionManager_submitTransaction(t *testing.T) {
	type fields struct {
		runningRequests            []*requestTransaction
		numberOfConcurrentRequests int
		transactionId              int32
		workLog                    list.List
		executor                   pool.Executor
	}
	type args struct {
		handle *requestTransaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "submit it",
			args: args{
				handle: &requestTransaction{
					operation: func() {
						// doesn't matter
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &requestTransactionManager{
				runningRequests:            tt.fields.runningRequests,
				numberOfConcurrentRequests: tt.fields.numberOfConcurrentRequests,
				transactionId:              tt.fields.transactionId,
				workLog:                    tt.fields.workLog,
				executor:                   tt.fields.executor,
			}
			r.submitTransaction(tt.args.handle)
		})
	}
}

func Test_requestTransaction_AwaitCompletion(t1 *testing.T) {
	type fields struct {
		parent           *requestTransactionManager
		transactionId    int32
		operation        pool.Runnable
		completionFuture pool.CompletionFuture
		transactionLog   zerolog.Logger
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockSetup func(t *testing.T, fields *fields, args *args)
		wantErr   bool
	}{
		{
			name: "just wait",
			fields: fields{
				parent: &requestTransactionManager{
					runningRequests: []*requestTransaction{
						func() *requestTransaction {
							r := &requestTransaction{}
							go func() {
								time.Sleep(100 * time.Millisecond)
								// We fake an ending transaction like that
								r.transactionId = 1
							}()
							return r
						}(),
					},
				},
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancelFunc := context.WithCancel(context.Background())
					cancelFunc()
					return ctx
				}(),
			},
			mockSetup: func(t *testing.T, fields *fields, args *args) {
				completionFuture := NewMockCompletionFuture(t)
				expect := completionFuture.EXPECT()
				expect.AwaitCompletion(mock.Anything).Return(nil)
				fields.completionFuture = completionFuture
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup(t1, &tt.fields, &tt.args)
			}
			t := &requestTransaction{
				parent:           tt.fields.parent,
				transactionId:    tt.fields.transactionId,
				operation:        tt.fields.operation,
				completionFuture: tt.fields.completionFuture,
				transactionLog:   tt.fields.transactionLog,
			}
			if err := t.AwaitCompletion(tt.args.ctx); (err != nil) != tt.wantErr {
				t1.Errorf("AwaitCompletion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestTransaction_EndRequest(t1 *testing.T) {
	type fields struct {
		parent           *requestTransactionManager
		transactionId    int32
		operation        pool.Runnable
		completionFuture pool.CompletionFuture
		transactionLog   zerolog.Logger
		completed        bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "just end it",
			fields: fields{
				parent: &requestTransactionManager{},
			},
			wantErr: true,
		},
		{
			name: "end it completed",
			fields: fields{
				parent:    &requestTransactionManager{},
				completed: true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &requestTransaction{
				parent:           tt.fields.parent,
				transactionId:    tt.fields.transactionId,
				operation:        tt.fields.operation,
				completionFuture: tt.fields.completionFuture,
				transactionLog:   tt.fields.transactionLog,
				completed:        tt.fields.completed,
			}
			if err := t.EndRequest(); (err != nil) != tt.wantErr {
				t1.Errorf("EndRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestTransaction_FailRequest(t1 *testing.T) {
	type fields struct {
		parent           *requestTransactionManager
		transactionId    int32
		operation        pool.Runnable
		completionFuture pool.CompletionFuture
		transactionLog   zerolog.Logger
		completed        bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockSetup func(t *testing.T, fields *fields, args *args)
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "just fail it",
			fields: fields{
				parent: &requestTransactionManager{},
			},
			mockSetup: func(t *testing.T, fields *fields, args *args) {
				completionFuture := NewMockCompletionFuture(t)
				expect := completionFuture.EXPECT()
				expect.Cancel(true, nil).Return()
				fields.completionFuture = completionFuture
			},
			wantErr: assert.Error,
		},
		{
			name: "just fail it (completed)",
			args: args{
				err: errors.New("nope"),
			},
			fields: fields{
				parent:    &requestTransactionManager{},
				completed: true,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup(t, &tt.fields, &tt.args)
			}
			r := &requestTransaction{
				parent:           tt.fields.parent,
				transactionId:    tt.fields.transactionId,
				operation:        tt.fields.operation,
				completionFuture: tt.fields.completionFuture,
				transactionLog:   tt.fields.transactionLog,
				completed:        tt.fields.completed,
			}
			tt.wantErr(t, r.FailRequest(tt.args.err), "FailRequest() error = %v", tt.args.err)
		})
	}
}

func Test_requestTransaction_String(t1 *testing.T) {
	type fields struct {
		parent           *requestTransactionManager
		transactionId    int32
		operation        pool.Runnable
		completionFuture pool.CompletionFuture
		transactionLog   zerolog.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "give a string",
			want: "Transaction{tid:0}",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &requestTransaction{
				parent:           tt.fields.parent,
				transactionId:    tt.fields.transactionId,
				operation:        tt.fields.operation,
				completionFuture: tt.fields.completionFuture,
				transactionLog:   tt.fields.transactionLog,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestTransaction_Submit(t1 *testing.T) {
	type fields struct {
		parent           *requestTransactionManager
		transactionId    int32
		operation        pool.Runnable
		completionFuture pool.CompletionFuture
		transactionLog   zerolog.Logger
		completed        bool
	}
	type args struct {
		operation RequestTransactionRunnable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "submit something",
			fields: fields{
				parent: &requestTransactionManager{},
			},
			args: args{
				operation: func(_ RequestTransaction) {
					// NOOP
				},
			},
		},
		{
			name: "submit something again",
			fields: fields{
				parent: &requestTransactionManager{},
				operation: func() {
					// NOOP
				},
			},
			args: args{
				operation: func(_ RequestTransaction) {
					// NOOP
				},
			},
		},
		{
			name: "submit completed",
			fields: fields{
				parent: &requestTransactionManager{},
				operation: func() {
					// NOOP
				},
				completed: true,
			},
			args: args{
				operation: func(_ RequestTransaction) {
					// NOOP
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &requestTransaction{
				parent:           tt.fields.parent,
				transactionId:    tt.fields.transactionId,
				operation:        tt.fields.operation,
				completionFuture: tt.fields.completionFuture,
				transactionLog:   tt.fields.transactionLog,
				completed:        tt.fields.completed,
			}
			t.Submit(tt.args.operation)
			t.operation()
		})
	}
}

func Test_requestTransactionManager_Close(t *testing.T) {
	type fields struct {
		runningRequests                     []*requestTransaction
		numberOfConcurrentRequests          int
		transactionId                       int32
		workLog                             list.List
		executor                            pool.Executor
		shutdown                            bool
		traceTransactionManagerTransactions bool
		log                                 zerolog.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		setup   func(t *testing.T, fields *fields)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "close it",
			setup: func(t *testing.T, fields *fields) {
				executor := NewMockExecutor(t)
				executor.EXPECT().Close().Return(nil)
				fields.executor = executor
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, &tt.fields)
			}
			r := &requestTransactionManager{
				runningRequests:                     tt.fields.runningRequests,
				numberOfConcurrentRequests:          tt.fields.numberOfConcurrentRequests,
				transactionId:                       tt.fields.transactionId,
				workLog:                             tt.fields.workLog,
				executor:                            tt.fields.executor,
				shutdown:                            tt.fields.shutdown,
				traceTransactionManagerTransactions: tt.fields.traceTransactionManagerTransactions,
				log:                                 tt.fields.log,
			}
			tt.wantErr(t, r.Close(), fmt.Sprintf("Close()"))
		})
	}
}

func Test_requestTransactionManager_CloseGraceful(t *testing.T) {
	type fields struct {
		runningRequests                     []*requestTransaction
		numberOfConcurrentRequests          int
		transactionId                       int32
		workLog                             list.List
		executor                            pool.Executor
		shutdown                            bool
		traceTransactionManagerTransactions bool
		log                                 zerolog.Logger
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func(t *testing.T, fields *fields)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "close it",
			setup: func(t *testing.T, fields *fields) {
				executor := NewMockExecutor(t)
				executor.EXPECT().Close().Return(nil)
				fields.executor = executor
			},
			wantErr: assert.NoError,
		},
		{
			name: "close it with timeout",
			args: args{
				timeout: 20 * time.Millisecond,
			},
			setup: func(t *testing.T, fields *fields) {
				executor := NewMockExecutor(t)
				executor.EXPECT().Close().Return(nil)
				fields.executor = executor
			},
			wantErr: assert.NoError,
		},
		{
			name: "close it with timeout fires",
			fields: fields{
				runningRequests: []*requestTransaction{
					{},
				},
			},
			args: args{
				timeout: 20 * time.Millisecond,
			},
			setup: func(t *testing.T, fields *fields) {
				executor := NewMockExecutor(t)
				executor.EXPECT().Close().Return(nil)
				fields.executor = executor
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, &tt.fields)
			}
			r := &requestTransactionManager{
				runningRequests:                     tt.fields.runningRequests,
				numberOfConcurrentRequests:          tt.fields.numberOfConcurrentRequests,
				transactionId:                       tt.fields.transactionId,
				workLog:                             tt.fields.workLog,
				executor:                            tt.fields.executor,
				shutdown:                            tt.fields.shutdown,
				traceTransactionManagerTransactions: tt.fields.traceTransactionManagerTransactions,
				log:                                 tt.fields.log,
			}
			tt.wantErr(t, r.CloseGraceful(tt.args.timeout), fmt.Sprintf("CloseGraceful(%v)", tt.args.timeout))
		})
	}
}