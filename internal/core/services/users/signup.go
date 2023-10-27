package users

import (
	"context"
	"fmt"

	"github.com/samverrall/errmaps/internal/core/domain/user"
	"github.com/samverrall/errmaps/internal/core/services"
	"github.com/samverrall/errmaps/pkg/errsx"
)

type SignupReq struct {
	Email    string
	Name     string
	Password string
}

type SignupResp struct {
	UserID int
}

func (s *Service) Signup(ctx context.Context, req SignupReq) (*SignupResp, error) {
	var input struct {
		email    user.Email
		name     user.Name
		password user.Password
	}
	{
		var err error
		var errs errsx.Map

		input.email, err = user.NewEmail(req.Email)
		if err != nil {
			errs.Set("email", err)
		}

		input.name, err = user.NewName(req.Name)
		if err != nil {
			errs.Set("name", err)
		}

		input.password, err = user.NewPassword(req.Password)
		if err != nil {
			errs.Set("password", err)
		}

		// If we have a non-nil errs map here, we know we have some errors.
		if errs != nil {
			return nil, fmt.Errorf("%w: %w", services.ErrBadRequest, errs)
		}
	}

	// we have valid input, actually sign up

	return nil, nil
}

func (s *Service) SignupNotAsGood(ctx context.Context, req SignupReq) (*SignupResp, error) {
	var input struct {
		email    user.Email
		name     user.Name
		password user.Password
	}
	{
		var err error

		input.email, err = user.NewEmail(req.Email)
		if err != nil {
			return nil, err
		}

		input.name, err = user.NewName(req.Name)
		if err != nil {
			return nil, err
		}

		input.password, err = user.NewPassword(req.Password)
		if err != nil {
			return nil, err
		}

	}

	// we have valid input, actually sign up

	return nil, nil
}
