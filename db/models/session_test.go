// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testSessions(t *testing.T) {
	t.Parallel()

	query := Sessions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSessionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = session.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Sessions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSessionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SessionSlice{session}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSessionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SessionExists(tx, session.ID)
	if err != nil {
		t.Errorf("Unable to check if Session exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SessionExistsG to return true, but got false.")
	}
}
func testSessionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	sessionFound, err := FindSession(tx, session.ID)
	if err != nil {
		t.Error(err)
	}

	if sessionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSessionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Sessions(tx).Bind(session); err != nil {
		t.Error(err)
	}
}

func testSessionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Sessions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSessionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	sessionOne := &Session{}
	sessionTwo := &Session{}
	if err = randomize.Struct(seed, sessionOne, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}
	if err = randomize.Struct(seed, sessionTwo, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = sessionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = sessionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Sessions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSessionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	sessionOne := &Session{}
	sessionTwo := &Session{}
	if err = randomize.Struct(seed, sessionOne, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}
	if err = randomize.Struct(seed, sessionTwo, sessionDBTypes, false, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = sessionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = sessionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func sessionBeforeInsertHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterInsertHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterSelectHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionBeforeUpdateHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterUpdateHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionBeforeDeleteHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterDeleteHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionBeforeUpsertHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func sessionAfterUpsertHook(e boil.Executor, o *Session) error {
	*o = Session{}
	return nil
}

func testSessionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Session{}
	o := &Session{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, sessionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Session object: %s", err)
	}

	AddSessionHook(boil.BeforeInsertHook, sessionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	sessionBeforeInsertHooks = []SessionHook{}

	AddSessionHook(boil.AfterInsertHook, sessionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	sessionAfterInsertHooks = []SessionHook{}

	AddSessionHook(boil.AfterSelectHook, sessionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	sessionAfterSelectHooks = []SessionHook{}

	AddSessionHook(boil.BeforeUpdateHook, sessionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	sessionBeforeUpdateHooks = []SessionHook{}

	AddSessionHook(boil.AfterUpdateHook, sessionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	sessionAfterUpdateHooks = []SessionHook{}

	AddSessionHook(boil.BeforeDeleteHook, sessionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	sessionBeforeDeleteHooks = []SessionHook{}

	AddSessionHook(boil.AfterDeleteHook, sessionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	sessionAfterDeleteHooks = []SessionHook{}

	AddSessionHook(boil.BeforeUpsertHook, sessionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	sessionBeforeUpsertHooks = []SessionHook{}

	AddSessionHook(boil.AfterUpsertHook, sessionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	sessionAfterUpsertHooks = []SessionHook{}
}
func testSessionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSessionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx, sessionColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSessionToOneUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Session
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SessionSlice{&local}
	if err = local.L.LoadUser(tx, false, (*[]*Session)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSessionToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Session
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sessionDBTypes, false, strmangle.SetComplement(sessionPrimaryKeyColumns, sessionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Sessions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}
func testSessionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = session.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSessionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SessionSlice{session}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSessionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Sessions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	sessionDBTypes = map[string]string{`ID`: `char`, `LastSeen`: `timestamp`, `LoginTime`: `timestamp`, `UserID`: `char`}
	_              = bytes.MinRead
)

func testSessionsUpdate(t *testing.T) {
	t.Parallel()

	if len(sessionColumns) == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	if err = session.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSessionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(sessionColumns) == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	session := &Session{}
	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, session, sessionDBTypes, true, sessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(sessionColumns, sessionPrimaryKeyColumns) {
		fields = sessionColumns
	} else {
		fields = strmangle.SetComplement(
			sessionColumns,
			sessionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(session))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SessionSlice{session}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSessionsUpsert(t *testing.T) {
	t.Parallel()

	if len(sessionColumns) == len(sessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	session := Session{}
	if err = randomize.Struct(seed, &session, sessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = session.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Session: %s", err)
	}

	count, err := Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &session, sessionDBTypes, false, sessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Session struct: %s", err)
	}

	if err = session.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Session: %s", err)
	}

	count, err = Sessions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
