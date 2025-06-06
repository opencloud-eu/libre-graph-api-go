/*
Libre Graph API

Testing RoleManagementApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package libregraph

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/opencloud-eu/libre-graph-api-go"
)

func Test_libregraph_RoleManagementApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleManagementApiService GetPermissionRoleDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleId string

		resp, httpRes, err := apiClient.RoleManagementApi.GetPermissionRoleDefinition(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleManagementApiService ListPermissionRoleDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleManagementApi.ListPermissionRoleDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
