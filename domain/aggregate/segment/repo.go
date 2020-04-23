package segment

import (
	"context"
	"github.com/JerryZhou343/leaf-go/domain/aggregate/segment/entity"
)

type Repo interface {
	GetAllLeafAllocs(ctx context.Context) ([]*entity.LeafAlloc, error)
	UpdateMaxIdAndGetLeafAlloc(ctx context.Context, tag string) (*entity.LeafAlloc, error)
	UpdateMaxIdByCustomStepAndGetLeafAlloc(ctx context.Context, leafAlloc *entity.LeafAlloc) (*entity.LeafAlloc, error)
	GetAllTags(ctx context.Context) ([]string, error)
}
