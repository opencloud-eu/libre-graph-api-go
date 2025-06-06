/*
Libre Graph API

Testing ActivitiesApiService

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

func Test_libregraph_ActivitiesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActivitiesApiService GetActivities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ActivitiesApi.GetActivities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
