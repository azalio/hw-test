package hw03frequencyanalysis

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	name     string
	text     string
	expected []string
}

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var (
	textOne           = "one one one one one one one one one one one one one one one one"
	textNum           = "one two three four five six seven eight nine ten eleven"
	textWord          = "word"
	textStrangeSpaces = `space space Space  Space`
)

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	testTable := []testCase{

		{
			name: "just one",
			text: textOne,
			expected: []string{
				"one",
			},
		},
		{
			name: "numbers",
			text: textNum,
			expected: []string{
				"eight",
				"eleven",
				"five",
				"four",
				"nine",
				"one",
				"seven",
				"six",
				"ten",
				"three",
			},
		},
		{
			name: "just word",
			text: textWord,
			expected: []string{
				"word",
			},
		},
		{
			name: "strage spaces",
			text: textStrangeSpaces,
			expected: []string{
				"space",
			},
		},
		{
			name: "vinne-puh",
			text: text,
			expected: []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			},
		},
	}

	for _, test := range testTable {
		actual := Top10(test.text)
		require.Equal(t, test.expected, actual, test.name)
	}
}

func Test_clearWords(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "just one",
			args: args{[]string{"one"}},
			want: []string{"one"},
		},
		{
			name: "just two",
			args: args{[]string{"one", "two"}},
			want: []string{"one", "two"},
		},
		{
			name: "just comma",
			args: args{[]string{"one,"}},
			want: []string{"one"},
		},
		{
			name: "just comma and exclamation",
			args: args{[]string{"!one,"}},
			want: []string{"one"},
		},
		{
			name: "just comma and exclamation and hypen",
			args: args{[]string{"!one-one,"}},
			want: []string{"one-one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearWords(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clearWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
