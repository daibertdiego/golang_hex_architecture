package arithmetic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAdapter_Addition(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "Valid positive addition",
			args: args{
				a: 12,
				b: 1,
			},
			want:    13,
			wantErr: false,
		},
		{
			name: "Valid negative addition",
			args: args{
				a: -12,
				b: -1,
			},
			want:    -13,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := NewAdapter()
			got, err := arith.Addition(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Adapter.Addition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func TestAdapter_Subtraction(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		Answer    int32
		Operation string
		CreatedAt time.Time
	}
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "Valid positive subtraction",
			args: args{
				a: 12,
				b: 1,
			},
			want:    11,
			wantErr: false,
		},
		{
			name: "Valid negative subtraction",
			args: args{
				a: -12,
				b: -1,
			},
			want:    -11,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := &Adapter{
				ID:        tt.fields.ID,
				Answer:    tt.fields.Answer,
				Operation: tt.fields.Operation,
				CreatedAt: tt.fields.CreatedAt,
			}
			got, err := arith.Subtraction(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Adapter.Subtraction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func TestAdapter_Multiplication(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		Answer    int32
		Operation string
		CreatedAt time.Time
	}
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "Valid positive multiplication",
			args: args{
				a: 12,
				b: 2,
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "Valid negative multiplication",
			args: args{
				a: -12,
				b: -1,
			},
			want:    12,
			wantErr: false,
		},
		{
			name: "Valid multiplication by 0",
			args: args{
				a: 100,
				b: 0,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := &Adapter{
				ID:        tt.fields.ID,
				Answer:    tt.fields.Answer,
				Operation: tt.fields.Operation,
				CreatedAt: tt.fields.CreatedAt,
			}
			got, err := arith.Multiplication(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Adapter.Multiplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func TestAdapter_Division(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		Answer    int32
		Operation string
		CreatedAt time.Time
	}
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "Valid positive division",
			args: args{
				a: 12,
				b: 2,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "Valid negative division",
			args: args{
				a: -12,
				b: 2,
			},
			want:    -6,
			wantErr: false,
		},
		{
			name: "Invalid division by 0",
			args: args{
				a: 100,
				b: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := &Adapter{
				ID:        tt.fields.ID,
				Answer:    tt.fields.Answer,
				Operation: tt.fields.Operation,
				CreatedAt: tt.fields.CreatedAt,
			}
			got, err := arith.Division(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Adapter.Division() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}
