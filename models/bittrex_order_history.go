// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// BittrexOrderHistory is an object representing the database table.
type BittrexOrderHistory struct {
	ID                int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	UUID              string            `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	Exchange          string            `boil:"exchange" json:"exchange" toml:"exchange" yaml:"exchange"`
	Timestamp         time.Time         `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	OrderType         int               `boil:"order_type" json:"order_type" toml:"order_type" yaml:"order_type"`
	Limit             types.Decimal     `boil:"limit" json:"limit" toml:"limit" yaml:"limit"`
	Quantity          types.Decimal     `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	QuantityRemaining types.Decimal     `boil:"quantity_remaining" json:"quantity_remaining" toml:"quantity_remaining" yaml:"quantity_remaining"`
	Commission        types.Decimal     `boil:"commission" json:"commission" toml:"commission" yaml:"commission"`
	Price             types.Decimal     `boil:"price" json:"price" toml:"price" yaml:"price"`
	PricePerUnit      types.Decimal     `boil:"price_per_unit" json:"price_per_unit" toml:"price_per_unit" yaml:"price_per_unit"`
	IsConditional     bool              `boil:"is_conditional" json:"is_conditional" toml:"is_conditional" yaml:"is_conditional"`
	Condition         null.String       `boil:"condition" json:"condition,omitempty" toml:"condition" yaml:"condition,omitempty"`
	ConditionTarget   types.NullDecimal `boil:"condition_target" json:"condition_target,omitempty" toml:"condition_target" yaml:"condition_target,omitempty"`
	ImmediateOrCancel bool              `boil:"immediate_or_cancel" json:"immediate_or_cancel" toml:"immediate_or_cancel" yaml:"immediate_or_cancel"`
	Closed            time.Time         `boil:"closed" json:"closed" toml:"closed" yaml:"closed"`
	TimeInForceTypeID int               `boil:"time_in_force_type_id" json:"time_in_force_type_id" toml:"time_in_force_type_id" yaml:"time_in_force_type_id"`
	TimeInForce       null.String       `boil:"time_in_force" json:"time_in_force,omitempty" toml:"time_in_force" yaml:"time_in_force,omitempty"`

	R *bittrexOrderHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bittrexOrderHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BittrexOrderHistoryColumns = struct {
	ID                string
	UUID              string
	Exchange          string
	Timestamp         string
	OrderType         string
	Limit             string
	Quantity          string
	QuantityRemaining string
	Commission        string
	Price             string
	PricePerUnit      string
	IsConditional     string
	Condition         string
	ConditionTarget   string
	ImmediateOrCancel string
	Closed            string
	TimeInForceTypeID string
	TimeInForce       string
}{
	ID:                "id",
	UUID:              "uuid",
	Exchange:          "exchange",
	Timestamp:         "timestamp",
	OrderType:         "order_type",
	Limit:             "limit",
	Quantity:          "quantity",
	QuantityRemaining: "quantity_remaining",
	Commission:        "commission",
	Price:             "price",
	PricePerUnit:      "price_per_unit",
	IsConditional:     "is_conditional",
	Condition:         "condition",
	ConditionTarget:   "condition_target",
	ImmediateOrCancel: "immediate_or_cancel",
	Closed:            "closed",
	TimeInForceTypeID: "time_in_force_type_id",
	TimeInForce:       "time_in_force",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var BittrexOrderHistoryWhere = struct {
	ID                whereHelperint
	UUID              whereHelperstring
	Exchange          whereHelperstring
	Timestamp         whereHelpertime_Time
	OrderType         whereHelperint
	Limit             whereHelpertypes_Decimal
	Quantity          whereHelpertypes_Decimal
	QuantityRemaining whereHelpertypes_Decimal
	Commission        whereHelpertypes_Decimal
	Price             whereHelpertypes_Decimal
	PricePerUnit      whereHelpertypes_Decimal
	IsConditional     whereHelperbool
	Condition         whereHelpernull_String
	ConditionTarget   whereHelpertypes_NullDecimal
	ImmediateOrCancel whereHelperbool
	Closed            whereHelpertime_Time
	TimeInForceTypeID whereHelperint
	TimeInForce       whereHelpernull_String
}{
	ID:                whereHelperint{field: "`bittrex_order_history`.`id`"},
	UUID:              whereHelperstring{field: "`bittrex_order_history`.`uuid`"},
	Exchange:          whereHelperstring{field: "`bittrex_order_history`.`exchange`"},
	Timestamp:         whereHelpertime_Time{field: "`bittrex_order_history`.`timestamp`"},
	OrderType:         whereHelperint{field: "`bittrex_order_history`.`order_type`"},
	Limit:             whereHelpertypes_Decimal{field: "`bittrex_order_history`.`limit`"},
	Quantity:          whereHelpertypes_Decimal{field: "`bittrex_order_history`.`quantity`"},
	QuantityRemaining: whereHelpertypes_Decimal{field: "`bittrex_order_history`.`quantity_remaining`"},
	Commission:        whereHelpertypes_Decimal{field: "`bittrex_order_history`.`commission`"},
	Price:             whereHelpertypes_Decimal{field: "`bittrex_order_history`.`price`"},
	PricePerUnit:      whereHelpertypes_Decimal{field: "`bittrex_order_history`.`price_per_unit`"},
	IsConditional:     whereHelperbool{field: "`bittrex_order_history`.`is_conditional`"},
	Condition:         whereHelpernull_String{field: "`bittrex_order_history`.`condition`"},
	ConditionTarget:   whereHelpertypes_NullDecimal{field: "`bittrex_order_history`.`condition_target`"},
	ImmediateOrCancel: whereHelperbool{field: "`bittrex_order_history`.`immediate_or_cancel`"},
	Closed:            whereHelpertime_Time{field: "`bittrex_order_history`.`closed`"},
	TimeInForceTypeID: whereHelperint{field: "`bittrex_order_history`.`time_in_force_type_id`"},
	TimeInForce:       whereHelpernull_String{field: "`bittrex_order_history`.`time_in_force`"},
}

// BittrexOrderHistoryRels is where relationship names are stored.
var BittrexOrderHistoryRels = struct {
}{}

// bittrexOrderHistoryR is where relationships are stored.
type bittrexOrderHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*bittrexOrderHistoryR) NewStruct() *bittrexOrderHistoryR {
	return &bittrexOrderHistoryR{}
}

// bittrexOrderHistoryL is where Load methods for each relationship are stored.
type bittrexOrderHistoryL struct{}

var (
	bittrexOrderHistoryAllColumns            = []string{"id", "uuid", "exchange", "timestamp", "order_type", "limit", "quantity", "quantity_remaining", "commission", "price", "price_per_unit", "is_conditional", "condition", "condition_target", "immediate_or_cancel", "closed", "time_in_force_type_id", "time_in_force"}
	bittrexOrderHistoryColumnsWithoutDefault = []string{"uuid", "exchange", "timestamp", "order_type", "limit", "quantity", "quantity_remaining", "commission", "price", "price_per_unit", "is_conditional", "condition", "condition_target", "immediate_or_cancel", "closed", "time_in_force_type_id", "time_in_force"}
	bittrexOrderHistoryColumnsWithDefault    = []string{"id"}
	bittrexOrderHistoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// BittrexOrderHistorySlice is an alias for a slice of pointers to BittrexOrderHistory.
	// This should generally be used opposed to []BittrexOrderHistory.
	BittrexOrderHistorySlice []*BittrexOrderHistory
	// BittrexOrderHistoryHook is the signature for custom BittrexOrderHistory hook methods
	BittrexOrderHistoryHook func(context.Context, boil.ContextExecutor, *BittrexOrderHistory) error

	bittrexOrderHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bittrexOrderHistoryType                 = reflect.TypeOf(&BittrexOrderHistory{})
	bittrexOrderHistoryMapping              = queries.MakeStructMapping(bittrexOrderHistoryType)
	bittrexOrderHistoryPrimaryKeyMapping, _ = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, bittrexOrderHistoryPrimaryKeyColumns)
	bittrexOrderHistoryInsertCacheMut       sync.RWMutex
	bittrexOrderHistoryInsertCache          = make(map[string]insertCache)
	bittrexOrderHistoryUpdateCacheMut       sync.RWMutex
	bittrexOrderHistoryUpdateCache          = make(map[string]updateCache)
	bittrexOrderHistoryUpsertCacheMut       sync.RWMutex
	bittrexOrderHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bittrexOrderHistoryBeforeInsertHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryBeforeUpdateHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryBeforeDeleteHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryBeforeUpsertHooks []BittrexOrderHistoryHook

var bittrexOrderHistoryAfterInsertHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryAfterSelectHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryAfterUpdateHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryAfterDeleteHooks []BittrexOrderHistoryHook
var bittrexOrderHistoryAfterUpsertHooks []BittrexOrderHistoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BittrexOrderHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BittrexOrderHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BittrexOrderHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BittrexOrderHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BittrexOrderHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BittrexOrderHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BittrexOrderHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BittrexOrderHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BittrexOrderHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bittrexOrderHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBittrexOrderHistoryHook registers your hook function for all future operations.
func AddBittrexOrderHistoryHook(hookPoint boil.HookPoint, bittrexOrderHistoryHook BittrexOrderHistoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bittrexOrderHistoryBeforeInsertHooks = append(bittrexOrderHistoryBeforeInsertHooks, bittrexOrderHistoryHook)
	case boil.BeforeUpdateHook:
		bittrexOrderHistoryBeforeUpdateHooks = append(bittrexOrderHistoryBeforeUpdateHooks, bittrexOrderHistoryHook)
	case boil.BeforeDeleteHook:
		bittrexOrderHistoryBeforeDeleteHooks = append(bittrexOrderHistoryBeforeDeleteHooks, bittrexOrderHistoryHook)
	case boil.BeforeUpsertHook:
		bittrexOrderHistoryBeforeUpsertHooks = append(bittrexOrderHistoryBeforeUpsertHooks, bittrexOrderHistoryHook)
	case boil.AfterInsertHook:
		bittrexOrderHistoryAfterInsertHooks = append(bittrexOrderHistoryAfterInsertHooks, bittrexOrderHistoryHook)
	case boil.AfterSelectHook:
		bittrexOrderHistoryAfterSelectHooks = append(bittrexOrderHistoryAfterSelectHooks, bittrexOrderHistoryHook)
	case boil.AfterUpdateHook:
		bittrexOrderHistoryAfterUpdateHooks = append(bittrexOrderHistoryAfterUpdateHooks, bittrexOrderHistoryHook)
	case boil.AfterDeleteHook:
		bittrexOrderHistoryAfterDeleteHooks = append(bittrexOrderHistoryAfterDeleteHooks, bittrexOrderHistoryHook)
	case boil.AfterUpsertHook:
		bittrexOrderHistoryAfterUpsertHooks = append(bittrexOrderHistoryAfterUpsertHooks, bittrexOrderHistoryHook)
	}
}

// One returns a single bittrexOrderHistory record from the query.
func (q bittrexOrderHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BittrexOrderHistory, error) {
	o := &BittrexOrderHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bittrex_order_history")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BittrexOrderHistory records from the query.
func (q bittrexOrderHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (BittrexOrderHistorySlice, error) {
	var o []*BittrexOrderHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BittrexOrderHistory slice")
	}

	if len(bittrexOrderHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BittrexOrderHistory records in the query.
func (q bittrexOrderHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bittrex_order_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bittrexOrderHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bittrex_order_history exists")
	}

	return count > 0, nil
}

// BittrexOrderHistories retrieves all the records using an executor.
func BittrexOrderHistories(mods ...qm.QueryMod) bittrexOrderHistoryQuery {
	mods = append(mods, qm.From("`bittrex_order_history`"))
	return bittrexOrderHistoryQuery{NewQuery(mods...)}
}

// FindBittrexOrderHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBittrexOrderHistory(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BittrexOrderHistory, error) {
	bittrexOrderHistoryObj := &BittrexOrderHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bittrex_order_history` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bittrexOrderHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bittrex_order_history")
	}

	return bittrexOrderHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BittrexOrderHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bittrex_order_history provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bittrexOrderHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bittrexOrderHistoryInsertCacheMut.RLock()
	cache, cached := bittrexOrderHistoryInsertCache[key]
	bittrexOrderHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bittrexOrderHistoryAllColumns,
			bittrexOrderHistoryColumnsWithDefault,
			bittrexOrderHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bittrex_order_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bittrex_order_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bittrex_order_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bittrexOrderHistoryPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into bittrex_order_history")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bittrexOrderHistoryMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bittrex_order_history")
	}

