package keys

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/tests"

	"github.com/stretchr/testify/assert"
)

func Test_runAddCmdBasic(t *testing.T) {
	cmd := addKeyCommand()
	assert.NotNil(t, cmd)
	mockIn, _, _ := tests.ApplyMockIO(cmd)

	// Prepare a keybase
	kbHome, kbCleanUp := tests.NewTestCaseDir(t)
	assert.NotNil(t, kbHome)
	defer kbCleanUp()
	viper.Set(cli.HomeFlag, kbHome)

	/// Test Text
	viper.Set(cli.OutputFlag, OutputFormatText)
	// Now enter password

	mockIn.Reset("test1234\ntest1234\n")
	err = runAddCmd(cmd, []string{"keyname1"})
	assert.NoError(t, err)

	/// Test Text - Replace? >> FAIL
	viper.Set(cli.OutputFlag, OutputFormatText)

	// Now enter password
	mockIn.Reset("test1234\ntest1234\n")
	err = runAddCmd(cmd, []string{"keyname1"})
	assert.Error(t, err)

	/// Test Text - Replace? Answer >> PASS
	viper.Set(cli.OutputFlag, OutputFormatText)

	// Now enter password
	mockIn.Reset("y\ntest1234\ntest1234\n")
	err = runAddCmd(cmd, []string{"keyname1"})
	assert.NoError(t, err)

	// Check JSON
	viper.Set(cli.OutputFlag, OutputFormatJSON)

	// Now enter password
	mockIn.Reset("test1234\ntest1234\n")
	err = runAddCmd(cmd, []string{"keyname2"})
	assert.NoError(t, err)
}

type MockResponseWriter struct {
	dataHeaderStatus int
	dataBody         []byte
}

func (MockResponseWriter) Header() http.Header {
	panic("Unexpected call!")
}

func (w *MockResponseWriter) Write(data []byte) (int, error) {
	w.dataBody = append(w.dataBody, data...)
	return len(data), nil
}

func (w *MockResponseWriter) WriteHeader(statusCode int) {
	w.dataHeaderStatus = statusCode
}

func TestCheckAndWriteErrorResponse(t *testing.T) {
	mockRW := MockResponseWriter{}

	mockRW.WriteHeader(100)
	assert.Equal(t, 100, mockRW.dataHeaderStatus)

	detected := CheckAndWriteErrorResponse(&mockRW, http.StatusBadRequest, errors.New("some ERROR"))
	require.True(t, detected)
	require.Equal(t, http.StatusBadRequest, mockRW.dataHeaderStatus)
	require.Equal(t, "some ERROR", string(mockRW.dataBody[:]))

	mockRW = MockResponseWriter{}
	detected = CheckAndWriteErrorResponse(&mockRW, http.StatusBadRequest, nil)
	require.False(t, detected)
	require.Equal(t, 0, mockRW.dataHeaderStatus)
	require.Equal(t, "", string(mockRW.dataBody[:]))
}
