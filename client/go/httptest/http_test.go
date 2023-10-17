package httptest

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"testing"

	"github.com/fredyzh/user_manager/client/go/grpctest"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

var addr_user string = "http://44.204.19.88:7778"
var addr_token string = "http://44.204.19.88:7779"
var objId = "6512e9bdc7b5e49b51365733"

func TestHTTPServerHealth(t *testing.T) {
	should := require.New(t)
	_, getErr := http.Get(addr_user + "/user/v1/service/health")

	should.NoError(getErr)
}
func TestHTTPTokenServerHealth(t *testing.T) {
	should := require.New(t)
	_, getErr := http.Get(addr_token + "/jwt/v1/service/health")
	should.NoError(getErr)
}

func TestHTTPServerCreateUser(t *testing.T) {

	jsonBody, err := protojson.Marshal(&grpctest.TestCreateUserReq)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, getErr := http.Post(fmt.Sprintf("%s/user/v1/service/createuser", addr_user), "application/json", bodyReader)
	if getErr != nil {
		slog.Error(getErr.Error())
		return
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		slog.Error(readErr.Error())
		return
	}

	log.Println(string(body))

	if body == nil {
		t.FailNow()
	}
}

func TestHTTPServerGetUser(t *testing.T) {

	resp, getErr := http.Get(fmt.Sprintf("%s/user/v1/service/getuser/{%s}", addr_user, objId))
	if getErr != nil {
		slog.Error(getErr.Error())
		return
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		slog.Error(readErr.Error())
		return
	}

	log.Println(string(body))

	if body == nil {
		t.FailNow()
	}
}
