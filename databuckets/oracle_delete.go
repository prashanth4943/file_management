package databuckets

import (
	// "bytes"
	// "context"

	// "io/ioutil"
	// "dms-backend/internal/db"

	// "os"
	// "github.com/joho/godotenv"
	"context"
	"fmt"

	// "github.com/google/uuid"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func DeleteFromOracle(Email *string) error {
	provider := common.DefaultConfigProvider()

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
	if err != nil {
		panic(err)
	}

	bucketName := "test_bucket_1"
	namespace := "axpzvvirgqhg"
	objectName := fmt.Sprintf("%s.png", *Email)

	request := objectstorage.DeleteObjectRequest{
		NamespaceName: common.String(namespace),
		BucketName:    common.String(bucketName),
		ObjectName:    common.String(objectName),
	}
	_, err = client.DeleteObject(context.Background(), request)
	if err != nil {
		return fmt.Errorf("failed to delete object from Oracle: %v", err)
	}
	fmt.Printf("File %s deleted successfully from bucket %s\n", objectName, bucketName)
	return nil
}
