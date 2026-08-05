package tfs_storage

import (
	"context"
	"sync"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// TfsStorage tfs存储实现
type TfsStorage struct {
	mu sync.RWMutex
}

var _ storage.Storage = &TfsStorage{}

func NewTfsStorage(ctx context.Context) (*TfsStorage, error) {
	return &TfsStorage{}, nil
}

const StorageName = "tfs-storage"

func (x *TfsStorage) GetName() string { return StorageName }
func (x *TfsStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{storage.CapabilityCAS, storage.CapabilityReliableTime}
}
func (x *TfsStorage) Init(ctx context.Context) error { return nil }
func (x *TfsStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}
func (x *TfsStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}
func (x *TfsStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}
func (x *TfsStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return "", storage_lock.ErrLockNotFound
}
func (x *TfsStorage) GetTime(ctx context.Context) (time.Time, error) { return time.Now(), nil }
func (x *TfsStorage) Close(ctx context.Context) error { return nil }
func (x *TfsStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}
