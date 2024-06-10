package goemoji

import (
	"reflect"
	"testing"
)

const mockStrategyReturn = "🍎"

type MockStrategy struct{}

func (i MockStrategy) Emojify(input string, minimumWordLength int, emojiTags map[string][]string, emojiSet map[string]bool) (output string) {
	return mockStrategyReturn
}

func TestEmojifier_ContainsEmoji(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		e    *Emojifier
		args args
		want bool
	}{
		{
			name: "what a delicious apple",
			e: &Emojifier{
				strategy:          MockStrategy{},
				minimumWordLength: 4,
				emojiTags:         defaultEmojiTags,
				emojiSet:          defaultEmojiSet,
			},
			args: args{
				text: "what a delicious 🍎",
			},
			want: true,
		}, {
			name: "does not contain emoji",
			e: &Emojifier{
				strategy:          MockStrategy{},
				minimumWordLength: 4,
				emojiTags:         defaultEmojiTags,
				emojiSet:          defaultEmojiSet,
			},
			args: args{
				text: "what a delicious apple",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ContainsEmoji(tt.args.text); got != tt.want {
				t.Errorf("Emojifier.ContainsEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmojifier_ExtractEmojis(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		e    *Emojifier
		args args
		want []string
	}{
		{
			name: "extract emojis",
			e: &Emojifier{
				strategy:          MockStrategy{},
				minimumWordLength: 4,
				emojiTags:         defaultEmojiTags,
				emojiSet:          defaultEmojiSet,
			},
			args: args{
				text: "what a delicious 🍎🍏🍍",
			},
			want: []string{"🍎", "🍏", "🍍"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ExtractEmojis(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Emojifier.ExtractEmojis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmojifier_Emojify(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		e    *Emojifier
		args args
		want string
	}{
		{
			name: "test emojify",
			e: &Emojifier{
				strategy:          MockStrategy{},
				minimumWordLength: 4,
				emojiTags:         defaultEmojiTags,
				emojiSet:          defaultEmojiSet,
			},
			args: args{
				text: "what a delicious apple",
			},
			want: mockStrategyReturn,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Emojify(tt.args.text); got != tt.want {
				t.Errorf("Emojifier.Emojify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmojifier(t *testing.T) {
	emojifier, err := NewEmojifier(MockStrategy{}, 4)

	if err != nil {
		t.Errorf("NewEmojifier() error = %v", err)
	}
	if emojifier == nil {
		t.Errorf("NewEmojifier() emojifier = %v", emojifier)
	}
}
