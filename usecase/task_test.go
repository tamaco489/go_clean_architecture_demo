package usecase_test

import (
	"testing"

	"github.com/clean_architecture_beta/model"
	"github.com/clean_architecture_beta/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *model.Task) (int, error) {
	args := m.Called(task)

	// Int は、指定されたインデックスの引数を取得する。
	// 引数がない場合、または引数の型が間違っている場合はパニックになります。
	return args.Int(0), args.Error(1)
}

func (m *MockTaskRepository) Read(id int) (*model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	testCase := []struct {
		title string
		id    int
		err   error
	}{
		{
			title: "test title",
			id:    1,
			err:   nil,
		},
		{
			title: "",
			id:    0,
			err:   model.ErrTaskTitleEmpty,
		},
	}
	for _, tc := range testCase {
		// DBアクセスせず、ID:1, error: nilを返すモックを設定
		mockRepo.On("Create", mock.Anything).Return(tc.id, tc.err)

		id, err := taskUsecase.CreateTask(tc.title)
		assert.Equal(t, tc.id, id)   // 返ってきたidが期待通りであるかを確認
		assert.Equal(t, tc.err, err) // 返ってきたerrorが期待通りであるかを確認
	}
}
