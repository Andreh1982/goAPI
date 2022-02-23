package controllers_test

import (
	"goAPI/controllers"
	"net/http"

	// "net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.Home(tt.args.w, tt.args.r)
		})
	}
}

func TestTodasPersonalidades(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.TodasPersonalidades(tt.args.w, tt.args.r)
		})
	}
}

func TestRetornaUmaPersonalidade(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		Id       int
		Nome     string
		Historia string
		args     args
	}{
		// TODO: Add test cases.
		{1, "nome 1", "longa historia", args{}},
		{2, "nome 2", "não tão longa historia", args{}},
		{3, "nome 3", "uma media historia", args{}},
		{4, "nome 4", "longa historia", args{}},
		{5, "nome 5", "não tão longa historia", args{}},
		{6, "nome 6", "uma media historia", args{}},
	}

	// writer := args{}

	for _, tt := range tests {
		t.Run(tt.Nome, func(t *testing.T) {
			controllers.RetornaUmaPersonalidade(tt.args.w, tt.args.r)
		})
	}
}

func TestCriaUmaNovaPersonalidade(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.CriaUmaNovaPersonalidade(tt.args.w, tt.args.r)
		})
	}
}

func TestDeletaUmaPersonalidade(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.DeletaUmaPersonalidade(tt.args.w, tt.args.r)
		})
	}
}

func TestEditaPersonalidade(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.EditaPersonalidade(tt.args.w, tt.args.r)
		})
	}
}
