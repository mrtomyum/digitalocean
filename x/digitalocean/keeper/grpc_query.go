package keeper

import (
	"github.com/mrtomyum/digitalocean/x/digitalocean/types"
)

var _ types.QueryServer = Keeper{}
