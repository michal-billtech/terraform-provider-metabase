package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/flovouin/terraform-provider-metabase/metabase"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Converts a possibly `nil` string to a Terraform `String` type.
func stringValueOrNull[T ~string](v *T) types.String {
	if v == nil {
		return types.StringNull()
	}

	return types.StringValue(string(*v))
}

// Converts a possibly `nil` integer to a Terraform `Int64` type.
func int64ValueOrNull(v *int) types.Int64 {
	if v == nil {
		return types.Int64Null()
	}

	return types.Int64Value(int64(*v))
}

// Returns the value of a Terraform `String` type, or `nil` if it is null.
func valueStringOrNull(v types.String) *string {
	if v.IsNull() {
		return nil
	}

	r := v.ValueString()
	return &r
}

// Returns the value of a Terraform `Int64` type, or `nil` if it is null.
func valueInt64OrNull(v types.Int64) *int {
	if v.IsNull() {
		return nil
	}

	r := int(v.ValueInt64())
	return &r
}

// Ensures that a Metabase response is not an error and has the expected status code. Otherwise, returns a diagnostic
// error.
func checkMetabaseResponse(r metabase.MetabaseResponse, err error, statusCodes []int, operation string) diag.Diagnostics {
	if err != nil {
		return diag.Diagnostics{
			diag.NewErrorDiagnostic(
				fmt.Sprintf("Unexpected error while calling the Metabase API for operation '%s'.", operation),
				err.Error(),
			),
		}
	}

	for _, s := range statusCodes {
		if r.StatusCode() == s {
			return diag.Diagnostics{}
		}
	}

	return diag.Diagnostics{
		diag.NewErrorDiagnostic(
			fmt.Sprintf("Unexpected response while calling the Metabase API for operation '%s'.", operation),
			fmt.Sprintf("Status code: %d, body: %s", r.StatusCode(), string(r.BodyString())),
		),
	}
}

// Performs the import operation for a resource identified using its `id` integer attribute.
func importStatePassthroughIntegerId(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Unable to convert ID to an integer.", req.ID)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
