package gokhttp_cookiejar

import (
	"context"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	gokhttp_requests "github.com/BRUHItsABunny/gOkHttp/requests"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJarOption(t *testing.T) {
	jar, err := NewCookieJar(".cookies", "", nil)
	require.NoError(t, err, "cookies.NewCookieJar: errored unexpectedly.")

	// To load cookies from existing .cookies into jar instance
	// err = jar.Load()
	// require.NoError(t, err, "jar.Load: errored unexpectedly.")

	hClient, err := gokhttp.NewHTTPClient(
		NewJarOption(jar),
		// gokhttp_client.NewProxyOption("http://127.0.0.1:8888"),
	)
	require.NoError(t, err, "NewHTTPClient: errored unexpectedly.")

	for i := 0; i <= 1; i++ {
		req, err := gokhttp_requests.MakeGETRequest(context.Background(), "https://github.com")
		require.NoError(t, err, "requests.MakeGETRequest: errored unexpectedly.")

		_, err = hClient.Do(req)
		require.NoError(t, err, "hClient.Do: errored unexpectedly.")

		time.Sleep(time.Duration(500) * time.Millisecond)
	}

	err = jar.Save()
	require.NoError(t, err, "jar.Save: errored unexpectedly.")
}
