package subaccount

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	apisv1alpha1 "github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	"github.com/sap/crossplane-provider-btp/btp"
	"github.com/sap/crossplane-provider-btp/internal"
	accountclient "github.com/sap/crossplane-provider-btp/internal/openapi_clients/btp-accounts-service-api-go/pkg"
)

const (
	errNotSubaccount        = "managed resource is not a Subaccount custom resource"
	subaccountStateDeleting = "DELETING"
	subaccountStateOk       = "OK"
)

type SubaccountClient struct {
	btp btp.Client
}

func NewSubaccountClient(btp btp.Client) *SubaccountClient {
	return &SubaccountClient{
		btp: btp,
	}
}

func (c SubaccountClient) DescribeInstance(
	ctx context.Context,
	subaccount *apisv1alpha1.Subaccount,
) (*accountclient.SubaccountResponseObject, error) {
	response, _, err := c.btp.AccountsServiceClient.SubaccountOperationsAPI.GetSubaccounts(ctx).Execute()
	if err != nil {
		return nil, err
	}

	var foundAccount *accountclient.SubaccountResponseObject = nil
	btpSubaccounts := response.Value
	for _, account := range btpSubaccounts {
		if resourceMatchesResponse(subaccount, &account) {
			foundAccount = &account
			break
		}
	}

	return foundAccount, nil
}

func (c SubaccountClient) CreateInstance(
	ctx context.Context, subaccount *apisv1alpha1.Subaccount,
) (*accountclient.SubaccountResponseObject, error) {
	createdSubaccount, _, err := c.btp.AccountsServiceClient.SubaccountOperationsAPI.
		CreateSubaccount(ctx).
		CreateSubaccountRequestPayload(getPayloadFromResource(subaccount)).
		Execute()

	if err != nil {
		return nil, specifyAPIError(err)
	}

	// TODO Should this be done in the controller?
	// TODO Should we check if the subaccount is already created?
	if createdSubaccount == nil {
		return nil, errors.New("subaccount creation failed")
	}
	guid := createdSubaccount.Guid
	subaccount.Status.AtProvider.SubaccountGuid = &guid
	subaccount.Status.AtProvider.Status = createdSubaccount.StateMessage
	subaccount.Status.AtProvider.ParentGuid = &createdSubaccount.ParentGUID

	return createdSubaccount, nil
}

func (c SubaccountClient) DeleteInstance(
	ctx context.Context, subaccount *apisv1alpha1.Subaccount,
) error {

	subaccountId := *subaccount.Status.AtProvider.SubaccountGuid

	response, raw, err := c.btp.AccountsServiceClient.SubaccountOperationsAPI.DeleteSubaccount(ctx, subaccountId).Execute()

	if err != nil {
		return errors.Wrap(err, "deletion of subaccount failed")
	}

	// Nothing to do, deletion is pending or already done
	if raw.StatusCode == 404 || response.State == subaccountStateDeleting {
		return nil
	}

	// If the subaccount is in any other state, we assume that the deletion failed
	return errors.New(fmt.Sprintf("Deletion Pending: Current status: %s", response.State))
}

func (c SubaccountClient) UpdateInstance(ctx context.Context, subaccount *apisv1alpha1.Subaccount) error {
	if directoryParentChanged(&subaccount.Spec.ForProvider, &subaccount.Status.AtProvider) {
		return c.moveSubaccount(ctx, subaccount)
	}
	return c.updateSubaccount(ctx, subaccount)
}

func (c SubaccountClient) updateSubaccount(ctx context.Context, subaccount *apisv1alpha1.Subaccount) error {

	guid := subaccount.Status.AtProvider.SubaccountGuid

	label := addOperatorLabel(subaccount)

	params := accountclient.UpdateSubaccountRequestPayload{
		BetaEnabled:       &subaccount.Spec.ForProvider.BetaEnabled,
		Description:       &subaccount.Spec.ForProvider.Description,
		DisplayName:       subaccount.Spec.ForProvider.DisplayName,
		Labels:            &label,
		UsedForProduction: &subaccount.Spec.ForProvider.UsedForProduction,
	}

	_, _, err := c.btp.AccountsServiceClient.SubaccountOperationsAPI.
		UpdateSubaccount(ctx, internal.Val(guid)).
		UpdateSubaccountRequestPayload(params).
		Execute()

	if err != nil {
		return errors.Wrap(err, "update of subaccount failed")
	}
	return nil
}

