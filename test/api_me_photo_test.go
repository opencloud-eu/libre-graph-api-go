/*
Libre Graph API

Testing MePhotoApiService

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

func Test_libregraph_MePhotoApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MePhotoApiService DeleteOwnUserPhoto", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MePhotoApi.DeleteOwnUserPhoto(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MePhotoApiService GetOwnUserPhoto", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MePhotoApi.GetOwnUserPhoto(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MePhotoApiService UpdateOwnUserPhotoPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MePhotoApi.UpdateOwnUserPhotoPatch(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MePhotoApiService UpdateOwnUserPhotoPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MePhotoApi.UpdateOwnUserPhotoPut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
