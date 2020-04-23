package segment

import (
	"context"
	"github.com/JerryZhou343/leaf-go/domain/aggregate/segment"
	"github.com/JerryZhou343/leaf-go/domain/aggregate/segment/entity"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	"time"
)

const (
	_allLeafAllocs         = "select biz_tag,max_id,step,description,update_time from leaf_alloc"
	_selectLeafAlloc       = "select biz_tag,max_id,step,description,update_time from leaf_alloc where biz_tag = ?"
	_allTag                = "select biz_tag from leaf_alloc"
	_updateMaxId           = "UPDATE leaf_alloc SET max_id = max_id + step,update_time = ? WHERE biz_tag = ?"
	_updateMaxIdCustomStep = "UPDATE leaf_alloc SET max_id = max_id + ?, update_time = ? WHERE biz_tag = ?"
)

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) segment.Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetAllLeafAllocs(ctx context.Context) (ret []*entity.LeafAlloc, err error) {
	var (
		rows *sql.Rows
	)
	rows, err = r.db.Query(ctx, _allLeafAllocs)
	if err != nil {
		log.Errorv(ctx, log.KV("event", "mysql_query"), log.KV("error", err), log.KV("sql", _allLeafAllocs))
		return
	}
	defer rows.Close()
	for rows.Next() {
		leafAlloc := new(entity.LeafAlloc)
		if err = rows.Scan(&leafAlloc.BizTag,
			&leafAlloc.MaxID,
			&leafAlloc.Step,
			&leafAlloc.Description,
			&leafAlloc.UpdateTime); err != nil {
			log.Errorv(ctx, log.KV("event", "mysql_scan"), log.KV("error", err), log.KV("sql", _allLeafAllocs))
			return
		}

		ret = append(ret, leafAlloc)
	}

	log.Infov(ctx, log.KV("event", "mysql_query"), log.KV("row_num", len(ret)), log.KV("sql", _allLeafAllocs))
	return
}

func (r *repo) UpdateMaxIdAndGetLeafAlloc(ctx context.Context, tag string) (ret *entity.LeafAlloc, err error) {
	var (
		tx  *sql.Tx
		row *sql.Row
	)
	tx, err = r.db.Begin(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec(_updateMaxId, time.Now(), tag)
	if err != nil {
		log.Errorv(ctx, log.KV("event", "mysql_exec"), log.KV("error", err), log.KV("sql", _updateMaxId))
		return
	}
	row = tx.QueryRow(_selectLeafAlloc, tag)
	ret = new(entity.LeafAlloc)
	if err = row.Scan(&ret.BizTag, &ret.MaxID, &ret.Step, &ret.Description, &ret.UpdateTime); err != nil {
		log.Errorv(ctx, log.KV("event", "mysql_scan"), log.KV("error", err), log.KV("sql", _updateMaxId))
		return
	}

	return
}

func (r *repo) UpdateMaxIdByCustomStepAndGetLeafAlloc(ctx context.Context, leafAlloc *entity.LeafAlloc) (ret *entity.LeafAlloc, err error) {

	var (
		tx  *sql.Tx
		row *sql.Row
	)
	tx, err = r.db.Begin(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec(_updateMaxIdCustomStep, leafAlloc.Step, time.Now(), leafAlloc.BizTag)
	if err != nil {
		log.Errorv(ctx, log.KV("event", "mysql_exec"), log.KV("error", err), log.KV("sql", _updateMaxId))
		return
	}
	row = tx.QueryRow(_selectLeafAlloc, leafAlloc.BizTag)
	ret = new(entity.LeafAlloc)
	if err = row.Scan(&ret.BizTag, &ret.MaxID, &ret.Step, &ret.Description, &ret.UpdateTime); err != nil {
		log.Errorv(ctx, log.KV("event", "mysql_scan"), log.KV("error", err), log.KV("sql", _updateMaxId))
		return
	}

	return
}

func (r *repo) GetAllTags(ctx context.Context) (ret []string, err error) {
	var (
		rows *sql.Rows
	)
	rows, err = r.db.Query(ctx, _allTag)
	if err != nil {
		log.Errorv(ctx, log.KV("event", "mysql_query"), log.KV("error", err), log.KV("sql", _allTag))
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tmp string
		if err = rows.Scan(&tmp); err != nil {
			log.Errorv(ctx, log.KV("event", "mysql_scan"), log.KV("error", err), log.KV("sql", _allTag))
			return
		}
		ret = append(ret, tmp)
	}
	log.Infov(ctx, log.KV("event", "mysql_query"), log.KV("row_num", len(ret)), log.KV("sql", _allTag))
	return
}
