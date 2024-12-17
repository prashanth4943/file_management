package databuckets

import (
	// "bytes"
	// "context"
	"io"
	// "io/ioutil"
	// "dms-backend/internal/db"

	"os"

	// "os"
	// "github.com/joho/godotenv"
	"context"
	"fmt"

	// "github.com/google/uuid"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func DownloadFromOracle(uuid string, Email string) (string, error) {

	provider := common.DefaultConfigProvider()

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
	if err != nil {
		panic(err)
	}

	bucketName := "test_bucket_1"
	namespace := "axpzvvirgqhg"
	objectName := fmt.Sprintf("%s.png", Email)

	request := objectstorage.GetObjectRequest{
		NamespaceName: common.String(namespace),
		BucketName:    common.String(bucketName),
		ObjectName:    common.String(objectName),
	}

	response, err := client.GetObject(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("failed to get object from Oracle: %v", err)
	}
	defer response.Content.Close()

	savePath := fmt.Sprintf("./%s", objectName)
	file, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("failed to create local file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Content)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	fmt.Printf("File downloaded successfully and saved to %s\n", savePath)

	return savePath, nil
}
