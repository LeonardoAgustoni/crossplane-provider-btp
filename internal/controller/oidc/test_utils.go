package oidc

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrNoResource = errors.New("Can't find resource")
)

type TrackerMock struct {
	wasCalled bool
}

func (t *TrackerMock) Track(ctx context.Context, mg resource.Managed) error {
	t.wasCalled = true
	return nil
}

func MockTracker() resource.Tracker {
	return &TrackerMock{}
}

func MockCertLookup(certs []corev1.Secret, deleteRecorder *string) *test.MockClient {
	return &test.MockClient{
		MockGet: func(ctx context.Context, key client.ObjectKey, obj client.Object) error {
			secret, ok := obj.(*corev1.Secret)
			if !ok {
				return errors.New("Unexpected lookup")
			}
			for _, v := range certs {
				if key.Name == v.Name {
					secret.Name = v.Name
					secret.Data = v.Data
					return nil
				}
			}
			return ErrNoResource
		},
		MockDelete: func(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
			if deleteRecorder != nil {
				*deleteRecorder = obj.GetName()
			}
			return nil
		},
	}
}
