// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AuthFallback is an object representing the database table.
type AuthFallback struct {
	ID           int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name         string `boil:"name" json:"name" toml:"name" yaml:"name"`
	PasswordHash []byte `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`

	R *authFallbackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authFallbackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthFallbackColumns = struct {
	ID           string
	Name         string
	PasswordHash string
}{
	ID:           "id",
	Name:         "name",
	PasswordHash: "password_hash",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var AuthFallbackWhere = struct {
	ID           whereHelperint64
	Name         whereHelperstring
	PasswordHash whereHelper__byte
}{
	ID:           whereHelperint64{field: "\"auth_fallback\".\"id\""},
	Name:         whereHelperstring{field: "\"auth_fallback\".\"name\""},
	PasswordHash: whereHelper__byte{field: "\"auth_fallback\".\"password_hash\""},
}

// AuthFallbackRels is where relationship names are stored.
var AuthFallbackRels = struct {
}{}

// authFallbackR is where relationships are stored.
type authFallbackR struct {
}

// NewStruct creates a new relationship struct
func (*authFallbackR) NewStruct() *authFallbackR {
	return &authFallbackR{}
}

// authFallbackL is where Load methods for each relationship are stored.
type authFallbackL struct{}

var (
	authFallbackAllColumns            = []string{"id", "name", "password_hash"}
	authFallbackColumnsWithoutDefault = []string{}
	authFallbackColumnsWithDefault    = []string{"id", "name", "password_hash"}
	authFallbackPrimaryKeyColumns     = []string{"id"}
)

type (
	// AuthFallbackSlice is an alias for a slice of pointers to AuthFallback.
	// This should generally be used opposed to []AuthFallback.
	AuthFallbackSlice []*AuthFallback
	// AuthFallbackHook is the signature for custom AuthFallback hook methods
	AuthFallbackHook func(context.Context, boil.ContextExecutor, *AuthFallback) error

	authFallbackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authFallbackType                 = reflect.TypeOf(&AuthFallback{})
	authFallbackMapping              = queries.MakeStructMapping(authFallbackType)
	authFallbackPrimaryKeyMapping, _ = queries.BindMapping(authFallbackType, authFallbackMapping, authFallbackPrimaryKeyColumns)
	authFallbackInsertCacheMut       sync.RWMutex
	authFallbackInsertCache          = make(map[string]insertCache)
	authFallbackUpdateCacheMut       sync.RWMutex
	authFallbackUpdateCache          = make(map[string]updateCache)
	authFallbackUpsertCacheMut       sync.RWMutex
	authFallbackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authFallbackBeforeInsertHooks []AuthFallbackHook
var authFallbackBeforeUpdateHooks []AuthFallbackHook
var authFallbackBeforeDeleteHooks []AuthFallbackHook
var authFallbackBeforeUpsertHooks []AuthFallbackHook

var authFallbackAfterInsertHooks []AuthFallbackHook
var authFallbackAfterSelectHooks []AuthFallbackHook
var authFallbackAfterUpdateHooks []AuthFallbackHook
var authFallbackAfterDeleteHooks []AuthFallbackHook
var authFallbackAfterUpsertHooks []AuthFallbackHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthFallback) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthFallback) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthFallback) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthFallback) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthFallback) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthFallback) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthFallback) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthFallback) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthFallback) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authFallbackAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthFallbackHook registers your hook function for all future operations.
func AddAuthFallbackHook(hookPoint boil.HookPoint, authFallbackHook AuthFallbackHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authFallbackBeforeInsertHooks = append(authFallbackBeforeInsertHooks, authFallbackHook)
	case boil.BeforeUpdateHook:
		authFallbackBeforeUpdateHooks = append(authFallbackBeforeUpdateHooks, authFallbackHook)
	case boil.BeforeDeleteHook:
		authFallbackBeforeDeleteHooks = append(authFallbackBeforeDeleteHooks, authFallbackHook)
	case boil.BeforeUpsertHook:
		authFallbackBeforeUpsertHooks = append(authFallbackBeforeUpsertHooks, authFallbackHook)
	case boil.AfterInsertHook:
		authFallbackAfterInsertHooks = append(authFallbackAfterInsertHooks, authFallbackHook)
	case boil.AfterSelectHook:
		authFallbackAfterSelectHooks = append(authFallbackAfterSelectHooks, authFallbackHook)
	case boil.AfterUpdateHook:
		authFallbackAfterUpdateHooks = append(authFallbackAfterUpdateHooks, authFallbackHook)
	case boil.AfterDeleteHook:
		authFallbackAfterDeleteHooks = append(authFallbackAfterDeleteHooks, authFallbackHook)
	case boil.AfterUpsertHook:
		authFallbackAfterUpsertHooks = append(authFallbackAfterUpsertHooks, authFallbackHook)
	}
}

