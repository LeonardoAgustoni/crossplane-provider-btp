/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this RoleCollection.
func (mg *RoleCollection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecret,
		Extract:      SubaccountApiCredentialSecret(),
		Reference:    mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef,
		Selector:     mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSelector,
		To: reference.To{
			List:    &SubaccountApiCredentialList{},
			Managed: &SubaccountApiCredential{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecret")
	}
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecret = rsp.ResolvedValue
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecretNamespace,
		Extract:      SubaccountApiCredentialSecretSecretNamespace(),
		Reference:    mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef,
		Selector:     mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSelector,
		To: reference.To{
			List:    &SubaccountApiCredentialList{},
			Managed: &SubaccountApiCredential{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecretNamespace")
	}
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecretNamespace = rsp.ResolvedValue
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleCollectionAssignment.
func (mg *RoleCollectionAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecret,
		Extract:      SubaccountApiCredentialSecret(),
		Reference:    mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef,
		Selector:     mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSelector,
		To: reference.To{
			List:    &SubaccountApiCredentialList{},
			Managed: &SubaccountApiCredential{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecret")
	}
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecret = rsp.ResolvedValue
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecretNamespace,
		Extract:      SubaccountApiCredentialSecretSecretNamespace(),
		Reference:    mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef,
		Selector:     mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSelector,
		To: reference.To{
			List:    &SubaccountApiCredentialList{},
			Managed: &SubaccountApiCredential{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecretNamespace")
	}
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialSecretNamespace = rsp.ResolvedValue
	mg.Spec.XSUAACredentialsReference.SubaccountApiCredentialRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubaccountApiCredential.
func (mg *SubaccountApiCredential) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubaccountID),
		Extract:      v1alpha1.SubaccountUuid(),
		Reference:    mg.Spec.ForProvider.SubaccountRef,
		Selector:     mg.Spec.ForProvider.SubaccountSelector,
		To: reference.To{
			List:    &v1alpha1.SubaccountList{},
			Managed: &v1alpha1.Subaccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubaccountID")
	}
	mg.Spec.ForProvider.SubaccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubaccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubaccountID),
		Extract:      v1alpha1.SubaccountUuid(),
		Reference:    mg.Spec.InitProvider.SubaccountRef,
		Selector:     mg.Spec.InitProvider.SubaccountSelector,
		To: reference.To{
			List:    &v1alpha1.SubaccountList{},
			Managed: &v1alpha1.Subaccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubaccountID")
	}
	mg.Spec.InitProvider.SubaccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubaccountRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubaccountTrustConfiguration.
func (mg *SubaccountTrustConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubaccountID),
		Extract:      v1alpha1.SubaccountUuid(),
		Reference:    mg.Spec.ForProvider.SubaccountRef,
		Selector:     mg.Spec.ForProvider.SubaccountSelector,
		To: reference.To{
			List:    &v1alpha1.SubaccountList{},
			Managed: &v1alpha1.Subaccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubaccountID")
	}
	mg.Spec.ForProvider.SubaccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubaccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubaccountID),
		Extract:      v1alpha1.SubaccountUuid(),
		Reference:    mg.Spec.InitProvider.SubaccountRef,
		Selector:     mg.Spec.InitProvider.SubaccountSelector,
		To: reference.To{
			List:    &v1alpha1.SubaccountList{},
			Managed: &v1alpha1.Subaccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubaccountID")
	}
	mg.Spec.InitProvider.SubaccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubaccountRef = rsp.ResolvedReference

	return nil
}
