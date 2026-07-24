package bankingbridge

import (
	"testing"

	"github.com/hanzo-fi/go-libs/v5/pkg/observe/log"
	"github.com/hanzo-fi/payments/ee/plugins/bankingbridge/client"
	"github.com/hanzo-fi/payments/pkg/domain/plugins"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type PluginTestSuite struct {
	suite.Suite

	client *client.MockClient
	plugin *Plugin
}

func TestPluginTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(PluginTestSuite))
}

func (suite *PluginTestSuite) SetupTest() {
	logger := logging.Testing()
	ctrl := gomock.NewController(suite.T())
	suite.client = client.NewMockClient(ctrl)
	suite.plugin = &Plugin{
		Plugin: plugins.NewBasePlugin(),
		name:   "test",
		logger: logger,
		client: suite.client,
	}
}
