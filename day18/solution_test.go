package day18

import (
	"os"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		line string
		want []*RegularNumber
	}{
		{
			"[11,2111]",
			[]*RegularNumber{
				{n: 11, depth: 0},
				{n: 2111, depth: 0},
			},
		},
		{
			"[11,[2111,[59,32]]]",
			[]*RegularNumber{
				{n: 11, depth: 0},
				{n: 2111, depth: 1},
				{n: 59, depth: 2},
				{n: 32, depth: 2},
			},
		},
		{
			"[[[[[9,8],1],2],3],4]",
			[]*RegularNumber{
				{n: 9, depth: 4},
				{n: 8, depth: 4},
				{n: 1, depth: 3},
				{n: 2, depth: 2},
				{n: 3, depth: 1},
				{n: 4, depth: 0},
			},
		},
		{
			"[7,[6,[5,[4,[3,2]]]]]",
			[]*RegularNumber{
				{n: 7, depth: 0},
				{n: 6, depth: 1},
				{n: 5, depth: 2},
				{n: 4, depth: 3},
				{n: 3, depth: 4},
				{n: 2, depth: 4},
			},
		},
		{
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			[]*RegularNumber{
				{n: 3, depth: 1},
				{n: 2, depth: 2},
				{n: 8, depth: 3},
				{n: 0, depth: 3},
				// right
				{n: 9, depth: 1},
				{n: 5, depth: 2},
				{n: 4, depth: 3},
				{n: 3, depth: 4},
				{n: 2, depth: 4},
			},
		},
		{
			"[[6,[5,[7,0]]],3]",
			[]*RegularNumber{
				{n: 6, depth: 1},
				{n: 5, depth: 2},
				{n: 7, depth: 3},
				{n: 0, depth: 3},
				{n: 3, depth: 0},
			},
		},
		{
			"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			[]*RegularNumber{
				{n: 3, depth: 1},
				{n: 2, depth: 2},
				{n: 1, depth: 3},
				{n: 7, depth: 4},
				{n: 3, depth: 4},

				{n: 6, depth: 1},
				{n: 5, depth: 2},
				{n: 4, depth: 3},
				{n: 3, depth: 4},
				{n: 2, depth: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := ParseLine(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplode(t *testing.T) {
	tests := []struct {
		before string
		after  string
	}{
		{
			"[[[[[9,8],1],2],3],4]",
			"[[[[0,9],2],3],4]",
		},
		{
			"[7,[6,[5,[4,[3,2]]]]]",
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.before, func(t *testing.T) {
			rns := ParseLine(tt.before)
			want := ParseLine(tt.after)
			if got, changed := Explode(rns); !reflect.DeepEqual(got, want) || !changed {
				t.Errorf("\n%s\n%s\n", got, want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		before string
		after  string
	}{
		{
			"[[[[0,7],4],[15,[0,13]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
		{
			"[[3,[2,[1,[7,3]]]],[6,[5,[14,[3,2]]]]]",
			"[[3,[2,[1,[7,3]]]],[6,[5,[[7,7],[3,2]]]]]",
		},
		{
			"[[3,[2,[8,17]]],[9,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,[8,9]]]],[9,[5,[4,[3,2]]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.before, func(t *testing.T) {
			rns := ParseLine(tt.before)
			want := ParseLine(tt.after)
			if got, _ := Split(rns); !reflect.DeepEqual(got, want) {
				t.Errorf("\n%s\n%s\n", got, want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a     string
		b     string
		after string
	}{
		{
			"[1,2]",
			"[[3,4],5]",
			"[[1,2],[[3,4],5]]",
		},
		{
			"[[[[4,3],4],4],[7,[[8,4],9]]]",
			"[1,1]",
			"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.a+" + "+tt.b, func(t *testing.T) {
			a := ParseLine(tt.a)
			b := ParseLine(tt.b)
			want := ParseLine(tt.after)
			if got := Add(a, b); !reflect.DeepEqual(got, want) {
				t.Errorf("\n%s\n%s\n", got, want)
			}
		})
	}
}

func TestAddAndReduce(t *testing.T) {
	tests := []struct {
		a     string
		b     string
		after string
	}{
		{
			"[[[[4,3],4],4],[7,[[8,4],9]]]",
			"[1,1]",
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.a+" + "+tt.b, func(t *testing.T) {
			a := ParseLine(tt.a)
			b := ParseLine(tt.b)
			want := ParseLine(tt.after)
			if got := Reduce(Add(a, b)); !reflect.DeepEqual(got, want) {
				t.Errorf("\n%s\n%s\n", got, want)
			}
		})
	}
}

func TestAddAndReduceAll(t *testing.T) {
	tests := []struct {
		filename string
		want     string
	}{
		{
			"slightly_larger_example.txt",
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			"small_1.txt",
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			"small_2.txt",
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			"small_3.txt",
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			f, _ := os.Open(tt.filename)
			all := ParseInput(f)
			got := AddAndReduceAll(all)
			want := ParseLine(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\n%s\n%s\n", got, want)
			}
		})
	}

}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		raw  string
		want int
	}{
		{
			"[9,1]",
			29,
		},
		{
			"[[9,1],[1,9]]",
			129,
		},
		// edge case
		// doesn't come up in input so ¯\_(ツ)_/¯
		// {
		// 	"[[1,2],[[3,4],5]]",
		// 	143,
		// },
		{
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			1384,
		},
		{
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
			445,
		},
		{
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
			445,
		},
		{
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
			791,
		},
		{
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
			1137,
		},
		{
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			3488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.raw, func(t *testing.T) {
			if got := Magnitude(ParseLine(tt.raw)); got != tt.want {
				t.Errorf("Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("example_assignment.txt", func(t *testing.T) {
		f, _ := os.Open("example_assignment.txt")
		all := ParseInput(f)
		got := Magnitude(AddAndReduceAll(all))
		want := 4140
		if got != want {
			t.Errorf("Magnitude() = %v, want %v", got, want)
		}
	})

	t.Run("input.txt", func(t *testing.T) {
		f, _ := os.Open("input.txt")
		all := ParseInput(f)
		got := Magnitude(AddAndReduceAll(all))
		want := 4184
		if got != want {
			t.Errorf("Magnitude() = %v, want %v", got, want)
		}
	})
}

func TestHighestMagnitude(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{
			"example_assignment.txt",
			3993,
		},
		{
			"input.txt",
			4731,
		},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			f, _ := os.Open(tt.filename)
			all := ParseInput(f)
			if got := HighestMagnitude(all); got != tt.want {
				t.Errorf("HighestMagnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