func (c SubaccountClient) moveSubaccount(ctx context.Context, subaccount *apisv1alpha1.Subaccount) error {
	guid := subaccount.Status.AtProvider.SubaccountGuid
	targetId := subaccount.Spec.ForProvider.DirectoryGuid
	// if not specified we need to set the global account as parent
	if emptyDirectoryRef(&subaccount.Spec.ForProvider) {
		targetId = internal.Val(subaccount.Status.AtProvider.GlobalAccountGUID)
		if targetId == "" {
			return errors.New("targetId must be set for move subaccount api call")
		}
	}

	_, _, err := c.btp.AccountsServiceClient.SubaccountOperationsAPI.
		MoveSubaccount(ctx, internal.Val(guid)).
		MoveSubaccountRequestPayload(
			accountclient.MoveSubaccountRequestPayload{TargetAccountGUID: targetId}).
		Execute()

	if err != nil {
		return errors.Wrap(err, "moving subaccount failed")
	}
	return nil
}

func directoryParentChanged(spec *apisv1alpha1.SubaccountParameters, status *apisv1alpha1.SubaccountObservation) bool {
	supposeGlobal := emptyDirectoryRef(spec)
	// With no directory specified we expect it to be in global account
	if supposeGlobal {
		return !reflect.DeepEqual(status.ParentGuid, status.GlobalAccountGUID)
	}
	return !reflect.DeepEqual(status.ParentGuid, &spec.DirectoryGuid)
}

func emptyDirectoryRef(spec *apisv1alpha1.SubaccountParameters) bool {
	return spec.DirectoryRef == nil && spec.DirectorySelector == nil && spec.DirectoryGuid == ""
}

// If region and subdomain are equal, then the subaccount resorce is the same as the one in the BTP
func resourceMatchesResponse(subaccount *apisv1alpha1.Subaccount, account *accountclient.SubaccountResponseObject) bool {
	return strings.Compare(
		subaccount.Spec.ForProvider.Subdomain, account.Subdomain,
	) == 0 && strings.Compare(subaccount.Spec.ForProvider.Region, account.Region) == 0
}

func getPayloadFromResource(subaccount *apisv1alpha1.Subaccount) accountclient.CreateSubaccountRequestPayload {
	subaccountSpec := subaccount.Spec

	label := addOperatorLabel(subaccount)

	return accountclient.CreateSubaccountRequestPayload{
		BetaEnabled:       &subaccountSpec.ForProvider.BetaEnabled,
		Description:       &subaccountSpec.ForProvider.Description,
		DisplayName:       subaccountSpec.ForProvider.DisplayName,
		Labels:            &label,
		Region:            subaccountSpec.ForProvider.Region,
		SubaccountAdmins:  subaccountSpec.ForProvider.SubaccountAdmins,
		Subdomain:         &subaccountSpec.ForProvider.Subdomain,
		UsedForProduction: &subaccountSpec.ForProvider.UsedForProduction,
		ParentGUID:        &subaccountSpec.ForProvider.DirectoryGuid,
	}
}

func addOperatorLabel(subaccount *apisv1alpha1.Subaccount) map[string][]string {
	if subaccount.Spec.ForProvider.Labels == nil {
		return map[string][]string{}
	}
	labels := map[string][]string{}
	internal.CopyMaps(labels, subaccount.Spec.ForProvider.Labels)
	labels[apisv1alpha1.SubaccountOperatorLabel] = []string{string(subaccount.UID)}
	return labels
}

func specifyAPIError(err error) error {
	if genericErr, ok := err.(*accountclient.GenericOpenAPIError); ok {
		if accountError, ok := genericErr.Model().(accountclient.ApiExceptionResponseObject); ok {
			return fmt.Errorf("API Error: %v, Code %v", internal.Val(accountError.Error.Message), internal.Val(accountError.Error.Code))
		}
		if genericErr.Body() != nil {
			return fmt.Errorf("API Error: %s", string(genericErr.Body()))
		}
	}
	return err
}
