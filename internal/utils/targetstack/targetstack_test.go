package targetstack

import (
	"fmt"
	"testing"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
)

func Test_stack_Push(t *testing.T) {
	type args struct {
		target models.TargetParameters
	}
	tests := []struct {
		name string
		s    FlowStack
		args args
		want FlowStack
	}{
		{
			name: "sanity_unittest_targetstack_push",
			s:    []models.TargetParameters{},
			args: args{
				target: models.TargetParameters{
					Method: "",
					URL:    "",
					Body:   nil,
					Header: nil,
				},
			},
			want: []models.TargetParameters{
				{
					Method: "",
					URL:    "",
					Body:   nil,
					Header: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Push(tt.args.target); len(got) != len(tt.want) {
				t.Errorf("Push() = %d elements, want %d elements", len(got), len(tt.want))
			}
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    FlowStack
		want FlowStack
	}{
		{
			name: "sanity_unittest_targetstack_pop",
			s: []models.TargetParameters{
				{
					Method: "",
					URL:    "",
					Body:   nil,
					Header: nil,
				},
			},
			want: []models.TargetParameters{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.s.Pop()

			if len(got) != len(tt.want) {
				t.Errorf("Pop() = %d elements, want %d elements", len(got), len(tt.want))
			}
		})
	}
}

func Test_Push_Push_Pop_Pop_Sequence(t *testing.T) {
	var stack FlowStack
	stack = stack.Push(models.TargetParameters{
		URL: "delete url",
	})
	fmt.Println(stack)
	stack = stack.Push(models.TargetParameters{
		URL: "pt post connect url",
	})
	fmt.Println(stack)
	stack = stack.Push(models.TargetParameters{
		URL: "caller connect url",
	})
	fmt.Println(stack)
	stack = stack.Push(models.TargetParameters{
		URL: "caller validate url",
	})
	fmt.Println(stack)
	stack = stack.Push(models.TargetParameters{
		URL: "post allocation url",
	})
	fmt.Println(stack)
	stack, _ = stack.Pop()
	stack, _ = stack.Pop()
	stack, _ = stack.Pop()
	stack, _ = stack.Pop()
	if a, got := stack.Pop(); got.URL != "delete url" {
		t.Errorf("expected: url: delete url, got: %s, stack: %v", got.URL, a)
	}
	fmt.Println(stack)
}