// One returns a single authFallback record from the query.
func (q authFallbackQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthFallback, error) {
	o := &AuthFallback{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_fallback")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuthFallback records from the query.
func (q authFallbackQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthFallbackSlice, error) {
	var o []*AuthFallback

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthFallback slice")
	}

	if len(authFallbackAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuthFallback records in the query.
func (q authFallbackQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_fallback rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authFallbackQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_fallback exists")
	}

	return count > 0, nil
}

// AuthFallbacks retrieves all the records using an executor.
func AuthFallbacks(mods ...qm.QueryMod) authFallbackQuery {
	mods = append(mods, qm.From("\"auth_fallback\""))
	return authFallbackQuery{NewQuery(mods...)}
}

// FindAuthFallback retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthFallback(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AuthFallback, error) {
	authFallbackObj := &AuthFallback{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_fallback\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authFallbackObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_fallback")
	}

	return authFallbackObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthFallback) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_fallback provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authFallbackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authFallbackInsertCacheMut.RLock()
	cache, cached := authFallbackInsertCache[key]
	authFallbackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authFallbackAllColumns,
			authFallbackColumnsWithDefault,
			authFallbackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authFallbackType, authFallbackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authFallbackType, authFallbackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"auth_fallback\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"auth_fallback\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"auth_fallback\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, authFallbackPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into auth_fallback")
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

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authFallbackMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for auth_fallback")
	}

CacheNoHooks:
	if !cached {
		authFallbackInsertCacheMut.Lock()
		authFallbackInsertCache[key] = cache
		authFallbackInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuthFallback.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthFallback) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authFallbackUpdateCacheMut.RLock()
	cache, cached := authFallbackUpdateCache[key]
	authFallbackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authFallbackAllColumns,
			authFallbackPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update auth_fallback, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_fallback\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, authFallbackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authFallbackType, authFallbackMapping, append(wl, authFallbackPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update auth_fallback row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for auth_fallback")
	}

	if !cached {
		authFallbackUpdateCacheMut.Lock()
		authFallbackUpdateCache[key] = cache
		authFallbackUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q authFallbackQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for auth_fallback")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for auth_fallback")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthFallbackSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authFallbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"auth_fallback\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, authFallbackPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in authFallback slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all authFallback")
	}
	return rowsAff, nil
}

// Delete deletes a single AuthFallback record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthFallback) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuthFallback provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authFallbackPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_fallback\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from auth_fallback")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for auth_fallback")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authFallbackQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no authFallbackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auth_fallback")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_fallback")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthFallbackSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authFallbackBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authFallbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"auth_fallback\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, authFallbackPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from authFallback slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_fallback")
	}

	if len(authFallbackAfterDeleteHooks) != 0 {
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
func (o *AuthFallback) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthFallback(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthFallbackSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthFallbackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authFallbackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"auth_fallback\".* FROM \"auth_fallback\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, authFallbackPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthFallbackSlice")
	}

	*o = slice

	return nil
}

// AuthFallbackExists checks if the AuthFallback row exists.
func AuthFallbackExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"auth_fallback\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_fallback exists")
	}

	return exists, nil
}
