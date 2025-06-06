/*
Libre Graph API

Testing DriveItemApiService

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

func Test_libregraph_DriveItemApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DriveItemApiService DeleteDriveItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var itemId string

		httpRes, err := apiClient.DriveItemApi.DeleteDriveItem(context.Background(), driveId, itemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DriveItemApiService GetDriveItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var itemId string

		resp, httpRes, err := apiClient.DriveItemApi.GetDriveItem(context.Background(), driveId, itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DriveItemApiService UpdateDriveItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var driveId string
		var itemId string

		resp, httpRes, err := apiClient.DriveItemApi.UpdateDriveItem(context.Background(), driveId, itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