CacheNoHooks:
	if !cached {
		bittrexOrderHistoryInsertCacheMut.Lock()
		bittrexOrderHistoryInsertCache[key] = cache
		bittrexOrderHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BittrexOrderHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BittrexOrderHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bittrexOrderHistoryUpdateCacheMut.RLock()
	cache, cached := bittrexOrderHistoryUpdateCache[key]
	bittrexOrderHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bittrexOrderHistoryAllColumns,
			bittrexOrderHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bittrex_order_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bittrex_order_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bittrexOrderHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, append(wl, bittrexOrderHistoryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update bittrex_order_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bittrex_order_history")
	}

	if !cached {
		bittrexOrderHistoryUpdateCacheMut.Lock()
		bittrexOrderHistoryUpdateCache[key] = cache
		bittrexOrderHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bittrexOrderHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bittrex_order_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bittrex_order_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BittrexOrderHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bittrexOrderHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bittrex_order_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bittrexOrderHistoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bittrexOrderHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bittrexOrderHistory")
	}
	return rowsAff, nil
}

var mySQLBittrexOrderHistoryUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BittrexOrderHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bittrex_order_history provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bittrexOrderHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBittrexOrderHistoryUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bittrexOrderHistoryUpsertCacheMut.RLock()
	cache, cached := bittrexOrderHistoryUpsertCache[key]
	bittrexOrderHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bittrexOrderHistoryAllColumns,
			bittrexOrderHistoryColumnsWithDefault,
			bittrexOrderHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bittrexOrderHistoryAllColumns,
			bittrexOrderHistoryPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert bittrex_order_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`bittrex_order_history`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bittrex_order_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for bittrex_order_history")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bittrexOrderHistoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(bittrexOrderHistoryType, bittrexOrderHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for bittrex_order_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for bittrex_order_history")
	}

CacheNoHooks:
	if !cached {
		bittrexOrderHistoryUpsertCacheMut.Lock()
		bittrexOrderHistoryUpsertCache[key] = cache
		bittrexOrderHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BittrexOrderHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BittrexOrderHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BittrexOrderHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bittrexOrderHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `bittrex_order_history` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bittrex_order_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bittrex_order_history")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bittrexOrderHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bittrexOrderHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bittrex_order_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bittrex_order_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BittrexOrderHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bittrexOrderHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bittrexOrderHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bittrex_order_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bittrexOrderHistoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bittrexOrderHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bittrex_order_history")
	}

	if len(bittrexOrderHistoryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BittrexOrderHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBittrexOrderHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BittrexOrderHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BittrexOrderHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bittrexOrderHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bittrex_order_history`.* FROM `bittrex_order_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bittrexOrderHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BittrexOrderHistorySlice")
	}

	*o = slice

	return nil
}

// BittrexOrderHistoryExists checks if the BittrexOrderHistory row exists.
func BittrexOrderHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bittrex_order_history` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bittrex_order_history exists")
	}

	return exists, nil
}
