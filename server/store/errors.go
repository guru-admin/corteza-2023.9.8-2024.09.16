package store

import "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/errors"

type (
	// ErrorHandler
	// each implementation can have internal error handler that can translate
	// impl. specific errors like transaction
	ErrorHandler func(error) error
)

var (
	ErrNotFound  = errors.Plain(errors.KindNotFound, "not found")
	ErrNotUnique = errors.Plain(errors.KindDuplicateData, "not unique")
)

func HandleError(err error, h ErrorHandler) error {
	if err == nil {
		return nil
	}

	if h != nil {
		err = h(err)
	}

	if _, wrapped := err.(*errors.Error); wrapped {
		return err
	}

	return errors.
		Store("store error: %v", err).
		Apply(errors.StackSkip(1)).
		Wrap(err)
}
