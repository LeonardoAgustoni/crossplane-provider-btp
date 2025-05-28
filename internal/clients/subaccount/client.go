package subaccount

import (
	"context"

	apisv1alpha1 "github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	accountclient "github.com/sap/crossplane-provider-btp/internal/openapi_clients/btp-accounts-service-api-go/pkg"
)

type Client interface {
	DescribeInstance(ctx context.Context, cr *apisv1alpha1.Subaccount) error
	CreateInstance(ctx context.Context, cr *apisv1alpha1.Subaccount) error
	DeleteInstance(ctx context.Context, cr *apisv1alpha1.Subaccount) error
	UpdateInstance(ctx context.Context, cr *apisv1alpha1.Subaccount) error
}

func GenerateObservation(
	subaccount *accountclient.SubaccountResponseObject,
) apisv1alpha1.SubaccountObservation {
	observation := apisv1alpha1.SubaccountObservation{}

	if subaccount == nil {
		return observation
	}

	observation.SubaccountGuid = &subaccount.Guid
	observation.Status = &subaccount.State
	observation.StatusMessage = subaccount.StateMessage
	observation.BetaEnabled = &subaccount.BetaEnabled
	observation.Labels = subaccount.Labels
	observation.Description = &subaccount.Description
	observation.Subdomain = &subaccount.Subdomain
	observation.DisplayName = &subaccount.DisplayName
	observation.Region = &subaccount.Region
	observation.UsedForProduction = &subaccount.UsedForProduction
	observation.ParentGuid = &subaccount.ParentGUID
	observation.GlobalAccountGUID = &subaccount.GlobalAccountGUID

	return observation
}
