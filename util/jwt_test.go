package util

import (
	"douyin-app/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWT(t *testing.T) {
	conf.InitConf("./conf/default_conf.yaml")
	InitJWTVal()
	for i := 0; i < 10000; i++ {
		token, err := GenerateToken(1, 1<<50)
		assert.Equal(t, err, nil)

		c, err := ParseToken(token)
		assert.Equal(t, err, nil)
		assert.Equal(t, c.UserId, int64(1<<50))
	}

}
