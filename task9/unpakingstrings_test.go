package main

import "testing"

func TestUnpackingString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "test-case 1",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			wantErr:  false,
		},
		{
			name:     "test-case 2",
			input:    "abcd",
			expected: "abcd",
			wantErr:  false,
		},
		{
			name:     "test-case 3",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "test-case 4",
			input:    "45",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "test-case 5",
			input:    "4a",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "test-case 6",
			input:    "qwe\\4\\5",
			expected: "qwe45",
			wantErr:  false,
		},
		{
			name:     "test-case 7",
			input:    "qwe\\45",
			expected: "qwe44444",
			wantErr:  false,
		},
		{
			name:     "test-case 8",
			input:    "a10b",
			expected: "aaaaaaaaaab",
			wantErr:  false,
		},
		{
			name:     "test-case 9",
			input:    "а2б3в",
			expected: "аабббв",
			wantErr:  false,
		},
		{
			name:     "test-case 10",
			input:    "a12b3",
			expected: "aaaaaaaaaaaabbb",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnpackingString(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackingString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if result != tt.expected {
				t.Errorf("UnpackingString() = %v, want %v", result, tt.expected)
			}
		})
	}
}
