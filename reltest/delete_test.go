package reltest

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	var (
		repo = New()
	)

	repo.ExpectDelete().For(&Book{ID: 1})
	assert.Nil(t, repo.Delete(context.TODO(), &Book{ID: 1}))
	repo.AssertExpectations(t)

	repo.ExpectDelete().For(&Book{ID: 1})
	assert.NotPanics(t, func() {
		repo.MustDelete(context.TODO(), &Book{ID: 1})
	})
	repo.AssertExpectations(t)
}

func TestDelete_forType(t *testing.T) {
	var (
		repo = New()
	)

	repo.ExpectDelete().ForType("reltest.Book")
	assert.Nil(t, repo.Delete(context.TODO(), &Book{ID: 1}))
	repo.AssertExpectations(t)

	repo.ExpectDelete().ForType("reltest.Book")
	assert.NotPanics(t, func() {
		repo.MustDelete(context.TODO(), &Book{ID: 1})
	})
	repo.AssertExpectations(t)
}

func TestDelete_error(t *testing.T) {
	var (
		repo = New()
	)

	repo.ExpectDelete().ConnectionClosed()
	assert.Equal(t, sql.ErrConnDone, repo.Delete(context.TODO(), &Book{ID: 1}))
	repo.AssertExpectations(t)

	repo.ExpectDelete().ConnectionClosed()
	assert.Panics(t, func() {
		repo.MustDelete(context.TODO(), &Book{ID: 1})
	})
	repo.AssertExpectations(t)
}

func TestDeleteAll(t *testing.T) {
	var (
		repo = New()
	)

	repo.ExpectDeleteAll().For(&[]Book{{ID: 1}})
	assert.Nil(t, repo.DeleteAll(context.TODO(), &[]Book{{ID: 1}}))
	repo.AssertExpectations(t)

	repo.ExpectDeleteAll().For(&[]Book{{ID: 1}})
	assert.NotPanics(t, func() {
		repo.MustDeleteAll(context.TODO(), &[]Book{{ID: 1}})
	})
	repo.AssertExpectations(t)
}
