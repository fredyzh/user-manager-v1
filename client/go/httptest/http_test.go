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
	"google.golang.org/protobuf/encoding/protojson"
)

var addr string = "http://localhost:7777"
var objId = "64e82761dbb126ebdb28d8c5"

func TestHTTPServerHealth(t *testing.T) {
	resp, getErr := http.Get(addr + "/user/v1/service/health")
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

func TestHTTPServerCreateUser(t *testing.T) {

	jsonBody, err := protojson.Marshal(&grpctest.TestCreateUserReq)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	bodyReader := bytes.NewReader(jsonBody)

	resp, getErr := http.Post(fmt.Sprintf("%s/user/v1/service/createuser", addr), "application/json", bodyReader)
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

	resp, getErr := http.Get(fmt.Sprintf("%s/user/v1/service/getuser/{%s}", addr, objId))
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
