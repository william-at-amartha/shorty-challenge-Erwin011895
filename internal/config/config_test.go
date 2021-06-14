package config

import (
	"testing"
)

func TestInit(t *testing.T) {
    t.Run("success", func(t *testing.T){
    	Init()
    	c := Get()
    	if c.Environment == "" {
    		t.Fatalf(`failed to get env value of "environment"`)
    	}
    })
}

