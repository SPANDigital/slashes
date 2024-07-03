package slashes

import "testing"

func TestEnsureLeadingSlash(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no leading slash",
			args: args{
				dir: "a/b/c",
			},
			want: "/a/b/c",
		},
		{
			name: "leading slash",
			args: args{
				dir: "/a/b/c",
			},
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnsureLeadingSlash(tt.args.dir); got != tt.want {
				t.Errorf("EnsureLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnsureTrailingSlash(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no trailing slash",
			args: args{
				dir: "a/b/c",
			},
			want: "a/b/c/",
		},
		{
			name: "trailing slash",
			args: args{
				dir: "a/b/c/",
			},
			want: "a/b/c/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnsureTrailingSlash(tt.args.dir); got != tt.want {
				t.Errorf("EnsureTrailingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveLeadingSlash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no leading slash",
			args: args{
				path: "a/b/c",
			},
			want: "a/b/c",
		},
		{
			name: "leading slash",
			args: args{
				path: "/a/b/c",
			},
			want: "a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLeadingSlash(tt.args.path); got != tt.want {
				t.Errorf("RemoveLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveTrailingSlash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no trailing slash",
			args: args{
				path: "a/b/c",
			},
			want: "a/b/c",
		},
		{
			name: "trailing slash",
			args: args{
				path: "a/b/c/",
			},
			want: "a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveTrailingSlash(tt.args.path); got != tt.want {
				t.Errorf("RemoveTrailingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
