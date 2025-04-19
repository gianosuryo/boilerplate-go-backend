package config

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"os"
	"sarana-dafa-ai-service/storage/env"
	"time"

	memorystore "cloud.google.com/go/redis/apiv1"
	redispb "cloud.google.com/go/redis/apiv1/redispb"
	redis "github.com/redis/go-redis/v9"
)

// Memorystore Redis instance.
func connectToDatabase(w io.Writer, projectID, location, instanceID string) (*redis.Client, error) {

	// Instantiate a Redis administrative client
	ctx := context.Background()

	var opt redis.Options
	opt.PoolSize = 1
	opt.MinIdleConns = 1
	opt.PoolTimeout = 0
	// opt.IdleTimeout = 20 * time.Second
	opt.DialTimeout = 2 * time.Second

	if os.Getenv(env.REDIS_TYPE) == "address" {
		opt.Addr = fmt.Sprintf("%s:%s", os.Getenv(env.REDIS_HOST), os.Getenv(env.REDIS_PORT))
		if len(os.Getenv(env.REDIS_USERNAME)) > 0 {
			opt.Username = os.Getenv(env.REDIS_USERNAME)
		}
		if len(os.Getenv(env.REDIS_PASSWORD)) > 0 {
			opt.Password = os.Getenv(env.REDIS_PASSWORD)
		}
	}

	if os.Getenv(env.REDIS_TYPE) == "certificate" {
		adminClient, err := memorystore.NewCloudRedisClient(ctx)
		if err != nil {
			return nil, err
		}
		defer adminClient.Close()

		req := &redispb.GetInstanceRequest{
			Name: fmt.Sprintf("projects/%s/locations/%s/instances/%s", projectID, location, instanceID),
		}

		instance, err := adminClient.GetInstance(ctx, req)
		if err != nil {
			return nil, err
		}

		// Load CA cert
		caCerts := instance.GetServerCaCerts()
		if len(caCerts) == 0 {
			return nil, errors.New("memorystore: no server CA certs for instance")
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(caCerts[0].Cert))

		opt.Addr = fmt.Sprintf("%s:%d", instance.Host, instance.Port)
		opt.TLSConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	}

	// Setup Redis Connection pool
	client := redis.NewClient(&opt)

	p, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(w, "Response:\n%s", p)

	return client, nil
}

func GetRedisInstace() (*redis.Client, error) {
	var buff bytes.Buffer
	return connectToDatabase(&buff, os.Getenv(env.REDIS_PROJECT_ID), os.Getenv(env.REDIS_LOCATION), os.Getenv(env.REDIS_INSTANCE_ID))
}
