package magnet

import "testing"

func TestBase32ToHex(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "right1",
			args: args{"W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK"},
			want: "b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa",
		},
		{
			name: "right2",
			args: args{"magnet:?xt=urn:btih:W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK"},
			want: "magnet:?xt=urn:btih:b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa",
		},
		{
			name: "too long",
			args: args{"magnet:?xt=urn:btih:W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK1"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base32ToHex(tt.args.src); got != tt.want {
				t.Errorf("Base32ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToBase32(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "right1",
			args: args{"b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa"},
			want: "W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK",
		},
		{
			name: "right2",
			args: args{"magnet:?xt=urn:btih:b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa"},
			want: "magnet:?xt=urn:btih:W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK",
		},
		{
			name: "too long",
			args: args{"magnet:?xt=urn:btih:b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa11111"},
			want: "",
		},
		{
			name: "too short",
			args: args{"magnet:?xt=urn:btih:b7092da3a99dd0fa5a701ebe2a6dcb897a6e50a"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToBase32(tt.args.src); got != tt.want {
				t.Errorf("HexToBase32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHash(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "32",
			args: args{"W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK"},
			want: "b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa",
		},
		{
			name: "40",
			args: args{"b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa"},
			want: "b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa",
		},
		{
			name: "52",
			args: args{"magnet:?xt=urn:btih:W4ES3I5JTXIPUWTQD27CU3OLRF5G4UFK"},
			want: "b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa",
		},
		{
			name: "60",
			args: args{"magnet:?xt=urn:btih:b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa"},
			want: "b7092da3a99dd0fa5a701ebe2a6dcb897a6e50aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHash(tt.args.src); got != tt.want {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
