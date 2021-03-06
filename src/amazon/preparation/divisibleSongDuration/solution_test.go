package divisibleSongDuration

import (
	"testing"
)

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1",
			args: args{[]int{30, 20, 150, 100, 40}},
			want: 3},
		{name: "test 2",
			args: args{[]int{60, 60, 60}},
			want: 3},
		{name: "test 3",
			args: args{time: []int{269, 230, 318, 468, 171, 158, 350, 60, 287, 27, 11, 384, 332, 267, 412, 478, 280, 303, 242, 378, 129,
				131, 164, 467, 345, 146, 264, 332, 276, 479, 284, 433, 117, 197, 430, 203, 100, 280, 145, 287, 91, 157, 5, 475, 288, 146, 370,
				199, 81, 428, 278, 2, 400, 23, 470, 242, 411, 470, 330, 144, 189, 204, 62, 318, 475, 24, 457, 83, 204, 322, 250, 478, 186, 467,
				350, 171, 119, 245, 399, 112, 252, 201, 324, 317, 293, 44, 295, 14, 379, 382, 137, 280, 265, 78, 38, 323, 347, 499, 238, 110,
				18, 224, 473, 289, 198, 106, 256, 279, 275, 349, 210, 498, 201, 175, 472, 461, 116, 144, 9, 221, 473}},
			want: 105},
		{name: "test 4",
			args: args{[]int{20, 40}},
			want: 1},
		{name: "test 5",
			args: args{[]int{336, 24, 100, 342, 274, 11, 43, 22, 416, 138, 384, 386, 70, 265, 59, 253, 344, 435, 400, 296, 192,
				143, 311, 424, 315, 63, 420, 254, 493, 431, 32, 394, 178, 51, 378, 335, 265, 92, 335, 325, 25, 355, 258, 298, 390, 399, 393,
				114, 149, 62, 299, 471, 286, 204, 163, 214, 15, 272, 315, 212, 272, 437, 339, 193, 125, 394, 62, 188, 154, 150, 109, 294,
				228, 200, 459, 42, 469, 132, 37, 460, 143, 1, 144, 127, 398, 82, 370, 464, 14, 85, 321, 358, 205, 14, 264, 289, 183, 93, 56,
				126, 413, 140, 441, 446, 445, 378, 258, 119, 385, 226, 8, 93, 476, 265, 115, 86, 360, 92, 396, 407, 458, 58, 65, 397, 381,
				32, 228, 37, 319, 220, 73, 328, 162, 458, 231, 219, 481, 387, 423, 256, 252, 36, 309, 395, 471, 4, 225, 146, 188, 182, 347,
				82, 21, 292, 91, 144, 387, 263, 206, 452, 197, 192, 324, 257, 370, 28, 440, 180, 294}},
			want: 245},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
