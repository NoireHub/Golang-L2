package main

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
)


func TestTask_stringUnpack(t *testing.T) {
	testCases := []struct {
		name string
		inputStr string
		exceptedOutput string
		exceptedError error
	}{
		{
			name: "valid string with numbers",
			inputStr: "a4bc2d5e",
			exceptedOutput: "aaaabccddddde",
			exceptedError: nil,
		},
		{
			name: "valid string with escape 1",
			inputStr: "qwe\\4\\5",
			exceptedOutput: "qwe45",
			exceptedError: nil,
		},
		{
			name: "valid string with escape 2",
			inputStr: "qwe\\\\5",
			exceptedOutput: "qwe\\\\\\\\\\",
			exceptedError: nil,
		},
		{
			name: "valid string without numbers",
			inputStr: "asdgss",
			exceptedOutput: "asdgss",
			exceptedError: nil,
		},
		{
			name: "empty string",
			inputStr: "",
			exceptedOutput: "",
			exceptedError: nil,
		},
		{
			name: "invalid string",
			inputStr: "45",
			exceptedOutput: "",
			exceptedError: errors.New("invalid string"),
		},
		{
			name: "invalid string with zero",
			inputStr: "a4bc2d0e",
			exceptedOutput: "",
			exceptedError: errors.New("invalid number"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			str, err:= stringUnpack(tc.inputStr)

			assert.Equal(t,str,tc.exceptedOutput)
			assert.Equal(t,err,tc.exceptedError)
		})


	}
}