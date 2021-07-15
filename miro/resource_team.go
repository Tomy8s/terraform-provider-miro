package miro

import (
	"context"
	
	"github.com/Miro-Ecosystem/go-miro/miro"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Team",
				Type: schema.TypeString,
				Required: true,
			},
		},
		CreateContext: resourceTeamCreate,
		ReadContext: resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
	}
}

func resourceTeamRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	var diags diag.Diagnostics
	
	authInfo, err := c.AuthzInfo.Get(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	team := authInfo.Team
	data.SetId(team.ID)
	data.Set("name", team.Name)

	return diags
}

func resourceTeamCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// The team cannot be created, only have its name updated
	return resourceTeamRead(ctx, data, meta)
}

func resourceTeamUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	name := data.Get("name").(string)
	authInfo, err := c.AuthzInfo.Get(ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	id := authInfo.Team.ID

	req := &miro.UpdateTeamRequest{
		Name: name,
	}

	_, err = c.Teams.Update(ctx, id, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTeamRead(ctx, data, meta)
}

func resourceTeamDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// The team cannot be destroyed
	var diags diag.Diagnostics
	
	return diags
}
