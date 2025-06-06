/*
Libre Graph API

Testing DrivesRootApiService

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

func Test_libregraph_DrivesRootApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DrivesRootApiService CreateDriveItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string

		resp, httpRes, err := apiClient.DrivesRootApi.CreateDriveItem(context.Background(), driveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService CreateLinkSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string

		resp, httpRes, err := apiClient.DrivesRootApi.CreateLinkSpaceRoot(context.Background(), driveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService DeletePermissionSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var permId string

		httpRes, err := apiClient.DrivesRootApi.DeletePermissionSpaceRoot(context.Background(), driveId, permId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService GetPermissionSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var permId string

		resp, httpRes, err := apiClient.DrivesRootApi.GetPermissionSpaceRoot(context.Background(), driveId, permId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService GetRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string

		resp, httpRes, err := apiClient.DrivesRootApi.GetRoot(context.Background(), driveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService InviteSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string

		resp, httpRes, err := apiClient.DrivesRootApi.InviteSpaceRoot(context.Background(), driveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService ListPermissionsSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string

		resp, httpRes, err := apiClient.DrivesRootApi.ListPermissionsSpaceRoot(context.Background(), driveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService SetPermissionPasswordSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var permId string

		resp, httpRes, err := apiClient.DrivesRootApi.SetPermissionPasswordSpaceRoot(context.Background(), driveId, permId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DrivesRootApiService UpdatePermissionSpaceRoot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var permId string

		resp, httpRes, err := apiClient.DrivesRootApi.UpdatePermissionSpaceRoot(context.Background(), driveId, permId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
