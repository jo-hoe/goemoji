package goemoji

import (
	"reflect"
	"testing"
)

var defaultEmojiMap = map[string][]string{"apple": {"🍎", "🍏"}, "green apple": {"🍏"}, "pineapple": {"🍍"}}

func TestInsertBeforeString_Emojify(t *testing.T) {
	type args struct {
		input    string
		emojiMap map[string][]string
	}
	tests := []struct {
		name       string
		i          InsertBeforeString
		args       args
		wantOutput string
	}{
		{
			name: "single word replacement",
			i:    InsertBeforeString{},
			args: args{
				input:    "they ate an apple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "🍎 they ate an apple",
		}, {
			name: "multi-word emoji replacement",
			i:    InsertBeforeString{},
			args: args{
				input:    "they ate a green apple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "🍏 they ate a green apple",
		}, {
			name: "multi-word multi-emoji replacement",
			i:    InsertBeforeString{},
			args: args{
				input:    "they ate an apple and a green apple and a pineapple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "🍎🍏🍍 they ate an apple and a green apple and a pineapple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := tt.i.Emojify(tt.args.input, tt.args.emojiMap); gotOutput != tt.wantOutput {
				t.Errorf("InsertBeforeString.Emojify() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestInsertAfterString_Emojify(t *testing.T) {
	type args struct {
		input    string
		emojiMap map[string][]string
	}
	tests := []struct {
		name       string
		i          InsertAfterString
		args       args
		wantOutput string
	}{
		{
			name: "single word replacement",
			i:    InsertAfterString{},
			args: args{
				input:    "they ate an apple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "they ate an apple 🍎",
		}, {
			name: "multi-word emoji replacement",
			i:    InsertAfterString{},
			args: args{
				input:    "they ate a green apple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "they ate a green apple 🍏",
		}, {
			name: "multi-word multi-emoji replacement",
			i:    InsertAfterString{},
			args: args{
				input:    "they ate an apple and a green apple and a pineapple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "they ate an apple and a green apple and a pineapple 🍎🍏🍍",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := tt.i.Emojify(tt.args.input, tt.args.emojiMap); gotOutput != tt.wantOutput {
				t.Errorf("InsertBeforeString.Emojify() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestReplaceSubstring_Emojify(t *testing.T) {
	type args struct {
		input    string
		emojiMap map[string][]string
	}
	tests := []struct {
		name       string
		i          ReplaceSubstring
		args       args
		wantOutput string
	}{
		{
			name: "single word replacement",
			i:    ReplaceSubstring{},
			args: args{
				input:    "they ate an apple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "they ate an 🍎",
		}, {
			name: "multi-word emoji replacement",
			i:    ReplaceSubstring{},
			args: args{
				input:    "they ate a green apple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "they ate a 🍏",
		}, {
			name: "multi-word multi-emoji replacement",
			i:    ReplaceSubstring{},
			args: args{
				input:    "they ate an apple and a green apple and a pineapple",
				emojiMap: defaultEmojiMap,
			},
			wantOutput: "they ate an 🍎 and a 🍏 and a 🍍",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := tt.i.Emojify(tt.args.input, tt.args.emojiMap); gotOutput != tt.wantOutput {
				t.Errorf("got '%v', want '%v'", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_combineTokens(t *testing.T) {
	type args struct {
		input    []string
		numWords int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single word",
			args: args{
				input:    []string{"The", "quick", "brown", "fox"},
				numWords: 1,
			},
			want: []string{"The", "quick", "brown", "fox"},
		}, {
			name: "multi-word",
			args: args{
				input:    []string{"The", "quick", "brown", "fox"},
				numWords: 2,
			},
			want: []string{"The quick", "quick brown", "brown fox"},
		}, {
			name: "non alphabetical",
			args: args{
				input:    []string{"Th-", "qu1ck", "br0wn", "f0x"},
				numWords: 2,
			},
			want: []string{"Th- qu1ck", "qu1ck br0wn", "br0wn f0x"},
		}, {
			name: "shorter than numWords",
			args: args{
				input:    []string{"The", "quick", "brown", "fox"},
				numWords: 5,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineTokens(tt.args.input, tt.args.numWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursiveTokenize(t *testing.T) {
	type args struct {
		input    string
		numWords int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test recursion",
			args: args{
				input:    "The quick brown fox",
				numWords: 2,
			},
			want: []string{"The quick", "quick brown", "brown fox", "The", "quick", "brown", "fox"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveTokenize(tt.args.input, tt.args.numWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursiveTokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
