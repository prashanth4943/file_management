package databuckets

import (
	// "bytes"
	// "context"
	// "io"
	// "io/ioutil"
	"dms-backend/internal/db"
	"log"
	"os"

	// "os"
	// "github.com/joho/godotenv"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func UploadToOracle(filePath string, Email string) (string, error) {

	provider := common.DefaultConfigProvider()

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
	if err != nil {
		panic(err)
	}

	bucketName := "test_bucket_1"
	namespace := "axpzvvirgqhg"
	// filePath := "example.txt"
	objectName := fmt.Sprintf("%s.png", Email)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file information: %v", err)
	}

	// Prepare PutObject request
	request := objectstorage.PutObjectRequest{
		NamespaceName: common.String(namespace),
		BucketName:    common.String(bucketName),
		ObjectName:    common.String(objectName),
		PutObjectBody: file,
		ContentLength: common.Int64(fileInfo.Size()),
	}

	// Upload the file
	response, err := client.PutObject(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to upload object: %v", err)
	}
	query := `
	INSERT INTO users (email, image_name, uuid)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE
	image_name = VALUES(image_name), uuid = VALUES(uuid);`

	uniqueID := uuid.New().String()
	_, err = db.DBInstance.Conn.Exec(query, Email, objectName, uniqueID)
	if err != nil {
		return "", fmt.Errorf("failed to insert/update record in the database: %v", err)
	}
	fmt.Printf("File uploaded successfully! ETag: %s\n", *response.ETag)

	return uniqueID, nil
	// return "adass", nil
}
