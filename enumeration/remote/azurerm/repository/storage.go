package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/azurerm/common"
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/cache"
	"github.com/sirupsen/logrus"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/go-autorest/autorest/azure"
)

type StorageRespository interface {
	ListAllStorageAccount() ([]*armstorage.StorageAccount, error)
	ListAllStorageContainer(account *armstorage.StorageAccount) ([]string, error)
}

type blobContainerListPager interface {
	pager
	PageResponse() armstorage.BlobContainersListResponse
}

// Interfaces are only used to create mock on Azure SDK
type blobContainerClient interface {
	List(resourceGroupName string, accountName string, options *armstorage.BlobContainersListOptions) blobContainerListPager
}

type blobContainerClientImpl struct {
	client *armstorage.BlobContainersClient
}

func (c blobContainerClientImpl) List(resourceGroupName string, accountName string, options *armstorage.BlobContainersListOptions) blobContainerListPager {
	return c.client.List(resourceGroupName, accountName, options)
}

type storageAccountListPager interface {
	pager
	PageResponse() armstorage.StorageAccountsListResponse
}

type storageAccountClient interface {
	List(options *armstorage.StorageAccountsListOptions) storageAccountListPager
}

type storageAccountClientImpl struct {
	client *armstorage.StorageAccountsClient
}

func (c storageAccountClientImpl) List(options *armstorage.StorageAccountsListOptions) storageAccountListPager {
	return c.client.List(options)
}

type storageRepository struct {
	storageAccountsClient storageAccountClient
	blobContainerClient   blobContainerClient
	cache                 cache.Cache
}

func NewStorageRepository(cred azcore.TokenCredential, options *arm.ClientOptions, config common.AzureProviderConfig, cache cache.Cache) *storageRepository {
	return &storageRepository{
		storageAccountClientImpl{client: armstorage.NewStorageAccountsClient(config.SubscriptionID, cred, options)},
		blobContainerClientImpl{client: armstorage.NewBlobContainersClient(config.SubscriptionID, cred, options)},
		cache,
	}
}

func (s *storageRepository) ListAllStorageAccount() ([]*armstorage.StorageAccount, error) {

	cacheKey := "ListAllStorageAccount"
	v := s.cache.GetAndLock(cacheKey)
	defer s.cache.Unlock(cacheKey)
	if v != nil {
		return v.([]*armstorage.StorageAccount), nil
	}

	pager := s.storageAccountsClient.List(nil)
	results := make([]*armstorage.StorageAccount, 0)
	for pager.NextPage(context.Background()) {
		resp := pager.PageResponse()
		if err := pager.Err(); err != nil {
			return nil, err
		}
		results = append(results, resp.StorageAccountsListResult.StorageAccountListResult.Value...)
	}

	if err := pager.Err(); err != nil {
		return nil, err
	}

	s.cache.Put(cacheKey, results)

	return results, nil
}

func (s *storageRepository) ListAllStorageContainer(account *armstorage.StorageAccount) ([]string, error) {

	cacheKey := fmt.Sprintf("ListAllStorageContainer_%s", *account.Name)
	if v := s.cache.Get(cacheKey); v != nil {
		return v.([]string), nil
	}

	res, err := azure.ParseResourceID(*account.ID)
	if err != nil {
		return nil, err
	}

	pager := s.blobContainerClient.List(res.ResourceGroup, *account.Name, nil)
	results := make([]string, 0)
	for pager.NextPage(context.Background()) {
		resp := pager.PageResponse()
		if err := pager.Err(); err != nil {
			if !shouldIgnoreStorageContainerError(err) {
				return nil, err
			}
		}
		for _, item := range resp.BlobContainersListResult.ListContainerItems.Value {
			results = append(results, fmt.Sprintf("%s%s", *account.Properties.PrimaryEndpoints.Blob, *item.Name))
		}
	}

	if err := pager.Err(); err != nil {
		if !shouldIgnoreStorageContainerError(err) {
			return nil, err
		}
	}

	s.cache.Put(cacheKey, results)

	return results, nil
}

func shouldIgnoreStorageContainerError(err error) bool {
	azureErr, ok := err.(azblob.ResponseError)
	if !ok {
		return false
	}
	unwrapped := azureErr.Unwrap().Error()
	if strings.Contains(unwrapped, "FeatureNotSupportedForAccount") {
		logrus.WithFields(logrus.Fields{
			"repository": "StorageRepository",
			"error":      err,
		}).Debug("Ignoring ListStorageContainer error ...")
		return true
	}
	return false
}
